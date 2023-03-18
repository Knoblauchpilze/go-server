package users

// https://github.com/stretchr/testify
import (
	"testing"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddUser_InvalidName(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	_, err := udb.AddUser("", "")
	assert.True(errors.IsErrorWithCode(err, errors.ErrInvalidUserName))
}

func TestAddUser_InvalidPassword(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	_, err := udb.AddUser("foo", "")
	assert.True(errors.IsErrorWithCode(err, errors.ErrInvalidPassword))
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
	assert.True(errors.IsErrorWithCode(err, errors.ErrUserAlreadyExists))
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

	wrongId := uuid.New()
	_, err := udb.GetUser(wrongId)
	assert.True(errors.IsErrorWithCode(err, errors.ErrNoSuchUser))
}

func TestGetUser_WrongId(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	udb.AddUser("foo", "haha")

	wrongId := uuid.New()
	_, err := udb.GetUser(wrongId)
	assert.True(errors.IsErrorWithCode(err, errors.ErrNoSuchUser))
}

func TestGetUser(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	id, _ := udb.AddUser("foo", "haha")

	user, err := udb.GetUser(id)
	assert.Nil(err)

	assert.Equal(user.Id, id)
	assert.Equal(user.Name, "foo")
	assert.Equal(user.Password, "haha")
}

func TestGetUserFromName_NoUsers(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()

	_, err := udb.GetUserFromName("foo")
	assert.True(errors.IsErrorWithCode(err, errors.ErrNoSuchUser))
}

func TestGetUserFromname_WrongName(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	udb.AddUser("foo", "haha")

	_, err := udb.GetUserFromName("")
	assert.True(errors.IsErrorWithCode(err, errors.ErrNoSuchUser))

	_, err = udb.GetUserFromName("food")
	assert.True(errors.IsErrorWithCode(err, errors.ErrNoSuchUser))
}

func TestGetUserFromName(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	id, _ := udb.AddUser("foo", "haha")

	user, err := udb.GetUserFromName("foo")
	assert.Nil(err)

	assert.Equal(user.Id, id)
	assert.Equal(user.Name, "foo")
	assert.Equal(user.Password, "haha")
}
