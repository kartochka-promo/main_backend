package delivery

import (
	interfaces "2020_1_drop_table/internal/app/admin"
	"2020_1_drop_table/internal/app/admin/models"
	"2020_1_drop_table/internal/pkg/responses"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type AdminStatisticService struct {
	statisticLogic interfaces.UseCaseStatistics
}

func (ass *AdminStatisticService) GetStatistic(rwContext echo.Context) error {
	var (
		limit  int
		offset int
		err    error
		stats  models.GetStatisticsOutput
	)

	if limit, err = strconv.Atoi(rwContext.QueryParam("limit")); err != nil {
		return rwContext.NoContent(http.StatusBadRequest)
	}

	if offset, err = strconv.Atoi(rwContext.QueryParam("offset")); err != nil {
		return rwContext.NoContent(http.StatusBadRequest)
	}

	if stats, err = ass.statisticLogic.GetStatisticData(limit, offset); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.JSON(http.StatusOK, &stats)
}

func (ass *AdminStatisticService) UpdateStatistic(rwContext echo.Context) error {
	var (
		err                 error
		updateStatisticJson models.UpdateStatistic
	)

	if err = rwContext.Bind(&updateStatisticJson); err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if err = ass.statisticLogic.UpdateStatistic(&updateStatisticJson); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.NoContent(http.StatusOK)
}

func (ass *AdminStatisticService) CreateStatistic(rwContext echo.Context) error {
	var (
		err                 error
		createStatisticJson models.StatisticsStruct
	)

	if err = rwContext.Bind(&createStatisticJson); err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if err = ass.statisticLogic.CreateStatisticData(&createStatisticJson); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.NoContent(http.StatusOK)
}

func (ass *AdminStatisticService) DeleteStatistic(rwContext echo.Context) error {
	var (
		err                 error
		deleteStatisticJson models.DeleteStatistic
	)

	if err = rwContext.Bind(&deleteStatisticJson); err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if err = ass.statisticLogic.DeleteStatistic(&deleteStatisticJson); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.NoContent(http.StatusOK)
}
