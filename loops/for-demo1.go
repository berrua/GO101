package loops

import "fmt"

func Demo1() {
	var metin string = "Hello!"
	i := 1

	for i <= 5 {
		fmt.Println(metin)
		i++
	}
	fmt.Println("Bitti.")
}
