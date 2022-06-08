package proxy_reader

import (
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"io"
	"time"
)

type ProxyReader struct {
	io.Reader
	processed   int64
	callable    rawTypes.ProgressCallable
	finish      chan bool
	refreshRate time.Duration
	totalBytes  int64
}
