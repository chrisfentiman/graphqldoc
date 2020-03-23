package parser

import (
	"strings"
	"unicode"
)

func firstToUpper(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func firstToLower(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func title(str string) string {
	words := strings.Fields(strings.ToLower(str))
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") {
			words[index] = word
		} else {
			words[index] = strings.Title(word)
		}
	}
	return strings.Join(words, " ")
}

// runeMap removes or replaces spaces, hyphens, and underscores from strings.
// Hyphens and Underscores will be removed and a space will be added
// Unless trim is true which trims all whitespace
func runeMap(str string, replaceWith []rune, trim bool) string {
	var previous rune
	isOneOf := func(r rune) bool {
		return (r == 95 || unicode.IsSpace(r) || unicode.Is(unicode.Hyphen, r))
	}
	return strings.Map(func(r rune) rune {
		defer func() {
			previous = r
		}()

		if isOneOf(r) {
			switch {
			case trim, !trim && isOneOf(previous):
				return -1
			case r == 95 && !trim && !isOneOf(previous):
				if !emptyRune(replaceWith) {
					return replaceWith[0]
				}
				return 32
			case unicode.IsSpace(r) && !trim && !isOneOf(previous):
				if !emptyRune(replaceWith) {
					return replaceWith[0]
				}
				return r
			case unicode.Is(unicode.Hyphen, r) && !trim && !isOneOf(previous):
				if !emptyRune(replaceWith) {
					return replaceWith[0]
				}
				return 32
			default:
				return r
			}
		}

		return r

	}, str)
}

func emptyRune(r []rune) bool {
	return len([]rune(r)) <= 0
}
