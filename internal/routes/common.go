package routes

import (
	"net/http"

	"github.com/KnoblauchPilze/go-server/pkg/rest"
	"github.com/KnoblauchPilze/go-server/pkg/types"
)

func getUserDataFromRequest(r *http.Request) (types.UserData, error) {
	var ud types.UserData
	err := rest.GetBodyFromHttpRequestAs(r, &ud)
	return ud, err
}
