package connection

import (
	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/google/uuid"
)

var serverURL = "http://localhost:3000"

type Session interface {
	SignUp(in types.UserData) error
	Login(in types.UserData) error

	Authenticate(token auth.Token) error

	ListUsers() ([]uuid.UUID, error)
}
