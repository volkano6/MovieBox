package controllers

type OkResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type HomeResponse struct {
	Status           string    `json:"status"`
	Logged           bool      `json:"logged"`
	CurrentMonth     string    `json:"month"`
	MonthlyPopular   []Movie   `json:"monthly_popular"`
	AllTimeFavorites []Movie   `json:"all_time_favorites"`
	LatestReviews    []Comment `json:"latest_reviews"`
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
	ID                  int
	Title               string
	Description         string
	ReleaseDate         string
	Poster              string
	Rating              float32
	Length              int
	Genre               string
	FavoriteCount       int
	WatchedCount        int
	MonthlyWatchedCount any
}

type Comment struct {
	UserID      int
	Username    string
	UserAvatar  any
	MovieID     int
	MovieTitle  string
	Comment     string
	CommentDate string
}

type Favorite struct {
	MovieID     	int
	UserID    		int
	FavoritedDate 	string
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

type ProfileResponse struct {
	Status 			string		`json:"status"`
	Logged			bool		`json:"logged"`
	User			User		`json:"user"`
	UserWatched 	[]Movie		`json:"user_watched"`	
	UserWatchlist 	[]Movie		`json:"user_watchlist"`
	UserComments 	[]Comment	`json:"user_comments"`
	UserFavorites	[]Movie	`json:"user_favorites"`
}
