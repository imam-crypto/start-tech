package transaction

import (
	"pustaka-api/campaign"
	"pustaka-api/user"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	User_id     int     `json:"user_id"`
	Campaign_id int     `json:"campaign_id"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	Code        string  `json:"code"`
	Campaign    campaign.Campaign
	User        user.User
	PaymentUrl  string
	CreatedAt   time.Time      `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
