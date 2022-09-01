package stringutil

import (
	"regexp"
	"strings"
	"unicode"
)

const (
	lower charType = iota
	upper
	number
	other
)

type charType uint

func SplitWordsByCamelCase(str string) []string {
	if str == "" {
		return []string{}
	}

	data := []rune(str)
	prevType := charTypeOf(data[0])
	tokenStart := 0
	var result []string
	for i := 1; i < len(data); i++ {
		// TODO - checking whether the rune is valid character
		currType := charTypeOf(data[i])
		if currType == prevType {
			continue
		}
		if currType == lower && prevType == upper {
			if tokenStart != i-1 {
				// new word
				result = append(result, string(data[tokenStart:i-1]))
				tokenStart = i - 1
			}
		} else {
			// new word
			result = append(result, string(data[tokenStart:i]))
			tokenStart = i
		}
		prevType = currType
	}
	result = append(result, string(data[tokenStart:]))
	return result
}

func charTypeOf(r rune) charType {
	if unicode.IsNumber(r) {
		return number
	}
	if unicode.IsUpper(r) {
		return upper
	}
	if unicode.IsLower(r) {
		return lower
	}
	return other
}

// IsBlank reports whether the string is all whitespace or empty string ("").
//	IsBlank("") == true
//	IsBlank(" ") == true
//	IsBlank("\t\n") == true
//	IsBlank("abc") == false
//	IsBlank("  abc  ") == false
func IsBlank(str string) bool {
	if len(str) == 0 {
		return true
	}
	return strings.IndexFunc(str, func(r rune) bool {
		return !unicode.IsSpace(r)
	}) == -1
}

func ContainsIgnoreCase(str string, sub []string) bool {
	r := regexp.MustCompile("(?i)"+strings.Join(sub, "|"))
	return r.MatchString(str)
}
