package handler

import (
	"popaket/auth"
	"popaket/entity"
	"popaket/helper"
	"popaket/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
	authService auth.Service
}

func NewUserHandler(userService service.UserService, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// CreateUser godoc
// @Summary Create new User account
// @Description Create new User account
// @Tags user
// @Accept json
// @Produce json
// @Param user body entity.AuthInputs true "create user"
// @Success 201 {object} helper.Response
// @Router /popaket/users/register [post]
func (h *userHandler) RegisterUserHandler(c *gin.Context) {
	var inputUser entity.AuthInputs

	if err := c.ShouldBindJSON(&inputUser); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	if err := helper.ValidatePassword(inputUser.Password); err != nil {

		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"error": "error validate password length < 6"})

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.RegisterUser(inputUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create new user", 201, "Status Created", newUser)
	c.JSON(201, response)
}

// Loginuser godoc
// @Summary Login for user
// @Description Login for user
// @Tags user
// @ID Authentication
// @Consume json
// @Produce json
// @Param user body entity.LoginAuthInputs true "Login user"
// @Success 200 {object} helper.Response
// @Router /popaket/users/login [post]
func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var inputLoginUser entity.LoginAuthInputs

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	token, err := h.authService.GenerateTokenUser(userData.ID)

	if err != nil {
		responseError := helper.APIResponse("input data required", 500, "bad request", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success login user", 200, "success", gin.H{"token": token, "id": userData.ID})
	c.JSON(200, response)
}

// JwtUser godoc
// @Security Authorization
// @Summary Get claim jwt user
// @Description Get claim jwt user
// @Tags user
// @Accept json
// @Produce json
// @success 200 {object} helper.Response
// @Router /popaket/users/jwt [get]
func (h *userHandler) JwtHandler(c *gin.Context) {
	id := c.MustGet("currentUser").(string)

	c.JSON(200, gin.H{"claims": id})
}