package transaction

import (
	"pustaka-api/campaign"
	"pustaka-api/payment"
	"strconv"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

type Service interface {
	GetTransactionByID(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionByUserID(UserID int) ([]Transaction, error)
	GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {
	// campaign, err := s.campaignRepository.FindbyId(input.ID)
	// if err != nil {
	// 	return []Transaction{}, err
	// }
	// return []Transaction{}, nil

	// if campaign.User_id != input.User.ID {
	// 	return []Transaction{}, errors.New("ojos")
	// }
	transactions, errTrans := s.repository.GetCampaignByID(input.ID)
	if errTrans != nil {
		return transactions, errTrans
	}
	return transactions, nil
}

func (s *service) GetTransactionByID(input GetCampaignTransactionInput) ([]Transaction, error) {

	transactions, errTrans := s.repository.GetCampaignByID(input.ID)
	if errTrans != nil {
		return transactions, errTrans
	}
	return transactions, nil

}

func (s *service) GetTransactionByUserID(UserID int) ([]Transaction, error) {

	transactions, errTrans := s.repository.GetByUserID(UserID)
	if errTrans != nil {
		return transactions, errTrans
	}
	return transactions, nil

}
func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	tr := Transaction{}
	tr.Campaign_id = input.CampaignID
	tr.Amount = float64(input.Amount)
	tr.User_id = input.User.ID
	tr.Code = "GDPNIHBOS007"
	tr.Status = "PENDING"

	newTransaction, err := s.repository.Save(tr)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.TransactionPayment{
		ID:     newTransaction.ID,
		Amount: int(newTransaction.Amount),
	}

	paymentUrl, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}
	newTransaction.PaymentUrl = paymentUrl

	// newTransaction, err = s.repository.Update(tr)
	// if err != nil {
	// 	return newTransaction, err
	// }

	return newTransaction, nil
}
func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetTransactionByID(transaction_id)
	if err != nil {
		return err
	}
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "PAID"
	} else if input.TransactionStatus == "SETTLEMENT" {
		transaction.Status = "PAID"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "CANCELLED"
	}

	updatedTr, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindbyId(updatedTr.Campaign_id)
	if err != nil {
		return err
	}
	if updatedTr.Status == "PAID" {
		campaign.Backer = campaign.Backer + 1
		campaign.Current_amount = campaign.Current_amount + int(updatedTr.Amount)

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}
	return nil
}
