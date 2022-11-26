package controllers

type OkResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func OkMessage(data any) OkResponse {
	return OkResponse{
		Status: "ok",
		Data:   data,
	}
}

func ErrorMessage(message string) ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: message,
	}
}

type User struct {
	ID          int
	Displayname string
	Email       string
	DateOfBirth string
	Avatar      string
	Bio         string
	Location    string
	Twitter     string
	Instagram   string
	Type        int
}

type Movie struct {
	ID          int
	Title       string
	ReleaseDate string
	Poster      string
	Rating      float32
}

type Comment struct {
	UserID      int
	Username    string
	MovieID     string
	Comment     string
	CommentDate string
}