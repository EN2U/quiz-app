package quizhttp

import (
	"github.com/gorilla/mux"
)

func NewRouter(handler *QuizHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/questions", handler.GetQuestions).Methods("GET")
	r.HandleFunc("/submit", handler.SubmitAnswers).Methods("POST")

	return r
}
