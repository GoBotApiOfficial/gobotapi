package utils

import "strings"

func IsServerError(rawResult []byte, err error) bool {
	res, err := ParseResult(rawResult, err, nil)
	if err != nil {
		if res != nil {
			return res.Error.Code >= 500 && res.Error.Code < 600
		} else {
			return strings.Contains(err.Error(), "graceful shutdown GOAWAY")
		}
	}
	return false
}
