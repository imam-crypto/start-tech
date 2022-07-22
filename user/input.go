package user

type RegisterUserInput struct {
	First_name string `json:"first_name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmaiInput struct {
	Email string `json:"email" binding:"required,email"`
}
