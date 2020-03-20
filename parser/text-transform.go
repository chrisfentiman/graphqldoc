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

func title(input string) string {
	words := strings.Fields(input)
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

// String Split & Converts
// Removes or replaces Spaces, hyphens, and underscores from strings
// Hyphens and Underscores will be removed and a space will be added
// Unless force is true. Force always removes spaces
func runeMap(str string, replaceWith []rune, force bool) string {
	return strings.Map(func(r rune) rune {
		if unicode.Is(unicode.Hyphen, r) || string(r) == "_" || unicode.IsSpace(r) {
			switch {
			case force:
				return -1
			case !emptyRune(replaceWith) && !force:
				return replaceWith[0]
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
