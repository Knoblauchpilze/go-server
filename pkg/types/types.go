package types

import "github.com/google/uuid"

type UserData struct {
	Name     string
	Password string
}

type SignUpResponse struct {
	ID uuid.UUID
}
