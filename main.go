package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")

	// read problems.csv file
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	defer file.Close()

	// read file contents
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	score := 0

	fmt.Println("--- Quiz ---")

	for _, record := range records {
		fmt.Print(record[0] + " = ")
		var answer string
		fmt.Scanln(&answer)

		// parse answer to int
		answerInt, err := strconv.Atoi(answer)
		if err != nil {
			fmt.Println("Error parsing answer")
			os.Exit(1)
		}

		// parse record[1] to int
		recordInt, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Println("Error parsing result")
			os.Exit(1)
		}

		if answerInt == recordInt {
			score++
		}
	}

	fmt.Println("Your score is", score)
}
