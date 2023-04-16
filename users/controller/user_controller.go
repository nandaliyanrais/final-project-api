package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"mygram-api/helpers"
	"mygram-api/models/domain"
	"mygram-api/models/request"
	"mygram-api/models/response"
	"mygram-api/users/service"
)

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type UserControllerService struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerService{UserService: userService}
}

// Register godoc
// @Summary Register a user
// @Description Create and store a new user
// @Tags users
// @Accept json
// @Produce json
// @Param json body request.UserRegisterRequest true "User Register Request"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /users/register [post]
func (userController *UserControllerService) Register(c *gin.Context) {

	var req request.UserRegisterRequest

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
			Errors: "Unsupported Content-Type",
		})
		return
	}

	user := domain.User{
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	}

	if err := userController.UserService.Register(&user); err != nil {
		fieldErrorResponse := make(map[string]interface{})

		if strings.Contains(err.Error(), "idx_users_email") {
			fieldErrorResponse["email"] = "Email is already used"
		}

		if strings.Contains(err.Error(), "idx_users_username") {
			fieldErrorResponse["username"] = "Username is already used"
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: fieldErrorResponse,
		})

		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Data: response.UserRegisterResponse{
			Age:      user.Age,
			Email:    user.Email,
			ID:       user.ID,
			Username: user.Username,
		},
	})
}

// Login godoc
// @Summary Login a user
// @Description Authentication a user and retrieve a token
// @Tags users
// @Accept json
// @Produce json
// @Param json body request.UserLoginRequest true "User Login Request"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /users/login [post]
func (userController *UserControllerService) Login(c *gin.Context) {

	var req request.UserLoginRequest

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
		c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, response.ErrorResponse{
			Code:   http.StatusUnsupportedMediaType,
			Status: "Unsupported Media Type",
			Errors: "Request content type must be either 'application/json' or 'application/x-www-form-urlencoded'",
		})

		return
	}

	user := domain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := userController.UserService.Login(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Errors: err.Error(),
		})

		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, response.SuccessResponse{
		Data: response.UserLoginResponse{
			Token: token,
		},
	})
}