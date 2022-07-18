package user

type UserFormatter struct {
	ID         int    `json:"id"`
	First_name string `json:"first_name"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		First_name: user.First_name,
		Email:      user.Email,
		Token:      token,
	}
	return formatter
}
