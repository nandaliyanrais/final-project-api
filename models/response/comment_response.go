package response

import "time"

// CommentCreateResponse represents the comment create response
type CommentCreateResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// CommentUserGetAllResponse represents the comment user get all response
type CommentUserGetAllResponse struct {
	Username string `json:"username"`
}

// CommentPhotoGetAllResponse represents the comment photo get all response
type CommentPhotoGetAllResponse struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

// CommentGetAllResponse represents the comment get all response
type CommentGetAllResponse struct {
	ID        uint                       `json:"id"`
	Message   string                     `json:"message"`
	PhotoID   uint                       `json:"photo_id"`
	UserID    uint                       `json:"user_id"`
	CreatedAt time.Time                  `json:"created_at"`
	UpdatedAt time.Time                  `json:"updated_at"`
	User      CommentUserGetAllResponse  `json:"user"`
	Photo     CommentPhotoGetAllResponse `json:"photo"`
}

// CommentGetOneResponse represents the comment get one response
type CommentGetOneResponse struct {
	ID        uint                       `json:"id"`
	Message   string                     `json:"message"`
	PhotoID   uint                       `json:"photo_id"`
	UserID    uint                       `json:"user_id"`
	CreatedAt time.Time                  `json:"created_at"`
	UpdatedAt time.Time                  `json:"updated_at"`
	User      CommentUserGetAllResponse  `json:"user"`
	Photo     CommentPhotoGetAllResponse `json:"photo"`
}

// CommentUpdateResponse represents the comment update response
type CommentUpdateResponse struct {
	ID        uint                       `json:"id"`
	Message   string                     `json:"message"`
	PhotoID   uint                       `json:"photo_id"`
	UserID    uint                       `json:"user_id"`
	UpdatedAt time.Time                  `json:"updated_at"`
}

// CommentDeleteResponse represents the comment delete response
type CommentDeleteResponse struct {
	Message string `json:"message"`
}
