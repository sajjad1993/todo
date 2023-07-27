package request

type SignUp struct {
	Name     string `json:"name" validation:"required" binding:"required"`
	Email    string `json:"email" validation:"required" binding:"required,email"`
	Password string `json:"password" validation:"required;" binding:"required,min=6"`
}

type Login struct {
	Email    string `json:"email" validation:"required;email" binding:"required,email"`
	Password string `json:"password" validation:"required;" binding:"required,min=6"`
}
