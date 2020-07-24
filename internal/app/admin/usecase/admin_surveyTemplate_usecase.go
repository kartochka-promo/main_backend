package usecase

import (
	"2020_1_drop_table/internal/app/admin/models"
	"2020_1_drop_table/internal/app/admin/repository"
)

type SurveyTemplateLogic struct {
	templateStorage repository.SurveyTemplateStorage
}

func (stl *SurveyTemplateLogic) GetTemplates(limit, offset int) (models.GetSurveyTemplates, error) {
	return stl.templateStorage.GetTemplates(limit, offset)
}

func (stl *SurveyTemplateLogic) CreateSurveyTemplate(template *models.CreateSurveyTemplate) error {
	return stl.templateStorage.CreateSurveyTemplate(template)
}

func (stl *SurveyTemplateLogic) UpdateSurveyTemplate(template *models.UpdateSurveyTemplate) error {
	return stl.templateStorage.UpdateSurveyTemplate(template)
}

func (stl *SurveyTemplateLogic) DeleteSurveyTemplate(template *models.DeleteSurveyTemplate) error {
	return stl.templateStorage.DeleteSurveyTemplate(template)
}
