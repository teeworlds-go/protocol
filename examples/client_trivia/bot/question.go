package bot

import (
	"context"
	"errors"
	"html"
	"math/rand/v2"
	"strconv"
	"strings"
	"sync"

	"github.com/teeworlds-go/protocol/examples/client_trivia/trivia"
)

func NewQuestionFactory() QuestionFactory {
	return QuestionFactory{}
}

type QuestionFactory struct {
	mu        sync.Mutex
	questions []Question
}

func (q *QuestionFactory) Next(ctx context.Context) (question Question, err error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.questions) == 0 {
		q.questions, err = q.fetch(ctx)
		if err != nil {
			return Question{}, err
		}

		if len(q.questions) == 0 {
			return Question{}, errors.New("no questions fetched")
		}
	}

	question = q.questions[0]
	q.questions = q.questions[1:]

	return question, nil
}

func (q *QuestionFactory) fetch(ctx context.Context) ([]Question, error) {
	questions, err := trivia.GetNextTriviaQuestions(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]Question, 0, len(questions))
	for _, r := range questions {
		result = append(result, NewQuestion(r))
	}
	return result, nil
}

type Question struct {
	questionText       string
	correctAnswerIndex int
	answers            []string
}

func (q Question) Question() string {
	return q.questionText
}

func (q Question) Answers() string {
	var sb strings.Builder
	for i, a := range q.answers {
		sb.WriteString(strconv.Itoa(i + 1))
		sb.WriteString(". ")
		sb.WriteString(a)
		if i < len(q.answers)-1 {
			sb.WriteString("  ")
		}
	}
	return sb.String()
}

func (q Question) IsCorrectAnswer(answer string) bool {

	// someone types "1, 2, 3, 4, etc"
	answer = strings.TrimSpace(answer)
	ui, err := strconv.ParseUint(answer, 10, 8)
	if err == nil {
		ui -= 1
		return int(ui) == q.correctAnswerIndex
	}

	// someone actually types the answer text
	for idx, choise := range q.answers {
		if strings.EqualFold(choise, answer) {
			return idx == q.correctAnswerIndex
		}
	}
	return false
}

func (q Question) CorrectAnswer() string {
	return q.answers[q.correctAnswerIndex]
}

func NewQuestion(r trivia.Result) Question {
	q := Question{
		questionText:       r.Question,
		correctAnswerIndex: 0,
		answers:            make([]string, 0, len(r.IncorrectAnswers)+1),
	}

	for _, answer := range r.IncorrectAnswers[:3] { // max 3 incorrect answers
		q.answers = append(q.answers, html.UnescapeString(answer))
	}
	q.answers = append(q.answers, html.UnescapeString(r.CorrectAnswer))

	rand.Shuffle(len(q.answers), func(i, j int) {
		q.answers[i], q.answers[j] = q.answers[j], q.answers[i]
	})

	for i, a := range q.answers {
		if a == r.CorrectAnswer {
			q.correctAnswerIndex = i
			break
		}
	}

	return q
}
