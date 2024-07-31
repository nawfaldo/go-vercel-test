package types

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type RegisterUserPayload struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserPayload struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
