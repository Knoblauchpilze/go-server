package main

import (
	"os"

	"github.com/KnoblauchPilze/go-server/internal/connection"
	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) == 1 {
		logrus.Errorf("No command specified, aborting")
		return
	}

	command := os.Args[1]

	switch command {
	case "signup":
		signUp()
	default:
		logrus.Errorf("Unknown command \"%s\"", command)
	}
}

func signUp() {
	userData := types.UserData{
		Name:     "toto",
		Password: "123456",
	}

	if len(os.Args) > 2 {
		userData.Name = os.Args[2]
	}

	if len(os.Args) > 3 {
		userData.Password = os.Args[3]
	}

	signUp, err := connection.SignUp(userData)
	if err != nil {
		logrus.Fatalf("Sign up failed: %v", err)
		return
	}

	logrus.Infof("Signed up with id %v!", signUp)
}
