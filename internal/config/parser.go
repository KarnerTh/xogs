package config

type Pipeline struct {
	Processors []Processor
}

type Processor struct {
	InputKey string
	Parser   *Parser
	Remapper *Remapper
}

type Remapper struct {
	TargetKey          string
	KeepSource         bool
	OverrideOnConflict bool
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
