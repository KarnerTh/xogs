package aggregator

import (
	"time"

	"github.com/KarnerTh/xogs/internal/observer"
)

var logNotifier = observer.New[Log]()

type Level int

const (
	LevelUnknown Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

type Log struct {
	Timestamp time.Time
	Level     Level
	Msg       string
	Data      map[string]string
}

type LineParser interface {
	Parse(input Input) (*Log, error)
}

func Aggregate(parser LineParser) observer.Subscriber[Log] {
	inputSubscriber := getInputSubscriber().Subscribe()

	go func() {
		for {
			input := <-inputSubscriber
			// TODO: handle error
			log, _ := parser.Parse(input)

			logNotifier.Publish(*log)
		}

	}()

	return logNotifier
}
