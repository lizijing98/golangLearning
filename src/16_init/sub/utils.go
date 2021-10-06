package sub

import "fmt"

func init() {
	fmt.Println("this is a init() in utils package")
}

func test_1() {
	fmt.Println("this is a test_1 in sub/utils")
}
