package engine

import (
	"github.com/qwganker/boring/comm/table"
)

type EngineType string

const (
	GoCron  EngineType = "gocron"
	Default EngineType = "gocron"
)

type SchedulerEngine interface {
	Run()
	Stop()
	AddTask(jobtask *table.TJobTask) (table.JobSchedID, error)
	RemoveTask(schedId table.JobSchedID) error
	HasTask(schedId table.JobSchedID) (bool, error)
}

func NewSchedulerEngine(engineType EngineType) SchedulerEngine {
	switch engineType {
	case GoCron:
		return newGoCronEngine()
	default:
		return newGoCronEngine()
	}

}
