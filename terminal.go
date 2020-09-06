package terminal

import "fmt"

func ResetStyle() {
	fmt.Print("\033[0m")
}
