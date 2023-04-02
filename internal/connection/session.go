package connection

import (
	"fmt"
	"time"

	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/KnoblauchPilze/go-server/pkg/connection"
	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/KnoblauchPilze/go-server/pkg/users"
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

func (si *sessionImpl) SignUp(in types.UserData) error {
	var out types.SignUpResponse

	signUpUrl := fmt.Sprintf("%s/signup", serverURL)

	rb := connection.NewHttpPostRequestBuilder()
	rb.SetUrl(signUpUrl)
	rb.SetBody("application/json", in)

	req, err := rb.Build()
	if err != nil {
		return err
	}
	resp, err := req.Perform()
	if err != nil {
		return err
	}

	err = rest.GetBodyFromHttpResponseAs(resp, &out)
	if err != nil {
		return err
	}

	si.userId = out.Id
	logrus.Infof("Signed up with id %v", si.userId)

	return nil
}

func (si *sessionImpl) Login(in types.UserData) error {
	var out types.LoginResponse

	loginUrl := fmt.Sprintf("%s/login", serverURL)

	rb := connection.NewHttpPostRequestBuilder()
	rb.SetUrl(loginUrl)
	rb.SetBody("application/json", in)

	req, err := rb.Build()
	if err != nil {
		return err
	}
	resp, err := req.Perform()
	if err != nil {
		return err
	}

	err = rest.GetBodyFromHttpResponseAs(resp, &out)
	if err != nil {
		return err
	}

	si.token = out.Token
	logrus.Infof("Logged in, active token is %+v", si.token)

	return nil
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

func (si *sessionImpl) ListUsers() ([]uuid.UUID, error) {
	var out []uuid.UUID

	listUsersURL := fmt.Sprintf("%s/users", serverURL)

	auth, err := si.generateAuthenticationHeader()
	if err != nil {
		return out, err
	}

	rb := connection.NewHttpGetRequestBuilder()
	rb.SetUrl(listUsersURL)
	rb.SetHeaders(map[string][]string{
		"Authorization": {auth},
	})

	req, err := rb.Build()
	if err != nil {
		return out, err
	}
	resp, err := req.Perform()
	if err != nil {
		return out, err
	}

	err = rest.GetBodyFromHttpResponseAs(resp, &out)
	if err != nil {
		return out, errors.WrapCode(err, errors.ErrGetRequestFailed)
	}

	return out, nil
}

func (si *sessionImpl) ListUser(id uuid.UUID) (users.User, error) {
	var out users.User

	listUserURL := fmt.Sprintf("%s/users/%s", serverURL, id)

	auth, err := si.generateAuthenticationHeader()
	if err != nil {
		return out, err
	}

	rb := connection.NewHttpGetRequestBuilder()
	rb.SetUrl(listUserURL)
	rb.SetHeaders(map[string][]string{
		"Authorization": {auth},
	})

	req, err := rb.Build()
	if err != nil {
		return out, err
	}
	resp, err := req.Perform()
	if err != nil {
		return out, err
	}

	err = rest.GetBodyFromHttpResponseAs(resp, &out)
	if err != nil {
		return out, errors.WrapCode(err, errors.ErrGetRequestFailed)
	}

	return out, nil
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
