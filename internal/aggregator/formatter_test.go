package aggregator

import (
	"testing"

	"github.com/KarnerTh/xogs/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestFormatter(t *testing.T) {
	t.Parallel()
	t.Run("time formatter", func(t *testing.T) {
		t.Parallel()
		// Arrange
		data := map[string]string{"time": "2024-12-17T14:00:29.490Z"}

		// Act
		err := format(data, "time", config.Formatter{
			Time: &config.TimeFormater{
				From: "2006-01-02T15:04:05.999999999Z07:00",
				To:   "15:04:05",
			},
		})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "14:00:29", data["time"])
	})
	t.Run("time formatter - invalid layout", func(t *testing.T) {
		t.Parallel()
		// Arrange
		data := map[string]string{"time": "2024-12-17T14:00:29.490Z"}

		// Act
		err := format(data, "time", config.Formatter{
			Time: &config.TimeFormater{
				From: "XYZ",
				To:   "15:04:05",
			},
		})

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, "2024-12-17T14:00:29.490Z", data["time"])
	})
}
