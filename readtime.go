package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type readTime struct {
	minutes int
	seconds int
}

func main() {
	filePath := os.Args[1]
	input, err := ioutil.ReadFile(filePath)
	check(err)

	text := string(input)
	wordCount := countWords(text)
	estimatedTime := estimateReadingTime(text, wordCount)

	minuteOrMinutes := isItMinuteOrMinutes(estimatedTime.minutes)

	fmt.Println("Total words: ", wordCount)
	fmt.Println("Reading time:", "this will take about", estimatedTime.minutes,
		minuteOrMinutes, "and", estimatedTime.seconds, "seconds to be read.")
}

func countWords(text string) int {
	return len(strings.Fields(text))
}

/*
	I'm rounding up reading time here, based on
	the average reading speed of an adult (~275 WPM)

	Src: https://help.medium.com/hc/en-us/articles/214991667-Read-time
*/
func estimateReadingTime(text string, wordCount int) readTime {
	rawTime := float64(wordCount) / float64(275)
	// estimated time in minutes
	integerPart := int(rawTime)
	// estimated time in seconds
	decimalPart := (rawTime - float64(integerPart)) * 60

	readingTime := readTime{
		minutes: roundUp(float64(integerPart)),
		seconds: roundUp(decimalPart)}

	return readingTime
}

/*
	Golang doesn't have a native math.Round implementation yet :/
	Src: https://github.com/golang/go/issues/4594
*/
func roundUp(a float64) int {
	return int(a + 0.5)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isItMinuteOrMinutes(timeInMinutes int) string {
	minuteOrMinutes := "minutes"

	if timeInMinutes == 1 {
		minuteOrMinutes = "minute"
	}

	return minuteOrMinutes
}
