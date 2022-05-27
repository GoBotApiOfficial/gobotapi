package types

type InputURL string

func (r InputURL) FileName() string {
	return string(r)
}

func (r InputURL) IsURL() bool {
	return true
}

func (r InputURL) Content() []byte {
	return nil
}
