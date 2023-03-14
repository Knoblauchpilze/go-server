package main

import (
	"os"

	"github.com/KnoblauchPilze/go-server/internal/connection"
	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/sirupsen/logrus"
)

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
