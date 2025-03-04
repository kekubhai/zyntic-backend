package handlers

import (
	"github.com/gofiber/fiber/v2"
	models "github.com/kekubhai/zyntic/cmd/api"
)
type EventHandler struct {
	repository models.EventRepository
}
func (h *EventHandler) GetMany(ctx *fiber.Ctx)error{
	context,cancel:=
}
func (h *EventHandler) GetOne(ctx *fiber.Ctx)error{
	return nil
}
func (h *EventHandler) CreateOne(ctx *fiber.Ctx)error{
	return nil
}

func NewEventHandlers(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository: repository,
	}
	router.Get("/", handler.GetMany)
	router.Get("/", handler.CreateOne)
	router.Get("/", handler.GetOne)
}
