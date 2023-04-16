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
	"mygram-api/comments/service"
	photoService "mygram-api/photos/service"
)

type CommentController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CommentControllerService struct {
	CommentService service.CommentService
	PhotoService   photoService.PhotoService
}

func NewCommentController(commentService service.CommentService, photoService photoService.PhotoService) CommentController {
	return &CommentControllerService{CommentService: commentService, PhotoService: photoService}
}

// Create comment godoc
// @Summary Create a comment
// @Description Create and store a comment with authentication user
// @Tags comments
// @Accept json
// @Produce json
// @Param json body request.CommentCreateRequest true "Add Comment"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Security Bearer
// @Router /comments [post]
func (commentController *CommentControllerService) Create(c *gin.Context) {

	var req request.CommentCreateRequest

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	contentType := c.Request.Header.Get("Content-Type")

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

	photoID := req.PhotoID

	if _, err := commentController.PhotoService.GetOne(photoID); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, response.ErrorResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Errors: gin.H{
				"message": "Record not found",
			},
		})

		return
	}

	comment := domain.Comment{
		Message: req.Message,
		PhotoID: req.PhotoID,
		UserID:  userID,
	}

	if err := commentController.CommentService.Create(&comment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Data: response.CommentCreateResponse{
			ID:        comment.ID,
			UserID:    comment.UserID,
			PhotoID:   comment.PhotoID,
			Message:   comment.Message,
			CreatedAt: comment.CreatedAt,
		},
	})
}

// GetAll comment godoc
// @Summary Get all comments
// @Description Get all comments with authentication user
// @Tags comments
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Security Bearer
// @Router /comments [get]
func (commentController *CommentControllerService) GetAll(c *gin.Context) {

	comments, err := commentController.CommentService.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
	}

	commentsResponse := []response.CommentGetAllResponse{}
	for _, comment := range comments {
		commentsResponse = append(commentsResponse, response.CommentGetAllResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			User: response.CommentUserGetAllResponse{
				Username: comment.User.Username,
			},
			Photo: response.CommentPhotoGetAllResponse{
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
			},
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: commentsResponse,
	})
}

// GetOne comment godoc
// @Summary Get one comment
// @Description Get one comment by id with authentication user
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path int true "Comment ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Security Bearer
// @Router /comments/{commentId} [get]
func (commentController *CommentControllerService) GetOne(c *gin.Context) {
	commentID, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	comment, err := commentController.CommentService.GetOne(uint(commentID))
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Errors: err.Error(),
		})
		return
	}

	commentResponse := response.CommentGetAllResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		User: response.CommentUserGetAllResponse{
			Username: comment.User.Username,
		},
		Photo: response.CommentPhotoGetAllResponse{
			Title:    comment.Photo.Title,
			Caption:  comment.Photo.Caption,
			PhotoUrl: comment.Photo.PhotoUrl,
		},
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: commentResponse,
	})
}

// Update godoc
// @Summary Update a comment
// @Description Update a comment by id with authentication user
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path int true "Comment ID"
// @Param json body request.CommentUpdateRequest true "Update Comment"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Security Bearer
// @Router /comments/{commentId} [put]
func (commentController *CommentControllerService) Update(c *gin.Context) {
	var (
		req            request.CommentUpdateRequest
		updatedComment domain.Comment
		err            error
	)

	commentID, _ := strconv.ParseUint(c.Param("commentId"), 10, 32)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	contentType := c.Request.Header.Get("Content-Type")

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
		req.Message = c.PostForm("message")
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

	comment := domain.Comment{
		ID:     uint(commentID),
		UserID: userID,
	}

	if req.Message != "" {
		comment.Message = req.Message
	}

	if updatedComment, err = commentController.CommentService.Update(comment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: response.CommentUpdateResponse{
			ID:        updatedComment.ID,
			Message:   updatedComment.Message,
			PhotoID:   updatedComment.PhotoID,
			UserID:    updatedComment.UserID,
			UpdatedAt: updatedComment.UpdatedAt,
		},
	})
}

// Delete godoc
// @Summary Delete a comment
// @Description Delete a comment by id with authentication user
// @Tags comments
// @Accept json
// @Produce json
// @Param commentId path int true "Comment ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Security Bearer
// @Router /comments/{commentId} [delete]
func (commentController *CommentControllerService) Delete(c *gin.Context) {

	commentID, _ := strconv.ParseUint(c.Param("commentId"), 10, 32)

	if err := commentController.CommentService.Delete(uint(commentID)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: response.CommentDeleteResponse{
			Message: "Comment deleted successfully",
		},
	})
}
