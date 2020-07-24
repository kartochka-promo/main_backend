package models

type GetSurveyTemplate struct {
	CafeID              int    `json:"cafe_id"`
	SurveyTemplateValue string `json:"survey_template"`
	CafeOwnerId         int    `json:"cafe_owner_id"`
}
type GetSurveyTemplates struct {
	Templates []GetSurveyTemplate `json:"templates"`
}

type DeleteSurveyTemplate struct {
	CafeID              int    `json:"cafe_id"`
	SurveyTemplateValue string `json:"survey_template"`
	CafeOwnerId         int    `json:"cafe_owner_id"`
}

type UpdateSurveyTemplate struct {
	CafeID                 int    `json:"cafe_id"`
	NewSurveyTemplateValue string `json:"new_survey_template"`
}

type CreateSurveyTemplate struct {
	CafeID              int    `json:"cafe_id"`
	SurveyTemplateValue string `json:"survey_template"`
	CafeOwnerId         int    `json:"cafe_owner_id"`
}
