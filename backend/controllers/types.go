package controllers

type OkResponse struct {
	Status string `json:"status"`
	Logged bool   `json:"logged"`
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

type ProfileResponse struct {
	Status        string    `json:"status"`
	Logged        bool      `json:"logged"`
	User          User      `json:"user"`
	UserWatched   []Movie   `json:"user_watched"`
	UserWatchlist []Movie   `json:"user_watchlist"`
	UserComments  []Comment `json:"user_comments"`
	UserFavorites []Movie   `json:"user_favorites"`
}

type MovieResponse struct {
	Status   string    `json:"status"`
	Logged   bool      `json:"logged"`
	Movie    Movie     `json:"movie"`
	Comments []Comment `json:"comments"`
}

type SearchResponse struct {
	Status string  `json:"status"`
	Logged bool    `json:"logged"`
	Movies []Movie `json:"movies"`
	Users  []User  `json:"users"`
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
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Displayname string `json:"displayname"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Avatar      string `json:"avatar"`
	Bio         string `json:"bio"`
	Location    string `json:"location"`
	Twitter     string `json:"twitter"`
	Instagram   string `json:"instagram"`
	Type        int    `json:"type"`
}

type Movie struct {
	ID                  int      `json:"id"`
	Title               string   `json:"title"`
	Description         string   `json:"description"`
	ReleaseDate         string   `json:"release_date"`
	Poster              string   `json:"poster"`
	Rating              float32  `json:"rating"`
	Length              int      `json:"length"`
	Genres              []string `json:"genres"`
	FavoriteCount       int      `json:"favorite_count"`
	WatchedCount        int      `json:"watched_count"`
	MonthlyWatchedCount any      `json:"monthly_watched_count"`
}

type Comment struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"user_name"`
	UserAvatar  any    `json:"user_avatar"`
	MovieID     int    `json:"movie_id"`
	MovieTitle  string `json:"movie_title"`
	Comment     string `json:"comment"`
	CommentDate string `json:"comment_date"`
}

type CommentPost struct {
	Comment string `json:"comment" binding:"required"`
}

type Favorite struct {
	MovieID       int    `json:"movie_id"`
	UserID        int    `json:"user_id"`
	FavoritedDate string `json:"favorited_date"`
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
