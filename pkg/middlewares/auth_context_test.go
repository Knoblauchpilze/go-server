package middlewares

import (
	"fmt"
	"testing"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestParseAuthenticationHeader_InvalidHeader(t *testing.T) {
	assert := assert.New(t)

	genErr := errors.New("invalid authentication header")

	authData := ""
	_, err := parseAuthenticationHeader(authData)
	assert.Equal(err, genErr)

	authData = "toto tata"
	_, err = parseAuthenticationHeader(authData)
	assert.Equal(err, genErr)
}

func TestParseAuthenticationHeader_InvalidAuthenticationScheme(t *testing.T) {
	assert := assert.New(t)

	genErr := errors.New("invalid authentication header")

	authData := "wrong-scheme haha hihi"
	_, err := parseAuthenticationHeader(authData)
	assert.Equal(err, genErr)
}

func TestParseAuthenticationHeader_InvalidPropSyntax(t *testing.T) {
	assert := assert.New(t)

	id := "74b02b94-ca64-11ed-ba48-18c04d0e6a41"

	authData := "bearer haha hihi"
	_, err := parseAuthenticationHeader(authData)
	expectedErr := fmt.Sprintf("%s (cause: Ill-formed prop in Authorization header: \"haha\")", genErrMsg)
	assert.Equal(err.Error(), expectedErr)

	authData = "bearer haha= hihi"
	_, err = parseAuthenticationHeader(authData)
	expectedErr = fmt.Sprintf("%s (cause: Ill-formed prop in Authorization header: \"haha=\")", genErrMsg)
	assert.Equal(err.Error(), expectedErr)

	authData = "bearer =haha hihi"
	_, err = parseAuthenticationHeader(authData)
	expectedErr = fmt.Sprintf("%s (cause: Ill-formed prop in Authorization header: \"=haha\")", genErrMsg)
	assert.Equal(err.Error(), expectedErr)

	authData = "bearer haha=22 hihi=45"
	_, err = parseAuthenticationHeader(authData)
	expectedErr = fmt.Sprintf("%s (cause: no \"%s\" key in Authorization header)", genErrMsg, authenticationUserKey)
	assert.Equal(err.Error(), expectedErr)

	authData = "bearer user=22 hihi=45"
	_, err = parseAuthenticationHeader(authData)
	expectedErr = fmt.Sprintf("%s (cause: failed to parse user id \"22\" (cause: invalid UUID length: 2))", genErrMsg)
	assert.Equal(err.Error(), expectedErr)

	authData = fmt.Sprintf("bearer user=%s hihi=45", id)
	_, err = parseAuthenticationHeader(authData)
	expectedErr = fmt.Sprintf("%s (cause: no \"%s\" key in Authorization header)", genErrMsg, authenticationTokenKey)
	assert.Equal(err.Error(), expectedErr)
}

func TestParseAuthenticationHeader_Success(t *testing.T) {
	assert := assert.New(t)

	id, _ := uuid.Parse("74b02b94-ca64-11ed-ba48-18c04d0e6a41")
	val := "dummy-token"

	authData := fmt.Sprintf("bearer user=%s token=%s", id, val)
	token, err := parseAuthenticationHeader(authData)
	assert.Nil(err)
	assert.Equal(token.User, id)
	assert.Equal(token.Value, val)
}
