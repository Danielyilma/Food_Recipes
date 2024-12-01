package models

type LoginForm struct {
	Email    *string `json:"email" validate:"email,required"`
	Password *string `json:"password" validate:"required"`
}

// type LoginRequest struct {
// 	Arg struct {
// 		Email    string `json:"email" validate:"email,required"`
// 		Password string `json:"password" validate:"required"`
// 	} `json:"arg"`
// }
