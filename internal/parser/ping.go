package parser

import (
	"regexp"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/google/uuid"
)

type curlParser struct{}

func newPingParser() curlParser {
	return curlParser{}
}

func (p curlParser) Parse(input aggregator.Input) (*aggregator.Log, error) {
	pattern := regexp.MustCompile("time=(?P<time>.*)")
	matches := pattern.FindStringSubmatch(input.Value)

	data := map[string]string{}

	if matches != nil {
		data["time"] = matches[pattern.SubexpIndex("time")]
	}

	return &aggregator.Log{
		Id:   uuid.New().String(),
		Raw:  input.Value,
		Data: data,
	}, nil
}
