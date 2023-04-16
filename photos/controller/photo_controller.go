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
	"mygram-api/photos/service"
)


type PhotoController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type PhotoControllerService struct {
	PhotoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) PhotoController {
	return &PhotoControllerService{PhotoService: photoService}
}

// Create photo godoc
// @Summary Create a photo
// @Description Create and store a new photo with authentication user
// @Tags photos
// @Accept json
// @Produce json
// @Param json body request.PhotoCreateRequest true "Add Photo"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Security Bearer
// @Router /photos [post]
func (photoController *PhotoControllerService) Create(c *gin.Context) {

	var req request.PhotoCreateRequest

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	contentType := helpers.GetContentType(c)

	if contentType != "application/json" && contentType != "application/x-www-form-urlencoded" {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: "Invalid Content-Type. Expected application/json or application/x-www-form-urlencoded.",
		})
		return
	}

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
		if err := c.ShouldBind(&req); err != nil {
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
	}

	photo := domain.Photo{
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
		UserID:   userID,
	}

	if err := photoController.PhotoService.Create(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Data: response.PhotoCreateResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
		},
	})
}

// GetAll photo godoc
// @Summary Get all photos
// @Description Get all photos with authentication user
// @Tags photos
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Security Bearer
// @Router /photos [get]
func (photoController *PhotoControllerService) GetAll(c *gin.Context) {

	photos, err := photoController.PhotoService.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
	}

	photosResponse := []response.PhotoGetAllResponse{}
	for _, photo := range photos {
		photosResponse = append(photosResponse, response.PhotoGetAllResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: response.PhotoUserGetAllReponse{
				Username: photo.User.Username,
			},
		})
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data:   photosResponse,
	})
}

// GetOne photo godoc
// @Summary Get one photo
// @Description Get one photo by id with authentication user
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Photo ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Security Bearer
// @Router /photos/{id} [get]
func (photoController *PhotoControllerService) GetOne(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, response.ErrorResponse{
            Code:   http.StatusBadRequest,
            Status: "Bad Request",
            Errors: err.Error(),
        })
        return
    }

    photo, err := photoController.PhotoService.GetOne(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, response.ErrorResponse{
            Code:   http.StatusNotFound,
            Status: "Not Found",
            Errors: err.Error(),
        })
        return
    }

    photoResponse := response.PhotoGetOneResponse{
        ID:        photo.ID,
        Title:     photo.Title,
        Caption:   photo.Caption,
        PhotoUrl:  photo.PhotoUrl,
        UserID:    photo.UserID,
        CreatedAt: photo.CreatedAt,
        UpdatedAt: photo.UpdatedAt,
        User: response.PhotoUserGetAllReponse{
            Username: photo.User.Username,
        },
    }

    c.JSON(http.StatusOK, response.SuccessResponse{
        Data:   photoResponse,
    })
}

// Update photo godoc
// @Summary Update a photo
// @Description Update a photo by id with authentication user
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Photo ID"
// @Param json body request.PhotoUpdateRequest true "Photo Update Request"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Security Bearer
// @Router /photos/{id} [put]
func (photoController *PhotoControllerService) Update(c *gin.Context) {
	var (
		req          request.PhotoUpdateRequest
		updatedPhoto domain.Photo
		err          error
	)

	photoID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

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
		req.Title = c.PostForm("title")
		req.Caption = c.PostForm("caption")
		req.PhotoUrl = c.PostForm("photo_url")
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: "Unsupported content type",
		})
		return
	}

	photo := domain.Photo{
		ID:       uint(photoID),
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoUrl: req.PhotoUrl,
	}

	if updatedPhoto, err = photoController.PhotoService.Update(photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: response.PhotoUpdateResponse{
			ID:        updatedPhoto.ID,
			UserID:    updatedPhoto.UserID,
			Title:     updatedPhoto.Title,
			PhotoUrl:  updatedPhoto.PhotoUrl,
			Caption:   updatedPhoto.Caption,
			UpdatedAt: updatedPhoto.UpdatedAt,
		},
	})
}

// Delete photo godoc
// @Summary Delete a photo
// @Description Delete a photo by id with authentication user
// @Tags photos
// @Accept json
// @Produce json
// @Param id path int true "Photo ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Security Bearer
// @Router /photos/{id} [delete]
func (photoController *PhotoControllerService) Delete(c *gin.Context) {
	photoID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := photoController.PhotoService.Delete(uint(photoID)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: response.PhotoDeleteResponse{
			Message: "Photo deleted successfully",
		},
	})
}