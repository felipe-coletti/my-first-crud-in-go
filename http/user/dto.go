package user

import "my-first-crud-in-go/internal/user"

type UserRequest struct {
	DisplayName string `json:"display_name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID          string `json:"id" binding:"len=36"`
	DisplayName string `json:"display_name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
}

func ToUserResponse(domainInterface user.UserDomainInterface) UserResponse {
	return UserResponse{
		ID: "",
		DisplayName: domainInterface.GetDisplayName(),
		Username: domainInterface.GetUsername(),
		Email: domainInterface.GetEmail(),
	}
}
