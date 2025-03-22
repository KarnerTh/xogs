package aggregator

import (
	"fmt"

	"github.com/KarnerTh/xogs/internal/config"
)

func remap(data map[string]string, inputKey string, config config.Remapper) error {
	if len(inputKey) == 0 {
		return fmt.Errorf("remapper needs input key")
	}

	if data == nil {
		return nil
	}

	value := data[inputKey]
	data[config.TargetKey] = value

	if !config.KeepSource {
		delete(data, inputKey)
	}

	return nil
}
