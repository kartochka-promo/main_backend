package models

import "time"

type StatisticsStruct struct {
	JsonData   string    `json:"json_data"`
	Time       time.Time `json:"time"`
	ClientUUID string    `json:"clientUuid"`
	StaffId    int       `json:"staffId"`
	CafeId     int       `json:"cafeId"`
}

type GetStatisticsOutput struct {
	Stats []StatisticsStruct `json:"stats"`
}

type UpdateStatistic struct {
	NewData StatisticsStruct `json:"new_data"`
	OldData StatisticsStruct `json:"old_data"`
}

type DeleteStatistic struct {
	DeletedData StatisticsStruct `json:"deleted_data"`
}
