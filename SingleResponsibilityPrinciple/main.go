package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadFile(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func CountLines(data []byte) int {
	count := 0
	for _, b := range data {
		if b == '\n' {
			count++
		}
	}
	return count
}

func main() {
	filename := "example.txt"
	data, err := ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	count := CountLines(data)
	fmt.Printf("El archivo %s tiene %d líneas\n", filename, count)
}
