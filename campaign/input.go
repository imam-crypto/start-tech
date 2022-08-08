package campaign

type GetCampaignInput struct {
	ID int `uri:"ID" binding:"required"`
}
