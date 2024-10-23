package tasks

import "github.com/avila-r/gor"

var (
	StatusPending = "pending"
	StatusDone    = "done"
)

var (
	// Status Not Found
	ErrTaskNotFound = gor.NewError(gor.StatusNotFound, "task not found")

	// Status Conflict
	ErrTaskAlreadyExists = gor.NewError(gor.StatusConflict, "already exists a task with provided id")

	// Status No Content
	ErrNoTasksToList = gor.NewError(gor.StatusNotFound, "no tasks to list")

	// Status Bad Request
	ErrInvalidID = gor.NewError(gor.StatusBadRequest, "invalid id")
)
