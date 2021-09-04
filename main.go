package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happen", err)
		log.Println("err happen (log):", err)
		//log.Fatalln(err)
		log.Panic(err) // panic println to stdout more info (stack trace) than just cal Fatal
	}
}
