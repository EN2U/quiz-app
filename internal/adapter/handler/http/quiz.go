package quizhttp

import (
	"encoding/json"
	"math"
	"net/http"
	"quiz-app/internal/core/service"
)

type QuizHandler struct {
	service *service.QuizService
}

func NewQuizHandler(service *service.QuizService) *QuizHandler {
	return &QuizHandler{service: service}
}

func (h *QuizHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions := h.service.GetQuestions()
	json.NewEncoder(w).Encode(questions)
}

func (h *QuizHandler) SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	var answers map[int]int
	err := json.NewDecoder(r.Body).Decode(&answers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	correct, total := h.service.SubmitAnswers(answers)
	performance := h.service.GetPerformance(correct, total)
	response := map[string]int{
		"correct":     correct,
		"total":       total,
		"performance": int(math.Round(performance)),
	}

	json.NewEncoder(w).Encode(response)
}
