package list

import (
	"github.com/KnoblauchPilze/go-server/internal/connection"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "List info about a specific user",
	Run:   userCmdBody,
}

func init() {
	userCmd.Flags().StringVar(&userArg, "user", "", "the id of the user")
	userCmd.Flags().StringVar(&tokenArg, "token", "", "the token of the user")
}

func userCmdBody(cmd *cobra.Command, args []string) {
	token, err := buildTokenFromFlags()
	if err != nil {
		logrus.Errorf("invalid parameters to list users (%v)", err)
		return
	}

	sess := connection.NewSession()
	if err := sess.Authenticate(token); err != nil {
		logrus.Fatalf("Failed to list users: %+v", err)
		return
	}

	data, err := sess.ListUser(token.User)
	if err != nil {
		logrus.Fatalf("Failed to list users: %+v", err)
		return
	}

	logrus.Infof("Users: %+v", data)
}
