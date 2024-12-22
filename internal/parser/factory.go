package parser

import (
	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
)

func GetParser(profile config.Profile) aggregator.LineParser {
	switch {
	case profile.Parser.Regex != nil:
		return newRegexParser(profile.Parser.Regex.Values)
	case profile.Parser.Logfmt != nil:
		return newLogfmtParser()
	case profile.Parser.Json != nil:
		return newJsonParser()
	default:
		return newLogfmtParser() // TODO: sane default?
	}
}
