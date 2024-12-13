package aggregator

import (
	"time"

	"github.com/KarnerTh/xogs/internal/observer"
)

var logNotifier = observer.New[Notification]()
var filterNotifier = observer.New[string]()

type Level int

const (
	LevelNone Level = iota
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

type Notification struct {
	NewEntry *Log
	BaseList []Log
}

type LineParser interface {
	Parse(input Input) (*Log, error)
}

type LogRepository interface {
	Add(log Log) error
	GetAll() []Log
}

type Aggregator struct {
	parser LineParser
	repo   LogRepository
	filter string
}

func NewAggregator(lineParser LineParser, logRepository LogRepository) Aggregator {
	return Aggregator{
		parser: lineParser,
		repo:   logRepository,
	}
}

func (a *Aggregator) Aggregate() (observer.Subscriber[Notification], observer.Publisher[string]) {
	inputSubscription := getInputSubscriber().Subscribe()
	filterSubscription := filterNotifier.Subscribe()

	go func() {
		for {
			select {
			case input := <-inputSubscription:
				// TODO: handle error
				log, _ := a.parser.Parse(input)
				a.repo.Add(*log)

				if checkLogFilter(*log, a.filter) {
					logNotifier.Publish(Notification{NewEntry: log})
				}
			case filter := <-filterSubscription:
				a.filter = filter
				if len(filter) == 0 {
					logNotifier.Publish(Notification{BaseList: a.repo.GetAll()})
					continue
				}

				logList := []Log{}
				for _, v := range a.repo.GetAll() {
					if checkLogFilter(v, filter) {
						logList = append(logList, v)
					}
				}
				logNotifier.Publish(Notification{BaseList: logList})
			}
		}
	}()

	return logNotifier, filterNotifier
}
