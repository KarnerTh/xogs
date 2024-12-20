package persistence

import (
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

func (r *inMemory) Get(filter aggregator.Filter) []aggregator.Log {
	logs := []aggregator.Log{}
	for _, log := range r.logs {
		if filter.Matches(log) {
			logs = append(logs, log)
		}
	}
	return logs
}
