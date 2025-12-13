package table

type TAlertType struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string     `gorm:"column:name;not null;type=text"`
	Code      string     `gorm:"column:code;not null;type=text"`
	CreatedAt *LocalTime `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *LocalTime `gorm:"column:updated_at;autoUpdateTime"`
}

func (TAlertType) TableName() string {
	return "t_alert_type"
}
