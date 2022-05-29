package types

type InputPath string

func (r InputPath) FileName() string {
	return string(r)
}

func (r InputPath) Content() []byte {
	return nil
}
