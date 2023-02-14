package prettyprint

import "fmt"

type color string

// Define colors
var (
	GREEN color = "\033[32m"
)

// CPrintF is a function that prints a colored string and arbitrary a formatted string.
func CPrintF(color color, format string, colored string, a ...any) {
	fmt.Printf("%s%s\033[0m %s", color, colored, fmt.Sprintf(format, a...))
}
