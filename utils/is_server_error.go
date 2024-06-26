package utils

import (
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"strings"
)

func IsServerError(rawResult []byte, err error, method rawTypes.Method) bool {
	res, err := ParseResult(rawResult, err, method)
	if err != nil {
		if res != nil {
			return res.Error.Code >= 500 && res.Error.Code < 600
		} else {
			return strings.Contains(err.Error(), "graceful shutdown GOAWAY")
		}
	}
	return false
}
