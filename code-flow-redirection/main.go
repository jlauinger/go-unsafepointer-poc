package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"unsafe"
)

func address(i interface{}) int {
	addr, err := strconv.ParseUint(fmt.Sprintf("%p", i), 0, 0)
	if err != nil {
		panic(err)
	}
	return int(addr)
}

func arrayCopy(dest, src *[64]byte) {
	for i := 0; i < 64; i++ {
		(*dest)[i] = (*src)[i]
	}
}

func win() {
	fmt.Println("win!")
}

type fixedLayout struct {
	exploit [64]byte
	harmlessData [8]byte
}

func main() {
	// use struct to enforce a fixed layout of local variables on the stack to make life easier in GDB
	theData := fixedLayout{
		harmlessData: [8]byte{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'},
	}

	addressBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(addressBuf, uint32(address(win)))

	theData.exploit = [64]byte{
		// padding offset (structured to be recognizable in GDB)
		'A', 'A', 'A', 'A', 'B', 'B', 'B', 'B',
		'C', 'C', 'C', 'C', 'D', 'D', 'D', 'D',
		//'E', 'E', 'E', 'E', 'F', 'F', 'F', 'F',
		addressBuf[0], addressBuf[1], addressBuf[2], addressBuf[3], 0, 0, 0, 0,
		'G', 'G', 'G', 'G', 'H', 'H', 'H', 'H',
		'I', 'I', 'I', 'I', 'J', 'J', 'J', 'J',
		'K', 'K', 'K', 'K', 'L', 'L', 'L', 'L',
		'M', 'M', 'M', 'M', 'N', 'N', 'N', 'N',
		'O', 'O', 'O', 'O', 'P', 'P', 'P', 'P'}

	arrayCopy((*[64]byte)(unsafe.Pointer(&theData.harmlessData)), &theData.exploit)
}
