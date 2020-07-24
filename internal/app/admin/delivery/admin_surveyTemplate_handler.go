package delivery

import (
	interfaces "2020_1_drop_table/internal/app/admin"
	"2020_1_drop_table/internal/app/admin/models"
	"2020_1_drop_table/internal/pkg/responses"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type AdminSurveyTemplateService struct {
	surveyTemplateLogic interfaces.UseCaseSurveyTemplate
}

func (ass *AdminSurveyTemplateService) GetTemplates(rwContext echo.Context) error {
	var (
		limit     int
		offset    int
		err       error
		templates models.GetSurveyTemplates
	)

	if limit, err = strconv.Atoi(rwContext.QueryParam("limit")); err != nil {
		return rwContext.NoContent(http.StatusBadRequest)
	}

	if offset, err = strconv.Atoi(rwContext.QueryParam("offset")); err != nil {
		return rwContext.NoContent(http.StatusBadRequest)
	}

	if templates, err = ass.surveyTemplateLogic.GetTemplates(limit, offset); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.JSON(http.StatusOK, &templates)
}

func (ass *AdminSurveyTemplateService) UpdateTemplate(rwContext echo.Context) error {
	var (
		err                error
		updateTemplateJson models.UpdateSurveyTemplate
	)

	if err = rwContext.Bind(&updateTemplateJson); err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if err = ass.surveyTemplateLogic.UpdateSurveyTemplate(&updateTemplateJson); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.NoContent(http.StatusOK)
}

func (ass *AdminSurveyTemplateService) CreateTemplate(rwContext echo.Context) error {
	var (
		err                error
		createTemplateJson models.CreateSurveyTemplate
	)

	if err = rwContext.Bind(&createTemplateJson); err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if err = ass.surveyTemplateLogic.CreateSurveyTemplate(&createTemplateJson); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.NoContent(http.StatusOK)
}

func (ass *AdminSurveyTemplateService) DeleteTemplate(rwContext echo.Context) error {
	var (
		err                error
		deleteTemplateJson models.DeleteSurveyTemplate
	)

	if err = rwContext.Bind(&deleteTemplateJson); err != nil {
		return rwContext.NoContent(http.StatusNotAcceptable)
	}

	if err = ass.surveyTemplateLogic.DeleteSurveyTemplate(&deleteTemplateJson); err != nil {
		return rwContext.JSON(http.StatusBadRequest, &responses.HttpError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return rwContext.NoContent(http.StatusOK)
}
