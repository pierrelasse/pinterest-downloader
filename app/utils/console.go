package utils

import "os"

func Console_clear() {
	Console_write("\033[H\033[2J")
}

func Console_write(str string) {
	os.Stdout.Write([]byte(str))
}

func Console_writeln(str string) {
	os.Stdout.Write([]byte(str + "\n"))
}
