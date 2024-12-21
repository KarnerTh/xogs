package aggregator

import "strings"

type Filter struct {
	StringTokens []string
	DataTokens   map[string]string
}

func parseFilter(input string) Filter {
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

	return Filter{
		StringTokens: stringTokens,
		DataTokens:   dataTokens,
	}
}

func (filter Filter) Matches(log Log) bool {
	match := true

	for _, v := range filter.StringTokens {
		match = match && strings.Contains(log.Raw, v)
	}

	for k, v := range filter.DataTokens {
		match = match && strings.Contains(log.Data[k], v)
	}

	return match
}
