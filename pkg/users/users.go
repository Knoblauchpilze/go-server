package users

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Password string
}

type UserDb interface {
	AddUser(name string, password string) (uuid.UUID, error)
	GetUser(id uuid.UUID) (User, error)
	GetUserFromName(name string) (User, error)
	GetUsers() []uuid.UUID
}

type UserDbImpl struct {
	lock  sync.Mutex
	users map[string]string
	ids   map[uuid.UUID]string
	names map[string]uuid.UUID
}

var ErrUserAlreadyExists = fmt.Errorf("user already exists")
var ErrInvalidUserName = fmt.Errorf("user name is invalid")
var ErrInvalidPassword = fmt.Errorf("password is invalid")
var ErrUserCreationFailure = fmt.Errorf("internal error while creating user")
var ErrNoSuchUser = fmt.Errorf("no such user")

func NewUserDb() UserDb {
	return &UserDbImpl{
		users: make(map[string]string),
		ids:   make(map[uuid.UUID]string),
		names: make(map[string]uuid.UUID),
	}
}

func (udb *UserDbImpl) AddUser(name string, password string) (uuid.UUID, error) {
	if len(name) == 0 {
		return uuid.UUID{}, ErrInvalidUserName
	}
	if len(password) == 0 {
		return uuid.UUID{}, ErrInvalidPassword
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
	udb.ids[id] = name
	udb.names[name] = id

	return id, nil
}

func (udb *UserDbImpl) GetUser(id uuid.UUID) (User, error) {
	udb.lock.Lock()
	defer udb.lock.Unlock()

	user := User{
		ID: id,
	}

	name, ok := udb.ids[id]
	if !ok {
		return user, ErrNoSuchUser
	}

	password, ok := udb.users[name]
	if !ok {
		return user, ErrNoSuchUser
	}

	user.Name = name
	user.Password = password

	return user, nil
}

func (udb *UserDbImpl) GetUserFromName(name string) (User, error) {
	user := User{
		Name: name,
	}

	if len(name) == 0 {
		return user, ErrNoSuchUser
	}

	udb.lock.Lock()
	defer udb.lock.Unlock()

	id, ok := udb.names[name]
	if !ok {
		return user, ErrNoSuchUser
	}

	password, ok := udb.users[name]
	if !ok {
		return user, ErrNoSuchUser
	}

	user.ID = id
	user.Password = password

	return user, nil
}

func (udb *UserDbImpl) GetUsers() []uuid.UUID {
	udb.lock.Lock()
	defer udb.lock.Unlock()

	users := make([]uuid.UUID, 0, len(udb.ids))
	for id := range udb.ids {
		users = append(users, id)
	}

	return users
}
