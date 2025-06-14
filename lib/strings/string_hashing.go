package strings

import (
	"fmt"
	"strings"
)

func HashStringsToShortestDistincts(stringGetter func(int) string, totalStrings int) []string {
	hashes := make([]string, 0)
	for currStringIndex := 0; currStringIndex < totalStrings; currStringIndex++ {
		currString := stringGetter(currStringIndex)
		candidate := ""
		for upper := 0; upper <= 1; upper++ {
			for indexInLong := 0; indexInLong < len(currString); indexInLong++ {
				if upper == 0 {
					candidate = strings.ToLower(string(currString[indexInLong]))
				} else {
					candidate = strings.ToUpper(string(currString[indexInLong]))
				}
				// check if candidate is in other strings
				alreadyUsed := false
				for indexInHash := range hashes {
					if hashes[indexInHash] == candidate {
						alreadyUsed = true
						break
					}
				}
				if !alreadyUsed {
					hashes = append(hashes, candidate)
					break
				}
			}
			if len(hashes) == currStringIndex+1 {
				break
			}
		}
		if len(hashes) < currStringIndex+1 {
			panic(fmt.Sprintf("Can't assign, problem with %dth %s", currStringIndex, currString))
		}
	}
	return hashes
}

func SelectIndexFromStringsByHash(stringGetter func(int) string, totalStrings int, hash string) int {
	hashes := HashStringsToShortestDistincts(stringGetter, totalStrings)
	for i := range hashes {
		if hashes[i] == hash {
			return i
		}
	}
	return -1
}
