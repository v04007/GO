package main

import "fmt"

func main() {
	println("hello")
	str := "hello world!"
	fmt.Printf("%d\n", str) //go vet   检查发现可能的bug或者可疑的构造

	m := make(map[string]string)
	m["a"] = "A"
	fmt.Println(m)
}

func add(a, b int) int {
	return a + b
}
