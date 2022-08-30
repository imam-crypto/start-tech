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
