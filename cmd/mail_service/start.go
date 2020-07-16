package main

import (
	"2020_1_drop_table/configs"
	mailGRPCServer "2020_1_drop_table/internal/microservices/mail/delivery/grpc/server"
	mailRepo "2020_1_drop_table/internal/microservices/mail/repository"
	mailUseCase "2020_1_drop_table/internal/microservices/mail/usecase"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

func main() {
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
	mailRepository := mailRepo.NewPostgresMailRepository(conn)
	mailUsecase := mailUseCase.NewMailUsecase(&mailRepository, timeoutContext)

	list, err := net.Listen("tcp", configs.GRPCEmailUrl)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}
	server := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)
	mailGRPCServer.NewMailServerGRPC(server, mailUsecase)
	log.Info().Msgf("GRPC mail server started at %s", configs.GRPCEmailUrl)
	err = server.Serve(list)
	if err != nil {
		log.Error().Msgf(err.Error())
		return
	}
}
