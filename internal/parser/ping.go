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
	timePattern := regexp.MustCompile("time=(?P<time>.*)")
	matches := timePattern.FindStringSubmatch(input.Value)

	data := map[string]string{}
	if matches != nil {
		data["time"] = matches[timePattern.SubexpIndex("time")]
	}

	return &aggregator.Log{
		Level:     aggregator.LevelNone,
		Timestamp: input.Timestamp,
		Msg:       input.Value,
		Data:      data,
	}, nil
}
