package users

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type UserDb interface {
	AddUser(name string, password string) (uuid.UUID, error)
}

type UserDbImpl struct {
	lock  sync.Mutex
	users map[string]string
	ids   map[string]uuid.UUID
}

var ErrUserAlreadyExists = fmt.Errorf("user already exists")
var ErrInvalidUserName = fmt.Errorf("user name is invalid")
var ErrInvalidPasswordName = fmt.Errorf("password is invalid")
var ErrUserCreationFailure = fmt.Errorf("internal error while creating user")

func NewUserDb() UserDb {
	return &UserDbImpl{
		users: make(map[string]string),
		ids:   make(map[string]uuid.UUID),
	}
}

func (udb *UserDbImpl) AddUser(name string, password string) (uuid.UUID, error) {
	if len(name) == 0 {
		return uuid.UUID{}, ErrInvalidUserName
	}
	if len(password) == 0 {
		return uuid.UUID{}, ErrInvalidPasswordName
	}

	udb.lock.Lock()
	defer udb.lock.Unlock()

	if _, ok := udb.users[name]; ok {
		return uuid.UUID{}, ErrUserAlreadyExists
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, ErrUserCreationFailure
	}

	udb.users[name] = password
	udb.ids[name] = id

	return id, nil
}
