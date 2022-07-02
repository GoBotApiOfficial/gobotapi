package utils

import (
	"encoding/json"
	"errors"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

func ParseResult(rawResult []byte, err error, method rawTypes.Method) (*rawTypes.Result, error) {
	if err != nil {
		var result rawTypes.ErrorMessage
		errMarshal := json.Unmarshal(rawResult, &result)
		if errMarshal != nil {
			return nil, err
		}
		return &rawTypes.Result{
			Kind:  types.TypeErrorMessage,
			Error: result,
		}, errors.New(result.Description)
	}
	methodResult, err := method.ParseResult(rawResult)
	if err != nil {
		return nil, err
	}
	return methodResult, nil
}
