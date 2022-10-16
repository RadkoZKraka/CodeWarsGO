package CodeWarsGO

import (
	"fmt"
	"strings"
)

func main() {
	//var res = wave("siemano elo co tam")hh
	fmt.Println(RemoveChar("witam co tu sie"))
}
func wave(words string) []string {
	// Your code here and happy coding!
	var splice = strings.Split(words, " ")

	return splice

}
func RemoveChar(word string) string {
	//word = word[:len(word)]
	return word[1 : len(word)-1]
}
func CloseCompare(a, b, margin float64) int {
	if a-b > margin {
		return -1
	}
	if a+b > margin {
		return 1
	}
	return 0
}
func Rps(p1, p2 string) string {
	if p1 == scissors && p2 == rock {
		return "Player 2 won!"
	}
	if p1 == scissors && p2 == paper {
		return "Player 1 won!"
	}
	if p1 == paper && p2 == rock {
		return "Player 1 won!"
	}
	return "Draw!"
}

const ( // iota is reset to 0
	scissors = "scissors" // c0 == 0
	rock     = "rock"     // c1 == 1
	paper    = "paper"    // c2 == 2
)
