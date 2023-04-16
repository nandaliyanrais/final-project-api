package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"mygram-api/helpers"
	"mygram-api/models/domain"
	"mygram-api/models/request"
	"mygram-api/models/response"
	"mygram-api/social_medias/service"
)

type SocialMediaController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type SocialMediaControllerService struct {
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaController(socialMediaService service.SocialMediaService) SocialMediaController {
	return &SocialMediaControllerService{SocialMediaService: socialMediaService}
}

// Create social media godoc
// @Summary Add a social media
// @Description Create and store a social media with authentication user
// @Tags Social media
// @Accept json
// @Produce json
// @Param json body request.SocialMediaCreateRequest true "Add Social Media"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Security Bearer
// @Router /social-media [post]
func (socialMediaController *SocialMediaControllerService) Create(c *gin.Context) {

	var req request.SocialMediaCreateRequest
	
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	contentType := helpers.GetContentType(c)

	if contentType == "application/json" {
		if err := c.ShouldBindJSON(&req); err != nil {
			validationError := err.(validator.ValidationErrors)
			fieldErrorResponse := make(map[string]interface{})

			for _, v := range validationError {
				fieldErrorResponse[strings.ToLower(v.Field())] = helpers.GetValidationErrorMsg(v)
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
				Errors: fieldErrorResponse,
			})

			return
		}
	} else if contentType == "application/x-www-form-urlencoded" {
		req.Name = c.PostForm("name")
		req.SocialMediaUrl = c.PostForm("social_media_url")
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: "Unsupported Content-Type",
		})

		return
	}

	socialMedia := domain.SocialMedia{
		Name:           req.Name,
		SocialMediaUrl: req.SocialMediaUrl,
		UserID:         userID,
	}

	if err := socialMediaController.SocialMediaService.Create(&socialMedia); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Data: response.SocialMediaCreateResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserID:         socialMedia.UserID,
			CreatedAt:      socialMedia.CreatedAt,
		},
	})
}

// GetAll social media godoc
// @Summary Get all social media
// @Description Get all social media with authentication user
// @Tags Social media
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Security Bearer
// @Router /social-media [get]
func (socialMediaController *SocialMediaControllerService) GetAll(c *gin.Context) {

	var (
		socialMedias         []domain.SocialMedia
		err                  error
		socialMediasResponse = []response.SocialMediaGetAllResponse{}
	)

	if socialMedias, err = socialMediaController.SocialMediaService.GetAll(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})

		return
	}

	for _, socialMedia := range socialMedias {
		socialMediasResponse = append(socialMediasResponse, response.SocialMediaGetAllResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserID:         socialMedia.UserID,
			CreatedAt:      socialMedia.CreatedAt,
			UpdatedAt:      socialMedia.UpdatedAt,
			User: response.SocialMediaUserGetAllResponse{
				Username: socialMedia.User.Username,
			},
		})
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: gin.H{
			"social_medias": socialMediasResponse,
		},
	})
}

// GetOne social media godoc
// @Summary Get one social media
// @Description Get one social media by id with authentication user
// @Tags Social media
// @Accept json
// @Produce json
// @Param id path int true "Social Media ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Security Bearer
// @Router /social-media/{id} [get]
func (socialMediaController *SocialMediaControllerService) GetOne(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	socialMedia, err := socialMediaController.SocialMediaService.GetOne(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Errors: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: gin.H{
			"social_media": response.SocialMediaGetOneResponse{
				ID:             socialMedia.ID,
				Name:           socialMedia.Name,
				SocialMediaUrl: socialMedia.SocialMediaUrl,
				UserID:         socialMedia.UserID,
				CreatedAt:      socialMedia.CreatedAt,
				UpdatedAt:      socialMedia.UpdatedAt,
				User: response.SocialMediaUserGetOneResponse{
					Username: socialMedia.User.Username,
				},
			},
		},
	})
}

// Update social media godoc
// @Summary Update a social media
// @Description Update a social media by id with authentication user
// @Tags Social media
// @Accept json
// @Produce json
// @Param id path int true "Social Media ID"
// @Param json body request.SocialMediaUpdateRequest true "Update Social Media"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Security Bearer
// @Router /social-media/{id} [put]
func (socialMediaController *SocialMediaControllerService) Update(c *gin.Context) {

	var (
		req                request.SocialMediaUpdateRequest
		updatedSocialMedia domain.SocialMedia
		err                error
	)
	
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	
	contentType := helpers.GetContentType(c)
	
	if contentType == "application/json" {
		if err = c.ShouldBindJSON(&req); err != nil {
			validationError := err.(validator.ValidationErrors)
			fieldErrorResponse := make(map[string]interface{})
	
			for _, v := range validationError {
				fieldErrorResponse[strings.ToLower(v.Field())] = helpers.GetValidationErrorMsg(v)
			}
	
			c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
				Errors: fieldErrorResponse,
			})
	
			return
		}
	} else if contentType == "application/x-www-form-urlencoded" {
		req.Name = c.PostForm("name")
		req.SocialMediaUrl = c.PostForm("social_media_url")
	} else {
		
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: gin.H{
				"message": "Invalid Content-Type header value. Only application/json or application/x-www-form-urlencoded are allowed.",
			},
		})
		return

	}
	
	socialMedia := domain.SocialMedia{
		ID:     uint(id),
		UserID: userID,
	}
	
	if req.Name != "" {
		socialMedia.Name = req.Name
	}
	
	if req.SocialMediaUrl != "" {
		socialMedia.SocialMediaUrl = req.SocialMediaUrl
	}
	
	if updatedSocialMedia, err = socialMediaController.SocialMediaService.Update(socialMedia); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
	
		return
	}
	
	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: response.SocialMediaUpdateResponse{
			ID:             updatedSocialMedia.ID,
			Name:           updatedSocialMedia.Name,
			SocialMediaUrl: updatedSocialMedia.SocialMediaUrl,
			UserID:         updatedSocialMedia.UserID,
			UpdatedAt:      updatedSocialMedia.UpdatedAt,
		},
	})
}

// Delete social media godoc
// @Summary Delete a social media
// @Description Delete a social media by id with authentication user
// @Tags Social media
// @Accept json
// @Produce json
// @Param id path int true "Social Media ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Security Bearer
// @Router /social-media/{id} [delete]
func (socialMediaController *SocialMediaControllerService) Delete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := socialMediaController.SocialMediaService.Delete(uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: response.SocialMediaDeleteResponse{
			Message: "Social Media Deleted Successfully",
		},
	})
}