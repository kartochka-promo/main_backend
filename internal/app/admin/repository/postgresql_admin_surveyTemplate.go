package repository

import (
	"2020_1_drop_table/internal/app/admin/models"
	"github.com/jackc/pgx"
)

type SurveyTemplateStorage struct {
	dbPool *pgx.ConnPool
}

func (sts *SurveyTemplateStorage) GetTemplates(limit, offset int) (models.GetSurveyTemplates, error) {
	var (
		err          error
		response     models.GetSurveyTemplates
		query        = `SELECT cafeID, surveyTemplate,cafeOwnerId FROM surveyTemplate LIMIT $1 OFFSET $2;`
		storageQuery *pgx.Rows
	)
	defer func() {
		if storageQuery != nil {
			storageQuery.Close()
		}
	}()
	response.Templates = make([]models.GetSurveyTemplate, 0, 4)

	if storageQuery, err = sts.dbPool.Query(query, limit, offset); err != nil {
		return response, err
	}

	for storageQuery.Next() {
		var (
			templateInst models.GetSurveyTemplate
		)
		if err = storageQuery.Scan(&templateInst.CafeID, &templateInst.SurveyTemplateValue, &templateInst.CafeOwnerId); err != nil {
			return response, err
		}
		response.Templates = append(response.Templates, templateInst)
	}
	return response, err
}

func (sts *SurveyTemplateStorage) CreateSurveyTemplate(template *models.CreateSurveyTemplate) error {
	var (
		err   error
		query = `INSERT INTO surveytemplate (cafeid, surveytemplate,cafeOwnerId) VALUES ($1,$2,$3)`
	)
	_, err = sts.dbPool.Exec(query, template.CafeID, template.SurveyTemplateValue, template.CafeOwnerId)
	return err
}

func (sts *SurveyTemplateStorage) UpdateSurveyTemplate(template *models.UpdateSurveyTemplate) error {
	var (
		err   error
		query = `UPDATE surveytemplate set surveytemplate=$1 where cafeid=$2`
	)
	_, err = sts.dbPool.Exec(query, template.NewSurveyTemplateValue, template.CafeID)
	return err
}

func (sts *SurveyTemplateStorage) DeleteSurveyTemplate(template *models.DeleteSurveyTemplate) error {
	var (
		err   error
		query = `DELETE FROM surveytemplate WHERE surveytemplate=$1 AND cafeid=$2 AND cafeOwnerId=$3`
	)
	_, err = sts.dbPool.Exec(query, template.SurveyTemplateValue, template.CafeID, template.CafeOwnerId)
	return err
}
