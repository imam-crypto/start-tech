package transaction

import (
	"pustaka-api/campaign"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByID(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionByUserID(UserID int) ([]Transaction, error)
	GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
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
