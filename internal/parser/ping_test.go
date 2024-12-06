package parser

import (
	"testing"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/stretchr/testify/assert"
)

func TestParserPing(t *testing.T) {
	t.Run("header", func(t *testing.T) {
		// Arrange
		parser := newPingParser()
		line := "PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data."

		// Act
		log, err := parser.Parse(aggregator.Input{Value: line})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Msg)
	})

	t.Run("content", func(t *testing.T) {
		// Arrange
		parser := newPingParser()
		line := "64 bytes from 8.8.8.8: icmp_seq=2 ttl=118 time=17.5 ms"

		// Act
		log, err := parser.Parse(aggregator.Input{Value: line})

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, line, log.Msg)
		assert.Equal(t, "17.5 ms", log.Data["time"])
	})
}
