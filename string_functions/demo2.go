package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	kelime := "golang"
	fmt.Println(s.HasPrefix(kelime, "go"))   // başlangıç
	fmt.Println(s.HasSuffix(kelime, "lang")) // son
	fmt.Println(s.Index(kelime, "lang"))

	harfler := []string{"g", "o", "l", "a", "n", "g"}
	sonuc := s.Join(harfler, "*")
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "|", -1)) // tamamı
	fmt.Println(s.Replace(sonuc, "*", "|", 3))  // 3 tane

	fmt.Println(s.Split(sonuc, "*"))

	fmt.Println(s.Repeat(sonuc, 3))
}
