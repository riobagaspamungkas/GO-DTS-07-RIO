package main

import (
	"fmt"
	"unicode"
	"strconv"
)

// func isRussianUpper(text string) bool {
//     for _, r := range []rune(text) {
//         if r < '\u0410' || r > '\u042F' {
//             return false
//         }
//     }
//     return true
// }

func main()  {
	i := 21
	var j bool = true

	// menampilkan tipe data dari variabel i
	fmt.Printf("Nilai i adalah %v \n", i)
	fmt.Printf("Tipe data i adalah %T \n", i)
	
	// menampilkan tanda "%"
	fmt.Printf("Menampilkan karakter : ")
	fmt.Println("%")
	
	// menampilkan nilai boolean j
	fmt.Printf("Nilai boolean dari j adalah %t \n", j)
	
	// menampilkan unicode russian
	fmt.Printf("Cek unicode russian Я : ")
	fmt.Println(unicode.Is(unicode.Cyrillic, 'Я'))
	
	// menampilkan nilai base 10 : 21 dan base 8 : 25
	angka := int64(21)
	fmt.Printf("Nilai base 10 : ")
	fmt.Println(strconv.FormatInt(angka, 10))
	fmt.Printf("Nilai base 8 : ")
	fmt.Println(strconv.FormatInt(angka, 8))
	
	// menampilkan nilai base 16 : f
	fmt.Printf("Nilai base 16 : ")
	fmt.Println(strconv.FormatInt(15, 16))
	
	// menampilkan nilai base 16 : F13
	fmt.Printf("Nilai base 16 : ")
	fmt.Println(strconv.FormatInt(3859, 16))

	// menampilkan unicode karakter Я
	fmt.Printf("Nilai karakter unicode Я adalah %U \n", 'Я')

	// float
	var k float64 = 123.456
	fmt.Printf("Float :  %f \n", k)

	// float scientific
	fmt.Printf("Float scientific :  %E \n", k)
}