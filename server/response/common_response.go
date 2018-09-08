package response

const (
	ERR_OCCURED             = "An error occured whipe processing your request"
	INVALID_TASK_ID         = "Invalid task id"
	TASK_DOESNT_EXIST       = "Task you are looking for does not exists"
	INCORRECT_AUTH_TOKEN    = "Incorrect authorization token provided"
	JWT_MISSING             = "JWT secret token is missed"
	FORBIDDEN               = "Forbidden"
	USERNAME_PASS_INCORRECT = "Username or password is invalid"
)

type ResponseError struct {
	ErrorCodeId int    `json:"errorCodeId"`
	ServerError string `json:"serverError"`
	UserMessage string `json:"userMessage"`
}
