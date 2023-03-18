package connection

import (
	"fmt"

	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/sirupsen/logrus"
)

func (si *sessionImpl) Login(in types.UserData) error {
	var out types.LoginResponse

	url := fmt.Sprintf("%s/login", serverURL)
	resp, err := performPostRequest(url, map[string][]string{}, "application/json", in)
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
