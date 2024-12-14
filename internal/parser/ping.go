package parser

import (
	"regexp"

	"github.com/KarnerTh/xogs/internal/aggregator"
)

type curlParser struct{}

func newPingParser() curlParser {
	return curlParser{}
}

func (p curlParser) Parse(input aggregator.Input) (*aggregator.Log, error) {
	pattern := regexp.MustCompile("time=(?P<time>.*)")
	matches := pattern.FindStringSubmatch(input.Value)

	data := map[string]any{"timestamp": input.Timestamp}

	if matches != nil {
		data["time"] = matches[pattern.SubexpIndex("time")]
	}

	return &aggregator.Log{
		Original: input.Value,
		Data:     data,
	}, nil
}
