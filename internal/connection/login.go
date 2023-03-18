package connection

import (
	"fmt"

	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/types"
)

func Login(in types.UserData) (types.LoginResponse, error) {
	var out types.LoginResponse

	url := fmt.Sprintf("%s/login", serverURL)
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
