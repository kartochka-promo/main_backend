package mail

import (
	"2020_1_drop_table/internal/microservices/mail/models"
	"fmt"
	"github.com/k3a/html2text"
	"strings"
)

// InsertValuesInTemplate function that
// formatting email template with replacing all vars with values
//
// Example:
// <bold>Thank you for registration, {{username}}</bold>
//                          ↓↓↓
// <bold>Thank you for registration, Alex</bold>
func InsertValuesInTemplate(template models.EmailTemplate,
	emailContext map[string]string) models.EmailTemplate {

	replaceValues := make([]string, 0, len(emailContext)*2)
	for key, value := range emailContext {
		valueName := fmt.Sprintf("{{%s}}", key)

		// Protection from injections
		plain := html2text.HTML2Text(value)

		replaceValues = append(replaceValues, valueName)
		replaceValues = append(replaceValues, plain)
	}
	r := strings.NewReplacer(replaceValues...)
	var newMail models.EmailTemplate

	newMail.MailContent = r.Replace(template.MailContent)
	newMail.MailTitle = r.Replace(template.MailTitle)
	newMail.TemplateName = template.TemplateName
	return newMail
}
