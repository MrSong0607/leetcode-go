package tests

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	t.Log(`"flower", "flow", "flight" -->`, longestCommonPrefix([]string{"flower", "flow", "flight"}))
	t.Log(`"dog","racecar","car" -->`, longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		cur := strs[i]
		if len(strs[i]) < len(prefix) {
			prefix, cur = strs[i], prefix
		}

		for l := len(prefix); l >= 0; l-- {
			if prefix[0:l] == cur[0:l] {
				prefix = prefix[0:l]
				break
			}
		}
	}

	return prefix
}
