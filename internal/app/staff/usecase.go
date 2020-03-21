package staff

import (
	"2020_1_drop_table/internal/app/staff/models"
	"context"
)

type Usecase interface {
	Add(c context.Context, newStaff models.Staff) (models.SafeStaff, error)
	GetByID(c context.Context, id int) (models.SafeStaff, error)
	Update(c context.Context, newStaff models.SafeStaff) (models.SafeStaff, error)
	GetByEmailAndPassword(c context.Context, form models.LoginForm) (models.SafeStaff, error)
	GetFromSession(c context.Context) (models.SafeStaff, error)
}