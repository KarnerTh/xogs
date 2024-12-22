package config

type Parser struct {
	Regex *ParserRegex
}

type ParserRegex struct {
	Values []ParserRegexValue
}

type ParserRegexValue struct {
	Key   string
	Regex string
}
