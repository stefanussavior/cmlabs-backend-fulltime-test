package main

import (
	"fmt"
	"strings"
)

func Palindrome(input string) bool {
	kataAwal := strings.ToLower(input)
	for i := 0; i < len(kataAwal)/2; i++ {
		if kataAwal[i] != kataAwal[len(kataAwal)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var input string

	fmt.Println(Palindrome("SAIPPUAKIVIKAUPPIAS"))
	fmt.Println(Palindrome("Aibohphobia"))
	fmt.Println(Palindrome("Anna"))
	fmt.Println(Palindrome("Civic"))

	fmt.Println(Palindrome("My gym"))
	fmt.Println(Palindrome("No lemon, no melon"))

	fmt.Println("Masukkan Kata : ")
	fmt.Scanln(&input)
	fmt.Println(Palindrome(input))

}
