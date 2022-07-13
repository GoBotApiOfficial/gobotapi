package gobotapi

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi/logger"
	"reflect"
)

func NewLogger(writer Writer, config logger.Config) Logger {
	colorfulPrint := "%s[%s]%s %%s %%s"
	closeColor, debugColor, infoColor, warnColor, errorColor := "\033[0m", "\033[1;34m", "\033[1;32m", "\033[1;33m", "\033[1;31m"
	if !config.Colorful {
		closeColor, debugColor, infoColor, warnColor, errorColor = "", "", "", "", ""
	}
	debugStr := fmt.Sprintf(colorfulPrint, debugColor, "DEBUG", closeColor)
	infoStr := fmt.Sprintf(colorfulPrint, infoColor, "INFO", closeColor)
	warnStr := fmt.Sprintf(colorfulPrint, warnColor, "WARN", closeColor)
	errStr := fmt.Sprintf(colorfulPrint, errorColor, "ERROR", closeColor)

	return &logging{
		Writer:   writer,
		Config:   config,
		infoStr:  infoStr,
		warnStr:  warnStr,
		errStr:   errStr,
		debugStr: debugStr,
	}
}

func (l logging) Debug(client *Client, msg string, data ...any) {
	if l.LogLevel >= logger.Debug {
		if client != nil && client.me != nil {
			l.Printf(l.debugStr, fmt.Sprintf("[%d] %s", client.me.ID, msg), sprint(data...))
		} else {
			l.Printf(l.debugStr, msg, sprint(data...))
		}
	}
}

func (l logging) Info(client *Client, msg string, data ...any) {
	if l.LogLevel >= logger.Info {
		if client != nil && client.me != nil {
			l.Printf(l.infoStr, fmt.Sprintf("[%d] %s", client.me.ID, msg), sprint(data...))
		} else {
			l.Printf(l.infoStr, msg, sprint(data...))
		}
	}
}

func (l logging) Warn(client *Client, msg string, data ...any) {
	if l.LogLevel >= logger.Warn {
		if client != nil && client.me != nil {
			l.Printf(l.warnStr, fmt.Sprintf("[%d] %s", client.me.ID, msg), sprint(data...))
		} else {
			l.Printf(l.warnStr, msg, sprint(data...))
		}
	}
}

func (l logging) Error(client *Client, msg string, data ...any) {
	if l.LogLevel >= logger.Error {
		if client != nil && client.me != nil {
			l.Printf(l.errStr, fmt.Sprintf("[%d] %s", client.me.ID, msg), sprint(data...))
		} else {
			l.Printf(l.errStr, msg, sprint(data...))
		}
	}
}

type Logger interface {
	Debug(*Client, string, ...any)
	Info(*Client, string, ...any)
	Warn(*Client, string, ...any)
	Error(*Client, string, ...any)
}

type Writer interface {
	Printf(string, ...any)
}

type logging struct {
	Writer
	logger.Config
	debugStr, infoStr, warnStr, errStr string
}

func sprint(a ...any) string {
	for t := range a {
		if a[t] != nil {
			rV := reflect.TypeOf(a[t]).Kind()
			if rV == reflect.Ptr {
				a[t] = reflect.ValueOf(a[t]).Elem().Interface()
			}
			rV = reflect.TypeOf(a[t]).Kind()
			if rV == reflect.Struct || rV == reflect.Map {
				a[t], _ = logger.Serialize(a[t])
			}
		}
	}
	r := fmt.Sprintln(a...)
	r = r[:len(r)-1]
	return r
}
