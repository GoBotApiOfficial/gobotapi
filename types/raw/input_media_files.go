package raw

type InputMediaFiles interface {
	Files() map[string]InputFile
}
