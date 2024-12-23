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

func (p regexParser) Parse(line string) (*aggregator.Log, error) {
	data := map[string]string{}

	for _, v := range p.values {
		pattern, err := regexp.Compile(v.Regex)
		if err != nil {
			return nil, err
		}

		matches := pattern.FindStringSubmatch(line)
		if len(matches) > 0 {
			data[v.Key] = matches[1]
		}
	}

	return &aggregator.Log{
		Id:   uuid.New().String(),
		Raw:  line,
		Data: data,
	}, nil
}
