package raw

type InputFile interface {
	FileName() string
	Content() []byte
}
