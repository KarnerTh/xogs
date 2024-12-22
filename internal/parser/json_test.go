package parser

import (
	"testing"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/stretchr/testify/assert"
)

func TestParserJons(t *testing.T) {
	t.Parallel()
	parser := newJsonParser()

	t.Run("simple primitive values", func(t *testing.T) {
		t.Parallel()
		// Arrange
		line := `{"level": 7, "prod": true, "msg": "test log"}`

		// Act
		log, err := parser.Parse(aggregator.Input{Value: line})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Raw)
		assert.Equal(t, "7", log.Data["level"])
		assert.Equal(t, "true", log.Data["prod"])
		assert.Equal(t, "test log", log.Data["msg"])
	})

	t.Run("nested object", func(t *testing.T) {
		t.Parallel()
		// Arrange
		line := `{"some": {"value": "works"}}`

		// Act
		log, err := parser.Parse(aggregator.Input{Value: line})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Raw)
		assert.Equal(t, "works", log.Data["some.value"])
	})

	t.Run("deeply nested object", func(t *testing.T) {
		t.Parallel()
		// Arrange
		line := `{"some": {"value": {"that": {"is": {"deeply": {"nested": "works"}}}}}}`

		// Act
		log, err := parser.Parse(aggregator.Input{Value: line})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Raw)
		assert.Equal(t, "works", log.Data["some.value.that.is.deeply.nested"])
	})
}
