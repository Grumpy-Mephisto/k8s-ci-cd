package util

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

func ToJSONBytes(v interface{}) []byte {
	raw, _ := json.Marshal(v)
	return raw
}

func LoadJSONFile(filename string, data interface{}) error {
	repoFile := filepath.Clean(filename)
	if !strings.HasPrefix(repoFile, "data/") {
		return os.ErrInvalid
	}

	byContext, err := os.ReadFile(repoFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(byContext, data)
}
