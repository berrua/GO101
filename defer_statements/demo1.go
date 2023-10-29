package defer_statements

import "fmt"

// lIFO
func A() {
	fmt.Println("A çalıştı.")
}

func C() {
	fmt.Println("C çalıştı.")
}

func D() {
	fmt.Println("D çalıştı.")
}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("B çalıştı.")
}
