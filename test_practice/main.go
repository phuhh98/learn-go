package main

import (
	"encoding/json"
	"fmt"
	"os"
	"test_practice/quote"
	"test_practice/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
	bs, _ := json.Marshal(word.UseCount(quote.SunAlso))
	//str := fmt.Sprint(string(bs))

	f, _ := os.Create("sample.txt")
	defer f.Close()
	f.Write(bs)

}
