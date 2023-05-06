package main

import "fmt"

// type Swimming interface {
// 	Swim()
// }

// type Duck interface {
// 	Swimming
// }

// type Base struct {
// 	Name string
// }

// type Concrete struct {
// 	Base
// }

// func (c Concrete)SayHello(){
// 	fmt.Println(c.Name)
// 	fmt.Println(c.Base.Name)
// 	c.Base.SayHello()
// 	c.SayGoodBye()
// }

// func (b *Base)SayGoodBye(){
// 	fmt.Println("base",b.Name)
// }

// type Parent struct{}

// func (p Parent) SayHello() {
// 	fmt.Println("I am" + p.Name())
// }
// func (p Parent) Name() string {
// 	return "Parent"
// }

// type Son struct {
// 	Parent
// }

// func (s Son) Name() string {
// 	return "Son"
// }

// func main() {
// 	Son := Son{
// 		Parent{},
// 	}
// 	Son.SayHello() //会寻找父类 最后输出(p Parent) Name() result为 Parent
// }




