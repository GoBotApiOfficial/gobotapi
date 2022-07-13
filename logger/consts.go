package logger

const (
	reset      = "\033[0m"
	redBold    = "\033[1;31m"
	greenBold  = "\033[1;32m"
	yellowBold = "\033[1;33m"
	blueBold   = "\033[1;34m"
)

type LogLevel int

const (
	Silent LogLevel = iota
	Error
	Warn
	Info
	Debug
)
