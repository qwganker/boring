package table

type JobTaskState string

const (
	JobTaskStateRunning JobTaskState = "running"
	JobTaskStatePaused  JobTaskState = "paused"
	JobTaskStateStopped JobTaskState = "stopped"
	JobTaskStateAborted JobTaskState = "aborted"
)

type JobTaskType string

const (
	JobTaskTypeSQL  JobTaskType = "sql"
	JobTaskTypeHTTP JobTaskType = "http"
)

type JobSchedID string

type TJobTask struct {
	ID        int64        `gorm:"column:id;primaryKey;autoIncrement"`
	SchedID   JobSchedID   `gorm:"column:sched_id;type:text"`
	Type      JobTaskType  `gorm:"column:type;type:text"`
	Cron      string       `gorm:"column:cron;type:text"`
	Payload   string       `gorm:"column:payload;type:text"`
	State     JobTaskState `gorm:"column:state;type:text"`
	CreatedAt *LocalTime   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *LocalTime   `gorm:"column:updated_at;autoUpdateTime"`
}

func (TJobTask) TableName() string {
	return "t_job_task"
}
