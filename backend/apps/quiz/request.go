package quiz

type GenerateQuizRequest struct {
	Text string `json:"text" validate:"required"`
}