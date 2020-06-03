package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

func main() {
	go heapHeapHeap()

	readAndHaveFun()
}

func unsafeStringToBytes(s *string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(s))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}))
}

func readAndHaveFun() {
	reader := bufio.NewReader(os.Stdin)
	count := 1
	var firstChar byte

	for {
		s, _ := reader.ReadString('\n')
		if len(s) == 0 {
			continue
		}
		firstChar = s[0]

		// HERE BE DRAGONS
		bytes := unsafeStringToBytes(&s)

		_, _ = reader.ReadString('\n')

		if len(bytes) > 0 && bytes[0] != firstChar {
			fmt.Printf("win! after %d iterations\n", count)
			os.Exit(0)
		}

		count++
	}
}

func heapHeapHeap() {
	var a *[]byte
	for {
		tmp := make([]byte, 1000000, 1000000)
		a = &tmp
		_ = a
	}
}