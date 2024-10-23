package tasks

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/avila-r/gor"
)

type TaskHandler struct {
	Service *TaskService
}

var (
	DefaultHandler = NewHandler(DefaultService)
)

func NewHandler(s *TaskService) *TaskHandler {
	return &TaskHandler{s}
}

func (h *TaskHandler) Route(r fiber.Router) {
	// Get all
	r.Get("/", h.GetTasks)

	// Insert one
	r.Post("/", h.PostTask)

	// Get one
	r.Get("/id/:id", h.GetTask)

	// Delete one
	r.Delete("/id/:id", h.DeleteTask)
}

func (h *TaskHandler) GetTasks(c *fiber.Ctx) error {
	all, err := h.Service.List()

	if err != nil {
		return err
	}

	if len(all) == 0 {
		return ErrNoTasksToList
	}

	return c.JSON(all)
}

func (h *TaskHandler) PostTask(c *fiber.Ctx) error {
	new := struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}{}

	if err := c.BodyParser(&new); err != nil {
		return err
	}

	t, err := h.Service.Create(Task{
		Title:       new.Title,
		Description: new.Description,
	})

	if err != nil {
		return err
	}

	return c.Status(gor.StatusCreated.Code).JSON(t)
}

func (h *TaskHandler) GetTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return ErrInvalidID
	}

	t, err := h.Service.FindByID(int32(id))

	if err != nil {
		return err
	}

	return c.JSON(t)
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return ErrInvalidID
	}

	if err := h.Service.DeleteByID(int32(id)); err != nil {
		return err
	}

	return c.SendStatus(gor.StatusOK.Code)
}
