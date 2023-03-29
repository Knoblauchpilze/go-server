package list

import (
	"github.com/KnoblauchPilze/go-server/internal/connection"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "List all registered users",
	Run:   usersCmdBody,
}

func usersCmdBody(cmd *cobra.Command, args []string) {
	sess := connection.NewSession()

	token, err := fetchUserToken()
	if err != nil {
		logrus.Fatalf("Failed to list users: %+v", err)
		return
	}

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
