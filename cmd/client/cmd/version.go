package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var versionMajor = 0
var versionMinor = 1
var versionPatch = 0

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Client",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("Client v%d.%d.%d", versionMajor, versionMinor, versionPatch)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
