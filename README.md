# Terminal UI Package For Golang

The package includes colors, background colors and decorators for texts. Also clear the screen, output something on the display, and then restore the screen to the preserved state that it was in before the task was carried out.

### Usage
Color functions return string. The resulting strings can be used with the "fmt" package.

```sh
fmt.Println(color.Black("Black Text"))
fmt.Printf("%s", color.BackgroundRed("Red Backgrounded Text"))
font.Bold("Bold Text!")
```

#### Terminal control/Preserve screen
Clear the screen, output something on the display, and then restore the screen to the preserved state that it was in before the task was carried out.

###### Clear Screen
##

```sh
screen.ClearScreen()
```

###### Restore Screen
##

```sh
screen.RestoreScreen()
```
