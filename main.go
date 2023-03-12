package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
	}

	for j := 0; j < 11; j++ {
		if j == 5 {
			characters := []rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041E'}
			for pos, character := range characters {
				fmt.Printf("character %U '%c' starts at byte position %d \n", character, character, pos*2)
			}
		} else {
			fmt.Println("Nilai j =", j)
		}
	}
}