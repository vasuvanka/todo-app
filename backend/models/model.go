package models

//Login - login object
type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

//LoginResponse - login response
type LoginResponse struct {
	Token string `json:"token"`
}

//Response - error response
type Response struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message"`
}

//Signup - login object
type Signup struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

//JwtUserPayload - jwt user payload
type JwtUserPayload struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	Name string `json:"name"`
}

//ShareTodo - share todo
type ShareTodo struct {
	Email string `json:"email"`
}
