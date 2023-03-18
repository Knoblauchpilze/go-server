package connection

import (
	"fmt"

	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/types"
)

func SignUp(in types.UserData) (types.SignUpResponse, error) {
	var out types.SignUpResponse

	url := fmt.Sprintf("%s/signup", serverURL)
	resp, err := performPostRequest(url, map[string][]string{}, "application/json", in)
	if err != nil {
		return out, err
	}

	err = rest.GetBodyFromHttpResponseAs(resp, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}
