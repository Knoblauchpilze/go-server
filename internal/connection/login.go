package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/types"
)

func Login(in types.UserData) (types.LoginResponse, error) {
	data, err := json.Marshal(in)
	if err != nil {
		return types.LoginResponse{}, errors.WrapCode(err, errors.ErrPostInvalidData)
	}

	loginURL := fmt.Sprintf("%s/login", serverURL)
	resp, err := http.Post(loginURL, "application/json", bytes.NewReader(data))
	if err != nil {
		return types.LoginResponse{}, errors.WrapCode(err, errors.ErrPostRequestFailed)
	}

	var login types.LoginResponse
	err = rest.GetBodyFromHttpResponseAs(resp, &login)
	if err != nil {
		return types.LoginResponse{}, err
	}

	return login, nil
}
