package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/friarhob/advent-of-code-2024/utils"
)

func main() {
	inputFile := flag.String("input", "", "Path to the input file")
	flag.Parse()

	if *inputFile == "" {
		log.Fatalf("Error: Input file not specified. Use -input <filename> to specify the input file.")
	}

	if _, err := os.Stat(*inputFile); os.IsNotExist(err) {
		log.Fatalf("Error: Input file %s does not exist.", *inputFile)
	}

	lines, err := utils.ReadLines(*inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	part1Result := SolvePart1(lines)
	fmt.Printf("Part 1: %v\n", part1Result)

	part2Result := SolvePart2(lines)
	fmt.Printf("Part 2: %v\n", part2Result)
}
