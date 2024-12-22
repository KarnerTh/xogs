package parser

import (
	"encoding/json"
	"strconv"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/google/uuid"
)

type jsonParser struct{}

func newJsonParser() jsonParser {
	return jsonParser{}
}

func (p jsonParser) Parse(input aggregator.Input) (*aggregator.Log, error) {
	data := map[string]string{}

	var parsed map[string]any
	json.Unmarshal([]byte(input.Value), &parsed)

	// possible json types https://pkg.go.dev/encoding/json#Unmarshal
	for key, value := range parsed {
		switch v := value.(type) {
		case string:
			data[key] = v
		case float64:
			data[key] = strconv.FormatFloat(v, 'f', -1, 64)
		case bool:
			data[key] = strconv.FormatBool(v)
		}
	}

	return &aggregator.Log{
		Id:   uuid.New().String(),
		Raw:  input.Value,
		Data: data,
	}, nil
}
