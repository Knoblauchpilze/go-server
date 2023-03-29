package list

import (
	"os"
	"time"

	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/KnoblauchPilze/go-server/pkg/errors"
	"github.com/google/uuid"
)

func fetchUserToken() (auth.Token, error) {
	var token auth.Token

	if len(os.Args) < 5 {
		return token, errors.New("no user token provided")
	}

	if id, err := uuid.Parse(os.Args[3]); err != nil {
		return token, errors.Wrap(err, "invalid user id provided")
	} else {
		token.User = id
	}

	token.Value = os.Args[4]
	token.Expiration = time.Now().Add(1 * time.Minute)

	return token, nil
}
