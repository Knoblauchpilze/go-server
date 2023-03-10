package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/KnoblauchPilze/go-server/pkg/types"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var serverURL = "http://localhost:3000"

var errSignUpFailed = fmt.Errorf("sign up failed")

func main() {
	id, err := signUp()
	if err != nil {
		logrus.Fatalf("Sign up failed: %v", err)
		return
	}

	logrus.Infof("Signed up with id %v!", id)
}

func signUp() (uuid.UUID, error) {

	user := "toto"
	if len(os.Args) > 1 {
		user = os.Args[1]
	}

	password := "123456"
	if len(os.Args) > 2 {
		password = os.Args[2]
	}

	in := types.UserData{
		Name:     user,
		Password: password,
	}
	data, _ := json.Marshal(in)

	singUpURL := fmt.Sprintf("%s/signup", serverURL)
	resp, err := http.Post(singUpURL, "application/json", bytes.NewReader(data))
	if err != nil {
		logrus.Errorf("Sign up failed: %v", err)
		return uuid.UUID{}, errSignUpFailed
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Failed to read sign up response: %v", err)
		return uuid.UUID{}, errSignUpFailed
	}

	var login types.SignUpResponse
	if err = json.Unmarshal(data, &login); err != nil {
		logrus.Errorf("Failed to parse sign up response: %v", err)
		return uuid.UUID{}, errSignUpFailed
	}

	return login.ID, nil
}
