package response

import (
	"time"
)

// PhotoCreateResponse represents the photo create response
type PhotoCreateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// PhotoUserGetAllReponse represents the photo user get all response
type PhotoUserGetAllReponse struct {
	Username string `json:"username"`
}

// PhotoGetAllResponse represents the photo get all response
type PhotoGetAllResponse struct {
	ID        uint                   `json:"id"`
	Title     string                 `json:"title"`
	Caption   string                 `json:"caption"`
	PhotoUrl  string                 `json:"photo_url"`
	UserID    uint                   `json:"user_id"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	User      PhotoUserGetAllReponse `json:"user"`
}

// PhotoGetOneResponse represents the photo get one response
type PhotoGetOneResponse struct {
	ID        uint                   `json:"id"`
	Title     string                 `json:"title"`
	Caption   string                 `json:"caption"`
	PhotoUrl  string                 `json:"photo_url"`
	UserID    uint                   `json:"user_id"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	User      PhotoUserGetAllReponse `json:"user"`
}

// PhotoUpdateResponse represents the photo update response
type PhotoUpdateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PhotoDeleteResponse represents the photo delete response
type PhotoDeleteResponse struct {
	Message string `json:"message"`
}