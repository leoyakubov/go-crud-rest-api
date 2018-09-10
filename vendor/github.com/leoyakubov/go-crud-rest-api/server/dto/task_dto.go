package dto

import "time"

type TaskDto struct {
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Priority    *int       `json:"priority"`
	CompletedAt *time.Time `json:"completedAt"`
	IsDeleted   bool       `json:"isDeleted"`
	IsCompleted bool       `json:"isCompleted"`
}
