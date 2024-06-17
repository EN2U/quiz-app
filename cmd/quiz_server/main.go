package main

import (
	"log"
	"net/http"
	"quiz-app/internal/adapter/handler/http"
	"quiz-app/internal/adapter/repository"
	"quiz-app/internal/core/service"
)

func main() {
	repo := quizrepository.NewQuizRepository()
	quizService := service.NewQuizService(repo)
	quizHandler := quizhttp.NewQuizHandler(quizService)
	router := quizhttp.NewRouter(quizHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
