package util

import (
	"encoding/json"
	"os"
)

func ToJSONBytes(v interface{}) []byte {
	raw, _ := json.Marshal(v)
	return raw
}

func LoadJSONFile(filename string, data interface{}) error {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileData, data); err != nil {
		return err
	}

	return nil
}
