package input

import "quiz-app/internal/core/domain"

type QuizInput interface {
	GetQuestions() []domain.QuestionEntity
	SubmitAnswers(answers map[int]int) int
	GetAnswers() []domain.AnswersEntity
}
