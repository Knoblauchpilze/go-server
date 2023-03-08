package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken_InvalidPassword(t *testing.T) {
	assert := assert.New(t)

	auth := NewAuth()
	id, _ := uuid.NewUUID()

	_, err := auth.GenerateToken(id, "")
	assert.Equal(err, ErrInvalidPassword)
}

func TestGenerateToken(t *testing.T) {
	assert := assert.New(t)

	auth := NewAuth()
	id, _ := uuid.NewUUID()

	token, err := auth.GenerateToken(id, "foo")

	assert.Nil(err)
	assert.Equal(token.ID, id)
	assert.GreaterOrEqual(len(token.Value), 1)
	assert.True(time.Now().Before(token.Expiration))

	_, err = auth.GenerateToken(id, "foo")
	assert.Equal(err, ErrTokenAlreadyExists)
}

func TestGetToken(t *testing.T) {
	assert := assert.New(t)

	auth := NewAuth()
	id, _ := uuid.NewUUID()

	check, _ := auth.GenerateToken(id, "foo")
	token, err := auth.GetToken(id)

	assert.Nil(err)
	assert.Equal(token.ID, check.ID)
	assert.Equal(token.Value, check.Value)
	assert.Equal(token.Expiration, check.Expiration)
}

func TestGetToken_InvalidID(t *testing.T) {
	assert := assert.New(t)

	auth := NewAuth()
	id, _ := uuid.NewUUID()
	id2, _ := uuid.NewUUID()

	auth.GenerateToken(id, "foo")
	_, err := auth.GetToken(id2)
	assert.Equal(err, ErrNoSuchToken)
}
