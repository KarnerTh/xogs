package parser

import (
	"maps"

	"github.com/KarnerTh/xogs/internal/aggregator"
	"github.com/KarnerTh/xogs/internal/config"
	"github.com/google/uuid"
)

type pipeline struct {
	processors []config.Processor
}

// a pipeline combines one or more processors
func newPipeline(processors []config.Processor) pipeline {
	return pipeline{processors: processors}
}

func (p pipeline) Parse(line string) (*aggregator.Log, error) {
	data := map[string]string{}

	for _, processor := range p.processors {
		stepData := line
		if len(processor.InputKey) != 0 {
			stepData = data[processor.InputKey]
			// remove intermediate data from result data
			delete(data, processor.InputKey)
		}

		parser := getParser(processor.Parser)
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
