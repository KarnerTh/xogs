package config

type Parser struct {
	Regex   *ParserRegex
	Logfmt  *ParserLogfmt
	Json    *ParserJson
	Combine *ParserCombine
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

type ParserCombine struct {
	Steps []ParserCombineSteps
}

type ParserCombineSteps struct {
	InputKey string
	Parser   Parser
}
