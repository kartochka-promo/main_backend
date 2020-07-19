package staff

import (
	"2020_1_drop_table/internal/microservices/staff/models"
	"context"
)

type Usecase interface {
	Register(c context.Context, newStaff models.Staff, emailSecretKey string) (models.SafeStaff, error)
	GetByID(c context.Context, id int) (models.SafeStaff, error)
	Update(c context.Context, newStaff models.SafeStaff) (models.SafeStaff, error)
	GetByEmailAndPassword(c context.Context, form models.LoginForm) (models.SafeStaff, error)
	GetFromSession(c context.Context) (models.SafeStaff, error)
	GetQrForStaff(ctx context.Context, idCafe int, position string) (string, error)
	IsOwner(c context.Context, staffId int) (bool, error)
	DeleteQrCodes(uString string) error
	GetCafeId(c context.Context, uuid string) (int, error)
	GetStaffId(c context.Context) (int, error)
	GetStaffListByOwnerId(ctx context.Context, ownerId int) (map[string][]models.StaffByOwnerResponse, error)
	DeleteStaffById(ctx context.Context, staffId int) error
	CheckIfStaffInOwnerCafes(ctx context.Context, requestUser models.SafeStaff, staffId int) (bool, error)
	UpdatePosition(ctx context.Context, id int, position string) error
	SendRegisterEmail(ctx context.Context, email string) error
	ConfirmEmailToStaff(ctx context.Context, email, secretKey string) error
}
