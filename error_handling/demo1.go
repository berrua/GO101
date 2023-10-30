package error_handling

import (
	"fmt"
	"os"
)

// type assertion
func Demo1() {
	f, err := os.Open("demo1.txt")
	//nil
	if err != nil {
		// kontrol
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı:", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı.")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
