package font

func Bold(text string) string {
	styledText := "\u001b[1m" + text + "\033[0m"
	return styledText
}

func Underline(text string) string {
	styledText := "\u001b[4m" + text + "\033[0m"
	return styledText
}

func Reversed(text string) string {
	styledText := "\u001b[7m" + text + "\033[0m"
	return styledText
}
