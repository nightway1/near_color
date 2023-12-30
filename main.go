package nco

import (
	"fmt"

	"github.com/nightway1/near_color/internal"
)

// Print

/*
Printf formats according to a format specifier and writes to standard output.
Spaces are added between operands when neither is a string.
It returns the number of bytes written and any write error encountered.

Additionally, for more advanced formatting and customization, you can incorporate color-coded printing into your Printf statements.
This can be achieved using ANSI escape codes or specialized libraries designed for colored output.
Integrating colors into your formatted print statements can enhance the visual appeal and clarity of the displayed information.
*/
func Cprintf(format string, a ...interface{}) (int, error) {
	var output string
	for _, v := range a {
		output += v.(string) + " "
	}

	return fmt.Printf(internal.ParseString(format), internal.ParseString(output))
}

/*
Print formats using the default formats for its operands and writes to standard output.
Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.

Additionally, for enhanced display, you can use color-coded printing.
To achieve this, you may employ ANSI escape codes or utilize specific libraries designed for colored output.
Adding colors to your print statements can improve readability and visual representation of information.
*/
func Cprint(a ...interface{}) (int, error) {
	return Cprintf("%s", a...)

}

/*
Println formats using the default formats for its operands and writes to standard output.
Spaces are always added between operands, and a newline is appended.
It returns the number of bytes written and any write error encountered.

For an added dimension to your output, consider incorporating colored printing into your Println statements. Y
ou can achieve this by using ANSI escape codes or specialized libraries designed for colored output.
Adding colors to your print statements not only improves visual representation but also enhances the overall readability of t
he displayed information.
*/
func Cprintln(a ...interface{}) (int, error) {
	return Cprintf("%s\n", a...)
}

// Sprint

/*
Sprintf formats according to a format specifier and returns the resulting string.
Additionally, you can enhance the visual presentation by incorporating style formatting using ANSI escape codes.
*/
func Csprintf(format string, a ...interface{}) string {
	var output string
	for _, v := range a {
		output += v.(string) + " "
	}

	return fmt.Sprintf(internal.ParseString(format), internal.ParseString(output))
}

/*
Sprint formats using the default formats for its operands and returns the resulting string. S
paces are added between operands when neither is a string. Consider utilizing ANSI escape codes for basic color formatting.
*/
func Csprint(a ...any) string {
	return Csprintf("%s", a...)
}

/*
Sprintln formats using the default formats for its operands and returns the resulting string.
Spaces are always added between operands, and a newline is appended.
*/
func Csprintln(a ...any) string {
	fmt.Sprintln()
	return Csprintf("%s\n", a...)
}

/*
Toggles the visibility of the terminal cursor.
*/
func SetCursor(enable bool) {
	switch enable {
	case true:
		fmt.Print("\x1b[?25h")
	case false:
		fmt.Print("\x1b[?25l")
	}

}

/*
Activates the alternate screen buffer in the terminal, as if you where changing worlds like other world.
*/
func SetScreen(enable bool) {
	fmt.Print()
	switch enable {
	case true:
		fmt.Print("\x1b[?1049h")
	case false:
		fmt.Print("\x1b[?1049l")
	}
}
