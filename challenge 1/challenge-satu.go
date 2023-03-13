package main

import "fmt"

func main() {
	i := 21
	fmt.Printf("%v \n", i) // menampilkan nilai i : 21
	fmt.Printf("%T \n", i) // menampilkan tipe data dari variabel i
	fmt.Printf("%% \n")    // menampilkan tanda %

	j := true
	fmt.Printf("%v \n\n", j) // menampilkan nilai boolean j : true
	fmt.Printf("%b \n", i)   // menampilkan nilai base 2 dari i : 10101

	fmt.Printf("%c\n", '\u042F') // menampilkan unicode russia : Я

	base10 := 21
	fmt.Printf("%d \n", base10) // menampilkan nilai base 10 dari i : 21

	base8 := 25
	fmt.Printf("%o \n", base8) // menampilkan nilai base 8 dari i : 25

	base16 := 15
	fmt.Printf("%x \n", base16) // menampilkan nilai base 16 : f --> base 16 dari f itu angka berapa?
	fmt.Printf("%X \n", base16) // menampilkan nilai base 16 : F --> base 16 dari F itu angka berapa?

	unicode := 'Я'
	fmt.Printf("%U \n\n", unicode) // menampilkan unicode karakter Я (bahasa rusia "ya") : U+042F

	k := 123.456
	fmt.Printf("%f \n", k) // menampilkan float dari k : 123.456000
	fmt.Printf("%e \n", k) // menampilkan float scientific dari k: 1.234560E+02
}
