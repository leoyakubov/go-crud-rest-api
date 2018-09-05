package dto

type TaskDto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	Tasks = map[int]*TaskDto{}
	Seq   = 1
)
