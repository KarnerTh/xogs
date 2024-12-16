package aggregator

import "strings"

type filter struct {
	stringTokens []string
	dataTokens   map[string]string
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
		stringTokens: stringTokens,
		dataTokens:   dataTokens,
	}
}

func checkLogFilter(log Log, input string) bool {
	filter := tokenize(input)
	match := true

	for _, v := range filter.stringTokens {
		match = match && strings.Contains(log.Original, v)
	}

	for k, v := range filter.dataTokens {
		switch data := log.Data[k].(type) {
		case string:
			match = match && strings.Contains(data, v)
		default:
			match = match && data == v
		}
	}

	return match
}
