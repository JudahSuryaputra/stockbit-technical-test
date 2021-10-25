package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strings := []string{
		"kita",
		"atik",
		"tika",
		"aku",
		"kita",
		"makan",
		"kua",
	}

	groupedAnagram := groupAnagrams(strings)
	fmt.Println(groupedAnagram)
}

func groupAnagrams(strings []string) [][]string {
	data := make(map[string][]string)
	for _, v := range strings {
		key := SortString(v)
		data[key] = append(data[key], v)
	}
	var groupedAnagrams [][]string
	for key := range data {
		groupedAnagrams = append(groupedAnagrams, data[key])
	}
	return groupedAnagrams
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
