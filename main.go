package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	// Generate random number > 0 dan <= 10
	var getInput int
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Tebak angka 0 - 10")
	random_number := rand.Intn(10);
	fmt.Println("Masukkan angka yang anda tebak")	
	fmt.Scanln(&getInput)

	if (getInput > 10) {
		fmt.Println("Maksimal angka adalah 10 - FAILED")
	} else if (getInput < 0) {
		fmt.Println("Minimal angka adalah 0 - FAILED")
	} else if (getInput == random_number) {
		fmt.Println("Jawaban anda benar - SUCCESS")
	} else {
		fmt.Println("==================================")
		fmt.Println("FAILED - Angka yang benar adalah ", random_number)
	}
	
	
}