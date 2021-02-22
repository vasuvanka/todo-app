package models

//Login - login object
type Login struct {
	Email    string `vd:"email($)" json:"email"`
	Password string `vd:"len($)>1" json:"password"`
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
	Email    string `vd:"email($); msg:'Valid email required'" json:"email"`
	Password string `vd:"len($)>3; msg:'Password required'" json:"password"`
	Name     string `vd:"len($)>1; msg:'Name required'" json:"name"`
}

//JwtUserPayload - jwt user payload
type JwtUserPayload struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	Name string `json:"name"`
}

//ShareTodo - share todo
type ShareTodo struct {
	Email string `vd:"email($); msg:'Valid email required'" json:"email"`
}
