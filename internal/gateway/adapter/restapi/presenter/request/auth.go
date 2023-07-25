package request

type SignUp struct {
	Name     string `json:"name" validation:"required"`
	Email    string `json:"email" validation:"required;email"`
	Password string `json:"password" validation:"required;"`
}

type Login struct {
	Email    string `json:"email" validation:"required;email"`
	Password string `json:"password" validation:"required;"`
}
