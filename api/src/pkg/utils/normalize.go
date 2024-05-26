package utils

import (
	"strings"

	"golang.org/x/text/width"
)

func NormalizeString(str string) string {
	return width.Narrow.String(strings.ToLower(str))
}
