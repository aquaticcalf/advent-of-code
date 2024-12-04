package main

import (
	"fmt"
	"sort"
	"math"
)

func calculate_total_distance(left_list, right_list []int) int {
	sort.Ints(left_list)
	sort.Ints(right_list)

	total_distance := 0

	for i := 0; i < len(left_list); i++ {
		distance := int(math.Abs(float64(left_list[i] - right_list[i])))
		total_distance += distance
	}

	return total_distance
}

func main() {
	left_list, right_list, _ := read_lists_from_file("input.txt")

	total_distance := calculate_total_distance(left_list, right_list)
	fmt.Printf("total distance between the lists: %d\n", total_distance)
}