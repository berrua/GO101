package string_functions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	kelime := "golang"
	fmt.Println(s.Count(kelime, "g"))
	fmt.Println(s.Contains(kelime, "e"))
	fmt.Println(s.Index(kelime, "a"))

	sonuc := s.Index(kelime, "w")
	if sonuc != -1 {
		fmt.Println("W harfi var.")
	} else {
		fmt.Println("W harfi yok.")
	}

	fmt.Println(s.ToLower(kelime))
	fmt.Println(s.ToUpper(kelime))
}
