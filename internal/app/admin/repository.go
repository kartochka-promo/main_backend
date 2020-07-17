package admin

import "2020_1_drop_table/internal/app/admin/models"

type RepositoryAdmin interface {
	CreateAdmin(createAdminRequest *models.CreateOrUpdateAdmin) error
	GetAdminHashedPassword(getPasswordRequest *models.LogAdmin) (string, error)
	DeleteAdmin(deleteAdminRequest *models.DeleteAdmin) error
	UpdateAdmin(updateAdminRequest *models.CreateOrUpdateAdmin) error
}

type RepositoryCafe interface {
	GetCafes(limit, offset int) (models.GetCafes, error)
	GetCafe(getCafeID int) (models.GetCafe, error)
	UpdateCafe(updateCafeRequest *models.CreateOrUpdateCafe) error
	DeleteCafe(deleteCafeID int) error
}
