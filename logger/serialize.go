package logger

import (
	"encoding/json"
)

func Serialize(v any) (string, error) {
	marshal, err := json.MarshalIndent(internalSerialize(v), "", "    ")
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
