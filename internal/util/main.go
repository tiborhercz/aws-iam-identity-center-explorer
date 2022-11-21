package util

import (
	"bytes"
	"encoding/json"
)

func PrettifyJson(value []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, value, "", "\t")
	if err != nil {
		return "", err
	}

	return prettyJSON.String(), nil
}
