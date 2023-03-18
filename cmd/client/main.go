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
	case "login":
		login()
	case "list":
		list()
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

	data, err := connection.SignUp(userData)
	if err != nil {
		logrus.Fatalf("Sign up failed: %+v", err)
		return
	}

	logrus.Infof("Signed up with id %+v!", data)
}

func login() {
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

	data, err := connection.Login(userData)
	if err != nil {
		logrus.Fatalf("Login failed: %+v", err)
		return
	}

	logrus.Infof("Logged in and received token %+v!", data)
}

func list() {
	if len(os.Args) < 3 {
		logrus.Fatalf("Nothing to list")
		return
	}

	item := os.Args[2]
	switch item {
	case "users":
		listUsers()
	default:
		logrus.Fatalf("Unrecognized item to list: \"%v\"", item)
	}
}

func listUsers() {
	userData := types.UserData{
		Name:     "toto",
		Password: "123456",
	}

	data, err := connection.ListUsers(userData)
	if err != nil {
		logrus.Fatalf("Failed to list users: %+v", err)
		return
	}

	logrus.Infof("Users: %+v", data)
}
