package config

type Parser struct {
	Regex  *ParserRegex
	Logfmt *ParserLogfmt
	Json   *ParserJson
}

type ParserRegex struct {
	Values []ParserRegexValue
}

type ParserRegexValue struct {
	Key   string
	Regex string
}

type ParserLogfmt struct{}

type ParserJson struct{}
