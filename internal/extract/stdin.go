package extract

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/KarnerTh/xogs/internal/observer"
)

type InputData struct {
	Timestamp time.Time
	Value     string
}

var inputNotifier = observer.New[InputData]()

func GetInputSubscriber() observer.Subscriber[InputData] {
	if !hasStdinContent() {
		return nil
	}

	go readFromStdin()
	return inputNotifier
}

func readFromStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		inputNotifier.Publish(InputData{Timestamp: time.Now(), Value: line})
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
