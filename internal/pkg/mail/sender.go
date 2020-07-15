package mail

import (
	"2020_1_drop_table/internal/microservices/mail/models"
	"errors"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(recipient string, email models.EmailTemplate) error {
	MailSenderName := os.Getenv("MAIL_SENDER_NAME")
	MailSenderEmail := os.Getenv("MAIL_SENDER_EMAIL")
	from := mail.NewEmail(MailSenderName, MailSenderEmail)

	to := mail.NewEmail("", recipient)

	message := mail.NewSingleEmail(from, email.MailTitle, to, email.MailContent, email.MailContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		return err
	}
	if response.StatusCode != 202 {
		return errors.New(response.Body)
	}
	return nil
}
