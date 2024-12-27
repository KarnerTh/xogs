package config

type Pipeline struct {
	Processors []Processor
}

type Processor struct {
	Parser   Parser
	InputKey string
}

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
