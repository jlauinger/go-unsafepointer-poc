package main

import (
	"fmt"
	"unsafe"
)

type PinkStruct struct  {
	A uint8
	B int
	C int64
}

type VioletStruct struct  {
	A uint8
	B int64
	C int64
}

func main() {
	pink := PinkStruct{
		A: 1,
		B: 42,
		C: 9000,
	}

	violet := *(*VioletStruct)(unsafe.Pointer(&pink))

	fmt.Println(violet.A)
	fmt.Println(violet.B)
	fmt.Println(violet.C)
}

