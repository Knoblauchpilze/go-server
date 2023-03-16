package users

// https://github.com/stretchr/testify
import (
	"testing"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func checkErrorForCode(err error, code errors.ErrorCode) bool {
	if err == nil {
		return false
	}

	impl, ok := err.(errors.ErrorWithCode)
	if !ok {
		return false
	}

	return impl.Code() == code
}

func TestAddUser_InvalidName(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	_, err := udb.AddUser("", "")
	assert.True(checkErrorForCode(err, errors.ErrInvalidUserName))
}

func TestAddUser_InvalidPassword(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	_, err := udb.AddUser("foo", "")
	assert.True(checkErrorForCode(err, errors.ErrInvalidPassword))
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
	assert.True(checkErrorForCode(err, errors.ErrUserAlreadyExists))
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
	assert.True(checkErrorForCode(err, errors.ErrNoSuchUser))
}

func TestGetUser_WrongID(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	udb.AddUser("foo", "haha")

	wrongID := uuid.New()
	_, err := udb.GetUser(wrongID)
	assert.True(checkErrorForCode(err, errors.ErrNoSuchUser))
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

func TestGetUserFromName_NoUsers(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()

	_, err := udb.GetUserFromName("foo")
	assert.True(checkErrorForCode(err, errors.ErrNoSuchUser))
}

func TestGetUserFromname_WrongName(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	udb.AddUser("foo", "haha")

	_, err := udb.GetUserFromName("")
	assert.True(checkErrorForCode(err, errors.ErrNoSuchUser))

	_, err = udb.GetUserFromName("food")
	assert.True(checkErrorForCode(err, errors.ErrNoSuchUser))
}

func TestGetUserFromName(t *testing.T) {
	assert := assert.New(t)

	udb := NewUserDb()
	id, _ := udb.AddUser("foo", "haha")

	user, err := udb.GetUserFromName("foo")
	assert.Nil(err)

	assert.Equal(user.ID, id)
	assert.Equal(user.Name, "foo")
	assert.Equal(user.Password, "haha")
}
