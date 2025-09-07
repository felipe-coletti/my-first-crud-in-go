package response

type UserResponse struct {
	ID          string `json:"id" binding:"len=36"`
	DisplayName string `json:"display_name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
}
