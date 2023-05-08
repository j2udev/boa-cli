package templates

import (
	"bytes"
	"path"
	"unicode"
)

func pascalCase(str string) string {
	var buf bytes.Buffer
	shouldCapitalize := true
	for _, c := range str {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			if shouldCapitalize {
				buf.WriteRune(unicode.ToUpper(c))
				shouldCapitalize = false
			} else {
				buf.WriteRune(unicode.ToLower(c))
			}
		} else {
			shouldCapitalize = true
		}
	}
	return buf.String()
}

func lastIndex(slice []string) int {
	return len(slice) - 1
}

func moduleName(modulePath string) string {
	return path.Base(modulePath)
}
