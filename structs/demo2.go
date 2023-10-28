package structs

import "fmt"

type customer struct {
	name     string
	lastname string
	age      int
}

func (c customer) save() {
	fmt.Println("Eklendi:", c.name)
}

func (c customer) update() {
	fmt.Println("Güncellendi:", c.age)
}

func Demo2() {
	c := customer{name: "Ali", lastname: "Demir", age: 25}
	c.save()
	c.update()
}
