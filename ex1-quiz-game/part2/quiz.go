package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const defaultQuizName = "problems.csv"

func main() {
	quizFile := flag.String("file", defaultQuizName, "CSV Quiz file name")
	timerLength := flag.Int("timer", 30, "Timer for the quiz")
	randomizeQuiz := flag.Bool("rand", false, "Randomize the quiz order")
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
	if *randomizeQuiz {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(totalQuestions, func(i, j int) {
			questionsAndAnswers[i], questionsAndAnswers[j] = questionsAndAnswers[j], questionsAndAnswers[i]
		})
	}
	correctQuestions := 0
	fmt.Printf("You'll have %d seconds to answer all the questions\nPress enter whenever you're ready to begin the quiz:", *timerLength)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	timer := time.NewTimer(time.Duration(*timerLength) * time.Second)
	defer timer.Stop()
	go func() {
		<-timer.C
		fmt.Println("TIMES UP!")
		printFinalScore(correctQuestions, totalQuestions)
		os.Exit(0)
	}()
	for _, quesAndAns := range questionsAndAnswers {
		fmt.Printf("What is the answer to: %s\n", quesAndAns[0])
		var userInput string
		fmt.Scanf("%s", &userInput)
		if strings.TrimSpace(strings.ToLower(userInput)) == quesAndAns[1] {
			correctQuestions++
		}
	}
	printFinalScore(correctQuestions, totalQuestions)
}

func printFinalScore(correct, total int)  {
	fmt.Println("That's the end of the quiz!")
	fmt.Printf("You scored %d/%d\n", correct, total)
}