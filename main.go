package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	succesCounter  = 0
	failCounter    = 0
	fileNameVar    string
	timeOut        int
	startQuiz      string
	totalQuestions int
	randomize      bool
)

// Initializing flags
func init() {
	flag.StringVar(&fileNameVar, "filename", "problems.csv", "file from which to run the quiz")
	flag.IntVar(&timeOut, "time", 30, "time limit of the quiz in seconds")
	flag.BoolVar(&randomize, "rand", false, "Shuffle the quiz order each time it is run")
}

func main() {
	// Parsing flags
	flag.Parse()

	// Welcome message before the quiz
	fmt.Println("WELCOME TO THE QUIZ GAME \nTRY TO ANSWERS TO MANY QUESTION AS YOU CAN.")

	// prompt user to type ok to begin the quiz
	for {
		fmt.Printf("You have %d seconds to complete this quiz, type ok to start the quiz: ", timeOut)
		fmt.Scan(&startQuiz)
		if startQuiz == "ok" {
			fmt.Println("Quiz started...")
			break
		}
	}

	// Opening the quiz csv file
	file, err := os.Open(fileNameVar)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// read all records in the csv file
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	totalQuestions = len(records)
	if randomize {
		records = randomizeSlice(records)
		fmt.Println("randomized", records)
	}
	// start the timer, exit the programme after specified time default set 30 seconds
	go func() {
		time.Sleep(time.Duration(timeOut) * time.Second)
		fmt.Println("\ntime's up!")
		fmt.Printf("You have %d correct answers over %d questions \n", succesCounter, totalQuestions)
		os.Exit(1)
	}()
	// loop to the quiz records
	for _, rec := range records {
		var response string
		fmt.Printf("What is the result of %s = ", rec[0])
		fmt.Scan(&response)
		if strings.Replace(response, " ", "", -1) == strings.Replace(rec[1], " ", "", -1) {
			succesCounter++
			continue
		}
	}
}

// Shuffle the a two dimensional array and return the result
func randomizeSlice(slice [][]string) [][]string {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}
