package models

import "time"

type User struct {
	Id            int       `json:"id"`
	Username      *string   `json:"username"`
	Email         *string   `json:"email" validate:"email,required"`
	Password      *string   `json:"password" validate:"required"`
	Profile_image *string   `json:"profile_image"`
	Phone_number  *string   `json:"phone_number"`
	Created_at    time.Time `json:"created_at"` // DEFAULT NOW()
	Updated_at    time.Time `json:"updated_at"`
}
