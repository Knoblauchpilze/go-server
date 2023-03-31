package auth

import (
	"sync"
	"time"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/google/uuid"
)

type Token struct {
	User       uuid.UUID
	Value      string
	Expiration time.Time
}

type Auth interface {
	GenerateToken(user uuid.UUID, password string) (Token, error)
	GetToken(user uuid.UUID) (Token, error)
}

type AuthImpl struct {
	lock   sync.Mutex
	tokens map[uuid.UUID]Token
}

var TokenDefaultExpirationTime = 10 * time.Minute

func NewAuth() Auth {
	return &AuthImpl{
		tokens: make(map[uuid.UUID]Token),
	}
}

func (auth *AuthImpl) GenerateToken(user uuid.UUID, password string) (Token, error) {
	if len(password) == 0 {
		return Token{}, errors.NewCode(errors.ErrInvalidPassword)
	}

	auth.lock.Lock()
	defer auth.lock.Unlock()

	if _, ok := auth.tokens[user]; ok {
		return Token{}, errors.NewCode(errors.ErrTokenAlreadyExists)
	}

	token := Token{
		User:       user,
		Expiration: time.Now().Add(TokenDefaultExpirationTime),
		Value:      "dummy-token",
	}

	auth.tokens[user] = token

	return token, nil
}

func (auth *AuthImpl) GetToken(user uuid.UUID) (Token, error) {
	auth.lock.Lock()
	defer auth.lock.Unlock()

	token, ok := auth.tokens[user]
	if !ok {
		return Token{}, errors.NewCode(errors.ErrNoSuchToken)
	}

	return token, nil
}
