package gobotapi

import (
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/logger"
	"github.com/Laky-64/gologging"
	"reflect"
	"regexp"
	"strings"
)

const LoggerName = "gobotapi"

var tagRgx = regexp.MustCompile(`^([a-z_]{2,10}): `)

func NewLogger(level logger.LogLevel) Logger {
	l := gologging.GetLogger(LoggerName)
	l.SetLevel(gologging.DebugLevel)
	return &logging{
		level:  level,
		logger: l,
	}
}

func (l *logging) Debug(client *Client, msg string, data ...any) {
	l.log(logger.Debug, l.logger.Debug, client, msg, data...)
}

func (l *logging) Info(client *Client, msg string, data ...any) {
	l.log(logger.Info, l.logger.Info, client, msg, data...)
}

func (l *logging) Warn(client *Client, msg string, data ...any) {
	l.log(logger.Warn, l.logger.Warn, client, msg, data...)
}

func (l *logging) Error(client *Client, msg string, data ...any) {
	l.log(logger.Error, l.logger.Error, client, msg, data...)
}

func (l *logging) log(level logger.LogLevel, print func(...any), client *Client, msg string, data ...any) {
	if l.level < level {
		return
	}
	msg, data = censorSensitiveInfo(client, msg, data...)
	if len(data) > 0 {
		msg = fmt.Sprintf("%s %s", msg, sprint(data...))
	}
	if client != nil && client.me != nil {
		if tag := tagRgx.FindString(msg); len(tag) > 0 {
			msg = fmt.Sprintf("%s[%d] %s", tag, client.me.ID, strings.TrimPrefix(msg, tag))
		} else {
			msg = fmt.Sprintf("[%d] %s", client.me.ID, msg)
		}
	}
	print(msg)
}

type Logger interface {
	Debug(*Client, string, ...any)
	Info(*Client, string, ...any)
	Warn(*Client, string, ...any)
	Error(*Client, string, ...any)
}

type logging struct {
	level  logger.LogLevel
	logger *gologging.Logger
}

func censorSensitiveInfo(client *Client, a string, b ...any) (string, []any) {
	if client == nil || client.Token == "" {
		return a, b
	}
	tokenFixed := strings.Split(client.Token, ":")[1]
	for t := range b {
		if b[t] != nil {
			rV := reflect.TypeOf(b[t]).Kind()
			if rV == reflect.String {
				b[t] = strings.ReplaceAll(b[t].(string), tokenFixed, strings.Repeat("X", len(tokenFixed)))
			}
		}
	}
	a = strings.ReplaceAll(a, tokenFixed, strings.Repeat("X", len(tokenFixed)))
	return a, b
}

func sprint(a ...any) string {
	for t := range a {
		if a[t] != nil {
			rV := reflect.TypeOf(a[t]).Kind()
			if rV == reflect.Ptr {
				a[t] = reflect.ValueOf(a[t]).Elem().Interface()
			}
			rV = reflect.TypeOf(a[t]).Kind()
			if rV == reflect.Struct || rV == reflect.Map || rV == reflect.Slice {
				a[t], _ = logger.Serialize(a[t])
			}
		}
	}
	r := fmt.Sprintln(a...)
	r = r[:len(r)-1]
	return r
}
