package parser

import (
	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
)

type parserFactory struct{}

func NewParserFactory() aggregator.ParserFactory {
	return parserFactory{}
}

func (f parserFactory) GetParser(parser config.Parser) aggregator.LineParser {
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
