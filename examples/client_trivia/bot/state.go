package bot

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"strings"
)

type State struct {
	qf              *QuestionFactory
	running         bool
	currentQuestion Question
	scoreMap        map[string]int // score map
}

func NewState(qf *QuestionFactory) *State {
	return &State{
		qf:       qf,
		scoreMap: make(map[string]int),
	}
}

func (s *State) Running() bool {
	return s.running
}

func (s *State) Start(ctx context.Context) (chatLine string, err error) {
	if !s.running {
		q, err := s.qf.Next(ctx)
		if err != nil {
			return "", err
		}

		s.currentQuestion = q
		s.running = true
	}

	// if already running or just started, always print the question.
	return fmt.Sprintf("%s: %s",
		s.currentQuestion.Question(),
		s.currentQuestion.Answers()), nil
}

func (s *State) Answer(username, answer string) (string, bool) {
	if !s.running {
		return "", false
	}

	if !s.currentQuestion.IsCorrectAnswer(answer) {
		return "", false
	}

	s.scoreMap[username]++
	s.running = false

	return s.currentQuestion.CorrectAnswer(), true
}

func (s *State) Top() string {
	if len(s.scoreMap) == 0 {
		return "No scores yet"
	}

	list := make([]tuple, 0, len(s.scoreMap))
	for k, v := range s.scoreMap {
		list = append(list, tuple{PlayerName: k, Score: v})
	}

	slices.SortFunc(list, func(a, b tuple) int {
		return cmp.Compare(b.Score, a.Score)
	})

	var sb strings.Builder
	sb.Grow(16 * len(list))

	for i, t := range list {
		sb.WriteString(fmt.Sprintf("%d. %s: %d", i+1, t.PlayerName, t.Score))
		if i < len(list)-1 {
			sb.WriteString(", ")
		}
	}

	return sb.String()
}

func (s *State) Score(playerName string) string {
	score, ok := s.scoreMap[playerName]
	if !ok {
		return fmt.Sprintf("%s, you have no score", playerName)
	}
	return fmt.Sprintf("%s, your score is: %d", playerName, score)
}

type tuple struct {
	PlayerName string
	Score      int
}
