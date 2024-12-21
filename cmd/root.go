package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
	"github.com/KarnerTh/xogs/internal/parser"
	"github.com/KarnerTh/xogs/internal/persistence"
	"github.com/KarnerTh/xogs/internal/view"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var selectedProfile string
var rootCmd = &cobra.Command{
	Use: "xogs",
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Setup()
		profile, err := c.GetProfileByName(selectedProfile)
		if err != nil {
			profile = &config.DefaultProfile
		}

		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()

		parser := parser.GetParser(profile.Parser)
		logRepo := persistence.NewInMemory()
		agg := aggregator.NewAggregator(parser, logRepo)
		logSubscriber, filterPublisher := agg.Aggregate()
		logSubscription := logSubscriber.Subscribe()

		p := view.CreateRootProgram(profile.DisplayConfig, filterPublisher, logRepo)
		go func() {
			for {
				notification := <-logSubscription
				p.Send(notification)
			}

		}()
		if _, err := p.Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	rootCmd.Flags().StringVar(&selectedProfile, "profile", "", "the profile from the config that should be used")
	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
