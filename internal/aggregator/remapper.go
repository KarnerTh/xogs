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

	if data[config.TargetKey] != "" && !config.OverrideOnConflict {
		return nil
	}

	data[config.TargetKey] = data[inputKey]

	if !config.KeepSource {
		delete(data, inputKey)
	}

	return nil
}
