package cmd

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
	"github.com/KarnerTh/xogs/internal/parser"
	"github.com/KarnerTh/xogs/internal/persistence"
	"github.com/KarnerTh/xogs/internal/view"
	"github.com/spf13/cobra"
)

var selectedProfile string
var rootCmd = &cobra.Command{
	Use:  "xogs",
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Setup()
		if err := c.Validate(config.ValidationData{SelectedProfile: selectedProfile}); err != nil {
			log.Println(err)
			os.Exit(1)
		}

		profile, _ := c.GetProfileByName(selectedProfile)
		// f, err := tea.LogToFile("debug.log", "debug")
		// if err != nil {
		// 	fmt.Println("fatal:", err)
		// 	os.Exit(1)
		// }
		// defer f.Close()

		pipeline := aggregator.NewPipeline(profile.Pipeline.Processors, parser.NewParserFactory())
		logRepo := persistence.NewInMemory()
		agg := aggregator.NewAggregator(pipeline, logRepo)
		logSubscriber, filterPublisher := agg.Aggregate()
		logSubscription := logSubscriber.Subscribe()

		p := view.CreateRootProgram(profile.DisplayConfig, filterPublisher, logRepo)
		go func() {
			for {
				p.Send(<-logSubscription)
			}
		}()

		// handle file argument
		if len(args) != 0 {
			err := agg.AggregateFile(args[0])
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}

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
