package main

import "fmt"

func main() {
	// Loop untuk nilai i dari 0 hingga 4
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
	}

	// Loop untuk nilai j dari 0 hingga 10
	for j := 0; j <= 10; j++ {
		if j == 5 {
			// Jika nilai j sama dengan 5, maka print karakter pada byte position tertentu
			str := "САШАРВО"
			for a, b := range str {
				fmt.Printf("character %#U starts at byte position %d\n", b, a)
			}
			continue
		}
		fmt.Println("Nilai j =", j)
	}
}
