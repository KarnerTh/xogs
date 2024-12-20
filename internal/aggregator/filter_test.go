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
		filter := parseFilter(input)

		// Assert
		assert.Empty(t, filter.DataTokens)
		assert.NotEmpty(t, filter.StringTokens)
		assert.Len(t, filter.StringTokens, 3)
		assert.Equal(t, filter.StringTokens[0], "Test")
		assert.Equal(t, filter.StringTokens[1], "string")
		assert.Equal(t, filter.StringTokens[2], "ABC")
	})

	t.Run("Tokenize data", func(t *testing.T) {
		t.Parallel()
		// Arrange
		input := "level:warn customField:true"

		// Act
		filter := parseFilter(input)

		// Assert
		assert.Empty(t, filter.StringTokens)
		assert.NotEmpty(t, filter.DataTokens)
		assert.Len(t, filter.DataTokens, 2)
		assert.Equal(t,
			map[string]string{
				"level":       "warn",
				"customField": "true",
			},
			filter.DataTokens,
		)
	})

	t.Run("Tokenize data only when value is not empty", func(t *testing.T) {
		t.Parallel()
		// Arrange
		input := "level:"

		// Act
		filter := parseFilter(input)

		// Assert
		assert.Empty(t, filter.StringTokens)
		assert.Empty(t, filter.DataTokens)
	})
}

func TestFilter(t *testing.T) {
	t.Parallel()

	t.Run("Empty filter should match", func(t *testing.T) {
		t.Parallel()
		// Arrange
		log := Log{Original: "unit test msg"}
		filter := Filter{}

		// Act
		found := filter.Matches(log)

		// Assert
		assert.True(t, found)
	})

	t.Run("Filter simple text input", func(t *testing.T) {
		t.Parallel()
		t.Run("log matches", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Original: "unit test msg"}
			filter := Filter{StringTokens: []string{"unit"}}

			// Act
			found := filter.Matches(log)

			// Assert
			assert.True(t, found)
		})

		t.Run("log does not match", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Original: "unit test msg"}
			filter := Filter{StringTokens: []string{"shouldFail"}}

			// Act
			found := filter.Matches(log)

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
			filter := Filter{DataTokens: map[string]string{"dataA": "works"}}

			// Act
			found := filter.Matches(log)

			// Assert
			assert.True(t, found)
		})

		t.Run("data does not match", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Data: map[string]any{"dataA": "works"}}
			filter := Filter{DataTokens: map[string]string{"dataA": "shouldFail"}}

			// Act
			found := filter.Matches(log)

			// Assert
			assert.False(t, found)
		})

		t.Run("unknown data does not match", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Data: map[string]any{"dataA": "works"}}
			filter := Filter{DataTokens: map[string]string{"shouldFail": "works"}}

			// Act
			found := filter.Matches(log)

			// Assert
			assert.False(t, found)
		})

		t.Run("partial data token does match", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Data: map[string]any{"dataA": "someLongValue"}}
			filter := Filter{DataTokens: map[string]string{"dataA": "some"}}

			// Act
			found := filter.Matches(log)

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
			filter := Filter{DataTokens: map[string]string{
				"dataA": "works",
				"dataB": "works",
			}}

			// Act
			found := filter.Matches(log)

			// Assert
			assert.True(t, found)
		})

		t.Run("one of two matching tokens does not match", func(t *testing.T) {
			t.Parallel()
			// Arrange
			log := Log{Data: map[string]any{"dataA": "works", "dataB": "works"}}
			filter := Filter{DataTokens: map[string]string{
				"dataA": "works",
				"dataB": "shouldFail",
			}}

			// Act
			found := filter.Matches(log)

			// Assert
			assert.False(t, found)
		})
	})
}
