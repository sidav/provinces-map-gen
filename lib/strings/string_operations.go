package strings

import (
	"regexp"
	"strings"
)

func StringsAreRoughlyEqual(s1, s2 string) bool {
	s1Modded := strings.Replace(strings.ToLower(s1), " ", "", -1)
	s2Modded := strings.Replace(strings.ToLower(s2), " ", "", -1)
	return strings.Contains(s1Modded, s2Modded) || strings.Contains(s2Modded, s1Modded)
}

func CenterStringWithSpaces(str string, desiredLength int) string {
	padAmount := desiredLength - len(str)
	if padAmount < 1 {
		return str
	}
	padLeft := padAmount / 2
	padRight := padAmount / 2
	if padAmount%2 == 1 {
		padLeft++
	}
	return strings.Repeat(" ", padLeft) + str + strings.Repeat(" ", padRight)
}

func DewovelAndTrimString(s string, trimLength int) string {
	s = regexp.MustCompile("[AEIOUY ]|[aeiouy]").ReplaceAllString(s, "")
	if trimLength > 0 && len(s) > trimLength {
		s = s[:trimLength]
	}
	return s
}
