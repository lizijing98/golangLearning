package main

import "fmt"

func main() {
	p1 := testPtr1()
	fmt.Println("p1:", *p1)
	fmt.Println("p1 ptr:", p1)
}

func testPtr1() *string {
	name := "Lizijing"
	ptr0 := &name
	fmt.Println("*ptr0:", *ptr0)
	fmt.Println("ptr0 ptr:", ptr0)
	city := "南京"
	cityPtr := &city
	return cityPtr
}
