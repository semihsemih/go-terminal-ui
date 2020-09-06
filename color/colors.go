package color

func Black(text string) string {
	styledText := "\u001b[30m" + text + "\033[0m"
	return styledText
}

func BrightBlack(text string) string {
	styledText := "\u001b[30;1m" + text + "\033[0m"
	return styledText
}

func Red(text string) string {
	styledText := "\u001b[31m" + text + "\033[0m"
	return styledText
}

func BrightRed(text string) string {
	styledText := "\u001b[31;1m" + text + "\033[0m"
	return styledText
}

func Green(text string) string {
	styledText := "\u001b[32m" + text + "\033[0m"
	return styledText
}

func BrightGreen(text string) string {
	styledText := "\u001b[32;1m" + text + "\033[0m"
	return styledText
}

func Yellow(text string) string {
	styledText := "\u001b[33m" + text + "\033[0m"
	return styledText
}

func BrightYellow(text string) string {
	styledText := "\u001b[33;1m" + text + "\033[0m"
	return styledText
}

func Blue(text string) string {
	styledText := "\u001b[34m" + text + "\033[0m"
	return styledText
}

func BrightBlue(text string) string {
	styledText := "\u001b[34;1m" + text + "\033[0m"
	return styledText
}

func Magenta(text string) string {
	styledText := "\u001b[35m" + text + "\033[0m"
	return styledText
}

func BrightMagenta(text string) string {
	styledText := "\u001b[35;1m" + text + "\033[0m"
	return styledText
}

func Cyan(text string) string {
	styledText := "\u001b[36m" + text + "\033[0m"
	return styledText
}

func BrightCyan(text string) string {
	styledText := "\u001b[36;1m" + text + "\033[0m"
	return styledText
}

func White(text string) string {
	styledText := "\u001b[37m" + text + "\033[0m"
	return styledText
}

func BrightWhite(text string) string {
	styledText := "\u001b[37;1m" + text + "\033[0m"
	return styledText
}
