package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	FindByUserID(UserId int) ([]Campaign, error)
	FindById(input GetCampaignInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
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
func (s *service) FindById(input GetCampaignInput) (Campaign, error) {

	// var campaign Campaign

	campaign, err := s.repository.FindbyId(input.ID)
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
