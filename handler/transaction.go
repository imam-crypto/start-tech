package handler

import (
	"net/http"
	"pustaka-api/helper"
	"pustaka-api/transaction"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	// idString := c.Param("id")
	// id, _ := strconv.Atoi(idString)

	currentUser := c.MustGet("current_user").(user.User)

	input.ID = currentUser.ID

	tr, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {

		response := helper.APIResponse("Data Not Found", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// transaction.FormatterCampaignTransactions(transaction)
	response := helper.APIResponse("Campaign Transaction", http.StatusOK, "success", transaction.FormatterCampaignTransactions(tr))
	c.JSON(http.StatusOK, response)

}

func (h *transactionHandler) GeUserTransaction(c *gin.Context) {

	currentUser := c.MustGet("current_user").(user.User)
	userID := currentUser.ID

	userTr, errUser := h.service.GetTransactionByUserID(userID)
	if errUser != nil {

		response := helper.APIResponse("Data Not Found", http.StatusBadRequest, "failed", errUser)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Get User Transaction Success", http.StatusOK, "success", transaction.FormatterCampaignTransactions(userTr))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) Create(c *gin.Context) {

	var input transaction.CreateTransactionInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Binding Error", http.StatusUnprocessableEntity, "failed", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)
	userID := currentUser.ID

	input.User.ID = userID

	tr, err := h.service.CreateTransaction(input)
	if err != nil {

		response := helper.APIResponse("Data Not Found", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Create User Transaction Success", http.StatusOK, "success", transaction.FormatTransaction(tr))
	c.JSON(http.StatusOK, response)
}
func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput
	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Binding Error", http.StatusUnprocessableEntity, "failed", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.service.ProcessPayment(input)

	if err != nil {

		response := helper.APIResponse("Data Not Found", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, input)

}
