package types

import (
	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/google/uuid"
)

type UserData struct {
	Name     string
	Password string
}

type SignUpResponse struct {
	ID uuid.UUID
}

type LoginResponse struct {
	Token auth.Token
}
