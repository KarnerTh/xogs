package parser

import (
	"encoding/json"
	"fmt"
	"maps"
	"strconv"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/google/uuid"
)

type jsonParser struct{}

func newJsonParser() jsonParser {
	return jsonParser{}
}

func (p jsonParser) Parse(input aggregator.Input) (*aggregator.Log, error) {
	var parsed map[string]any
	json.Unmarshal([]byte(input.Value), &parsed)
	data := parseMapValues(parsed, "")

	return &aggregator.Log{
		Id:   uuid.New().String(),
		Raw:  input.Value,
		Data: data,
	}, nil
}

// possible json types https://pkg.go.dev/encoding/json#Unmarshal
func parseMapValues(value map[string]any, keyPrefix string) map[string]string {
	data := map[string]string{}
	for k, v := range value {
		key := k
		if len(keyPrefix) != 0 {
			key = fmt.Sprintf("%s.%s", keyPrefix, key)
		}

		switch value := v.(type) {
		case string:
			data[key] = value
		case float64:
			data[key] = strconv.FormatFloat(value, 'f', -1, 64)
		case bool:
			data[key] = strconv.FormatBool(value)
		case map[string]any:
			nestedData := parseMapValues(value, key)
			maps.Copy(data, nestedData)
		}
	}

	return data
}
