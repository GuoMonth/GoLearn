package main

import "fmt"

//打印9*9乘法表
func showNineNineTable() {

	const num = 9
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println("")
	}

}

func addNumber() {
	nums := []int{2, 3, 3}
	nums = append(nums, 6)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
}
