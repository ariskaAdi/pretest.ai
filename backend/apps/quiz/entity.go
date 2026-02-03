package quiz

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type QuestionEntity struct {
	Question    string   `json:"question"`
	Options     []string `json:"options"`
	Answer      string   `json:"answer"`
	Explanation string   `json:"explanation"`
}

func fetchPdf(ctx context.Context, url string) ([]byte, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch pdf: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}