package proxy_reader

import (
	"bytes"
	"io"
)

func (b *ProxyReader) ReadFrom(r io.Reader) (n int64, err error) {
	n, err = b.Reader.(*bytes.Buffer).ReadFrom(r)
	b.processed += n
	return
}
