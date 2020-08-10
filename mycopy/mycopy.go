package main

import (
	"log"
	"os"
)

var Path string = "./test_file"

func testFiles() {
	text := "test file"
	//b := make([]byte, 1024*1024) // заполнен нулями
	file, _ := os.Create(Path)
	_, err := file.Write([]byte(text))
	if err != nil {
		log.Panicf("failed to write: %v", err)
	}
	err = file.Close() // что бы очистить буферы ОС
	if err != nil {
		log.Panicf("failed to close: %v", err)
	}
}

func main() {
	testFiles()
}