package db_test

import (
	"testing"

	"github.com/avila-r/tasker/db"
)

func Test_Conn(t *testing.T) {
	database := db.Conn

	params := db.CreateTaskParams{
		Title: "Test PostgreSQL connection",
	}

	r, err := database.CreateTask(params)

	if err != nil {
		t.Errorf("error while creating task - %v", err.Error())
	}

	database.DeleteTask(r.ID)
}
