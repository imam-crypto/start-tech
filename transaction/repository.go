package transaction

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}
type Repository interface {
	GetCampaignByID(campaignID int) ([]Transaction, error)
	GetByUserID(UserId int) ([]Transaction, error)
	GetTransactionByID(ID int) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	Save(transaction Transaction) (Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GetCampaignByID(campaignID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Preload("User").Where("campaign_id =?", campaignID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) GetByUserID(UserId int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary=1").Where("user_id =?", UserId).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
func (r *repository) GetTransactionByID(ID int) (Transaction, error) {
	var transaction Transaction

	err := r.db.Where("id =?", ID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
