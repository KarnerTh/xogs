package aggregator

import "github.com/KarnerTh/xogs/internal/config"

func remap(data map[string]string, inputKey string, config config.Remapper) {
	if data == nil {
		return
	}

	value := data[inputKey]
	delete(data, inputKey)
	data[config.TargetKey] = value
}
