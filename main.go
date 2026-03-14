package main

import (
	"bytes"
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

	lines := getLinesChannel(file)
	for line := range lines {
		fmt.Printf("Read: %v\n", line)
	}

}
func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		str := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Done reading file content")
					os.Exit(0)
				}
			}

			data = data[:n]
			if nlIndex := bytes.IndexByte(data, '\n'); nlIndex != -1 {
				str += string(data[:nlIndex])
				out <- str
				data = data[nlIndex+1:]
				str = ""
			}

			str += string(data)
		}
	}()
	return out
}
