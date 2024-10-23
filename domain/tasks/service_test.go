package tasks_test

import (
	"testing"

	"github.com/avila-r/tasker/domain/tasks"
)

var (
	service = tasks.DefaultService
)

var (
	to_create = tasks.Task{
		Title:       "Test title",
		Description: "Test description",
	}
)

var (
	insert = func(t *testing.T) *tasks.Task {
		r, err := service.Create(to_create)

		t.Cleanup(func() {
			service.Db.DeleteTask(r.ID)
		})

		if err != nil {
			t.Errorf("error while creating task - %v", err.Error())
		}

		t.Logf("created task - %v", r.ToString())

		return r
	}
)

func Test_CreateTask(t *testing.T) {
	_ = insert(t)
}

func Test_FindByID(t *testing.T) {
	id := insert(t).ID

	if _, err := service.FindByID(id); err != nil {
		t.Errorf("unable to find task with id %v, error occurred - %v", id, err.Error())
	}
}

func Test_ExistsByID(t *testing.T) {
	id := insert(t).ID

	if exists := service.ExistsByID(id); !exists {
		_, err := service.FindByID(id)

		t.Errorf("task with id %v doesn't exists, error occurred - %v", id, err.Error())
	}
}

func Test_DeleteByID(t *testing.T) {
	id := insert(t).ID

	if err := service.DeleteByID(id); err != nil {
		t.Errorf("unable to delete task with id %v - %v", id, err.Error())
	}
}
