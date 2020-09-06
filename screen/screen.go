package screen

import "fmt"

func ClearScreen() {
	fmt.Print("\033[?1049h\033[H")
}

func RestoreScreen() {
	fmt.Print("\033[?1049l")
}

func ToExitKeyPress() {
	fmt.Println("\n\nPress any key to exit program")
	fmt.Scanf(" ")
}
