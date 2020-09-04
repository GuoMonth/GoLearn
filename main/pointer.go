package main

import "fmt"

func showPointer() {
	var a int = 34
	fmt.Println(&a, a)
	var ip *int
	ip = &a
	(*ip)++
	fmt.Println(ip, a)
}

func showPPointer() {
	var a int = 24
	var ptr *int
	var pptr **int
	var ppptr ***int

	ptr = &a
	pptr = &ptr
	ppptr = &pptr

	fmt.Println("ptr", ptr, *ptr)
	fmt.Println("pptr", pptr, *pptr)
	fmt.Println("ppptr", ppptr, *ppptr)
}
