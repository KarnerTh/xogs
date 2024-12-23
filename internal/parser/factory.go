package parser

import (
	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
)

func GetParser(parser config.Parser) aggregator.LineParser {
	switch {
	case parser.Regex != nil:
		return newRegexParser(parser.Regex.Values)
	case parser.Logfmt != nil:
		return newLogfmtParser()
	case parser.Json != nil:
		return newJsonParser()
	case parser.Combine != nil:
		return newCombineParser(parser.Combine.Steps)
	default:
		return newLogfmtParser() // TODO: sane default?
	}
}
