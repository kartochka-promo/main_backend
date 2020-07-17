package delivery

import (
	interfaces "2020_1_drop_table/internal/app/admin"
	"2020_1_drop_table/internal/app/admin/models"
	"2020_1_drop_table/internal/pkg/responses"
	"net/http"
	"time"

	//"2020_1_drop_table/internal/app/admin/models"
	sessionInterfaces "2020_1_drop_table/internal/app/admin_session"
	//"2020_1_drop_table/internal/pkg/responses"
	"github.com/labstack/echo"
	//"net/http"
	//"time"
)

type AdminMainService struct {
	adminLogic   interfaces.UseCaseAdmin
	sessionLogic sessionInterfaces.UseCase
}

func (ams AdminMainService) CreateAdmin(rwContext echo.Context) error {
	var (
		err             error
		createAdminJson = new(models.CreateOrUpdateAdmin)
		cookieValue     string
	)

	if err = rwContext.Bind(&createAdminJson); err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if err = ams.adminLogic.CreateAdmin(createAdminJson); err != nil {
		return rwContext.JSON(http.StatusConflict, &responses.HttpError{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	}

	if cookieValue, err = ams.sessionLogic.CreateSession(createAdminJson.Username); err != nil {
		return rwContext.JSON(http.StatusExpectationFailed, &responses.HttpError{
			Code:    http.StatusExpectationFailed,
			Message: err.Error(),
		})
	}
	rwContext.SetCookie(&http.Cookie{
		Path:    "/admin",
		Name:    "cookie_value",
		Value:   cookieValue,
		Expires: time.Now().Add(24 * time.Hour),
	})
	return rwContext.NoContent(http.StatusOK)
}

func (ams AdminMainService) Authentication(rwContext echo.Context) error {
	var (
		err           error
		authAdminJson = new(models.LogAdmin)
		cookieValue   string
	)

	if err = rwContext.Bind(&authAdminJson); err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if flag, err := ams.adminLogic.Authentication(authAdminJson); !flag || err != nil {
		responseMessage := "Password are not equal"
		if err != nil {
			responseMessage = err.Error()
		}
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: responseMessage,
		})
	}

	if cookieValue, err = ams.sessionLogic.CreateSession(authAdminJson.Username); err != nil {
		return rwContext.JSON(http.StatusExpectationFailed, &responses.HttpError{
			Code:    http.StatusExpectationFailed,
			Message: err.Error(),
		})
	}
	rwContext.SetCookie(&http.Cookie{
		Path:    "/admin",
		Name:    "cookie_value",
		Value:   cookieValue,
		Expires: time.Now().Add(24 * time.Hour),
	})
	return rwContext.NoContent(http.StatusOK)
}

func (ams AdminMainService) UpdateAdmin(rwContext echo.Context) error {
	var (
		err             error
		updateAdminJson = new(models.CreateOrUpdateAdmin)
	)
	updateAdminJson.Username = rwContext.Get("username").(string)
	if err = rwContext.Bind(&updateAdminJson); err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if err = ams.adminLogic.UpdateAdmin(updateAdminJson); err != nil {
		return rwContext.JSON(http.StatusConflict, &responses.HttpError{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	}
	return rwContext.NoContent(http.StatusOK)
}

func (ams AdminMainService) Logout(rwContext echo.Context) error {
	var (
		cookieValue *http.Cookie
		err         error
	)

	if cookieValue, err = rwContext.Cookie("cookie_value"); err != nil {
		return rwContext.JSON(http.StatusExpectationFailed, &responses.HttpError{
			Code:    http.StatusExpectationFailed,
			Message: err.Error(),
		})
	}

	if err = ams.sessionLogic.DeleteSession(cookieValue.Value); err != nil {
		return rwContext.JSON(http.StatusConflict, &responses.HttpError{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	}
	return rwContext.NoContent(http.StatusOK)
}

func NewAdminMainService(adminLogic interfaces.UseCaseAdmin) *AdminMainService {
	return &AdminMainService{adminLogic: adminLogic}
}
