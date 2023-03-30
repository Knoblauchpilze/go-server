package list

import (
	"time"

	"github.com/KnoblauchPilze/go-server/internal/connection"
	"github.com/KnoblauchPilze/go-server/pkg/auth"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "List all registered users",
	Args:  cobra.ExactArgs(2),
	Run:   usersCmdBody,
}

func usersCmdBody(cmd *cobra.Command, args []string) {
	id, err := uuid.Parse(args[0])
	if err != nil {
		logrus.Errorf("invalid user id provided (%v)", err)
		return
	}

	token := auth.Token{
		User:       id,
		Value:      args[1],
		Expiration: time.Now().Add(1 * time.Minute),
	}

	sess := connection.NewSession()
	if err := sess.Authenticate(token); err != nil {
		logrus.Fatalf("Failed to list users: %+v", err)
		return
	}

	data, err := sess.ListUsers()
	if err != nil {
		logrus.Fatalf("Failed to list users: %+v", err)
		return
	}

	logrus.Infof("Users: %+v", data)
}
