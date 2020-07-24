package usecase

import (
	interfaces "2020_1_drop_table/internal/app/admin"
	"2020_1_drop_table/internal/app/admin/models"
)

type StatisticsLogic struct {
	statisticStorage interfaces.RepositoryStatistics
}

func (sl *StatisticsLogic) CreateStatisticData(request *models.StatisticsStruct) error {
	return sl.statisticStorage.CreateStatisticData(request)
}

func (sl *StatisticsLogic) GetStatisticData(limit, offset int) (models.GetStatisticsOutput, error) {
	return sl.statisticStorage.GetStatisticData(limit, offset)
}

func (sl *StatisticsLogic) DeleteStatistic(request *models.DeleteStatistic) error {
	return sl.statisticStorage.DeleteStatistic(request)
}

func (sl *StatisticsLogic) UpdateStatistic(request *models.UpdateStatistic) error {
	return sl.statisticStorage.UpdateStatistic(request)
}
