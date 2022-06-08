package proxy_reader

func (b *ProxyReader) Close() {
	if b != nil && b.callable != nil {
		b.processed = b.totalBytes
		b.callable(b.processed, b.totalBytes)
		b.finish <- true
	}
}
