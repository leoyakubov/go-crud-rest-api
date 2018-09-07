package model

import "time"

type Task struct {
	BaseModel
	Title       string     `sql:"default: not null", json:"title"`
	Description *string    `sql:"default: null", json:"description"`
	Priority    *int       `sql:"default: null", json:"priority"`
	CompletedAt *time.Time `json:"completedAt"`
	IsCompleted bool       `json:"isCompleted"`
}

func (t *Task) AfterFind() error {
	return nil
}

// Callback before update model.Task
func (task *Task) BeforeUpdate() (err error) {
	//task.UpdatedAt = &time.Now()
	return
}

// Callback before create model.Task
func (task *Task) BeforeCreate() (err error) {
	//task.CreatedAt = time.Now()
	return
}
