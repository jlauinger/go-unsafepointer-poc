package main

import (
	"fmt"
	"unsafe"
)

func main() {
	harmlessData := [8]byte{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'}
	// might be e.g. private key data
	secret := [17]byte{'l', '3', '3', 't', '-', 'h', '4', 'x', 'x', '0', 'r', '-', 'w', '1', 'n', 's', '!'}

	// read from memory behind buffer
	var leakingInformation = (*[8+17]byte)(unsafe.Pointer(&harmlessData[0]))

	fmt.Println(string((*leakingInformation)[:]))

	// avoid optimization of variable
	if secret[0] == 42 {
		fmt.Println("do not optimize secret")
	}
}
