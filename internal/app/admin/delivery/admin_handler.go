package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	interfaces "2020_1_drop_table/internal/app/admin"
	models "2020_1_drop_table/internal/app/admin/models"
	responses "2020_1_drop_table/internal/pkg/responses"
)

type AdminService struct {
	adminLogic interfaces.UseCase
}

func (as AdminService) CreateAdmin(rwContext echo.Context) error {
	var (
		err             error
		createAdminJson = new(models.CreateOrUpdateAdmin)
	)
	err = rwContext.Bind(&createAdminJson)

	if err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if err = as.adminLogic.CreateAdmin(createAdminJson); err != nil {
		return rwContext.JSON(http.StatusConflict, &responses.HttpError{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	}
	// todo : add token from reddis
	return rwContext.NoContent(http.StatusOK)
}

func (as AdminService) Authentication(rwContext echo.Context) error {
	var (
		err           error
		authAdminJson = new(models.LogAdmin)
	)
	err = rwContext.Bind(&authAdminJson)

	if err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if flag, err := as.adminLogic.Authentication(authAdminJson); !flag || err != nil {
		responseMessage := "Password are not equal"
		if err != nil {
			responseMessage = err.Error()
		}
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: responseMessage,
		})
	}
	// todo : add token from redis
	return rwContext.NoContent(http.StatusOK)
}

func (as AdminService) UpdateAdmin(rwContext echo.Context) error {
	var (
		err             error
		updateAdminJson = new(models.CreateOrUpdateAdmin)
	)
	err = rwContext.Bind(&updateAdminJson)

	if err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if err = as.adminLogic.UpdateAdmin(updateAdminJson); err != nil {
		return rwContext.JSON(http.StatusConflict, &responses.HttpError{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	}
	return rwContext.NoContent(http.StatusOK)
}

func (as AdminService) Logout(rwContext echo.Context) error {
	// todo : delete token from redis
	return rwContext.NoContent(http.StatusOK)
}


func NewAdminService(adminLogic interfaces.UseCase) *AdminService {
	return &AdminService{adminLogic: adminLogic}
}
