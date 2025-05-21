package aggregator

import (
	"testing"

	"github.com/KarnerTh/xogs/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestRemapper(t *testing.T) {
	t.Parallel()
	t.Run("replace source", func(t *testing.T) {
		t.Parallel()
		// Arrange
		data := map[string]string{"before": "testValue"}

		// Act
		err := remap(data, "before", config.Remapper{TargetKey: "after"})

		// Assert
		assert.Nil(t, err)
		assert.Contains(t, data, "after")
		assert.NotContains(t, data, "before")
		assert.Equal(t, "testValue", data["after"])
	})

	t.Run("keep source", func(t *testing.T) {
		t.Parallel()
		// Arrange
		data := map[string]string{"before": "testValue"}

		// Act
		err := remap(data, "before", config.Remapper{TargetKey: "after", KeepSource: true})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "testValue", data["before"])
		assert.Equal(t, "testValue", data["after"])
	})

	t.Run("do not override on conflict", func(t *testing.T) {
		t.Parallel()
		// Arrange
		data := map[string]string{"before": "testValue", "after": "already_set"}

		// Act
		err := remap(data, "before", config.Remapper{TargetKey: "after"})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "already_set", data["after"])
	})

	t.Run("override on conflict", func(t *testing.T) {
		t.Parallel()
		// Arrange
		data := map[string]string{"before": "testValue", "after": "already_set"}

		// Act
		err := remap(data, "before", config.Remapper{TargetKey: "after", OverrideOnConflict: true})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "testValue", data["after"])
	})
}
