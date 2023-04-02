package list

import (
	"github.com/KnoblauchPilze/go-server/internal/session"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "List all registered users",
	Run:   usersCmdBody,
}

func init() {
	usersCmd.Flags().StringVar(&userArg, "user", "", "the id of the user")
	usersCmd.Flags().StringVar(&tokenArg, "token", "", "the token of the user")
}

func usersCmdBody(cmd *cobra.Command, args []string) {
	token, err := buildTokenFromFlags()
	if err != nil {
		logrus.Errorf("invalid parameters to list users (%v)", err)
		return
	}

	sess := session.NewManager()
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
