package utils

import "strings"

func RemoveParentheses(typeString string) string {
	typeString = strings.ReplaceAll(typeString, " ", "")
	typeString = strings.ReplaceAll(typeString, "(", "_lp_")
	typeString = strings.ReplaceAll(typeString, ")", "_rp_")
	typeString = strings.ReplaceAll(typeString, "[", "_lb_")
	typeString = strings.ReplaceAll(typeString, "]", "_rb_")
	typeString = strings.ReplaceAll(typeString, "\n", "")
	typeString = strings.ReplaceAll(typeString, "\t", "")
	typeString = strings.ReplaceAll(typeString, "{", "_lc_")
	typeString = strings.ReplaceAll(typeString, "}", "_rc_")
	typeString = strings.ReplaceAll(typeString, "<-", "_ar_")
	typeString = strings.ReplaceAll(typeString, "*", "_st_")
	return typeString
}

var HasCapChan = false

//Path from GOROOT to the current folder
var PackagePath = ""
