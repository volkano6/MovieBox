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
	Username    string
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
	Description string
	ReleaseDate string
	Poster      string
	Rating      float32
	Length		int
	Genre		string
}

type Comment struct {
	UserID      int
	Username    string
	MovieID     string
	Comment     string
	CommentDate string
}

type UserRegister struct {
	Username    string `json:"user_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}
type UserLogin struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}