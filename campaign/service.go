package campaign

type Service interface {
	FindByUserID(UserId int) ([]Campaign, error)
	FindById(input GetCampaignInput) (Campaign, error)
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
