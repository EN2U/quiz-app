package service_test

import (
	"quiz-app/internal/adapter/repository"
	"quiz-app/internal/core/service"
	"testing"
)

func TestQuizService(t *testing.T) {
	repository := quizrepository.NewQuizRepository()
	service := service.NewQuizService(repository)

	questions := service.GetQuestions()
	if len(questions) != 4 {
		t.Errorf("expected 2 questions, got %d", len(questions))
	}

	answers := map[int]int{
		1: 2,
		2: 1,
	}

	correct, total := service.SubmitAnswers(answers)
	if correct != 2 {
		t.Errorf("expected 2 correct answers, got %d", correct)
	}

	if total != 4 {
		t.Errorf("expected total questions to be 2, got %d", total)
	}

	performance := service.GetPerformance(correct, total)
	if performance != 40 {
		t.Errorf("expected performance to be 100, got %.2f", performance)
	}
}
