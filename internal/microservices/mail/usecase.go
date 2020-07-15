package mail

import "context"

type Usecase interface {
	SendEmail(c context.Context, recipient, templateName string, emailContext map[string]string) error
}
