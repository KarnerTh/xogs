package parser

import (
	"regexp"
	"strings"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/google/uuid"
)

type logfmtParser struct{}

func newLogfmtParser() logfmtParser {
	return logfmtParser{}
}

func (p logfmtParser) Parse(line string) (*aggregator.Log, error) {
	pattern := regexp.MustCompile(`(?P<key>[a-z]+)=(?P<value>(?:"(.*)")|(?:(?:([^\s]+))))`)
	matches := pattern.FindAllStringSubmatch(line, -1)

	data := map[string]string{}
	for i := range matches {
		key := matches[i][pattern.SubexpIndex("key")]
		value := matches[i][pattern.SubexpIndex("value")]
		data[key] = strings.ReplaceAll(value, `"`, "")
	}

	return &aggregator.Log{
		Id:   uuid.New().String(),
		Raw:  line,
		Data: data,
	}, nil
}
