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
	case "blockquote":
		return "blockquote"
	case "expandable_blockquote":
		return "expandable_blockquote"
	case "custom_emoji":
		return "tg-emoji"
	case "date_time":
		return "tg-time"
	default:
		return ""
	}
}
