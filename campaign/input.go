package campaign

import "pustaka-api/user"

type GetCampaignInput struct {
	ID int `uri:"ID" binding:"required"`
}

type CreateCampaignInput struct {
	Name              string `json:"name" binding:"required"`
	Short_description string `json:"short_description" binding:"required"`
	Description       string `json:"description" binding:"required"`
	Goal_amount       int    `json:"goal_amount" binding:"required"`
	Perks             string `json:"perks" binding:"required"`
	User              user.User
}
