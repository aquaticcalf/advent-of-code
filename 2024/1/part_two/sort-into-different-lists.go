package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_lists_from_file(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left_list, right_list []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in left list: %s", parts[0])
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number in right list: %s", parts[1])
		}

		left_list = append(left_list, left)
		right_list = append(right_list, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left_list, right_list, nil
}