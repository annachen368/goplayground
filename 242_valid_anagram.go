package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := [26]int{}

	for i := 0; i < len(s); i++ {
		letters[s[i]-'a']++
		letters[t[i]-'a']--
	}

	for i := 0; i < len(letters); i++ {
		if letters[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
