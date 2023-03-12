package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/sirupsen/logrus"
)

var ErrInvalidSignUpData = fmt.Errorf("invalid sign up data")
var ErrSignUpFailed = fmt.Errorf("sign up failed")

func SignUp(in types.UserData) (types.SignUpResponse, error) {
	data, err := json.Marshal(in)
	if err != nil {
		return types.SignUpResponse{}, ErrInvalidSignUpData
	}

	singUpURL := fmt.Sprintf("%s/signup", serverURL)
	resp, err := http.Post(singUpURL, "application/json", bytes.NewReader(data))
	if err != nil {
		logrus.Errorf("Sign up failed: %v", err)
		return types.SignUpResponse{}, ErrSignUpFailed
	}

	var login types.SignUpResponse
	err = rest.GetBodyFromResponseAs(resp, &login)
	if err != nil {
		logrus.Errorf("Sign up request failed: %v", err)
		return types.SignUpResponse{}, ErrSignUpFailed
	}

	return login, nil
}
