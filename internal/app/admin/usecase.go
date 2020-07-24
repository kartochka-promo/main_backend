package admin

import (
	models "2020_1_drop_table/internal/app/admin/models"
)

type UseCaseAdmin interface {
	CreateAdmin(admin *models.CreateOrUpdateAdmin) error
	Authentication(authAdminRequest *models.LogAdmin) (bool, error)
	UpdateAdmin(updateAdminRequest *models.CreateOrUpdateAdmin) error
}

type UseCaseCafe interface {
	GetCafes(limit, offset int) (models.GetCafes, error)
	GetCafe(getCafeID int) (models.GetCafe, error)
	UpdateCafe(updateCafeRequest *models.CreateOrUpdateCafe) error
	CreateCafe(createCafeRequest *models.CreateOrUpdateCafe) error
	DeleteCafe(deleteCafeID int) error
}

type UseCaseSurveyTemplate interface {
	UpdateSurveyTemplate(template *models.UpdateSurveyTemplate) error
	GetTemplates(limit, offset int) (models.GetSurveyTemplates, error)
	CreateSurveyTemplate(template *models.CreateSurveyTemplate) error
	DeleteSurveyTemplate(template *models.DeleteSurveyTemplate) error
}

type UseCaseStatistics interface {
	CreateStatisticData(request *models.StatisticsStruct) error
	GetStatisticData(limit, offset int) (models.GetStatisticsOutput, error)
	DeleteStatistic(request *models.DeleteStatistic) error
	UpdateStatistic(request *models.UpdateStatistic) error
}