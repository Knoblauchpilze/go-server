package connection

import (
	"fmt"
	"time"

	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type sessionImpl struct {
	userId uuid.UUID
	token  auth.Token
}

func NewSession() Session {
	return &sessionImpl{}
}

func (si *sessionImpl) Authenticate(token auth.Token) error {
	if len(token.Value) == 0 {
		return errors.NewCode(errors.ErrNotLoggedIn)
	}
	if time.Now().After(token.Expiration) {
		logrus.Infof("now: %v, token: %v", time.Now(), token.Expiration)
		return errors.NewCode(errors.ErrAuthenticationExpired)
	}

	si.token = token

	return nil
}

func (si *sessionImpl) generateAuthenticationHeader() (string, error) {
	if len(si.token.Value) == 0 {
		return "", errors.NewCode(errors.ErrNotLoggedIn)
	}
	if time.Now().After(si.token.Expiration) {
		return "", errors.NewCode(errors.ErrAuthenticationExpired)
	}

	auth := fmt.Sprintf("bearer user=%v token=%v", si.token.User, si.token.Value)

	return auth, nil
}
