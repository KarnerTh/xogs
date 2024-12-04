package cmd

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"github.com/KarnerTh/xogs/internal/view"
	tea "github.com/charmbracelet/bubbletea"
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
		go readLog(p)
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

func readLog(p *tea.Program) {
	if !hasStdinContent() {
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		p.Send(view.InputTest{Msg: line})
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error in scanning: ", err)
		os.Exit(1)
	}
}

func hasStdinContent() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("error reading stdin: ", err)
		return false
	}

	// source https://stackoverflow.com/a/26567513
	return (fi.Mode() & os.ModeCharDevice) == 0
}
