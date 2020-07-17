package usecase

import (
	interfaces "2020_1_drop_table/internal/app/admin"
	"2020_1_drop_table/internal/app/admin/models"
)

type AdminCafeLogic struct {
	cafeStorage interfaces.RepositoryCafe
}

func (acl *AdminCafeLogic) GetCafes(limit, offset int) (models.GetCafes, error) {
	return acl.cafeStorage.GetCafes(limit,offset)
}

func (acl *AdminCafeLogic) GetCafe(getCafeID int) (models.GetCafe, error) {
	return acl.cafeStorage.GetCafe(getCafeID)
}

func (acl *AdminCafeLogic) UpdateCafe(updateCafeRequest *models.CreateOrUpdateCafe) error {
	return acl.cafeStorage.UpdateCafe(updateCafeRequest)
}

func (acl *AdminCafeLogic) DeleteCafe(deleteCafeID int) error{
	return acl.cafeStorage.DeleteCafe(deleteCafeID)
}
