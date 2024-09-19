package trivia

type TriviaResponse struct {
	ResponseCode int64    `json:"response_code"`
	Results      []Result `json:"results"`
}

type Result struct {
	Type             Type       `json:"type"`
	Difficulty       Difficulty `json:"difficulty"`
	Category         string     `json:"category"`
	Question         string     `json:"question"`
	CorrectAnswer    string     `json:"correct_answer"`
	IncorrectAnswers []string   `json:"incorrect_answers"`
}

type Difficulty string

const (
	Medium Difficulty = "medium"
)

type Type string

const (
	Multiple Type = "multiple"
)
