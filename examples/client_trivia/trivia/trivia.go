package trivia

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	TriviaUrl = "https://opentdb.com/api.php?amount=50&difficulty=easy&type=multiple"
)

func GetNextTriviaQuestions(ctx context.Context) ([]Result, error) {
	return GetTriviaQuestions(ctx, TriviaUrl)
}

func GetTriviaQuestions(ctx context.Context, url string) ([]Result, error) {
	client := &http.Client{}

	var result TriviaResponse

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)
	request.Header.Set("Accept", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		data, _ := io.ReadAll(resp.Body) // explicitly ignore error
		return nil, fmt.Errorf("error while fetching trivia questions: %s: %s", resp.Status, string(data))
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Results, nil
}
