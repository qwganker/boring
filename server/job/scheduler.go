package job

import (
	"encoding/json"
	"log"

	"github.com/qwganker/boring/comm/table"
	"github.com/qwganker/boring/job/engine"
	"github.com/qwganker/boring/storage"
)

var (
	jobScheduler *JobScheduler = nil
)

type JobScheduler struct {
	storage *JobTaskStorage
	engine  engine.SchedulerEngine
}

func NewJobScheduler(storage *JobTaskStorage, engine engine.SchedulerEngine) *JobScheduler {
	return &JobScheduler{
		storage: storage,
		engine:  engine,
	}
}

func GetJobScheduler() *JobScheduler {
	if jobScheduler == nil {
		jobScheduler = NewJobScheduler(NewJobTaskStorage(storage.GetDBInstance()), engine.NewSchedulerEngine(engine.GoCron))
	}

	return jobScheduler
}

func (s *JobScheduler) CleanAll() error {
	jobTasklist, err := s.storage.GetTaskList()
	if err != nil {
		return err
	}

	for _, jobtask := range jobTasklist {
		s.UnRegisterTask(jobtask.SchedID)
	}

	return nil
}

func (s *JobScheduler) Run() {
	s.engine.Run()
}

func (s *JobScheduler) Stop() {
	s.engine.Stop()
}

func (s *JobScheduler) ReRegisterTask(schedId table.JobSchedID, jobType table.JobTaskType, cron string, payload interface{}) (table.JobSchedID, error) {
	exist, err := s.HasRegisterTask(schedId)
	if err != nil {
		return "", err
	}

	if !exist {
		return s.RegisterTask(jobType, cron, payload)
	}

	s.UnRegisterTask(schedId)
	return s.RegisterTask(jobType, cron, payload)
}

func (s *JobScheduler) RegisterTaskIfNotExist(schedId table.JobSchedID, jobType table.JobTaskType, cron string, payload interface{}) (table.JobSchedID, error) {
	exist, err := s.HasRegisterTask(schedId)
	if err != nil {
		return "", err
	}

	if !exist {
		return s.RegisterTask(jobType, cron, payload)
	}

	return schedId, nil
}

func (s *JobScheduler) RegisterTask(jobType table.JobTaskType, cron string, payload interface{}) (table.JobSchedID, error) {

	bytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("RegisterTask failed", err)
		return "", err
	}

	jobtask := &table.TJobTask{
		Type:    jobType,
		Cron:    cron,
		Payload: string(bytes),
		State:   table.JobTaskStateRunning,
	}

	schedId, err := s.engine.AddTask(jobtask)
	if err != nil {
		return "", err
	}

	jobtask.SchedID = schedId
	jobtask.State = table.JobTaskStateRunning

	s.storage.AddTask(jobtask)
	return schedId, nil
}

func (s *JobScheduler) HasRegisterTask(schedId table.JobSchedID) (bool, error) {
	return s.engine.HasTask(schedId)
}

func (s *JobScheduler) UnRegisterTask(schedId table.JobSchedID) error {
	s.engine.RemoveTask(schedId)
	s.storage.RemoveTask(schedId)

	return nil
}
