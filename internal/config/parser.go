package config

type Pipeline struct {
	Processors []Processor
}

type Processor struct {
	InputKey  string
	Parser    *Parser
	Remapper  *Remapper
	Formatter *Formatter
}

type Remapper struct {
	TargetKey          string
	KeepSource         bool
	OverrideOnConflict bool
}

type Formatter struct {
	Time *TimeFormater
}

type TimeFormater struct {
	From string
	To   string
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
