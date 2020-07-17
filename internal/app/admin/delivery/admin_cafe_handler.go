package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	interfaces "2020_1_drop_table/internal/app/admin"
	"2020_1_drop_table/internal/app/admin/models"
	"2020_1_drop_table/internal/pkg/responses"
)

type AdminCafeService struct {
	cafeLogic interfaces.UseCaseCafe
}

func (acs AdminCafeService) GetCafes(rwContext echo.Context) error {
	var (
		limit  int
		offset int
		err    error
		cafes  models.GetCafes
	)

	if limit, err = strconv.Atoi(rwContext.QueryParam("limit")); err != nil {
		return rwContext.NoContent(http.StatusBadRequest)
	}

	if offset, err = strconv.Atoi(rwContext.QueryParam("offset")); err != nil {
		return rwContext.NoContent(http.StatusBadRequest)
	}

	if cafes, err = acs.cafeLogic.GetCafes(limit, offset); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.JSON(http.StatusOK, &cafes)
}

func (acs AdminCafeService) GetCafe(rwContext echo.Context) error {
	var (
		cafeID int
		err    error
		cafe   models.GetCafe
	)

	if cafeID, err = strconv.Atoi(rwContext.QueryParam("cafeID")); err != nil {
		return rwContext.NoContent(http.StatusBadRequest)
	}

	if cafe, err = acs.cafeLogic.GetCafe(cafeID); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.JSON(http.StatusOK, &cafe)
}

func (acs AdminCafeService) UpdateCafe(rwContext echo.Context) error {
	var (
		err            error
		updateCafeJson models.CreateOrUpdateCafe
	)

	if err = rwContext.Bind(&updateCafeJson); err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if err = acs.cafeLogic.UpdateCafe(&updateCafeJson); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.NoContent(http.StatusOK)
}

func (acs AdminCafeService) DeleteCafe(rwContext echo.Context) error {
	var (
		cafeID int
		err    error
	)

	if cafeID, err = strconv.Atoi(rwContext.QueryParam("cafeID")); err != nil {
		return rwContext.NoContent(http.StatusBadRequest)
	}

	if err = acs.cafeLogic.DeleteCafe(cafeID); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.NoContent(http.StatusOK)
}
