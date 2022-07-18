package user

import "time"

type User struct {
	ID         int
	First_name string
	Email      string
	Password   string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}
