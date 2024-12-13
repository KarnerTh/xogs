package aggregator

import "strings"

type filter struct {
	level        Level
	stringTokens []string
	dataTokens   map[string]string
}

func (f filter) isEmpty() bool {
	return f.level == LevelNone &&
		len(f.stringTokens) == 0 &&
		len(f.dataTokens) == 0
}

func tokenize(input string) filter {
	tokens := strings.Split(input, " ")
	stringTokens := []string{}
	dataTokens := map[string]string{}

	for _, t := range tokens {
		if strings.Contains(t, ":") {
			parts := strings.Split(t, ":")
			key, value := parts[0], parts[1]
			if len(value) == 0 {
				continue
			}

			dataTokens[key] = value
		} else {
			if len(t) == 0 {
				continue
			}
			stringTokens = append(stringTokens, t)
		}
	}

	return filter{
		level:        LevelNone, // TODO: parse level
		stringTokens: stringTokens,
		dataTokens:   dataTokens,
	}
}

func checkLogFilter(log Log, input string) bool {
	filter := tokenize(input)

	for _, v := range filter.stringTokens {
		if strings.Contains(log.Msg, v) {
			return true
		}
	}

	for k, v := range filter.dataTokens {
		if log.Data[k] == v {
			return true
		}
	}

	return filter.isEmpty()
}
