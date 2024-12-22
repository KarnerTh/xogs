package parser

import (
	"regexp"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
	"github.com/google/uuid"
)

type regexParser struct {
	values []config.ParserRegexValue
}

func newRegexParser(values []config.ParserRegexValue) regexParser {
	return regexParser{
		values: values,
	}
}

func (p regexParser) Parse(input aggregator.Input) (*aggregator.Log, error) {
	data := map[string]string{}

	for _, v := range p.values {
		pattern, err := regexp.Compile(v.Regex)
		if err != nil {
			return nil, err
		}

		matches := pattern.FindStringSubmatch(input.Value)
		if len(matches) > 0 {
			data[v.Key] = matches[1]
		}
	}

	return &aggregator.Log{
		Id:   uuid.New().String(),
		Raw:  input.Value,
		Data: data,
	}, nil
}
