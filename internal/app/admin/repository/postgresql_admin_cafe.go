package repository

import (
	"2020_1_drop_table/internal/app/admin/models"
	"github.com/jackc/pgx"
	"strconv"
)

type AdminCafeStorage struct {
	dbPool *pgx.ConnPool
}

func (acs AdminCafeStorage) GetCafes(limit, offset int) (models.GetCafes, error) {
	var (
		getCafesResponse      models.GetCafes
		getCafeStorageRequest = "SELECT CafeID, CafeName, Address, Description, StaffID, OpenTime, CloseTime, Photo, location_str" +
			" FROM Cafe ORDER BY CafeID LIMIT $1 OFFSET $2;"
		err              error
		databaseResponse *pgx.Rows
	)
	getCafesResponse.Cafes = make([]models.GetCafe, 0, 0)
	if databaseResponse, err = acs.dbPool.Query(getCafeStorageRequest, limit, offset); err != nil {
		return getCafesResponse, err
	}

	for databaseResponse.Next() {
		var (
			cafeInstance models.GetCafe
		)
		if err = databaseResponse.Scan(&cafeInstance.CafeID, &cafeInstance.CafeName, &cafeInstance.Address,
			&cafeInstance.Description, &cafeInstance.StaffID, &cafeInstance.OpenTime, &cafeInstance.CloseTime,
			&cafeInstance.Photo, &cafeInstance.Location); err != nil {
			return getCafesResponse, err
		}
		getCafesResponse.Cafes = append(getCafesResponse.Cafes, cafeInstance)
	}

	return getCafesResponse, err
}

func (acs AdminCafeStorage) GetCafe(getCafeID int) (models.GetCafe, error) {
	var (
		getCafesResponse      models.GetCafe
		getCafeStorageRequest = "SELECT CafeID, CafeName, Address, Description, StaffID, OpenTime, CloseTime, Photo, location_str" +
			" FROM Cafe WHERE CafeID = $1;"
		err error
	)
	err = acs.dbPool.QueryRow(getCafeStorageRequest, getCafeID).Scan(&getCafesResponse.CafeID, &getCafesResponse.CafeName, &getCafesResponse.Address,
		&getCafesResponse.Description, &getCafesResponse.StaffID, &getCafesResponse.OpenTime, &getCafesResponse.CloseTime,
		&getCafesResponse.Photo, &getCafesResponse.Location)
	return getCafesResponse, err
}

func (acs AdminCafeStorage) UpdateCafe(updateCafeRequest *models.CreateOrUpdateCafe) error {
	var (
		updateCafeStorageBodyRequest = "UPDATE Cafe SET "
		updateCafeStorageTailRequest = " WHERE CafeID = $1;"
		inputIterator                = 1
		inputData                    = make([]interface{}, 0, 0)
		err                          error
	)
	inputData = append(inputData, updateCafeRequest.CafeID)

	if updateCafeRequest.CafeName != "" {
		inputIterator++
		inputData = append(inputData, updateCafeRequest.CafeName)
		updateCafeStorageBodyRequest += "CafeName = $" + strconv.Itoa(inputIterator) + ","
	}
	if updateCafeRequest.Address != "" {
		inputIterator++
		inputData = append(inputData, updateCafeRequest.Address)

		updateCafeStorageBodyRequest += "Address = $" + strconv.Itoa(inputIterator) + ","
	}
	if updateCafeRequest.Description != "" {
		inputIterator++
		inputData = append(inputData, updateCafeRequest.Description)

		updateCafeStorageBodyRequest += "Description = $" + strconv.Itoa(inputIterator) + ","
	}
	if updateCafeRequest.StaffID != 0 {
		inputIterator++
		inputData = append(inputData, updateCafeRequest.StaffID)
		updateCafeStorageBodyRequest += "StaffID = $" + strconv.Itoa(inputIterator) + ","
	}
	if updateCafeRequest.OpenTime.IsZero() {
		inputIterator++
		inputData = append(inputData, updateCafeRequest.OpenTime)
		updateCafeStorageBodyRequest += "OpenTime = $" + strconv.Itoa(inputIterator) + ","
	}
	if updateCafeRequest.CloseTime.IsZero() {
		inputIterator++
		inputData = append(inputData, updateCafeRequest.CloseTime)
		updateCafeStorageBodyRequest += "CloseTime = $" + strconv.Itoa(inputIterator) + ","
	}
	if updateCafeRequest.Photo != "" {
		inputIterator++
		inputData = append(inputData, updateCafeRequest.Photo)
		updateCafeStorageBodyRequest += "Photo = $" + strconv.Itoa(inputIterator) + ","
	}
	updateCafeStorageBodyRequest = updateCafeStorageBodyRequest[:len(updateCafeStorageBodyRequest)-1]

	_, err = acs.dbPool.Exec(updateCafeStorageBodyRequest+updateCafeStorageTailRequest, inputData...)
	return err
}

func (acs AdminCafeStorage) DeleteCafe(deleteCafeID int) error {
	var (
		deleteCafeStorageRequest = "DELETE FROM Cafe WHERE CafeID = $1;"
	)
	_, err := acs.dbPool.Exec(deleteCafeStorageRequest, deleteCafeID)
	return err
}
