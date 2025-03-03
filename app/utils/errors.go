package utils

import "fmt"

func Err(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}
