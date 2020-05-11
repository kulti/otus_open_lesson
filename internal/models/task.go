package models

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
