package parser

import "github.com/KarnerTh/xogs/internal/aggregator"

type Parser = int

const (
	ParserPing Parser = iota
	ParserLogfmt
)

func GetParser(parser Parser) (aggregator.LineParser, error) {
	switch parser {
	case ParserPing:
		return newPingParser(), nil
	case ParserLogfmt:
		return newLogfmtParser(), nil
	default:
		return nil, nil
	}
}
