package customer

import (
	"context"
)

type Usecase interface {
	GetPoints(ctx context.Context, uuid string) (int, error)
	SetPoints(ctx context.Context, uuid string, points int) error
}
