package main

import (
	"fmt"
	"strconv"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		letters := make([]int, 26)
		s := strs[i]

		for j := 0; j < len(s); j++ {
			letters[s[j]-'a']++
		}

		var sb strings.Builder
		for j := 0; j < len(letters); j++ {
			sb.WriteString(strconv.Itoa(letters[j]))
			sb.WriteByte('#')
		}

		key := sb.String()
		m[key] = append(m[key], s)
	}

	var res [][]string

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
