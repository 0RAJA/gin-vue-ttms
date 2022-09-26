package request

type SendEmailCode struct {
	Email string `json:"email" binding:"email"`
}
