package engine

import (
	"log"
	"strconv"

	"github.com/qwganker/boring/comm/table"
	"github.com/robfig/cron/v3"
)

type GoCronEngine struct {
	cron *cron.Cron
}

func newGoCronEngine() *GoCronEngine {
	return &GoCronEngine{
		cron: cron.New(cron.WithSeconds()),
	}
}

func (e *GoCronEngine) Run() {
	e.cron.Run()
}

func (e *GoCronEngine) Stop() {
	e.cron.Stop()
}

func (s *GoCronEngine) RemoveTask(schedId table.JobSchedID) error {

	id, err := strconv.Atoi(string(schedId))
	if err != nil {
		return err
	}

	s.cron.Remove(cron.EntryID(id))
	return nil
}

func (s *GoCronEngine) AddTask(jobtask *table.TJobTask) (table.JobSchedID, error) {
	id, err := s.cron.AddFunc(jobtask.Cron, func() {
		e := NewJobEngineExecutor()
		e.Exec(jobtask)
	})

	if err != nil {
		log.Printf("Failed to schedule task: %v", err)
		return "", err
	}

	return table.JobSchedID(strconv.Itoa(int(id))), nil
}

func (s *GoCronEngine) HasTask(schedId table.JobSchedID) (bool, error) {
	if schedId == "" {
		return false, nil
	}

	id, err := strconv.Atoi(string(schedId))
	if err != nil {
		return false, err
	}

	for _, entry := range s.cron.Entries() {
		if cron.EntryID(id) == entry.ID {
			return true, nil
		}
	}

	return false, nil
}
