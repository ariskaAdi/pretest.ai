package quiz

type QuizResponse struct {
	Title     string           `json:"title"`
	Questions []QuestionEntity `json:"questions"`
}

type UsageInfo struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type QuizResultWithStats struct {
	Quiz  *QuizResponse `json:"quiz"`
	Usage UsageInfo     `json:"usage"`
}