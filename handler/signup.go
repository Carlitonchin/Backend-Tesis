package handler

type signup_req struct {
	Email string `json:"email" binding:"required, email"`
	Pass  string `json:"pass" binding:"required,gte=6,lte=30"`
}
