package user

type UserDTOResponse struct {
	Message string  `json:"message"`
	User    UserDTO `json:"user"`
}

type UserLoginDTO struct {
	ID       uint   `json:"id,omitempty"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Token    string `json:"token,omitempty"`
}

func (u *UserLogin) ToDTO() *UserLoginDTO {
	return &UserLoginDTO{
		Email: u.Email,
	}
}
