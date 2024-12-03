package cmd

import (
	"log/slog"
	"os"

	"github.com/KarnerTh/xogs/internal/view"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "xogs",
	Run: func(cmd *cobra.Command, args []string) {
		view.ShowRoot()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(0)
	}
}
