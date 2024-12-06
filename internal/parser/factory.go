package parser

import "github.com/KarnerTh/xogs/internal/aggregator"

type Parser = int

const (
	ParserPing Parser = iota
)

func GetParser(parser Parser) (aggregator.LineParser, error) {
	switch parser {
	case ParserPing:
		return newPingParser(), nil
	default:
		return nil, nil
	}
}
