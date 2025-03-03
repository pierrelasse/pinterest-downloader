package utils

import (
	"fmt"
	"io"
)

/*
Formats a string

Formatting Verbs:
  - %v: Default format for the value.
  - %#v: Go-syntax representation of the value.
  - %T: Type of the value.
  - %t: Boolean (true or false).
  - %d: Decimal representation of an integer.
  - %b: Binary representation of an integer.
  - %o: Octal representation of an integer.
  - %x: Hexadecimal representation (lowercase) of an integer.
  - %X: Hexadecimal representation (uppercase) of an integer.
  - %q: Quoted string representation for rune or byte.
  - %f: Decimal point and digits for floating-point numbers.
  - %e: Scientific notation (lowercase) for floating-point numbers.
  - %E: Scientific notation (uppercase) for floating-point numbers.
  - %g: Compact representation (either %e or %f).
  - %G: Compact representation (uppercase).
  - %s: String representation.
  - %p: Pointer address in hexadecimal.

Width and Precision:
  - %[width]d: Minimum width for integers.
  - %[width].[precision]f: Minimum width and precision for floating-point numbers.
*/
func Fmt(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

func FmtW(w io.Writer, format string, a ...any) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
}
