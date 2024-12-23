package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserJson(t *testing.T) {
	t.Parallel()
	parser := newJsonParser()

	t.Run("simple primitive values", func(t *testing.T) {
		t.Parallel()
		// Arrange
		line := `{"level": 7, "prod": true, "msg": "test log"}`

		// Act
		log, err := parser.Parse(line)

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
		log, err := parser.Parse(line)

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
		log, err := parser.Parse(line)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Raw)
		assert.Equal(t, "works", log.Data["some.value.that.is.deeply.nested"])
	})

	t.Run("array with objects", func(t *testing.T) {
		t.Parallel()
		// Arrange
		line := `{"values": [{"item": "one"},{"item": "two"}]}`

		// Act
		log, err := parser.Parse(line)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Raw)
		assert.Equal(t, "one", log.Data["values[0].item"])
		assert.Equal(t, "two", log.Data["values[1].item"])
	})

	t.Run("array with strings", func(t *testing.T) {
		t.Parallel()
		// Arrange
		line := `{"values": ["one","two"]}`

		// Act
		log, err := parser.Parse(line)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Raw)
		assert.Equal(t, "one", log.Data["values[0]"])
		assert.Equal(t, "two", log.Data["values[1]"])
	})

	t.Run("array of array with strings", func(t *testing.T) {
		t.Parallel()
		// Arrange
		line := `{"values": [["one", "two"],["three", "four"]]}`

		// Act
		log, err := parser.Parse(line)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Raw)
		assert.Equal(t, "one", log.Data["values[0][0]"])
		assert.Equal(t, "two", log.Data["values[0][1]"])
		assert.Equal(t, "three", log.Data["values[1][0]"])
		assert.Equal(t, "four", log.Data["values[1][1]"])
	})

	t.Run("array of array with strings and objects", func(t *testing.T) {
		t.Parallel()
		// Arrange
		line := `{"values": [["one", "two"],[{"item": "three"}, {"item": "four"}]]}`

		// Act
		log, err := parser.Parse(line)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Raw)
		assert.Equal(t, "one", log.Data["values[0][0]"])
		assert.Equal(t, "two", log.Data["values[0][1]"])
		assert.Equal(t, "three", log.Data["values[1][0].item"])
		assert.Equal(t, "four", log.Data["values[1][1].item"])
	})
}
