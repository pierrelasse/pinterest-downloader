package utils

import "os"

func Process_pid() int {
	return os.Getpid()
}
