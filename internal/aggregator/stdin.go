package aggregator

import (
	"bufio"
	"fmt"
	"os"

	"github.com/KarnerTh/xogs/internal/observer"
)

var stdinNotifier = observer.New[Input]()

type Input struct {
	Value string
}

func getStdinSubscriber() observer.Subscriber[Input] {
	if hasStdinContent() {
		go readFromStdin()
	}

	return stdinNotifier
}

func readFromStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		stdinNotifier.Publish(Input{Value: line})
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
