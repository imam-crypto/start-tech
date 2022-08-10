package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/auth"
	"pustaka-api/helper"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
)

type userhandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userhandler {
	return &userhandler{userService, authService}
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
	token, er := h.authService.GenerateToken(NewUser.ID)
	if er != nil {
		response := helper.APIResponse("Register Failed", http.StatusBadRequest, er.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(NewUser, token)

	response := helper.APIResponse("Account has been created", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userhandler) Login(c *gin.Context) {

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "failed", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loogedinUser, er := h.userService.Login(input)

	if er != nil {

		errorMessage := gin.H{"errors": er.Error()}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "failed", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, er := h.authService.GenerateToken(loogedinUser.ID)
	if er != nil {
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, er.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatUser(loogedinUser, token)

	response := helper.APIResponse("Login Succesfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userhandler) CheckEmailAvaibility(c *gin.Context) {

	var input user.CheckEmaiInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email Checking Failed", http.StatusBadRequest, "failed", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, er := h.userService.IsEmailAvailable(input)

	if er != nil {
		errorMessage := gin.H{"errors": "server error"}

		response := helper.APIResponse("Email Checking Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}
	metaMessage := "Email Has Been Registered"

	if isEmailAvailable {
		metaMessage = "Email Is Available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "failed", data)
	c.JSON(http.StatusOK, response)

}
func (h *userhandler) UploadAvatar(c *gin.Context) {

	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload Failed", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("current_user").(user.User)

	userID := currentUser.ID

	path := fmt.Sprintf("user/images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload Failed", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload Failed", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Upload Succesfuly", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
