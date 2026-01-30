package quiz

import (
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	svc *service
}

func newHandler(svc *service) *handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) GenerateQuiz(c *fiber.Ctx) error {
	var req GenerateQuizRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	if req.Text == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "text is required",
		})
	}

	ctx := c.Context()

	resp, err := h.svc.GenerateQuiz(ctx, req.Text)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    resp,
	})
}