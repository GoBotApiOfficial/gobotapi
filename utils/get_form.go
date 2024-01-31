package utils

import (
	"encoding/json"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"strings"
)

func GetForm(method rawTypes.Method) (map[string]string, error) {
	var jsonByte map[string]json.RawMessage
	bytes, err := json.Marshal(method)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &jsonByte)
	if err != nil {
		return nil, err
	}
	formData := make(map[string]string)
	for k, v := range jsonByte {
		formData[k] = string(v)
		if strings.HasPrefix(formData[k], "\"") {
			var fixedString string
			_ = json.Unmarshal(v, &fixedString)
			formData[k] = fixedString
		}
	}
	return formData, nil
}
