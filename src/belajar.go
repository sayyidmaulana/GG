package main

import "fmt"

func main() {
	fmt.Println("Hello World", "How", "Are", "You", "?")
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(seventh, eight, ninth, fmt, seventh, third, one, isFriday, twoPointTwo, say)
}