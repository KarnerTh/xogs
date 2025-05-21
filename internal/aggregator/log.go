package aggregator

type LogData = map[string]string

type Log struct {
	Id   string
	Data LogData
	Raw  string
}
