package gotopt

import (
	"os"
	"strings"
)

// strncmp compares up to num characters of the string str1 to those of the
// string str2.
//
// This function starts comparing the first character of each string.
// If they are equal to each other, it continues with the following pairs
// until the characters differ, until a terminating null-character is reached,
// or until num characters match in both strings, whichever happens first.
//
// If the first character that does not match has a lower value in str1 than
// in str2 then a negative number is returned.
//
// If the contents of both strings are equal than zero is returned.
//
// If the first character that does not match has a greater value in str1 than
// in str2 then a positive number is returned.
func strncmp(str1, str2 string, num int) int {
	lstr1 := len(str1)
	lstr2 := len(str2)

	for x := 0; x < num; x++ {
		if (lstr1 < x && lstr2 >= x) || (lstr2 < x && lstr1 >= x) {
			break
		}
		if str1[x] < str2[x] {
			return -1
		} else if str1[x] > str2[x] {
			return 1
		}
	}

	return 0
}

// strncmpb is like strncmp except true or false is returned based on whether
// strncmp would return that the strings are equal or different.
func strncmpb(str1, str2 string, num int) bool {
	return strncmp(str1, str2, num) == 0
}

func envVarExists(name string) bool {
	for _, ev := range os.Environ() {
		if strings.HasPrefix(ev, name) {
			return true
		}
	}
	return false
}

func toIntFromBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func parseLongOptSize(s string, nextChar int) (nameEnd, nameLen int) {
	nameLen = len(s)
	if nextChar > nameLen-1 {
		return -1, nameLen
	}

	s = s[nextChar:]
	nameLen = len(s)

	if nameEnd = strings.IndexByte(s, '='); nameEnd > -1 {
		nameLen = nameEnd
	}

	return
}
