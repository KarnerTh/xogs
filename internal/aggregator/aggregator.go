package aggregator

import (
	"github.com/KarnerTh/xogs/internal/observer"
)

var logNotifier = observer.New[Notification]()
var filterNotifier = observer.New[string]()

type Notification struct {
	NewEntry *Log
	BaseList []Log
}

type LineParser interface {
	Parse(input Input) (*Log, error)
}

type LogRepository interface {
	Add(log Log) error
	Get(filter Filter) []Log
}

type Aggregator struct {
	parser LineParser
	repo   LogRepository
	filter Filter
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

				if a.filter.Matches(*log) {
					logNotifier.Publish(Notification{NewEntry: log})
				}
			case filter := <-filterSubscription:
				a.filter = parseFilter(filter)
				logList := a.repo.Get(a.filter)
				logNotifier.Publish(Notification{BaseList: logList})
			}
		}
	}()

	return logNotifier, filterNotifier
}
