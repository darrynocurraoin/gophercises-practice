package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	)

const defaultQuizName = "problems.csv"

func main() {
	quizFile := flag.String("file", defaultQuizName, "CSV Quiz file name")
	flag.Parse()

	if *quizFile == defaultQuizName {
		*quizFile = "../" + *quizFile
	}
	file, err := os.Open(*quizFile)
	if err != nil {
		fmt.Printf("error reading file: %s\n", err)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	questionsAndAnswers, _ := r.ReadAll()
	totalQuestions := len(questionsAndAnswers)
	correctQuestions := 0
	for _, quesAndAns := range questionsAndAnswers {
		fmt.Printf("What is the answer to: %s\n", quesAndAns[0])
		var userInput string
		fmt.Scanf("%s", &userInput)
		if userInput == quesAndAns[1] {
			correctQuestions++
		}
	}
	fmt.Println("That's the end of the quiz!")
	fmt.Printf("You scored %d/%d\n", correctQuestions, totalQuestions)
}
