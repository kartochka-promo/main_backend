package usecase

import (
	"2020_1_drop_table/internal/microservices/mail"
	mailPkg "2020_1_drop_table/internal/pkg/mail"
	"context"
	"time"
)

type mailUsecase struct {
	mailRepo       mail.Repository
	contextTimeout time.Duration
}

func NewMailUsecase(m mail.Repository, timeout time.Duration) mail.Usecase {
	return &mailUsecase{
		mailRepo:       m,
		contextTimeout: timeout,
	}
}

func (m *mailUsecase) SendEmail(c context.Context, recipient, templateName string,
	emailContext map[string]string) error {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	template, err := m.mailRepo.GetMailByName(ctx, templateName)
	if err != nil {
		return err
	}

	email := mailPkg.InsertValuesInTemplate(template, emailContext)
	err = mailPkg.SendEmail(recipient, email)
	return err
}
