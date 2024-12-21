package config

type ReservedValueKey = string

const (
	ValueKeyRaw ReservedValueKey = "__raw"
)

type DisplayConfig struct {
	Columns []ColumnConfig
}

type ColumnConfig struct {
	Title    string
	Width    float32
	ValueKey string
}

