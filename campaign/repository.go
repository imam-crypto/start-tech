package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByIdUser(UserID int) ([]Campaign, error)
	FindbyId(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImageAsNonPrimary(campaignID int) (bool, error)
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

	err := r.db.Preload("User").Preload("CampaignImages").Where("id =?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}
func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
func (r *repository) Update(campaign Campaign) (Campaign, error) {

	err := r.db.Save(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil

}

func (r *repository) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error

	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}
func (r *repository) MarkAllImageAsNonPrimary(campaignID int) (bool, error) {
	err := r.db.Model(CampaignImage{}).Where("campaign_id =?", campaignID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}
	return true, nil
}
