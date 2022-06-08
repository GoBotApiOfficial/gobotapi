package types

type InputBytes struct {
	Name string
	Data []byte
}

func (r InputBytes) FileName() string {
	return r.Name
}

func (r InputBytes) Content() []byte {
	return r.Data
}
