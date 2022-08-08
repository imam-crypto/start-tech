package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByIdUser(UserID int) ([]Campaign, error)
	FindbyId(ID int) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) FindByIdUser(UserID int) ([]Campaign, error) {
	var campaigns []Campaign

	errCam := r.db.Where("user_id = ?", UserID).Preload("CampaignImage", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if errCam != nil {
		return campaigns, errCam
	}
	return campaigns, nil
}

func (r *repository) FindbyId(ID int) (Campaign, error) {

	var campaign Campaign

	err := r.db.Preload("users").Preload("campaign_images").Where("id =?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
