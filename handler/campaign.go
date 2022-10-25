package handler

import (
	"fmt"
	"net/http"
	"start-tech/campaign"
	"start-tech/helper"
	"start-tech/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.FindByUserID(userID)

	if err != nil {

		response := helper.APIResponse("Data Not Found", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Data Not Found", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {

	// var input campaign.GetCampaignInput
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	// fmt.Println("id baru ", id)
	// err := c.ShouldBindUri(&input)

	// if err != nil {

	// 	response := helper.APIResponse("Data Not Found", http.StatusBadRequest, "failed", nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }
	campaignDetail, err := h.service.FindById(id)

	if err != nil {

		response := helper.APIResponse("Data Not Found", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Data Campaign", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {

	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		response := helper.APIResponse("Failed ", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("current_user").(user.User)

	input.User = currentUser
	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {

		response := helper.APIResponse("Failed Create Campaign", http.StatusUnprocessableEntity, "failed Create Campaign", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := helper.APIResponse("Success ", http.StatusOK, "Success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}
func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	// var inputID campaign.GetCampaignInput

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	fmt.Println("id baru update", id)

	// err := c.ShouldBindUri(&inputID)
	// fmt.Println("[][][][", c.ShouldBindUri(&inputID))
	if err != nil {

		response := helper.APIResponse("Failed Update Campaign ", http.StatusBadRequest, "failed", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)

	if err != nil {

		response := helper.APIResponse("Failed Updated Data ", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("current_user").(user.User)

	inputData.User = currentUser

	updateCampaign, err := h.service.UpdateCampaign(id, inputData)
	if err != nil {

		response := helper.APIResponse("Failed Updated Data ", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success ", http.StatusOK, "Success", campaign.FormatCampaign(updateCampaign))
	c.JSON(http.StatusOK, response)
}
func (h *campaignHandler) UploadImage(c *gin.Context) {

	var input campaign.InputCreateImage

	err := c.ShouldBind(&input)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload Image Failed", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("image")

	currentUser := c.MustGet("current_user").(user.User)

	input.User.ID = currentUser.ID
	userID := currentUser.ID

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload Image Failed", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("user/images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload Image Failed", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.CreateImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload Image Failed", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Upload Image Succesfuly", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
func (h *campaignHandler) Tes(c *gin.Context) {

	fmt.Print("tes")
}
