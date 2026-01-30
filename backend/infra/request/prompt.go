package infrarequest

var (
	GenerateQuizPrompt = `
You are an academic assessment generator.

Rules:
- Output MUST be valid JSON
- Do NOT include markdown
- Do NOT include any text outside JSON

JSON schema:
{
  "title": "string",
  "questions": [
    {
      "question": "string",
      "options": ["string"],
      "answer": "string",
      "explanation": "string"
    }
  ]
}

Text:
%s
`
)
