package dto

type CreateUserRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginUserRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseDTO struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
