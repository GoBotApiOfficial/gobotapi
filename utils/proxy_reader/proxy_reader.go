package proxy_reader

import (
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"io"
	"time"
)

func NewProxyReader(buf io.Reader, refreshRate time.Duration, totalBytes int64, callable rawTypes.ProgressCallable) *ProxyReader {
	b := &ProxyReader{
		Reader:      buf,
		callable:    callable,
		finish:      make(chan bool),
		refreshRate: refreshRate,
		totalBytes:  totalBytes,
	}
	if callable != nil {
		go func() {
			for {
				select {
				case <-b.finish:
					return
				case <-time.After(b.refreshRate):
					if b.processed < b.totalBytes {
						b.callable(b.processed, b.totalBytes)
					}
				}
			}
		}()
	}
	return b
}
