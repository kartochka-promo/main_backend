package admin

import (
	models "2020_1_drop_table/internal/app/admin/models"
)

type UseCaseAdmin interface {
	CreateAdmin(admin *models.CreateOrUpdateAdmin) error
	Authentication(authAdminRequest *models.LogAdmin) (bool, error)
	UpdateAdmin(updateAdminRequest *models.CreateOrUpdateAdmin) error
}

type UseCaseCafe interface {
	GetCafes(limit, offset int) (models.GetCafes, error)
	GetCafe(getCafeID int) (models.GetCafe, error)
	UpdateCafe(updateCafeRequest *models.CreateOrUpdateCafe) error
	DeleteCafe(deleteCafeID int) error
}
