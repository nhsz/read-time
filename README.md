# read-time

:clock1::book: Estimate reading time from text file, written in Go

## Usage 

Remember this is a Go script, so you need to [install Go first!](https://golang.org/dl/)

`go run readtime.go <FILE PATH>`

## Example

If both script and file are on the same directory: 

```
> go run readtime.go example.txt
Total words:  455
Reading time: this will take about 1 minute and 39 seconds to be read.
```

Otherwise, you have to provide the full path for text file.
