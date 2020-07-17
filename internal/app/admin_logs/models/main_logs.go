package models

import (
	"2020_1_drop_table/internal/app/admin/models"
	"time"
)

type Columns struct {
	Columns []string `json:"columns"`
}

type CafeLogModel struct {
	NewData       models.CreateOrUpdateCafe `json:"newData"`
	OldData       models.GetCafe            `json:"oldData"`
	Columns       Columns                   `json:"columns"`
	Operator      string                    `json:"operator"`
	OperationTime time.Time                 `json:"operation_time"`
}

type CafeLogsModel struct {
	Logs []CafeLogModel `json:"logs"`
}
