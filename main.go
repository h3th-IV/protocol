package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("Unable to open file")
	}

	for {
		data := make([]byte, 8)
		n, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Done reading file content")
				os.Exit(0)
			}
		}
		fmt.Printf("Read: %d\n", len(data[:n]))
	}
}
