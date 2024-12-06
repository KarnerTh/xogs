package parser

import (
	"github.com/KarnerTh/xogs/internal/aggregator"
)

type curlParser struct{}

func newCurlParser() curlParser {
	return curlParser{}
}

func (p curlParser) Parse(input aggregator.Input) (*aggregator.Log, error) {
	return &aggregator.Log{
		Timestamp: input.Timestamp,
		Msg:       input.Value,
	}, nil
}
