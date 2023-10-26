package slices

import "fmt"

func Demo1() {
	isimler := make([]string, 3)

	fmt.Println(isimler)

	isimler[0] = "demo"
	isimler[1] = "go"
	isimler[2] = "golang"
	isimler = append(isimler, "demo2")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
