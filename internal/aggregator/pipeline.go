package aggregator

import (
	"log/slog"
	"maps"

	"github.com/KarnerTh/xogs/internal/config"
	"github.com/google/uuid"
)

type pipeline struct {
	processors    []config.Processor
	parserFactory ParserFactory
}

type ParserFactory interface {
	GetParser(parser config.Parser) LineParser
}

func NewPipeline(processors []config.Processor, parserFactory ParserFactory) pipeline {
	return pipeline{processors: processors, parserFactory: parserFactory}
}

func (p pipeline) Parse(line string) (*Log, error) {
	data := map[string]string{}

	for _, processor := range p.processors {
		switch {
		case processor.Parser != nil:
			stepData := line
			if len(processor.InputKey) != 0 {
				stepData = data[processor.InputKey]
				// remove intermediate data from result data
				delete(data, processor.InputKey)
			}
			parser := p.parserFactory.GetParser(*processor.Parser)
			// TODO: error handling
			log, _ := parser.Parse(stepData)
			maps.Copy(data, log.Data)
		case processor.Remapper != nil:
			err := remap(data, processor.InputKey, *processor.Remapper)
			if err != nil {
				slog.Error("Remapping failed", slog.Any("error", err))
			}
		case processor.Formatter != nil:
			err := format(data, processor.InputKey, *processor.Formatter)
			if err != nil {
				slog.Error("Formatting failed", slog.Any("error", err))
			}
		}
	}

	return &Log{
		Id:   uuid.New().String(),
		Raw:  line,
		Data: data,
	}, nil
}
