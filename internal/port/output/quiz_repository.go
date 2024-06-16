package output

import "quiz-app/internal/core/domain"

type QuizRepository interface {
	FetchQuestions() []domain.QuestionEntity
	SaveResult(score int)
	GetResults() []int
}
