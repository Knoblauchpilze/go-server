package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Client allows to interact with the server (duh)",
	Long:  "Makes submitting commands to the server easy and fast",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("failed to execute root command (%v)", err)
		os.Exit(1)
	}
}
