package users

// https://github.com/stretchr/testify
import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddUser_InvalidName(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	_, err := udb.AddUser("", "")
	assert.Equal(err, ErrInvalidUserName)
}

func TestAddUser_InvalidPassword(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	_, err := udb.AddUser("foo", "")
	assert.Equal(err, ErrInvalidPasswordName)
}

func TestAddUser(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	id, err := udb.AddUser("foo", "haha")
	assert.Nil(err)

	check, err := uuid.Parse(id.String())
	assert.Nil(err)
	assert.Equal(id, check)
}

func TestAddUser_Duplicated(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	_, err := udb.AddUser("foo", "haha")
	assert.Nil(err)

	_, err = udb.AddUser("foo", "haha")
	assert.Equal(err, ErrUserAlreadyExists)
}

func TestGetUsers(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()

	ids := udb.GetUsers()
	assert.Equal(len(ids), 0)

	id, _ := udb.AddUser("foo", "haha")

	ids = udb.GetUsers()
	assert.Equal(len(ids), 1)
	assert.Equal(ids[0], id)
}

func TestGetUser_NoUsers(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()

	wrongID := uuid.New()
	_, err := udb.GetUser(wrongID)
	assert.Equal(err, ErrNoSuchUser)
}

func TestGetUser_WrongID(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	udb.AddUser("foo", "haha")

	wrongID := uuid.New()
	_, err := udb.GetUser(wrongID)
	assert.Equal(err, ErrNoSuchUser)
}

func TestGetUser(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	id, _ := udb.AddUser("foo", "haha")

	user, err := udb.GetUser(id)
	assert.Nil(err)

	assert.Equal(user.ID, id)
	assert.Equal(user.Name, "foo")
	assert.Equal(user.Password, "haha")
}
