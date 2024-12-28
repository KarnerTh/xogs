package aggregator

import (
	"bufio"
	"log"
	"os"

	"github.com/KarnerTh/xogs/internal/observer"
)

var logNotifier = observer.New[Notification]()
var filterNotifier = observer.New[string]()

type Notification struct {
	NewEntry *Log
	BaseList []Log
}

type LineParser interface {
	Parse(line string) (*Log, error)
}

type LogRepository interface {
	Add(log Log) error
	Get(filter Filter) ([]Log, error)
	GetById(id string) (*Log, error)
}

type Aggregator struct {
	pipeline LineParser
	repo     LogRepository
	filter   Filter
}

func NewAggregator(pipeline LineParser, logRepository LogRepository) Aggregator {
	return Aggregator{
		pipeline: pipeline,
		repo:     logRepository,
	}
}

func (a *Aggregator) Aggregate() (observer.Subscriber[Notification], observer.Publisher[string]) {
	inputSubscription := getStdinSubscriber().Subscribe()
	filterSubscription := filterNotifier.Subscribe()

	go func() {
		for {
			select {
			case input := <-inputSubscription:
				// TODO: handle error
				log, _ := a.pipeline.Parse(input.Value)
				a.repo.Add(*log)

				if a.filter.Matches(*log) {
					logNotifier.Publish(Notification{NewEntry: log})
				}
			case filter := <-filterSubscription:
				a.filter = parseFilter(filter)
				logList, err := a.repo.Get(a.filter)
				if err != nil {
					log.Printf(err.Error())
				}

				logNotifier.Publish(Notification{BaseList: logList})
			}
		}
	}()

	return logNotifier, filterNotifier
}

func (a *Aggregator) AggregateFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// TODO: handle error
		log, _ := a.pipeline.Parse(line)
		a.repo.Add(*log)
	}

	logList, err := a.repo.Get(a.filter)
	if err != nil {
		log.Printf(err.Error())
	}

	logNotifier.Publish(Notification{BaseList: logList})
	return nil
}
