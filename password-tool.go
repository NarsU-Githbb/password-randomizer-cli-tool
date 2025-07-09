package main

import (
	//"flag"
	"fmt"
)

func main() {

	upper := []rune{}
	for i := 'A'; i <= 'Z'; i++ {
		upper = append(upper, i)
	}

	lower := []rune{}
	for i := 'a'; i <= 'z'; i++ {
		lower = append(lower, i)
	}

	numnum := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	symbol := []rune{'@', '#', '$', '%', '^', '&', '*', '_', '-', '+', '=', '|', ':', ';', '<', '>', '.', '?', '/', ','}

	fmt.Println(string(upper))
	fmt.Println(string(lower))
	fmt.Println(string(numnum))
	fmt.Println(string(symbol))

	randomizer(lower, upper, numnum, symbol)
}

func randomizer(lower, upper, numnum, symbol []rune) string {

}
