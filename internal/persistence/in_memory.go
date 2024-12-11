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

func (r *inMemory) GetAll() []aggregator.Log {
	return r.logs
}
