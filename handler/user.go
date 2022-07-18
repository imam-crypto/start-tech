package handler

import (
	"net/http"
	"pustaka-api/helper"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
)

type userhandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userhandler {
	return &userhandler{userService}
}
func (h *userhandler) RegisterUser(c *gin.Context) {
	// menangkap input dari user
	// mapping input dari user ke struct RegisterUser
	// struct di atas di parsing menjadi parameter ke service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "failed", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	NewUser, er := h.userService.RegisterUser(input)
	if er != nil {
		response := helper.APIResponse("Register Failed", http.StatusBadRequest, "failed", nil)

		c.JSON(http.StatusBadRequest, response)
		return
		// c.JSON(http.StatusBadRequest, nil)
	}

	formatter := user.FormatUser(NewUser, "tokentokentokenlistrik")

	response := helper.APIResponse("Account has been created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
