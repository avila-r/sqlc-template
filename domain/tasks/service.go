package tasks

import (
	"github.com/avila-r/tasker/db"
)

type TaskService struct {
	Db *db.Connection
}

var (
	DefaultService = NewService(db.Conn)
)

func NewService(db *db.Connection) *TaskService {
	return &TaskService{db}
}

func (s *TaskService) Create(task Task) (*Task, error) {
	if exists := s.ExistsByID(task.ID); exists {
		return nil, ErrTaskAlreadyExists
	}

	params := ToParams(task)

	result, err := s.Db.CreateTask(params)

	return FromRow(result), err
}

func (s *TaskService) FindByID(id int32) (*Task, error) {
	task, err := s.Db.GetTaskByID(id)

	if err != nil {
		return nil, ErrTaskNotFound
	}

	return FromResult(task), nil
}

func (s *TaskService) ExistsByID(id int32) bool {
	if t, err := s.FindByID(id); t == nil || err != nil {
		return false
	}

	return true
}

func (s *TaskService) DeleteByID(id int32) error {
	task, err := s.FindByID(id)

	if err != nil {
		return err
	}

	if err := s.Db.DeleteTask(task.ID); err != nil {
		return err
	}

	return nil
}

func (s *TaskService) List() ([]Task, error) {
	tasks := []Task{}

	result, err := s.Db.ListTasks()

	for _, r := range result {
		tasks = append(tasks, *FromResult(r))
	}

	return tasks, err
}
