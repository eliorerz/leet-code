package findAndReplacePattern

func getPattern(word string) []int {
	var pattern []int            // O(n)
	table := make(map[uint8]int) // O(n)

	j := 0
	for i := 0; i < len(word); i++ {
		if _, ok := table[word[i]]; !ok {
			table[word[i]] = j
			j++
		}
		pattern = append(pattern, table[word[i]])
	}
	return pattern
}

// Time: O(n)
// Size: O(n^2)?
func FindAndReplacePattern(words []string, pattern string) []string {
	ptrn := getPattern(pattern)
	var res []string

	for _, word := range words {
		if len(word) != len(pattern) {
			continue
		}

		wordPattern := getPattern(word)
		found := true
		if len(wordPattern) != len(ptrn) {
			continue
		}

		for i := 0; i < len(ptrn); i++ {
			if wordPattern[i] != ptrn[i] {
				found = false
				break
			}
		}
		if found {
			res = append(res, word)
		}
	}
	return res
}
