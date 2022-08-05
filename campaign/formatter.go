package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	ImagesUrl        string `json:"images_url"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.Short_description
	campaignFormatter.GoalAmount = campaign.Goal_amount
	campaignFormatter.CurrentAmount = campaign.Current_amount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImagesUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImagesUrl = campaign.CampaignImages[0].File_name
	}
	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	var campaignsFormatter []CampaignFormatter

	if len(campaigns) == 0 {
		return []CampaignFormatter{}
	}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}
