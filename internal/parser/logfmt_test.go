package parser

import (
	"testing"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/stretchr/testify/assert"
)

func TestParserLogfmt(t *testing.T) {
	t.Parallel()
	t.Run("info log with custom tags", func(t *testing.T) {
		t.Parallel()
		// Arrange
		parser := newLogfmtParser()
		line := `level=info tag=0815 env=prod msg="test log"`

		// Act
		log, err := parser.Parse(aggregator.Input{Value: line})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Original)
		assert.Equal(t,
			map[string]any{
				"level": "info",
				"tag":   "0815",
				"env":   "prod",
				"msg":   "test log",
			},
			log.Data,
		)
	})
}
