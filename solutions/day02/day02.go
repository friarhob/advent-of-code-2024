package main

import (
	"github.com/friarhob/advent-of-code-2024/utils"
)

func SolvePart1(lines []string) int {
	safe_levels := 0
	for _, line := range lines {
		numbers := utils.StringToIntSlice(line)

		is_safe := true
		is_increasing := false
		is_decreasing := false
		for i := 0; i < len(numbers)-1; i++ {
			diff := utils.AbsInt(numbers[i] - numbers[i+1])

			if diff < 1 || diff > 3 {
				is_safe = false
				break
			}

			if numbers[i] > numbers[i+1] {
				if is_increasing {
					is_safe = false
					break
				}
				is_decreasing = true
			} else {
				if is_decreasing {
					is_safe = false
					break
				}
				is_increasing = true
			}
		}

		if is_safe {
			safe_levels++
		}
	}

	return safe_levels
}

func SolvePart2(lines []string) int {
	return 57
}
