package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

func UUID_newV4(upper bool) string {
	uuid := make([]byte, 16)
	if _, err := rand.Read(uuid); err != nil {
		return ""
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	str := fmt.Sprintf("%08x-%04x-%04x-%04x-%12x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
	if upper {
		str = strings.ToUpper(str)
	}
	return str
}
