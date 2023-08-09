package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	TIME_LIMIT = 30
)

func main() {
	timeLimit := flag.Int("limit", TIME_LIMIT, "the time limit for the quiz in seconds")
	flag.Parse()

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
	fmt.Print("Press enter to start")
	fmt.Scanln()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for _, record := range records {

		// print question
		fmt.Print(record[0] + " = ")
		ansChan := make(chan int)

		go func() {
			var answer string
			fmt.Scanln(&answer)
			// parse answer to int
			answerInt, err := strconv.Atoi(answer)
			if err != nil {
				fmt.Println("Error parsing answer")
				os.Exit(1)
			}
			ansChan <- answerInt
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			fmt.Println("Your score is", score)
			return
		case answerInt := <-ansChan:
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
	}

	fmt.Println("Your score is", score)
}
