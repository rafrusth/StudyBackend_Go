package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Deklarasi
	var username string = "presidentSkroob"
	var password int = 12345
	passwordString := strconv.Itoa(password)

	// Algoritma
	// don't edit below this line
	fmt.Println("Authorization: Basic", username+":"+passwordString)
}
