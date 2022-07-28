package anagram

// Time: O(s) + O(t) + O(1) = O(n)
// Size: O(s) + O(t) = O(n)
func IsUnicodeCharsAnagram(s string, t string) bool {
	table := make(map[uint8]int)
	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if _, ok := table[s[i]]; ok {
			table[s[i]] += 1
		} else {
			table[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := table[t[i]]; ok {
			if table[t[i]] == 0 {
				return false
			}
			table[t[i]] -= 1

		} else {
			return false
		}
	}

	return true
}
