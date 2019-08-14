package gnirts

import (
	"regexp"
	"strings"
)

// ReverseText reverses a string
func ReverseText(s string) string {

	r := []rune(s)

	// method 1
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)

	// method 2
	// var b strings.Builder

	// for i := len(r) - 1; i >= 0; i-- {
	// 	b.WriteRune(r[i])
	// }

	// return b.String()
}

// ReverseWords reverses words based on a delimeter, e.g. " "
func ReverseWords(s string, d string) string {
	words := strings.Split(s, d)
	reversed := reverseStrSlice(words)
	return strings.Join(reversed, d)
}

func reverseStrSlice(strings []string) []string {
	for i := 0; i < len(strings)/2; i++ {
		j := len(strings) - i - 1
		strings[i], strings[j] = strings[j], strings[i]
	}
	return strings
}

// ReverseSentence reverses words based on zero or more delimeters, e.g. {" ", "."}
// This is a solution for a previous version of the course material
func ReverseSentence(s string, delims []string) string {

	delimRegex := "[" + strings.Join(delims, "") + "]"

	re := regexp.MustCompile(`(.*?` + delimRegex + `)\s*`)
	matches := re.FindAllStringSubmatch(s, -1)
	var sentences = []string{}

	for _, match := range matches {
		sentences = append(sentences, match[1]) // match[1] is the submatch
	}

	reversed := reverseStrSlice(sentences)

	return strings.Join(reversed, " ")
}
