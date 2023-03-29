package list

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "List info about a specific user",
	Run:   userCmdBody,
}

func userCmdBody(cmd *cobra.Command, args []string) {
	logrus.Warnf("should list user")
}
