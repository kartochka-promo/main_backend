package mail

import (
	"2020_1_drop_table/internal/microservices/mail/models"
	"context"
)

type Repository interface {
	InsertNewTemplate(ctx context.Context, newEmailTemplate models.EmailTemplate) (models.EmailTemplate, error)
	UpdateTemplate(ctx context.Context, newEmailTemplate models.EmailTemplate) (models.EmailTemplate, error)
	GetMailByName(ctx context.Context, templateName string) (models.EmailTemplate, error)
}
