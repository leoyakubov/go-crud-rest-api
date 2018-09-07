package errors

const (
	ERR_OCCURED       = "An error occured whipe processing your request"
	INVALID_TASK_ID   = "Invalid task id"
	TASK_DOESNT_EXIST = "Task you are looking for does not exists"
)

type ResponseError struct {
	ErrorCodeId int    `json:"errorCodeId"`
	DevMessage  string `json:"devMessage"`
	UserMessage string `json:"userMessage"`
}
