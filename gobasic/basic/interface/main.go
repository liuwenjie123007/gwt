package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("(name: %s - age: %s yeas)", p.name, strconv.Itoa(p.age))
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Denis", 70}

	for index, element := range list {
		// .(type) 语法只能在switch中使用
		// 以外的地方只能用comma-ok
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %v\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}
