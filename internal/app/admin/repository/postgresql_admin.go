package repository

import (
	"2020_1_drop_table/internal/app/admin/models"
	"github.com/jackc/pgx"
)

type AdminStorage struct {
	dbPool *pgx.ConnPool
}

func (as *AdminStorage) CreateAdmin(createAdminRequest *models.CreateOrUpdateAdmin) error {
	_, err := as.dbPool.Exec("INSERT INTO admins (username , password) VALUES ($1,$2)", createAdminRequest.Username, createAdminRequest.HashedPassword)
	return err
}

func (as *AdminStorage) GetAdminHashedPassword(getPasswordRequest *models.LogAdmin) (string, error) {
	var (
		password string
	)
	return password,
		as.dbPool.QueryRow(
			"SELECT password FROM admins WHERE username = $1",
			getPasswordRequest.Username,
		).Scan(&password)
}

func (as *AdminStorage) DeleteAdmin(deleteAdminRequest *models.DeleteAdmin) error {
	_, err := as.dbPool.Exec("DELETE FROM admins WHERE admin_id = $1", deleteAdminRequest.AdminId)
	return err
}

func (as *AdminStorage) UpdateAdmin(updateAdminRequest *models.CreateOrUpdateAdmin) error {
	_, err := as.dbPool.Exec(
		"UPDATE admins SET password = $1 WHERE username = $2",
		updateAdminRequest.HashedPassword,
		updateAdminRequest.Username,
	)
	return err
}