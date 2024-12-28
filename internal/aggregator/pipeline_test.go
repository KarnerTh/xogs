package aggregator_test

import (
	"fmt"
	"testing"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
	"github.com/KarnerTh/xogs/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	t.Parallel()
	t.Run("json with prefix", func(t *testing.T) {
		t.Parallel()
		parser := aggregator.NewPipeline([]config.Processor{
			{
				Parser: &config.Parser{
					Regex: &config.ParserRegex{
						Values: []config.ParserRegexValue{
							{Key: "service", Regex: `\[(.*)\]`},
							{Key: "log", Regex: `\[.*\]\s(.*)`},
						},
					},
				},
			},
			{
				InputKey: "log",
				Parser:   &config.Parser{Json: &config.ParserJson{}},
			},
		}, parser.NewParserFactory())

		// Arrange
		line := `[test-service] {"level":"info", "msg": "works"}`

		// Act
		log, err := parser.Parse(line)
		t.Log(fmt.Sprintf("data: %+v", log.Data))

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Raw)
		assert.Equal(t, "test-service", log.Data["service"])
		assert.Equal(t, "info", log.Data["level"])
		assert.Equal(t, "works", log.Data["msg"])

		// make sure intermediate data from pipeline is not persisted
		assert.NotContains(t, log.Data, "log")
	})
}
