package models

import (
	"encoding/json"
	"time"
)

type Plan struct {
	Id          uint64
	UserId      uint64
	Title       string
	Description string
	CreatedAt   time.Time
	DeadlineAt  time.Time
}

func NewPlan(id uint64, userId uint64, title string, description string, createdAt time.Time, deadlineAt time.Time) *Plan {
	return &Plan{Id: id, UserId: userId, Title: title, Description: description, CreatedAt: createdAt, DeadlineAt: deadlineAt}
}

func NewEmptyPlan() *Plan {
	return &Plan{}
}

func (p Plan) String() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
