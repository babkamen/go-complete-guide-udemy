package main

import (
	"fmt"
	"io"
	"os"
)

//print contents of the file
func main() {
	fileName := os.Args[1]
	PrintFile(fileName, os.Stdout)
}

func PrintFile(fileName string, output io.Writer) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("File contents")
	io.Copy(output, file)
}
