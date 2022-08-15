package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	FindByUserID(UserId int) ([]Campaign, error)
	FindById(ID int) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(ID int, inputData CreateCampaignInput) (Campaign, error)
	CreateImage(input InputCreateImage, fileLocarion string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByUserID(UserID int) ([]Campaign, error) {
	if UserID != 0 {
		campaigns, err := s.repository.FindByIdUser(UserID)

		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()

	if err != nil {
		return campaigns, err
	}
	return campaigns, err

}
func (s *service) FindById(ID int) (Campaign, error) {

	// var campaign Campaign

	campaign, err := s.repository.FindbyId(ID)
	if err != nil {
		return campaign, err
	}
	return campaign, err
}
func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}

	campaign.Name = input.Name
	campaign.Short_description = input.Short_description
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.Goal_amount = input.Goal_amount
	campaign.User_id = input.User.ID
	newSlug := fmt.Sprintf("%s %s %d", "PRABS", input.Name, input.User.ID)

	campaign.Slug = slug.Make(newSlug)

	newCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return newCampaign, err
	}
	return newCampaign, nil
}
func (s *service) UpdateCampaign(ID int, inputData CreateCampaignInput) (Campaign, error) {

	campaign, err := s.repository.FindbyId(ID)
	if err != nil {
		return campaign, err
	}
	if campaign.User_id != inputData.User.ID {
		return campaign, errors.New("not authorized for this campaign")
	}

	campaign.Name = inputData.Name
	campaign.Short_description = inputData.Short_description
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.Goal_amount = inputData.Goal_amount

	updatedCampaign, erUpdate := s.repository.Update(campaign)
	if erUpdate != nil {
		return campaign, err
	}
	return updatedCampaign, nil
}

func (s *service) CreateImage(input InputCreateImage, fileLocation string) (CampaignImage, error) {
	is_primary := 0

	campaign, err := s.repository.FindbyId(input.Campaign_id)
	if err != nil {
		return CampaignImage{}, err
	}
	if campaign.User_id != input.User.ID {
		return CampaignImage{}, errors.New("not an owner of the campaign")
	}

	if input.Is_primary {
		is_primary = 1
		_, err := s.repository.MarkAllImageAsNonPrimary(input.Campaign_id)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.Campaign_id = input.Campaign_id
	campaignImage.Is_primary = is_primary
	campaignImage.File_name = fileLocation

	newCampaignImage, errImage := s.repository.CreateImage(campaignImage)

	if errImage != nil {
		return newCampaignImage, errImage
	}

	return newCampaignImage, nil

}
