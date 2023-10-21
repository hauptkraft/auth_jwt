package controller

import (
	"go-jwt/data/response"
	"go-jwt/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository.UsersRepository
}

func NewUsersController(repository repository.UsersRepository) *UserController {
	return &UserController{userRepository: repository}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	// currentUser := ctx.MustGet("currentUser").(model.Users)
	users := controller.userRepository.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "пользоовательские данные успешно получены!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
