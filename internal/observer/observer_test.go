package observer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObserver(t *testing.T) {
	// Arrange
	o := New[string]()
	sub := o.Subscribe()

	// Act
	go o.Publish("works")

	// Assert
	result := <-sub
	assert.Equal(t, "works", result)
}
