package connection

import (
	"fmt"

	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/sirupsen/logrus"
)

func (si *sessionImpl) SignUp(in types.UserData) error {
	var out types.SignUpResponse

	url := fmt.Sprintf("%s/signup", serverURL)
	resp, err := performPostRequest(url, map[string][]string{}, "application/json", in)
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
