package view

import (
	"github.com/felipe-coletti/my-first-crud-in-go/src/model/domain"
	"github.com/felipe-coletti/my-first-crud-in-go/src/model/response"
)

func ToUserResponse(domainInterface domain.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID: "",
		DisplayName: domainInterface.GetDisplayName(),
		Username: domainInterface.GetUsername(),
		Email: domainInterface.GetEmail(),
	}
}