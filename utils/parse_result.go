package utils

import (
	"encoding/json"
	"errors"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

func ParseResult(rawResult []byte, err error, method rawTypes.Method) (*rawTypes.Result, error) {
	if err != nil {
		var result struct {
			Ok          bool   `json:"ok"`
			Description string `json:"description"`
		}
		errMarshal := json.Unmarshal(rawResult, &result)
		if errMarshal != nil {
			return nil, err
		}
		return nil, errors.New(result.Description)
	}
	methodResult, err := method.ParseResult(rawResult)
	if err != nil {
		return nil, err
	}
	return methodResult, nil
}
