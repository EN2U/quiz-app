package service

import (
	"quiz-app/internal/core/domain"
	"quiz-app/internal/core/port/output"
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

func (s *QuizService) GetPerformance(validAnswers int, totalAnswers int) float64 {
	newAnswerPerformance := float64(validAnswers) / float64(totalAnswers)
	previousAnswers := s.repo.GetAnswers()
	betterThan := 0

	for _, answer := range previousAnswers {
		answerPerformance := float64(answer.ValidAnswers) / float64(answer.ValidAnswers+answer.InvalidAnswers)
		if newAnswerPerformance > answerPerformance {
			betterThan++
		}
	}

	return (float64(betterThan) / float64(len(previousAnswers))) * 100
}
