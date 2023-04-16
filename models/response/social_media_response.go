package response

import "time"

// SocialMediaCreateResponse represents the social media create response
type SocialMediaCreateResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

// SocialMediaUserGetAllResponse represents the social media user get all response
type SocialMediaUserGetAllResponse struct {
	Username string `json:"username"`
}

// SocialMediaUserGetOneResponse represents the social media user get one response
type SocialMediaUserGetOneResponse struct {
	Username string `json:"username"`
}

// SocialMediaGetAllResponse represents the social media get all response
type SocialMediaGetAllResponse struct {
	ID             uint                          `json:"id"`
	Name           string                        `json:"name"`
	SocialMediaUrl string                        `json:"social_media_url"`
	UserID         uint                          `json:"user_id"`
	UpdatedAt      time.Time                     `json:"updated_at"`
	CreatedAt      time.Time                     `json:"created_at"`
	User           SocialMediaUserGetAllResponse `json:"user"`
}

// SocialMediaGetOneResponse represents the social media get one response
type SocialMediaGetOneResponse struct {
	ID             uint                          `json:"id"`
	Name           string                        `json:"name"`
	SocialMediaUrl string                        `json:"social_media_url"`
	UserID         uint                          `json:"user_id"`
	UpdatedAt      time.Time                     `json:"updated_at"`
	CreatedAt      time.Time                     `json:"created_at"`
	User           SocialMediaUserGetOneResponse `json:"user"`
}

// SocialMediaUpdateResponse represents the social media update response
type SocialMediaUpdateResponse struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// SocialMediaDeleteResponse represents the social media delete response
type SocialMediaDeleteResponse struct {
	Message string `json:"message"`
}
