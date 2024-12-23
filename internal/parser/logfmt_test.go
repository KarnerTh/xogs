package parser

import (
	"testing"

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
		log, err := parser.Parse(line)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Raw)
		assert.Equal(t,
			map[string]string{
				"level": "info",
				"tag":   "0815",
				"env":   "prod",
				"msg":   "test log",
			},
			log.Data,
		)
	})
}
