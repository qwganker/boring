package table

type TPrometheusConfig struct {
	ID          int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Remark      string     `gorm:"column:remark;not null;type=text"`
	Address     string     `gorm:"column:address;not null;type=text"`
	Username    string     `gorm:"column:username;type=text"`
	Password    string     `gorm:"column:password;type=text"`
	CtrlAddress string     `gorm:"column:ctrl_address;not null;type=text"`
	Config      string     `gorm:"column:config;type=text"`
	Rule        string     `gorm:"column:rule;type=text"`
	Enabled     string     `gorm:"column:enabled;;type:text"`
	CreatedAt   *LocalTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   *LocalTime `gorm:"column:updated_at;autoUpdateTime"`
}

func (TPrometheusConfig) TableName() string {
	return "t_prometheus_config"
}
