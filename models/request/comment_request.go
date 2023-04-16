package request

// CommentCreateRequest represents the comment create request
type CommentCreateRequest struct {
	Message string `binding:"required" json:"message" form:"message"`
	PhotoID uint   `json:"photo_id" form:"photo_id"`
}

// CommentUpdateRequest represents the comment update request
type CommentUpdateRequest struct {
	Message string `binding:"required" json:"message,omitempty" form:"message,omitempty"`
}