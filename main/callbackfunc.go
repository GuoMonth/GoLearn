package main

import "fmt"

type cb func(int) int

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("I am callBack func,%d", x)
	return x + 1
}
