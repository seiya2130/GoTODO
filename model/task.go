package model

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
