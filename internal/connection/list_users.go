package connection

import (
	"fmt"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/KnoblauchPilze/go-server/pkg/rest"
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
