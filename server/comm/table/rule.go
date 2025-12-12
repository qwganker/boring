package table

type TAlertRule struct {
	ID                 int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Title              string    `gorm:"column:title;not null;type=text"`
	Level              string    `gorm:"column:level;not null;type:text"`
	Type               string    `gorm:"column:type;not null;type:text"`
	Source             string    `gorm:"column:source;not null;type:text"`
	PromQLRule         string    `gorm:"column:promql_rule;not null;type:text"`
	Content            string    `gorm:"column:content;not null;type:text"`
	For                int64     `gorm:"column:for;not null;type:integer"`
	PromQLQuery        string    `gorm:"column:promql_query;type:text"`
	CustomLabels       string    `gorm:"column:custom_labels;type:text"`
	PrometheusConfigID int64     `gorm:"column:prometheus_config_id;not null;type:bigint"`
	NotifyID           int64     `gorm:"column:notify_id;type:bigint"`
	Enabled            string    `gorm:"column:enabled;not null;type:text;default:'1'"`
	CreatedAt          LocalTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          LocalTime `gorm:"column:updated_at;autoUpdateTime"`
}

func (TAlertRule) TableName() string {
	return "t_alert_rule"
}
