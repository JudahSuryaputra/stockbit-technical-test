package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "(backend developer)"
	fmt.Println(getStringInBracket(str))
}

func getStringInBracket(str string) string {
	if len(str) > 0 {
		openingBracket := strings.Index(str, "(")
		closingBracket := strings.Index(str, ")")
		if openingBracket == 0 && closingBracket == len(str)-1 {
			return strings.Trim(str, "()")
		}
	}
	return ""
}
