package proxy_reader

func (b *ProxyReader) Read(p []byte) (n int, err error) {
	n, err = b.Reader.Read(p)
	b.processed += int64(n)
	return
}
