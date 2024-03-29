package main

import (
	"2020_1_drop_table/configs"
	_appleHttpDeliver "2020_1_drop_table/internal/app/apple_passkit/delivery/http"
	_appleRepo "2020_1_drop_table/internal/app/apple_passkit/repository"
	_appleUsecase "2020_1_drop_table/internal/app/apple_passkit/usecase"
	cafeClient "2020_1_drop_table/internal/app/cafe/delivery/grpc/client"
	"2020_1_drop_table/internal/app/cafe/delivery/grpc/server"
	_cafeHttpDeliver "2020_1_drop_table/internal/app/cafe/delivery/http"
	_cafeRepo "2020_1_drop_table/internal/app/cafe/repository"
	_cafeUsecase "2020_1_drop_table/internal/app/cafe/usecase"
	server2 "2020_1_drop_table/internal/app/customer/delivery/grpc/server"
	_customerHttpDeliver "2020_1_drop_table/internal/app/customer/delivery/http"
	_customerRepo "2020_1_drop_table/internal/app/customer/repository"
	_customerUseCase "2020_1_drop_table/internal/app/customer/usecase"
	"2020_1_drop_table/internal/app/middleware"
	http2 "2020_1_drop_table/internal/app/statistics/delivery/http"
	"2020_1_drop_table/internal/app/statistics/repository"
	"2020_1_drop_table/internal/app/statistics/usecase"
	staffClient "2020_1_drop_table/internal/microservices/staff/delivery/grpc/client"
	"2020_1_drop_table/internal/pkg/apple_pass_generator"
	"2020_1_drop_table/internal/pkg/apple_pass_generator/meta"
	geo "2020_1_drop_table/internal/pkg/google_geocoder"
	"2020_1_drop_table/internal/pkg/metrics"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	redisStore "gopkg.in/boj/redistore.v1"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	//PromMetrics server
	metricsProm := metrics.RegisterMetrics(r)

	//Middleware
	var CookieStore, err = redisStore.NewRediStore(
		configs.RedisPreferences.Size,
		configs.RedisPreferences.Network,
		configs.RedisPreferences.Address,
		configs.RedisPreferences.Password,
		configs.RedisPreferences.SecretKey)

	middleware.NewMiddleware(r, CookieStore, metricsProm)

	timeoutContext := configs.Timeouts.ContextTimeout

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=%s host=%s",
		configs.PostgresPreferences.User,
		configs.PostgresPreferences.Password,
		configs.PostgresPreferences.DBName,
		configs.PostgresPreferences.Port,
		configs.PostgresPreferences.Host)

	conn, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}

	cafeRepo := _cafeRepo.NewPostgresCafeRepository(conn)
	grpcConn, err := grpc.Dial(configs.GRPCStaffUrl, grpc.WithInsecure())
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	grpcStaffClient := staffClient.NewStaffClient(grpcConn)

	geoCoder := geo.NewGoogleGeoCoder(configs.GoogleMapAPIKey, "ru", "ru")

	applePassGenerator := apple_pass_generator.NewGenerator(
		configs.AppleWWDR, configs.AppleCertificate, configs.AppleKey, configs.ApplePassword)

	customerRepo := _customerRepo.NewPostgresCustomerRepository(conn)
	applePassKitRepo := _appleRepo.NewPostgresApplePassRepository(conn)

	statRepo := repository.NewPostgresStatisticsRepository(conn)

	grpcCafeConn, err := grpc.Dial(configs.GRPCCafeUrl, grpc.WithInsecure())
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	grpcCafeClient := cafeClient.NewCafeClient(grpcCafeConn)

	statUcase := usecase.NewStatisticsUsecase(statRepo, grpcStaffClient, grpcCafeClient, timeoutContext)
	http2.NewStatisticsHandler(r, statUcase)
	customerUseCase := _customerUseCase.NewCustomerUsecase(customerRepo, applePassKitRepo, grpcStaffClient,
		timeoutContext, statUcase)
	_customerHttpDeliver.NewCustomerHandler(r, customerUseCase)

	customerUCase := _customerUseCase.NewCustomerUsecase(customerRepo, applePassKitRepo, grpcStaffClient,
		timeoutContext, statUcase)

	applePassKitUcase := _appleUsecase.NewApplePassKitUsecase(applePassKitRepo, cafeRepo, customerUCase,
		&applePassGenerator, timeoutContext, &meta.Meta{})

	_appleHttpDeliver.NewPassKitHandler(r, applePassKitUcase)

	cafeUsecase := _cafeUsecase.NewCafeUsecase(cafeRepo, grpcStaffClient, timeoutContext, geoCoder, applePassKitUcase)
	_cafeHttpDeliver.NewCafeHandler(r, cafeUsecase)

	go server.StartCafeGrpcServer(cafeUsecase, configs.GRPCCafeUrl)
	go server2.StartCustomerGrpcServer(customerUseCase, configs.GRPCCustomerUrl)

	//OPTIONS
	middleware.AddOptionsRequest(r)

	//static server
	r.PathPrefix(fmt.Sprintf("/%s/", configs.MediaFolder)).Handler(
		http.StripPrefix(fmt.Sprintf("/%s/", configs.MediaFolder),
			http.FileServer(http.Dir(configs.MediaFolder))))

	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         configs.MainService,
		WriteTimeout: configs.Timeouts.WriteTimeout,
		ReadTimeout:  configs.Timeouts.ReadTimeout,
	}
	fmt.Println("main server started at ", configs.MainService)
	log.Error().Msgf(srv.ListenAndServe().Error())
}
