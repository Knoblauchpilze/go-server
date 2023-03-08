package auth

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Token struct {
	ID         uuid.UUID
	Value      string
	Expiration time.Time
}

type Auth interface {
	GenerateToken(id uuid.UUID, password string) (Token, error)
	GetToken(id uuid.UUID) (Token, error)
}

type AuthImpl struct {
	lock   sync.Mutex
	tokens map[uuid.UUID]Token
}

var ErrInvalidPassword = fmt.Errorf("password is invalid")
var ErrTokenAlreadyExists = fmt.Errorf("token already exists")
var ErrTokenCreationFailure = fmt.Errorf("internal error while creating token")
var ErrNoSuchToken = fmt.Errorf("no such token")

var TokenDefaultExpirationTime = 1 * time.Minute

func NewAuth() Auth {
	return &AuthImpl{
		tokens: make(map[uuid.UUID]Token),
	}
}

func (auth *AuthImpl) GenerateToken(id uuid.UUID, password string) (Token, error) {
	if len(password) == 0 {
		return Token{}, ErrInvalidPassword
	}

	auth.lock.Lock()
	defer auth.lock.Unlock()

	if _, ok := auth.tokens[id]; ok {
		return Token{}, ErrTokenAlreadyExists
	}

	token := Token{
		ID:         id,
		Expiration: time.Now().Add(TokenDefaultExpirationTime),
		Value:      "dummy-token",
	}

	auth.tokens[id] = token

	return token, nil
}

func (auth *AuthImpl) GetToken(id uuid.UUID) (Token, error) {
	auth.lock.Lock()
	defer auth.lock.Unlock()

	token, ok := auth.tokens[id]
	if !ok {
		return Token{}, ErrNoSuchToken
	}

	return token, nil
}
