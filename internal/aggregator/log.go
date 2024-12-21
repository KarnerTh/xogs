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
	Id   string
	Data map[string]string
	Raw  string
}
