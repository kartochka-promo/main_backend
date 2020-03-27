package apple_passkit

import (
	"2020_1_drop_table/internal/app/apple_passkit/models"
	"context"
)

type Repository interface {
	Add(ctx context.Context, ap models.ApplePassDB) (models.ApplePassDB, error)
	GetByID(ctx context.Context, id int) (models.ApplePassDB, error)
	Update(ctx context.Context, newApplePass models.ApplePassDB) error
	UpdateDesign(ctx context.Context, Design string, id int) error
	Delete(ctx context.Context, id int) error
}