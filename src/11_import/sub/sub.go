package sub

//package utils // go 中同一个目录层级下不允许出现多个包名

func Sub(a, b int) int {
	test_1() //由于 test_1() 与 Sub() 在同一个包下面，所以可以直接使用，且不需要 sub.
	return a - b
}
