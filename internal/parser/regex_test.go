package parser

import (
	"testing"

	"github.com/KarnerTh/xogs/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestParserRegex(t *testing.T) {
	t.Parallel()
	t.Run("ping example", func(t *testing.T) {
		t.Parallel()
		parser := newRegexParser([]config.ParserRegexValue{
			{Key: "time", Regex: `time=(.*)`},
			{Key: "ttl", Regex: `ttl=(\d*)`},
		})

		t.Run("match", func(t *testing.T) {
			t.Parallel()
			// Arrange
			line := "64 bytes from 8.8.8.8: icmp_seq=2 ttl=118 time=17.5 ms"

			// Act
			log, err := parser.Parse(line)

			// Assert
			assert.Nil(t, err)
			assert.Equal(t, line, log.Raw)
			assert.Equal(t, "17.5 ms", log.Data["time"])
			assert.Equal(t, "118", log.Data["ttl"])
		})

		t.Run("no match (header line)", func(t *testing.T) {
			t.Parallel()
			// Arrange
			line := "PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data."

			// Act
			log, err := parser.Parse(line)

			// Assert
			assert.Nil(t, err)
			assert.Equal(t, line, log.Raw)
			assert.Empty(t, log.Data)
		})
	})
}
