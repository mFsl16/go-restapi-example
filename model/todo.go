package model

type Todo struct {
	Id         int    `json:"id"`
	Task       string `json:"task"`
	IsComplete bool   `json:"is_complete"`
	CreateAt   string `json:"create_at"`
}
