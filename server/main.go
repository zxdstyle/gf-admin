package main

import (
	_ "gf-admin/boot"
	_ "gf-admin/routes"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	g.Server().Run()
}

//type Eatable interface {
//	Eat()
//}
//
//type Animal struct {
//}
//
//func (a *Animal) Eat() {
//	println("Animal eat")
//}
//
//type Cat struct {
//	Animal
//}
//
//func (c *Cat) Eat() {
//	println("Cat eat")
//}
//
//func main() {
//	var eatable Eatable
//	eatable = &Animal{}
//	eatable.Eat()
//
//	eatable = &Cat{}
//	eatable.Eat()
//}
