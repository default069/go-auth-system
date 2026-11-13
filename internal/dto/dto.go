package dto

type SignUpRequest struct{
	Email 	 string 	`json:"email" binding:"required,email"`
	Name  	 string		`json:"name"  binding:"required,min=2,max=100"`
	Password string		`json:"password" binding:"required,min=8"`
}