package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type quiz struct {
	Question string
	Answer   string
}

func main() {
	greeting()
	csvQuest := readFile()
	problems := makeStruct(csvQuest)

	correct := 0

	correct = askQuestions(problems, correct)

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func greeting() {
	fmt.Println("Hello! Let's play a Quiz!")
	fmt.Println("Please enter the right answer and press ENTER")
	fmt.Println("You have a Time Limit - as soon as you press enter time counts.")
	fmt.Scanln()
	startTimer()
}

func readFile() [][]string {
	fileName := flag.String("csv", "questions.csv", "a CSV File in the format q,a")
	flag.Parse()
	csvFile, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}
	fmt.Println("Successfully opened your questions-dictionary!")
	defer csvFile.Close()
	csvQuest, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return csvQuest
}

func makeStruct(fileString [][]string) []quiz {
	ret := make([]quiz, len(fileString))
	for i, line := range fileString {
		ret[i] = quiz{
			Question: line[0],
			Answer:   line[1],
		}
	}
	return ret
}

func askQuestions(problems []quiz, correct int) int {
	for i, p := range problems {
		fmt.Printf("Problem Number %d: %s = \n", i+1, p.Question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.Answer {
			correct++
		}
	}
	return correct
}

func startTimer() {
	time.AfterFunc(10*time.Second, func() {
		fmt.Println("Time is over")
		os.Exit(1)
	})
}
