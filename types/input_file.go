package types

type InputFile struct {
	Name  string
	Bytes []byte
}

func (r InputFile) FileName() string {
	return r.Name
}

func (r InputFile) Content() []byte {
	return r.Bytes
}
