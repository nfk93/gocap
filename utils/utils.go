package utils

import str "strings"

func RemoveParentheses(typeString string) string {
	typeString = str.ReplaceAll(typeString, " ", "_sp_")
	typeString = str.ReplaceAll(typeString, "(", "_lp_")
	typeString = str.ReplaceAll(typeString, ")", "_rp_")
	typeString = str.ReplaceAll(typeString, "[", "_lb_")
	typeString = str.ReplaceAll(typeString, "]", "_rb_")

	return typeString
}
