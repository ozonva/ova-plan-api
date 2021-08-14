package plan

import (
	"encoding/json"
	"time"
)

type PlanInterface interface {
	String()
}

type Plan struct {
	Id          uint64
	UserId      uint64
	Title       string
	Description string
	CreatedAt   time.Time
	DeadlineAt  time.Time
}

func (p Plan) String() {
	panic("implement me")
}

func NewPlan(id uint64, userId uint64, title string, description string, createdAt time.Time, deadlineAt time.Time) *Plan {
	return &Plan{Id: id, UserId: userId, Title: title, Description: description, CreatedAt: createdAt, DeadlineAt: deadlineAt}
}

func String(p Plan) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
