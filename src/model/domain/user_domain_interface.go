package domain

type UserDomainInterface interface {
	GetDisplayName() string
	GetUsername() string
	GetEmail() string
	GetPassword() string
	EncryptPassword()
}

type userDomain struct {
	displayName string
	username    string
	email       string
	password    string
}

func NewUserDomain(
	displayName string,
	username string,
	email string,
	password string,
) UserDomainInterface {
	return &userDomain{
		displayName,
		username,
		email,
		password,
	}
}