package quizrepository

import (
	"quiz-app/internal/core/domain"
)

type QuizRepository struct {
	questions []domain.QuestionEntity
	results   []int
	answers	  []domain.AnswersEntity
}

func NewQuizRepository() *QuizRepository {
	return &QuizRepository{
		questions: GetQuestionsEntity(),
		results: []int{},
		answers: GetAnswersEntity(),
	}
}

func GetQuestionsEntity() []domain.QuestionEntity {
	return []domain.QuestionEntity{
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
		{
			ID:       3,
			Text:     "Which planet is closest to the Sun?",
			Options:  []string{"Venus", "Mars", "Mercury", "Jupiter"},
			AnswerID: 2,
		},
		{
			ID:       4,
			Text:     "What is the chemical element with the symbol 'O'?",
			Options:  []string{"Gold", "Osmium", "Oxygen", "Oganesson"},
			AnswerID: 2,
		},
	}
}

func GetAnswersEntity() []domain.AnswersEntity {
	return []domain.AnswersEntity{
		{
			ID: 1,
			ValidAnswers: 4,
			InvalidAnswers: 0,
		},
		{
			ID: 2,
			ValidAnswers: 0,
			InvalidAnswers: 4,
		},
		{
			ID: 3,
			ValidAnswers: 2,
			InvalidAnswers: 2,
		},
		{
			ID: 4,
			ValidAnswers: 3,
			InvalidAnswers: 1,
		},
		{
			ID: 5,
			ValidAnswers: 1,
			InvalidAnswers: 3,
		},
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

func (r *QuizRepository) GetAnswers() []domain.AnswersEntity{
	return r.answers
}
