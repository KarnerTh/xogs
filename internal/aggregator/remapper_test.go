package aggregator

import (
	"testing"

	"github.com/KarnerTh/xogs/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestRemapper(t *testing.T) {
	// Arrange
	data := map[string]string{"before": "testValue"}

	// Act
	remap(data, "before", config.Remapper{TargetKey: "after"})

	// Assert
	assert.Contains(t, data, "after")
	assert.NotContains(t, data, "before")
	assert.Equal(t, "testValue", data["after"])
}
