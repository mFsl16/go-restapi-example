package model

import (
	"database/sql"
)

type Todo struct {
	Id         int            `json:"id"`
	Task       string         `json:"task"`
	IsComplete bool           `json:"is_completed"`
	CreateAt   string         `json:"create_at"`
	UpdateAt   sql.NullString `json:"update_at"`
}
