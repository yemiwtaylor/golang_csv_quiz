package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"time"
)

var questionsAndAnswers [][]string
var totalScore int
var c chan string
var problems []problem

type problem struct {
	q string
	a string
}

func main() {
	questionsAndAnswers = readCsv("problems.csv")
	fmt.Println("Cores: ", runtime.NumCPU())

	problems = parseQuestions(questionsAndAnswers)

	timer := time.NewTimer(2 * time.Second)

	for i := range problems {
		question, correctAnswer := problems[i].q, problems[i].a
		fmt.Println(question)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			c <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Time ran out")
			return
		case answer := <-c:
			if answer == correctAnswer {
				totalScore++
			}
		}
	}
	fmt.Println("Total Scores: ", totalScore, "out of ", len(questionsAndAnswers))
}

func readCsv(fileName string) [][]string {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	qna, _ := reader.ReadAll()
	return qna
}

func parseQuestions(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return problems
}
