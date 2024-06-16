package service

import (
	"quiz-app/internal/core/domain"
	"quiz-app/internal/port/output"
)

type QuizService struct {
	repo output.QuizRepository
}

func NewQuizService(repo output.QuizRepository) *QuizService {
	return &QuizService{repo: repo}
}

func (service *QuizService) GetQuestions() []domain.QuestionEntity {
	return service.repo.FetchQuestions()
}

func (service *QuizService) SubmitAnswers(answers map[int]int) (int, int) {
	questions := service.repo.FetchQuestions()
	correctCount := 0

	for _, question := range questions {
		if answer, ok := answers[question.ID]; ok && answer == question.AnswerID {
			correctCount++
		}
	}

	service.repo.SaveResult(correctCount)

	totalQuestions := len(questions)
	return correctCount, totalQuestions
}

func (s *QuizService) GetPerformance(score int) float64 {
	results := s.repo.GetResults()
	betterThan := 0
	for _, result := range results {
		if score > result {
			betterThan++
		}
	}

	return float64(betterThan) / float64(len(results)) * 100
}
