package main

import (
	"fmt"
)

func calculate_similarity_score(left_list, right_list []int) int {
	right_frequency := make(map[int]int)
	for _, num := range right_list {
		right_frequency[num]++
	}

	similarity_score := 0
	for _, num := range left_list {
		similarity_score += num * right_frequency[num]
	}

	return similarity_score
}

func main() {
	left_list, right_list, _ := read_lists_from_file("input.txt")

	similarity_score := calculate_similarity_score(left_list, right_list)

	fmt.Printf("similarity score: %d\n", similarity_score)
}
