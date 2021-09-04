package main

import (
	"fmt"
)

func main() {

	var ans1, ans2, ans3 string
	fmt.Print("Name: ")
	_, err := fmt.Scan(&ans1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav food: ")
	// Scan scans text read from standard input, storing successive space-separated values
	//into successive arguments. Newlines count as space. It returns the number of items
	//successfully scanned. If that is less than the number of arguments, err will report why.
	_, err = fmt.Scan(&ans2)
	if err != nil {
		panic(err)
	}
	fmt.Print("Fav sport: ")
	_, err = fmt.Scan(&ans3)
	if err != nil {
		panic(err)
	}

	fmt.Println(ans1, ans2, ans3)
}
