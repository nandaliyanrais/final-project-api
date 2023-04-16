package request

// PhotoCreateRequest represents the photo create request
type PhotoCreateRequest struct {
	Title    string `binding:"required" json:"title" form:"title"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `binding:"required" json:"photo_url" form:"photo_url"`
}

// PhotoUpdateRequest represents the photo update request
type PhotoUpdateRequest struct {
	Title    string `binding:"required" json:"title,omitempty" form:"title,omitempty"`
	Caption  string `json:"caption,omitempty" form:"caption,omitempty"`
	PhotoUrl string `binding:"required" json:"photo_url,omitempty" form:"photo_url,omitempty"`
}