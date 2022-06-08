package types

type InputURL string

func (r InputURL) FileName() string {
	return string(r)
}

func (r InputURL) Content() []byte {
	return nil
}
