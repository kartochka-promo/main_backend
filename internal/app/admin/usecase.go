package admin

import (
	models "2020_1_drop_table/internal/app/admin/models"
)

type UseCase interface {
	CreateAdmin(admin *models.CreateOrUpdateAdmin) error
	Authentication(authAdminRequest *models.LogAdmin) (bool, error)
	UpdateAdmin(updateAdminRequest *models.CreateOrUpdateAdmin) error
}
