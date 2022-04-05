package utils

import "regexp"

func GetFirstParam(expr string, path string) string {
	re, _ := regexp.Compile(expr)
	values := re.FindStringSubmatch(path)
	return values[1]
}
