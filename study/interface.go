package study

import "fmt"

func InterfaceTest() {
	bb := Woman{}
	fmt.Println(bb.name(), bb.age())
	cc := Men{}
	fmt.Println(cc.name(), cc.age())
}

type Man interface {
	name() string
	age() int
}

type Woman struct {
}

type Men struct {
}

func (woman Woman) name() string {
	return "范冰冰"
}

func (men Men) name() string {
	return "李晨"
}

func (woman Woman) age() int {
	return 22
}

func (men Men) age() int {
	return 33
}
