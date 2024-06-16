package service_test

import (
	"fmt"
	quizRepository "quiz-app/internal/adapter/repository"
	"quiz-app/internal/core/service"
	"testing"
)

func TestQuizService(t *testing.T) {
	repository := quizRepository.NewQuizRepository()
	service := service.NewQuizService(repository)

	questions := service.GetQuestions()
	if len(questions) != 2 {
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

	if total != 2 {
		t.Errorf("expected total questions to be 2, got %d", total)
	}

	performance := service.GetPerformance(correct)
	fmt.Println(performance)
	if performance != 100 {
		t.Errorf("expected performance to be 100, got %.2f", performance)
	}
}
