package utils

import (
	"bytes"
	"encoding/json"
)

func Unmarshal(data []byte, v interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	return decoder.Decode(v)
}

func GetJson(v interface{}) string {
	js, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(js)
}
