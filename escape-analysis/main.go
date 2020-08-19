package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	stringResult := GetString()
	fmt.Printf("main:%s\n", stringResult) // expected (but failed) stdout is "abcdefgh"
}

func BytesToString(b []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	strHeader := reflect.StringHeader{
		Data: bytesHeader.Data,
		Len:  bytesHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&strHeader))
}

func GetString() string {
	b := []byte{97, 98, 99, 100, 101, 102, 103, 104}
	out := BytesToString(b)
	// At this point, Go escape analysis incorrectly infers that slice b is not used anymore, because it cannot
	// see that out is in fact a reference to it. This in turn stems from the broken reference link within BytesToString
	// therefore, b is placed on the stack when it really should have been allocated on the heap
	// printing out *within* this function still works well (kind of by accident), but the problem arises when out
	// is returned. After returning from this function its stack will be discarded, and out becomes a dangling pointer,
	// When main uses it it actually reads from invalid memory.
	fmt.Printf("GetString:%s\n", out) // expected stdout is "abcdefgh"
	return out
}
