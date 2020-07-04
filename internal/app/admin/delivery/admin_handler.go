package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	models "2020_1_drop_table/internal/app/admin/models"
)

type AdminService struct {
}

func (as *AdminService) CreateAdmin(rwContext echo.Context) error {
	var (
		err             error
		createAdminJson = new(models.CreateOrUpdateAdmin)
	)
	err = rwContext.Bind(&createAdminJson)

	if err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	// todo : impliment

	return rwContext.NoContent(http.StatusOK)
}
