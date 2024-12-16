package aggregator

type Level int

const (
	LevelNone Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

type Log struct {
	Data     map[string]any
	Original string
}

func (l Log) GetStringData(key string) string {
	value, ok := l.Data[key]
	if ok {
		return value.(string)
	} else {
		return ""
	}
}
