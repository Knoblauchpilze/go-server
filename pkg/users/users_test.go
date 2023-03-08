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
