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

func (p jsonParser) Parse(line string) (*aggregator.Log, error) {
	var parsed map[string]any
	_ = json.Unmarshal([]byte(line), &parsed)
	data := parseMapValues(parsed, "")

	return &aggregator.Log{
		Id:   uuid.New().String(),
		Raw:  line,
		Data: data,
	}, nil
}

// possible json types https://pkg.go.dev/encoding/json#Unmarshal
func parseMapValues(value map[string]any, keyPrefix string) map[string]string {
	data := map[string]string{}
	for k, v := range value {
		key := keyPrefix + k

		switch value := v.(type) {
		case string, float64, bool:
			data[key] = parsePrimitiveValues(value)
		case map[string]any:
			nestedKey := fmt.Sprintf("%s.", key)
			nestedData := parseMapValues(value, nestedKey)
			maps.Copy(data, nestedData)
		case []any:
			nestedData := parseArrayValues(value, key)
			maps.Copy(data, nestedData)
		}
	}

	return data
}

func parsePrimitiveValues(v any) string {
	switch value := v.(type) {
	case string:
		return value
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		return ""
	}
}

func parseArrayValues(value []any, key string) map[string]string {
	data := map[string]string{}
	for i, value := range value {
		switch nestedValue := value.(type) {
		case string, float64, bool:
			nestedKey := fmt.Sprintf("%s[%d]", key, i)
			data[nestedKey] = parsePrimitiveValues(nestedValue)
		case map[string]any:
			nestedKey := fmt.Sprintf("%s[%d].", key, i)
			nestedData := parseMapValues(nestedValue, nestedKey)
			maps.Copy(data, nestedData)
		case []any:
			nestedKey := fmt.Sprintf("%s[%d]", key, i)
			nestedData := parseArrayValues(nestedValue, nestedKey)
			maps.Copy(data, nestedData)
		}
	}

	return data
}
