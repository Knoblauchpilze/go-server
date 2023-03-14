package main

import (
	"os"

	"github.com/KnoblauchPilze/go-server/internal/connection"
	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/sirupsen/logrus"
)

// TODO: How to interpret the server response?
// Maybe as `GetBodyFromHttpResponseAs` is in the same package
// as the `ServerResponse` we can first parse the response as
// this and then convert the details (or whichever field we
// say contain the answer) to this.

func main() {
	userData := types.UserData{
		Name:     "toto",
		Password: "123456",
	}

	if len(os.Args) > 1 {
		userData.Name = os.Args[1]
	}

	if len(os.Args) > 2 {
		userData.Password = os.Args[2]
	}

	signUp, err := connection.SignUp(userData)
	if err != nil {
		logrus.Fatalf("Sign up failed: %v", err)
		return
	}

	logrus.Infof("Signed up with id %v!", signUp)
}
