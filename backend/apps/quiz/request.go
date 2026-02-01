package quiz

type GenerateQuizRequest struct {
	PdfUrl string `json:"pdfUrl" validate:"required"`
}