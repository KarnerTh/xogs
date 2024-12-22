package parser

import (
	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
)

func GetParser(profile config.Profile) aggregator.LineParser {
	if profile.Parser.Regex != nil {
		return newRegexParser(profile.Parser.Regex.Values)
	}

	return newLogfmtParser()
}
