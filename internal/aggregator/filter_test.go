package aggregator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	t.Parallel()
	t.Run("Tokenize simple strings", func(t *testing.T) {
		t.Parallel()
		// Arrange
		input := "Test string ABC"

		// Act
		filter := tokenize(input)

		// Assert
		assert.Empty(t, filter.dataTokens)
		assert.NotEmpty(t, filter.stringTokens)
		assert.Len(t, filter.stringTokens, 3)
		assert.Equal(t, filter.stringTokens[0], "Test")
		assert.Equal(t, filter.stringTokens[1], "string")
		assert.Equal(t, filter.stringTokens[2], "ABC")
	})

	t.Run("Tokenize data", func(t *testing.T) {
		t.Parallel()
		// Arrange
		input := "level:warn customField:true"

		// Act
		filter := tokenize(input)

		// Assert
		assert.Empty(t, filter.stringTokens)
		assert.NotEmpty(t, filter.dataTokens)
		assert.Len(t, filter.dataTokens, 2)
		assert.Equal(t,
			map[string]string{
				"level":       "warn",
				"customField": "true",
			},
			filter.dataTokens,
		)
	})

	t.Run("Tokenize data only when value is not empty", func(t *testing.T) {
		t.Parallel()
		// Arrange
		input := "level:"

		// Act
		filter := tokenize(input)

		// Assert
		assert.Empty(t, filter.stringTokens)
		assert.Empty(t, filter.dataTokens)
	})
}

func TestFilter(t *testing.T) {
	t.Parallel()

	t.Run("Empty filter should match", func(t *testing.T) {
		t.Parallel()
		// Arrange
		log := Log{Original: "unit test msg"}
		input := ""

		// Act
		found := checkLogFilter(log, input)

		// Assert
		assert.True(t, found)
	})

	t.Run("Filter simple text input", func(t *testing.T) {
		t.Parallel()
		t.Run("log matches", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Original: "unit test msg"}
			input := "unit"

			// Act
			found := checkLogFilter(log, input)

			// Assert
			assert.True(t, found)
		})

		t.Run("log does not match", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Original: "unit test msg"}
			input := "shouldFail"

			// Act
			found := checkLogFilter(log, input)

			// Assert
			assert.False(t, found)
		})
	})

	t.Run("Filter data tokens", func(t *testing.T) {
		t.Parallel()
		t.Run("data matches", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Data: map[string]any{"dataA": "works"}}
			input := "dataA:works"

			// Act
			found := checkLogFilter(log, input)

			// Assert
			assert.True(t, found)
		})

		t.Run("data does not match", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Data: map[string]any{"dataA": "works"}}
			input := "dataA:shouldFail"

			// Act
			found := checkLogFilter(log, input)

			// Assert
			assert.False(t, found)
		})

		t.Run("unknown data does not match", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Data: map[string]any{"dataA": "works"}}
			input := "shouldFail:works"

			// Act
			found := checkLogFilter(log, input)

			// Assert
			assert.False(t, found)
		})

		t.Run("partial data token does match", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Data: map[string]any{"dataA": "someLongValue"}}
			input := "dataA:some"

			// Act
			found := checkLogFilter(log, input)

			// Assert
			assert.True(t, found)
		})
	})

	t.Run("Filter multiple data tokens", func(t *testing.T) {
		t.Parallel()
		t.Run("two matching tokens matches", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Data: map[string]any{"dataA": "works", "dataB": "works"}}
			input := "dataA:works dataB:works"

			// Act
			found := checkLogFilter(log, input)

			// Assert
			assert.True(t, found)
		})

		t.Run("one of two matching tokens does not match", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Data: map[string]any{"dataA": "works", "dataB": "works"}}
			input := "dataA:works dataB:shouldFail"

			// Act
			found := checkLogFilter(log, input)

			// Assert
			assert.False(t, found)
		})
	})
}
