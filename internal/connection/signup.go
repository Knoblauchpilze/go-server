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

func SignUp(in types.UserData) (types.SignUpResponse, error) {
	data, err := json.Marshal(in)
	if err != nil {
		return types.SignUpResponse{}, errors.WrapCode(err, errors.ErrInvalidSignUpData)
	}

	singUpURL := fmt.Sprintf("%s/signup", serverURL)
	resp, err := http.Post(singUpURL, "application/json", bytes.NewReader(data))
	if err != nil {
		return types.SignUpResponse{}, errors.WrapCode(err, errors.ErrPostRequestFailed)
	}

	var login types.SignUpResponse
	err = rest.GetBodyFromHttpResponseAs(resp, &login)
	if err != nil {
		return types.SignUpResponse{}, err
	}

	return login, nil
}
