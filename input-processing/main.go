package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	textFile := "mobydick.txt"

	if len(os.Args) > 1 {
		textFile = os.Args[1]
	}

	f, err := os.Open(textFile)
	defer f.Close()
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF{
				return
			}
			log.Fatalf("Error reading line: %s", err)
		}
		if bytes.Contains(line, []byte("ERROR")) {
			log.Printf(string(line))
		}
	}

}
