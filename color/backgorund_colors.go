package color

func BackgroundBlack(text string) string {
	styledText := "\u001b[40m" + text + "\033[0m"
	return styledText
}

func BackgroundBrightBlack(text string) string {
	styledText := "\u001b[40;1m" + text + "\033[0m"
	return styledText
}

func BackgroundRed(text string) string {
	styledText := "\u001b[41m" + text + "\033[0m"
	return styledText
}

func BackgroundBrightRed(text string) string {
	styledText := "\u001b[41;1m" + text + "\033[0m"
	return styledText
}

func BackgroundGreen(text string) string {
	styledText := "\u001b[42m" + text + "\033[0m"
	return styledText
}

func BackgroundBrightGreen(text string) string {
	styledText := "\u001b[42;1m" + text + "\033[0m"
	return styledText
}

func BackgroundYellow(text string) string {
	styledText := "\u001b[43m" + text + "\033[0m"
	return styledText
}

func BackgroundBrightYellow(text string) string {
	styledText := "\u001b[43;1m" + text + "\033[0m"
	return styledText
}

func BackgroundBlue(text string) string {
	styledText := "\u001b[44m" + text + "\033[0m"
	return styledText
}

func BackgroundBrightBlue(text string) string {
	styledText := "\u001b[44;1m" + text + "\033[0m"
	return styledText
}

func BackgroundMagenta(text string) string {
	styledText := "\u001b[45m" + text + "\033[0m"
	return styledText
}

func BackgroundBrightMagenta(text string) string {
	styledText := "\u001b[45;1m" + text + "\033[0m"
	return styledText
}

func BackgroundCyan(text string) string {
	styledText := "\u001b[46m" + text + "\033[0m"
	return styledText
}

func BackgroundBrightCyan(text string) string {
	styledText := "\u001b[46;1m" + text + "\033[0m"
	return styledText
}

func BackgroundWhite(text string) string {
	styledText := "\u001b[47m" + text + "\033[0m"
	return styledText
}

func BackgroundBrightWhite(text string) string {
	styledText := "\u001b[47;1m" + text + "\033[0m"
	return styledText
}
