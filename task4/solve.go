package main

import "unicode"

func RemoveEven(arr []int) []int {
	answer := make([]int, 0)
	for _, elem := range arr {
		if elem % 2 == 1 {
			answer = append(answer, elem)
		}
	}
	return answer
}

func PowerGenerator(a int) func() int {
	x := 1
	return func() int {
		x *= a
		return x
	}
}

func DifferentWordsCount(s string) int {
	word := ""
	set := make(map[string]bool)
	answer := 0
	for _, c := range (s + " ") {
		if unicode.IsLetter(c) {
			word += string(unicode.ToLower(c))
		} else if word != "" {
			if !set[word] {
				answer += 1
			}
			set[word] = true
			word = ""
		}
	}
	return answer
}

