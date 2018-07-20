package tests

import (
	"strings"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	s := "Hello World"
	t.Log(s, lengthOfLastWord(s))

	s = ""
	t.Log(s, lengthOfLastWord(s))

	s = "a"
	t.Log(s, lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")

	if s == "" {
		return 0
	}

	index := strings.LastIndex(s, " ")
	if index == len(s)-1 {
		return 0
	}

	if index < 0 {
		index = 0
	}

	num := 0
	for ; index < len(s); index++ {
		if s[index] != ' ' {
			num++
		}
	}
	return num
}
