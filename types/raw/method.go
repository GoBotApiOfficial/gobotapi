package raw

type Method interface {
	MethodName() string
	Files() map[string]InputFile
	ParseResult(response []byte) (*Result, error)
	ProgressCallable() ProgressCallable
}
