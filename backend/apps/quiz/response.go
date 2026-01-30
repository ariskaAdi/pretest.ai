package quiz

type QuizResponse struct {
	Title     string           `json:"title"`
	Questions []QuestionEntity `json:"questions"`
}