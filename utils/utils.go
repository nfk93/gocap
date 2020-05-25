package utils

import (
	"strings"
)

func RemoveParentheses(typeString string) string {
	typeString = strings.ReplaceAll(typeString, " ", "_sp_")
	typeString = strings.ReplaceAll(typeString, "(", "_lp_")
	typeString = strings.ReplaceAll(typeString, ")", "_rp_")
	typeString = strings.ReplaceAll(typeString, "[", "_lb_")
	typeString = strings.ReplaceAll(typeString, "]", "_rb_")

	return typeString
}
