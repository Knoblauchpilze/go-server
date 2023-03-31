package connection

import (
	"fmt"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/users"
	"github.com/google/uuid"
)

func (si *sessionImpl) ListUsers() ([]uuid.UUID, error) {
	var out []uuid.UUID

	listUsersURL := fmt.Sprintf("%s/users", serverURL)

	auth, err := si.generateAuthenticationHeader()
	if err != nil {
		return out, err
	}

	headers := map[string][]string{
		"Authorization": {auth},
	}

	resp, err := performGetRequest(listUsersURL, headers)
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

	listUsersURL := fmt.Sprintf("%s/users/%s", serverURL, id)

	auth, err := si.generateAuthenticationHeader()
	if err != nil {
		return out, err
	}

	headers := map[string][]string{
		"Authorization": {auth},
	}

	resp, err := performGetRequest(listUsersURL, headers)
	if err != nil {
		return out, err
	}

	err = rest.GetBodyFromHttpResponseAs(resp, &out)
	if err != nil {
		return out, errors.WrapCode(err, errors.ErrGetRequestFailed)
	}

	return out, nil
}
