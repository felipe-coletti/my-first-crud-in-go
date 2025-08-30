package request

type UserRequest struct {
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
