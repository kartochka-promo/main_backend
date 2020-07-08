package usecase

import (
	interfaces "2020_1_drop_table/internal/app/admin"
	models "2020_1_drop_table/internal/app/admin/models"
	hasher "2020_1_drop_table/internal/pkg/hasher"
)

type AdminLogic struct {
	adminStorage interfaces.Repository
}

func (al *AdminLogic) hashPassword(currentPassword string) (string, error) {
	return hasher.HashAndSalt(nil, currentPassword)
}

func (al *AdminLogic) CreateAdmin(createAdminRequest *models.CreateOrUpdateAdmin) error {
	var (
		err error
	)
	if createAdminRequest.HashedPassword, err = al.hashPassword(createAdminRequest.Password);
		err != nil {
		return err
	}
	return al.adminStorage.CreateAdmin(createAdminRequest)
}

func (al *AdminLogic) Authentication(authAdminRequest *models.LogAdmin) (bool, error) {
	var (
		dbPass string
		err    error
	)
	if dbPass, err = al.adminStorage.GetAdminHashedPassword(authAdminRequest); err != nil {
		return false, err
	}

	if !hasher.CheckWithHash(dbPass, authAdminRequest.Password) {
		return false, err
	}
	return true, err
}

func (al *AdminLogic) UpdateAdmin(updateAdminRequest *models.CreateOrUpdateAdmin) error {
	return al.adminStorage.UpdateAdmin(updateAdminRequest)
}
