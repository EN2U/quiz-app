package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	submitCmd = &cobra.Command{
		Use:   "submit",
		Short: "Submit your answers",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please provide your answers in the format: submit 1=2 2=1")
				return
			}

			answers := make(map[int]int)
			for _, answer := range args {
				parts := strings.Split(answer, "=")
				if len(parts) != 2 {
					fmt.Printf("Invalid answer format: %s\n", answer)
					continue
				}

				questionID, err := strconv.Atoi(parts[0])
				if err != nil {
					fmt.Printf("Invalid question ID: %s\n", parts[0])
					continue
				}

				optionID, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Printf("Invalid option ID: %s\n", parts[1])
					continue
				}

				answers[questionID] = optionID
			}

			correct, total := quizService.SubmitAnswers(answers)
			fmt.Printf("You got %d out of %d questions correct.\n", correct, total)

			performance := quizService.GetPerformance(correct)
			fmt.Printf("You were better than %.2f%% of all quizzers.\n", performance)
		},
	}
)

func init() {
	rootCmd.AddCommand(submitCmd)
}
