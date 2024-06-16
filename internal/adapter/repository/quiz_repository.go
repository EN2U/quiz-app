package quizRepository

import (
	"quiz-app/internal/core/domain"
)

type QuizRepository struct {
	questions []domain.QuestionEntity
	results   []int
}

func NewQuizRepository() *QuizRepository {
	return &QuizRepository{
		questions: []domain.QuestionEntity{
			{
				ID:       1,
				Text:     "What is the capital of France?",
				Options:  []string{"Berlin", "Madrid", "Paris", "Lisbon"},
				AnswerID: 2,
			},
			{
				ID:       2,
				Text:     "What is 2+2?",
				Options:  []string{"3", "4", "5", "6"},
				AnswerID: 1,
			},
		},
		results: []int{},
	}
}

func (r *QuizRepository) FetchQuestions() []domain.QuestionEntity {
	return r.questions
}

func (r *QuizRepository) SaveResult(score int) {
	r.results = append(r.results, score)
}

func (r *QuizRepository) GetResults() []int {
	return r.results
}
