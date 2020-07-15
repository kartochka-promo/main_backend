package repository

import (
	"2020_1_drop_table/internal/microservices/mail/models"
	"context"
	"github.com/jmoiron/sqlx"
)

type PostgresMailRepository struct {
	conn *sqlx.DB
}

func NewPostgresMailRepository(conn *sqlx.DB) PostgresMailRepository {
	return PostgresMailRepository{
		conn: conn,
	}
}

func (p *PostgresMailRepository) InsertNewTemplate(ctx context.Context,
	newEmailTemplate models.EmailTemplate) (models.EmailTemplate, error) {
	query := `INSERT INTO EmailTemplate(
				 TemplateName, 
				 MailTitle, 
				 MailContent)
				 VALUES ($1, $2, $3)
				 RETURNING  TemplateName, MailTitle, MailContent`

	var dbEmailTemplate models.EmailTemplate
	err := p.conn.GetContext(ctx, &dbEmailTemplate, query,
		newEmailTemplate.TemplateName, newEmailTemplate.MailTitle, newEmailTemplate.MailContent)

	return dbEmailTemplate, err
}

func (p *PostgresMailRepository) UpdateTemplate(ctx context.Context,
	newEmailTemplate models.EmailTemplate) (models.EmailTemplate, error) {
	query := `UPDATE EmailTemplate SET
				 TemplateName=NotEmpty($1,TemplateName), 
				 MailTitle=NotEmpty($2,MailTitle),
				 MailContent=NotEmpty($3,MailContent)
				 WHERE TemplateName=$1
				 RETURNING  TemplateName, MailTitle, MailContent`

	var dbEmailTemplate models.EmailTemplate
	err := p.conn.GetContext(ctx, &dbEmailTemplate, query,
		newEmailTemplate.TemplateName, newEmailTemplate.MailTitle, newEmailTemplate.MailContent)

	return dbEmailTemplate, err
}

func (p *PostgresMailRepository) GetMailByName(ctx context.Context,
	templateName string) (models.EmailTemplate, error) {
	query := `SELECT TemplateName, MailTitle, MailContent FROM EmailTemplate WHERE TemplateName = $1`

	var dbEmailTemplate models.EmailTemplate
	err := p.conn.GetContext(ctx, &dbEmailTemplate, query, templateName)
	return dbEmailTemplate, err
}
