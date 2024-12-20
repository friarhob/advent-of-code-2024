package utils

import (
	"log"
	"strconv"
	"strings"
)

func StringToIntSlice(s string) []int {
	var numbers []int
	slice := strings.Fields(s)

	for _, n := range slice {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalln("Error converting string to int: ", err)
			panic(err)
		}
		numbers = append(numbers, num)
	}

	return numbers
}
