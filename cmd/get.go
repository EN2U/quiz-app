package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get the list of questions",
		Run: func(cmd *cobra.Command, args []string) {
			questions := quizService.GetQuestions()
			for _, q := range questions {
				fmt.Printf("Question %d: %s\n", q.ID, q.Text)
				for i, option := range q.Options {
					fmt.Printf("  %d: %s\n", i, option)
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
}
