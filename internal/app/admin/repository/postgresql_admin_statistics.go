package repository

import (
	"2020_1_drop_table/internal/app/admin/models"
	"github.com/jackc/pgx"
)

type StatisticsStorage struct {
	dbPool *pgx.ConnPool
}

func (ss *StatisticsStorage) CreateStatisticData(request *models.StatisticsStruct) error {
	var (
		err   error
		query = `insert into statistics_table (jsonData, time, clientUUID, staffId, cafeId) VALUES ($1,$2,$3,$4,$5)`
	)
	_, err = ss.dbPool.Exec(query, request.JsonData, request.Time, request.ClientUUID, request.StaffId, request.CafeId)
	return err
}

func (ss *StatisticsStorage) GetStatisticData(limit, offset int) (models.GetStatisticsOutput, error) {
	var (
		err          error
		query        = `SELECT * from statistics_table where staffID=$1 order by time LIMIT $2 OFFSET $3 `
		response     models.GetStatisticsOutput
		storageQuery *pgx.Rows
	)
	defer func() {
		if storageQuery != nil {
			storageQuery.Close()
		}
	}()
	response.Stats = make([]models.StatisticsStruct, 0, 8)

	if storageQuery, err = ss.dbPool.Query(query, limit, offset); err != nil {
		return response, err
	}

	for storageQuery.Next() {
		var (
			statInst models.StatisticsStruct
		)
		if err = storageQuery.Scan(&statInst.JsonData, &statInst.Time, &statInst.ClientUUID, &statInst.StaffId, &statInst.CafeId); err != nil {
			return response, err
		}
		response.Stats = append(response.Stats, statInst)
	}
	return response, err
}

func (ss *StatisticsStorage) DeleteStatistic(request *models.DeleteStatistic) error {
	var (
		query = `DELETE FROM statistics_table WHERE jsonData = $1 AND time = $2 AND clientUUID = $3 AND staffId = $4 AND cafeId = $5;`
		err   error
	)
	_, err = ss.dbPool.Exec(query, request.DeletedData.JsonData, request.DeletedData.Time, request.DeletedData.ClientUUID, request.DeletedData.StaffId, request.DeletedData.CafeId)
	return err
}

func (ss *StatisticsStorage) UpdateStatistic(request *models.UpdateStatistic) error {
	var (
		query = `UPDATE statistics_table SET jsonData = $1, time = $2, clientUUID = $3, staffId = $4, cafeId = $5 WHERE jsonData = $6 AND time = $7 AND clientUUID = $8 AND staffId = $9 AND cafeId = $10;`
		err   error
	)
	_, err = ss.dbPool.Exec(query, request.NewData.JsonData, request.NewData.Time,
		request.NewData.ClientUUID, request.NewData.StaffId, request.NewData.CafeId,
		request.OldData.JsonData, request.OldData.Time, request.OldData.ClientUUID,
		request.OldData.StaffId, request.OldData.CafeId)
	return err
}
