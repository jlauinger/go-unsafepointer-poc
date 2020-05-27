package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

func win() {
	fmt.Println("win!")
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	harmlessData := [8]byte{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'}

	confusedSlice := make([]byte, 512)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&confusedSlice))
	harmlessDataAddress := uintptr(unsafe.Pointer(&(harmlessData[0])))
	sliceHeader.Data = harmlessDataAddress

	_, _ = reader.Read(confusedSlice)

	// avoid optimization of win2
	if harmlessData[0] == 42 {
		win()
	}

	// input a padding and overflow $rip. Use address of win.
	// DEP will prevent to input shell code and jump to it. Instead: use ROP.
}
