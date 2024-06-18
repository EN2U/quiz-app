package service

import (
	"errors"
	"quiz-app/internal/core/domain"
	"quiz-app/internal/core/port/output"
)

type QuizService struct {
	repo output.QuizRepository
}

func NewQuizService(repo output.QuizRepository) *QuizService {
	return &QuizService{repo: repo}
}

func (service *QuizService) GetQuestions() ([]domain.QuestionEntity, error) {
	questions := service.repo.FetchQuestions()
	if len(questions) == 0 {
		return nil, errors.New("no questions available")
	}
	return questions, nil
}

func (service *QuizService) SubmitAnswers(answers map[int]int) (int, int, error) {
	questions := service.repo.FetchQuestions()
	if len(questions) == 0 {
		return 0, 0, errors.New("no questions available")
	}
	correctCount := 0

	for _, question := range questions {
		if answer, ok := answers[question.ID]; ok && answer == question.AnswerID {
			correctCount++
		}
	}

	service.repo.SaveResult(correctCount)

	totalQuestions := len(questions)
	return correctCount, totalQuestions, nil
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
