package parser

import (
	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
)

func NewPipeline(pipeline config.Pipeline) aggregator.LineParser {
	return newPipeline(pipeline.Processors)
}

func getParser(parser config.Parser) aggregator.LineParser {
	switch {
	case parser.Regex != nil:
		return newRegexParser(parser.Regex.Values)
	case parser.Logfmt != nil:
		return newLogfmtParser()
	case parser.Json != nil:
		return newJsonParser()
	default:
		return newLogfmtParser() // TODO: sane default?
	}
}
