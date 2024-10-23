package tasks

import (
	"fmt"
	"time"

	"github.com/avila-r/pgor"
	"github.com/avila-r/tasker/db"
)

type Task struct {
	ID          int32     `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Task) ToString() string {
	return fmt.Sprintf(
		"Task[ID: %d, Title: %s, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s]",
		t.ID,
		t.Title,
		t.Description,
		t.Status,
		t.CreatedAt.Format("2006-01-02 15:04:05"),
		t.UpdatedAt.Format("2006-01-02 15:04:05"),
	)
}

func FromRow(r db.CreateTaskRow) *Task {
	return &Task{
		ID:          r.ID,
		Title:       r.Title,
		Description: *pgor.FromText(r.Description),
		Status:      StatusPending, // Default status when created
		CreatedAt:   r.CreatedAt.Time,
		UpdatedAt:   r.UpdatedAt.Time,
	}
}

func FromResult(r db.Task) *Task {
	return &Task{
		ID:          r.ID,
		Title:       r.Title,
		Description: *pgor.FromText(r.Description),
		Status:      r.Status,
		CreatedAt:   r.CreatedAt.Time,
		UpdatedAt:   r.UpdatedAt.Time,
	}
}

func ToParams(t Task) db.CreateTaskParams {
	return db.CreateTaskParams{
		Title:       t.Title,
		Description: pgor.Text(t.Description),
	}
}
