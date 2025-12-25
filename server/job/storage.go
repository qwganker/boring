package job

import (
	"context"

	"github.com/qwganker/boring/comm/table"
	"gorm.io/gorm"
)

type JobTaskStorage struct {
	db *gorm.DB
}

func NewJobTaskStorage(db *gorm.DB) *JobTaskStorage {
	return &JobTaskStorage{
		db: db,
	}
}

func (s *JobTaskStorage) AddTask(jobtask *table.TJobTask) error {

	ctx := context.Background()

	// upsert
	if err := s.db.WithContext(ctx).Save(jobtask).Error; err != nil {
		return err
	}
	return nil
}

func (s *JobTaskStorage) UpdateTask(jobtask *table.TJobTask) error {

	ctx := context.Background()

	// upsert
	if err := s.db.WithContext(ctx).Save(jobtask).Error; err != nil {
		return err
	}
	return nil
}

func (s *JobTaskStorage) RemoveTask(schedId table.JobSchedID) error {
	if err := s.db.WithContext(context.Background()).Where("sched_id = ?", schedId).Delete(&table.TJobTask{}).Error; err != nil {
		return err
	}

	return nil
}

func (s *JobTaskStorage) GetTaskList() (jobtasklist []table.TJobTask, err error) {

	list := []table.TJobTask{}
	if err := s.db.WithContext(context.Background()).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (s *JobTaskStorage) GetTask(schedId table.JobSchedID) (*table.TJobTask, error) {
	task := table.TJobTask{}
	if err := s.db.WithContext(context.Background()).Where("sched_id = ?", schedId).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
