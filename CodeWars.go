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
