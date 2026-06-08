package shorturldto

type CreateShortURLRequest struct {
	LongURL string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id"`
}