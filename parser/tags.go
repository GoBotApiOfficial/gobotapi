package parser

func HtmlTag(tag string) string {
	switch tag {
	case "bold":
		return "b"
	case "italic":
		return "i"
	case "strikethrough":
		return "s"
	case "underline":
		return "u"
	case "code":
		return "code"
	case "pre":
		return "pre"
	case "spoiler":
		return "tg-spoiler"
	case "text_link", "text_mention":
		return "a"
	default:
		return ""
	}
}
