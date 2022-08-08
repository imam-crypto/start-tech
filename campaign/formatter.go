package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	ImagesUrl        string `json:"images_url"`
	Slug             string `json:"slug"`
	UserID           int    `json:"user_id"`
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

type CampaignDetailFormatter struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	GoalAmount       int      `json:"goal_amount"`
	CurrentAmount    int      `json:"current_amount"`
	ImagesUrl        string   `json:"images_url"`
	UserID           int      `json:"user_id"`
	Slug             string   `json:"slug"`
	Perks            []string `json:"perks"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}

	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.Short_description
	campaignDetailFormatter.GoalAmount = campaign.Goal_amount
	campaignDetailFormatter.CurrentAmount = campaign.Current_amount
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.UserID = campaign.User_id
	campaignDetailFormatter.ImagesUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImagesUrl = campaign.CampaignImages[0].File_name
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	return campaignDetailFormatter

}
