package admin

import "2020_1_drop_table/internal/app/admin/models"

type Repository interface {
	CreateAdmin(createAdminRequest *models.CreateOrUpdateAdmin) error
	GetAdminHashedPassword(getPasswordRequest *models.LogAdmin) (string, error)
	DeleteAdmin(deleteAdminRequest *models.DeleteAdmin) error
	UpdateAdmin(updateAdminRequest *models.CreateOrUpdateAdmin) error
}
