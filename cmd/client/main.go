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

	sess := connection.NewSession()

	command := os.Args[1]
	switch command {
	case "signup":
		signUp(sess)
	case "login":
		login(sess)
	case "list":
		listCLI(sess)
	default:
		logrus.Errorf("Unknown command \"%s\"", command)
	}
}

func signUp(sess connection.Session) {
	ud := fetchCredentialsOrDefault()
	if err := sess.SignUp(ud); err != nil {
		logrus.Fatalf("Sign up failed: %+v", err)
	}
}

func login(sess connection.Session) {
	ud := fetchCredentialsOrDefault()
	if err := sess.Login(ud); err != nil {
		logrus.Fatalf("Login failed: %+v", err)
	}
}

func fetchCredentialsOrDefault() types.UserData {
	ud := types.UserData{
		Name:     "toto",
		Password: "123456",
	}

	if len(os.Args) > 2 {
		ud.Name = os.Args[2]
	}

	if len(os.Args) > 3 {
		ud.Password = os.Args[3]
	}

	return ud
}

func listCLI(sess connection.Session) {
	signUp(sess)
	login(sess)

	if len(os.Args) < 3 {
		logrus.Fatalf("Nothing to list")
		return
	}

	item := os.Args[2]
	switch item {
	case "users":
		listUsers(sess)
	default:
		logrus.Fatalf("Unrecognized item to list: \"%v\"", item)
	}
}

func listUsers(sess connection.Session) {
	data, err := sess.ListUsers()
	if err != nil {
		logrus.Fatalf("Failed to list users: %+v", err)
		return
	}

	logrus.Infof("Users: %+v", data)
}
