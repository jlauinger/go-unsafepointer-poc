package main

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	stringResult := GetString()
	fmt.Printf("main:%s\n", stringResult) // expected (but failed) stdout is "abcdefgh"
	bytesResult := GetBytes()
	fmt.Printf("main:%s\n", bytesResult) // expected (but failed) stdout is "abcdefgh
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

func StringToBytes(s string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bytesHeader := reflect.SliceHeader{
		Data: strHeader.Data,
		Cap:  strHeader.Len,
		Len:  strHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bytesHeader))
}

func GetBytes() []byte {
	reader := bufio.NewReader(strings.NewReader("abcdefgh"))
	s, _ := reader.ReadString('\n')
	out := StringToBytes(s)
	// at this point, Go escape analysis incorrectly infers that string s is not used anymore, because it cannot
	// see that out is in fact a reference to it. This in turn stems from the broken reference link within StringToBytes
	// therefore, s is placed on the stack when it really should have been allocated on the heap
	// printing out *within* this function still works well (kind of by accident), but the problem arises when out
	// is returned. After returning from this function its stack will be discarded, and out becomes a dangling pointer,
	// When main uses it it actually reads from invalid memory.
	// Note: this needs the reader call. If s is created as a string literal (s := "abcdefgh"), then main will be able
	// to properly read the abcdefgh from the bytes slice. That is because in the string literal case, s will actually
	// neither live on the heap nor on the stack, it will be placed within the constant data section of the resulting
	// binary, and therefore the address out points to will remain valid after the stack is destroyed.
	fmt.Printf("GetBytes:%s\n", out) // expected stdout is "abcdefgh"
	return out
}