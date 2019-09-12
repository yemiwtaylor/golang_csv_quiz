package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

var totalScore int

type problem struct {
	q string
	a string
}

func main() {
	questionsAndAnswers := readCsv("problems.csv")

	problems := parseQuestions(questionsAndAnswers)

	timer := time.NewTimer(5 * time.Second)

	c := make(chan string)

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
			fmt.Println("Total Scores: ", totalScore, "out of ", len(questionsAndAnswers))
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
