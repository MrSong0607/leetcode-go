package tests

import "testing"

func TestImplementstrStr(t *testing.T) {
	haystack, needle := "hello", "ll"
	t.Log(haystack, needle, strStr(haystack, needle))

	haystack, needle = "aaaaa", "bba"
	t.Log(haystack, needle, strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for index := range haystack {
		if haystack[index] == needle[0] && len(haystack)-index >= len(needle) {
			for i, j := index, 0; j < len(needle); {
				if needle[j] != haystack[i] {
					goto noMatch
				}

				j++
				i++
			}

			return index

		noMatch:
			continue
		}
	}

	return -1
}
