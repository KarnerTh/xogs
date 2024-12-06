package parser

import "github.com/KarnerTh/xogs/internal/aggregator"

type Parser = int

const (
	ParserCurl Parser = iota
)

func GetParser(parser Parser) (aggregator.LineParser, error) {
	switch parser {
	case ParserCurl:
		return newCurlParser(), nil
	default:
		return nil, nil
	}
}
