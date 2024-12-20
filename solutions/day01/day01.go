package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func SolvePart1(lines []string) int {
	var first []int
	var second []int

	for _, value := range lines {
		numbers := strings.Split(value, "   ")

		n1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatalln("Error converting string to int: ", err)
			panic(err)
		}
		first = append(first, n1)

		n2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatalln("Error converting string to int: ", err)
			panic(err)
		}
		second = append(second, n2)
	}

	sort.Ints(first)
	sort.Ints(second)

	total_diff := 0
	for i := 0; i < len(first) && i < len(second); i++ {
		if first[i] < second[i] {
			total_diff += second[i] - first[i]
		} else {
			total_diff += first[i] - second[i]
		}
	}

	return total_diff
}

func SolvePart2(lines []string) int {
	counter := make(map[int]int)

	var first []int
	for _, value := range lines {
		numbers := strings.Split(value, "   ")

		n1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatalln("Error converting string to int: ", err)
			panic(err)
		}
		first = append(first, n1)

		n2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatalln("Error converting string to int: ", err)
			panic(err)
		}
		counter[n2]++
	}

	similarity := 0
	for i := 0; i < len(first); i++ {
		similarity += first[i] * counter[first[i]]
	}

	return similarity
}
