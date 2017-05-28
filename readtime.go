package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filePath := os.Args[1]
	input, err := ioutil.ReadFile(filePath)
	check(err)

	text := string(input)
	fmt.Println("Total words: ", countWords(text))
	fmt.Println("Reading time:", "this will take about", estimateReadingTime(text), "minute/s to be read.")
}

func countWords(text string) int {
	return len(strings.Fields(text))
}

func estimateReadingTime(text string) int {
	/*
		I'm rounding up reading time here,
		based on the average reading speed of an adult (~275 WPM)
	*/
	wordCount := countWords(text)
	readTimeInMinutes := roundUp(float32(wordCount) / 275)

	return readTimeInMinutes
}

/*
	Golang doesn't have a native math.Round implementation yet :/
	See https://github.com/golang/go/issues/4594
*/
func roundUp(a float32) int {
	return int(a + 0.6)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
