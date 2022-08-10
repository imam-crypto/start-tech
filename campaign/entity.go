package campaign

import (
	"pustaka-api/user"
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	ID                int            `json:"id" gorm:"primaryKey"`
	User_id           int            `json:"user_id"`
	Name              string         `json:"name"`
	Short_description string         `json:"short_description"`
	Description       string         `json:"description"`
	Slug              string         `json:"slug"`
	Perks             string         `json:"perks"`
	Backer            int            `json:"backer"`
	Goal_amount       int            `json:"goal_amount"`
	Current_amount    int            `json:"current_amount"`
	CreatedAt         time.Time      `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at"`
	CampaignImages    []CampaignImage
	User              user.User
}

type CampaignImage struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Campaign_id int            `json:"campaign_id"`
	File_name   string         `json:"file_name"`
	Is_primary  string         `json:"is_primary"`
	CreatedAt   time.Time      `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
