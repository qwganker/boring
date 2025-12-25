package table

type THttpTask struct {
	ID                 int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Remark             string     `gorm:"column:remark;not null;type=text"`
	Address            string     `gorm:"column:address;not null;type=text"`
	MetricDefine       string     `gorm:"column:metric_define;not null;type:text"`
	Cron               string     `gorm:"column:cron;not null;type=text"`
	PrometheusConfigID int64      `gorm:"column:prometheus_config_id;not null;type:bigint"`
	Enabled            string     `gorm:"column:enabled;not null;type:text;default:'1'"`
	CreatedAt          *LocalTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          *LocalTime `gorm:"column:updated_at;autoUpdateTime"`
}

func (THttpTask) TableName() string {
	return "t_http_task"
}
