/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"quiz-app/internal/adapter/repository"
	"quiz-app/internal/core/service"

	"github.com/spf13/cobra"
)

var quizRepository = quizrepository.NewQuizRepository()
var quizService = service.NewQuizService(quizRepository)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quiz-app",
	Short: "Quiz-App CLI",
	Long: `Quiz-App CLI is a command-line tool for interacting with a simple quiz application.
It allows users to answer quiz questions, submit their answers, and view their performance.

Quiz-App CLI supports fetching quiz questions, submitting answers, and comparing
your performance with others who have taken the quiz. Built with Go and Cobra, 
it provides a seamless experience for quiz enthusiasts.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
