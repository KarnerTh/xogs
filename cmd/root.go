package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/KarnerTh/xogs/internal/extract"
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
		inputSubscriber := extract.GetInputSubscriber().Subscribe()

		go func() {
			for {
				input := <-inputSubscriber
				p.Send(view.InputTest{Msg: input.Value})
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
