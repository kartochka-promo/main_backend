package repository

import (
	"2020_1_drop_table/app/staff/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type postgresStaffRepository struct {
	Conn *sqlx.DB
}

func NewPostgresStaffRepository(user string, password string, port string) (postgresStaffRepository, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=postgres sslmode=disable port=%s", user, password, port)
	conn, err := sqlx.Open("postgres", connStr)
	cafeStorage := postgresStaffRepository{conn}
	return cafeStorage, err
}

func (p *postgresStaffRepository) Add(ctx context.Context, st models.Staff) (models.Staff, error) {
	query := `INSERT into staff(name, email, password, editedat, photo, isowner) VALUES ($1,$2,$3,$4,$5,$6) RETURNING *`

	var dbStaff models.Staff
	err := p.Conn.GetContext(ctx, &dbStaff, query, st.Name, st.Email, st.Password, st.EditedAt, st.Photo, st.IsOwner)
	return dbStaff, err
}

func (p *postgresStaffRepository) GetByEmailAndPassword(ctx context.Context,
	email string, password string) (models.Staff, bool, error) {

	query := `SELECT * FROM Staff WHERE password=$1 AND email=$2`

	var dbStaff models.Staff
	err := p.Conn.GetContext(ctx, &dbStaff, query, password, email)

	switch err {
	case nil:
		return dbStaff, true, err
	case sql.ErrNoRows:
		return models.Staff{}, false, err
	default:
		return models.Staff{}, false, err
	}
}

func (p *postgresStaffRepository) GetById(ctx context.Context, id int) (models.Staff, error) {
	query := `SELECT * FROM Staff WHERE StaffID=$1`

	var dbStaff models.Staff
	err := p.Conn.GetContext(ctx, &dbStaff, query, id)

	if err != nil {
		return models.Staff{}, err
	}

	return dbStaff, nil
}

func (p *postgresStaffRepository) Update(ctx context.Context, newStaff models.SafeStaff) error {
	query := `UPDATE Staff SET name=$1,email=$2,editedat=$3,photo=$4 WHERE staffid = $5`
	_, err := p.Conn.ExecContext(ctx, query, newStaff.Name, newStaff.Email, newStaff.EditedAt,
		newStaff.Photo, newStaff.StaffID)

	return err
}
