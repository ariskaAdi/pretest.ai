package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type QuestionEntity struct {
    Question    string   `json:"question"`
    Options     []string `json:"options"`
    Answer      string   `json:"answer"`
    Explanation string   `json:"explanation"`
}

type Response struct {
    Title     string           `json:"title"`
    Questions []QuestionEntity `json:"questions"`
}

var genkitInstance *genkit.Genkit

// InitGenkit - dipanggil SEKALI saat aplikasi start
func InitGenkit(ctx context.Context) error {

	if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found")
    }
    
    // Get API key dari environment variable
    apiKey := os.Getenv("GOOGLE_AI_API_KEY")
    if apiKey == "" {
        return fmt.Errorf("GOOGLE_AI_API_KEY not set in environment")
    }


    g := genkit.Init(ctx,
        genkit.WithPlugins(&googlegenai.GoogleAI{
			APIKey: apiKey,
		}),
        genkit.WithDefaultModel("googleai/gemini-1.5-flash"),
    )
    
    genkitInstance = g
    return nil
}

// GenerateQuiz - fungsi generator yang bisa dipanggil langsung
func GenerateQuiz(ctx context.Context, text string) (Response, error) {
    prompt := fmt.Sprintf(`
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
`, text)

    resp, _, err := genkit.GenerateData[Response](
        ctx,
        genkitInstance,
        ai.WithPrompt(prompt),
    )
    if err != nil {
        return Response{}, fmt.Errorf("failed to generate quiz: %w", err)
    }

    return *resp, nil
}

// Handler
func GenerateQuizHandler(c *fiber.Ctx) error {
    var input struct {
        Text string `json:"text"`
    }

    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    if input.Text == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "text is required",
        })
    }

    ctx := c.Context()

    // ✅ Panggil fungsi generator langsung
    resp, err := GenerateQuiz(ctx, input.Text)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "success": true,
        "data":    resp,
    })
}

func main() {
    ctx := context.Background()
    
    // Init Genkit sekali saja
    if err := InitGenkit(ctx); err != nil {
        log.Fatal("Failed to init genkit:", err)
    }

    app := fiber.New()
    
    app.Post("/api/generate-quiz", GenerateQuizHandler)
    
    log.Fatal(app.Listen(":3000"))
}


/*
  "text": "Fotosintesis adalah proses pembuatan makanan oleh tumbuhan hijau dengan bantuan sinar matahari. Klorofil dalam daun menyerap cahaya matahari. Bahan baku fotosintesis adalah karbon dioksida (CO2) dan air (H2O). Hasil fotosintesis adalah glukosa (C6H12O6) dan oksigen (O2)."

  {
  "text": "Sistem Informasi Manajemen (SIM) adalah sistem berbasis komputer yang menyediakan informasi bagi beberapa pemakai dengan kebutuhan yang serupa. Para pemakai biasanya membentuk suatu entitas organisasi formal, perusahaan atau sub unit dibawahnya. Informasi menjelaskan perusahaan atau salah satu sistem utamanya mengenai apa yang terjadi di masa lalu, apa yang terjadi sekarang dan apa yang mungkin terjadi di masa yang akan datang. SIM menghasilkan informasi yang digunakan dalam pembuatan keputusan. SIM juga dapat mempersatukan beberapa fungsi informasi dengan program komputerisasi lainnya seperti e-commerce."
}

{
  "text": "Turunan fungsi adalah konsep fundamental dalam kalkulus yang mengukur laju perubahan suatu fungsi terhadap variabelnya. Jika f(x) adalah fungsi, maka turunannya ditulis f'(x) atau df/dx. Rumus dasar: turunan dari x^n adalah n*x^(n-1). Turunan fungsi konstan adalah 0. Aturan rantai digunakan untuk fungsi komposit."
}
*/