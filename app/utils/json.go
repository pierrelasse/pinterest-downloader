package utils

import "encoding/json"

type JSON = map[string]interface{}

func JSON_decode(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
