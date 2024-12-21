package parser

import "github.com/KarnerTh/xogs/internal/aggregator"

type Parser = string

const (
	ParserPing   Parser = "ping"
	ParserLogfmt Parser = "logfmt"
)

func GetParser(parser Parser) aggregator.LineParser {
	switch parser {
	case ParserPing:
		return newPingParser()
	case ParserLogfmt:
		return newLogfmtParser()
	default:
		// TODO: maybe different default?
		return newLogfmtParser()
	}
}
