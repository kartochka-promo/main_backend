package repository

import (
	"github.com/jackc/pgx"

	logModels "2020_1_drop_table/internal/app/admin_logs/models"
)

type LogsStorage struct {
	dbPool *pgx.ConnPool
}

func (cls *LogsStorage) LogData(adminName, tableName string, columns logModels.Columns, newData interface{}, oldData interface{}) error {
	_, err := cls.dbPool.Exec("INSERT INTO admins_operations"+
		" (operation_operator, operation_table, operation_columns, operation_old_data, operation_new_data)"+
		" VALUES ($1,$2,$3,$4,$5)", adminName, tableName, columns, oldData, newData)
	return err
}

func (cls *LogsStorage) GetCafeLogData(limit, offset int) (logModels.CafeLogsModel, error) {
	var (
		cafeLogs     logModels.CafeLogsModel
		err          error
		cafeLogQuery *pgx.Rows
	)
	cafeLogs.Logs = make([]logModels.CafeLogModel, 0, 0)

	if cafeLogQuery, err = cls.dbPool.Query("SELECT operation_operator, operation_time, operation_columns,"+
		" operation_new_data, operation_old_data FROM admins_operations"+
		" WHERE operation_table = 'Cafe' LIMIT $1 OFFSET $2", limit, offset); err != nil {
		return cafeLogs, err
	}

	for cafeLogQuery.Next() {
		var (
			cafeLogInstance logModels.CafeLogModel
		)
		if err = cafeLogQuery.Scan(&cafeLogInstance.Operator, &cafeLogInstance.OperationTime, &cafeLogInstance.Columns, &cafeLogInstance.NewData, &cafeLogInstance.OldData); err != nil {
			return cafeLogs, err
		}
		cafeLogs.Logs = append(cafeLogs.Logs, cafeLogInstance)
	}
	return cafeLogs, err
}
