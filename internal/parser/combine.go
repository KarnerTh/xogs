package parser

import (
	"maps"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
	"github.com/google/uuid"
)

type combineParser struct {
	steps []config.ParserCombineSteps
}

func newCombineParser(steps []config.ParserCombineSteps) combineParser {
	return combineParser{steps: steps}
}

func (p combineParser) Parse(line string) (*aggregator.Log, error) {
	data := map[string]string{}

	for _, step := range p.steps {
		stepData := line
		if len(step.InputKey) != 0 {
			stepData = data[step.InputKey]
			// remove intermediate data from result data
			delete(data, step.InputKey)
		}

		parser := GetParser(step.Parser)
		// TODO: error handling
		log, _ := parser.Parse(stepData)
		maps.Copy(data, log.Data)
	}

	return &aggregator.Log{
		Id:   uuid.New().String(),
		Raw:  line,
		Data: data,
	}, nil
}
