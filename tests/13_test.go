package tests

import (
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	t.Log("MCMXCIV", romanToInt("MCMXCIV"))
	t.Log("LVIII", romanToInt("LVIII"))
}

func romanToInt(s string) int {
	m := make(map[byte]int)
	m[byte('I')] = 1
	m[byte('V')] = 5
	m[byte('X')] = 10
	m[byte('L')] = 50
	m[byte('C')] = 100
	m[byte('D')] = 500
	m[byte('M')] = 1000

	length := len(s)
	num := 0
	for i := 0; i < length; i++ {
		if i < length-1 && m[s[i]] < m[s[i+1]] {
			num += (m[s[i+1]] - m[s[i]])
			i++
		} else {
			num += m[s[i]]
		}
	}

	return num
}
