package raw

type InputFile interface {
	FileName() string
	IsURL() bool
	Content() []byte
}
