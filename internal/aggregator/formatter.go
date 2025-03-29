package aggregator

import (
	"fmt"
	"time"

	"github.com/KarnerTh/xogs/internal/config"
)

func format(data map[string]string, inputKey string, config config.Formatter) error {
	if len(inputKey) == 0 {
		return fmt.Errorf("formatter needs input key")
	}

	switch {
	case config.Time != nil:
		value, err := time.Parse(config.Time.From, data[inputKey])
		if err != nil {
			return err
		}

		data[inputKey] = value.Format(config.Time.To)
	}

	return nil
}
