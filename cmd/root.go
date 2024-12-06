package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/parser"
	"github.com/KarnerTh/xogs/internal/view"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "xogs",
	Run: func(cmd *cobra.Command, args []string) {
		// f, err := tea.LogToFile("debug.log", "debug")
		// if err != nil {
		// 	fmt.Println("fatal:", err)
		// 	os.Exit(1)
		// }
		// defer f.Close()

		p := view.CreateRootProgram()
		// TODO: error handling
		parser, _ := parser.GetParser(parser.ParserPing)
		logSubscriber := aggregator.Aggregate(parser).Subscribe()

		go func() {
			for {
				log := <-logSubscriber
				p.Send(log)
			}

		}()
		if _, err := p.Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
