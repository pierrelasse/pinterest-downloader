package utils

import "encoding/json"

func JSON_decode(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
