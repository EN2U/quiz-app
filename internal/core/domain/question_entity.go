package domain

type QuestionEntity struct {
	ID       int
	Text     string
	Options  []string
	AnswerID int
}
