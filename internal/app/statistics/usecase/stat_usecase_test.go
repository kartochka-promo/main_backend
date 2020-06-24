package usecase

import (
	"2020_1_drop_table/configs"
	cafeClientGRPCMock "2020_1_drop_table/internal/app/cafe/delivery/grpc/client/mocks"
	"2020_1_drop_table/internal/app/statistics/mocks"
	"2020_1_drop_table/internal/app/statistics/models"
	staffClientGRPCMock "2020_1_drop_table/internal/microservices/staff/delivery/grpc/client/mocks"
	models2 "2020_1_drop_table/internal/microservices/staff/models"
	"context"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestGetSurvey(t *testing.T) {

	type CheckStructInput struct {
		Ctx     context.Context
		StaffID int
		Limit   int
		Since   int
	}
	type CheckStructOutput struct {
		Stat []models.StatisticsStruct
		Err  error
	}

	type testCaseStruct struct {
		InputData           CheckStructInput
		OutputData          CheckStructOutput
		StaffFromSession    models2.SafeStaff
		StaffFromSessionErr error
	}

	session := sessions.Session{Values: map[interface{}]interface{}{"userID": 228}}
	c := context.WithValue(context.Background(), configs.SessionStaffID, &session)

	firstStatData := []models.StatisticsStruct{{
		JsonData:   "valid",
		Time:       time.Time{},
		ClientUUID: "asdasd",
		StaffId:    229,
		CafeId:     2,
	}}

	testCases := []testCaseStruct{

		//all ok case
		{
			InputData: CheckStructInput{
				Ctx:     c,
				StaffID: 229,
				Limit:   5,
				Since:   0,
			},
			OutputData: CheckStructOutput{
				Err:  nil,
				Stat: firstStatData,
			},
			StaffFromSession: models2.SafeStaff{
				StaffID:  123,
				Name:     "",
				Email:    "",
				EditedAt: time.Time{},
				Photo:    "",
				IsOwner:  true,
				CafeId:   2,
				Position: "",
			},
		},
	}

	timeout := time.Second * 4
	statRepo := new(mocks.Repository)
	cafeRepo := new(cafeClientGRPCMock.CafeGRPCClientInterface)
	staffUsecase := new(staffClientGRPCMock.StaffClientInterface)
	s := NewStatisticsUsecase(statRepo, staffUsecase, cafeRepo, timeout)

	for _, testCase := range testCases {
		staffUsecase.On("GetFromSession", mock.AnythingOfType("*context.timerCtx")).Return(testCase.StaffFromSession, testCase.StaffFromSessionErr)
		statRepo.On("GetWorkerDataFromRepo", mock.AnythingOfType("*context.timerCtx"), testCase.InputData.StaffID, testCase.InputData.Limit, testCase.InputData.Since).Return(testCase.OutputData.Stat, testCase.OutputData.Err)
		stat, err := s.GetWorkerData(testCase.InputData.Ctx, testCase.InputData.StaffID, testCase.InputData.Limit, testCase.InputData.Since)
		assert.Equal(t, testCase.OutputData.Err, err)
		assert.Equal(t, testCase.OutputData.Stat, stat)
	}
}
