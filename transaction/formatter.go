package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"ID"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.First_name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func FormatterCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {

	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionsFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}
	return transactionsFormatter

}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}
type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatUserTransaction(tr Transaction) UserTransactionFormatter {

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = tr.Campaign.Name
	campaignFormatter.ImageUrl = ""
	if len(tr.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = tr.Campaign.CampaignImages[0].File_name
	}
	format := UserTransactionFormatter{
		ID:        tr.ID,
		Amount:    int(tr.Amount),
		Status:    tr.Status,
		CreatedAt: tr.CreatedAt,
		Campaign:  campaignFormatter,
	}
	return format

}
func FormatUserTransactions(tr []Transaction) []UserTransactionFormatter {
	if len(tr) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range tr {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}
	return transactionsFormatter
}

type TransactionFormatter struct {
	ID         int       `json:"ID"`
	CampaignID int       `json:"campaign_id"`
	UserID     int       `json:"user_id"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	PaymentUrl string    `json:"payment_url"`
	Name       string    `json:"name"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatTransaction(tr Transaction) TransactionFormatter {

	format := TransactionFormatter{}
	format.Name = tr.User.First_name
	format.CampaignID = tr.Campaign_id
	format.Status = tr.Status
	format.Code = tr.Code
	format.Amount = tr.Amount
	format.PaymentUrl = tr.PaymentUrl
	format.CreatedAt = tr.CreatedAt
	return format

}
