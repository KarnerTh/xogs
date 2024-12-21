package persistence

import (
	"fmt"

	"github.com/KarnerTh/xogs/internal/aggregator"
)

type inMemory struct {
	logs []aggregator.Log
}

func NewInMemory() aggregator.LogRepository {
	return &inMemory{}
}

func (r *inMemory) Add(logEntry aggregator.Log) error {
	r.logs = append(r.logs, logEntry)
	return nil
}

func (r *inMemory) Get(filter aggregator.Filter) ([]aggregator.Log, error) {
	logs := []aggregator.Log{}
	for _, log := range r.logs {
		if filter.Matches(log) {
			logs = append(logs, log)
		}
	}
	return logs, nil
}

func (r *inMemory) GetById(id string) (*aggregator.Log, error) {
	for _, log := range r.logs {
		if log.Id == id {
			return &log, nil
		}
	}
	return nil, fmt.Errorf("Log with id %s not found", id)
}
