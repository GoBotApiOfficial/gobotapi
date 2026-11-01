package logger

type LogLevel int

const (
	Silent LogLevel = iota
	Error
	Warn
	Info
	Debug
)
