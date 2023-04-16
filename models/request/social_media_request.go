package request

// SocialMediaCreateRequest represents the social media create request
type SocialMediaCreateRequest struct {
	Name           string `binding:"required" json:"name" form:"name"`
	SocialMediaUrl string `binding:"required" json:"social_media_url" form:"social_media_url"`
}

// SocialMediaUpdateRequest represents the social media update request
type SocialMediaUpdateRequest struct {
	Name           string `binding:"required" json:"name,omitempty" form:"name,omitempty"`
	SocialMediaUrl string `binding:"required" json:"social_media_url,omitempty" form:"social_media_url,omitempty"`
}