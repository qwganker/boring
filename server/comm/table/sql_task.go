package table

type TSqlTask struct {
	ID                 int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Remark             string     `gorm:"column:remark;not null;type=text"`
	DBType             string     `gorm:"column:db_type;not null;type=varchar(32)"`
	DSN                string     `gorm:"column:dsn;not null;type=text"`
	SQL                string     `gorm:"column:sql;not null;type=text"`
	MetricDefine       string     `gorm:"column:metric_define;not null;type:text"`
	Cron               string     `gorm:"column:cron;not null;type=text"`
	SchedID            JobSchedID `gorm:"column:sched_id;type:text"`
	PrometheusConfigID int64      `gorm:"column:prometheus_config_id;not null;type:bigint"`
	Enabled            string     `gorm:"column:enabled;not null;type:text;default:'1'"`
	CreatedAt          *LocalTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          *LocalTime `gorm:"column:updated_at;autoUpdateTime"`
}

func (TSqlTask) TableName() string {
	return "t_sql_task"
}
