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

	if req.PdfUrl == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "text is required",
		})
	}

	ctx := c.UserContext()

	resp, err := h.svc.GenerateQuiz(ctx, req)
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

func (h *handler) GenerateQuizFromPdfLocal(c *fiber.Ctx) error {
    file, err := c.FormFile("pdf")
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "PDF file is required"})
    }

    f, err := file.Open()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to open file"})
    }
    defer f.Close()

    fileBytes := make([]byte, file.Size)
    if _, err := f.Read(fileBytes); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to read file"})
    }

    ctx := c.UserContext()

    resp, err := h.svc.GenerateQuizFromBytes(ctx, fileBytes)
    
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"success": true, "data": resp})
}