package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

func test64bitdataendian() {
	fmt.Println()

	var x uint64 = 0x0102030405060708
	fmt.Printf("variable x = %#016x(hex)\n", x)
	fmt.Printf("first byte of x is %#v.\n", *(*uint8)(unsafe.Pointer((uintptr(unsafe.Pointer(&x)) + 0))))
	fmt.Printf("first byte of x is %#v.\n", *(*uint8)(unsafe.Pointer((uintptr(unsafe.Pointer(&x)) + 1))))
	fmt.Printf("first byte of x is %#v.\n", *(*uint8)(unsafe.Pointer((uintptr(unsafe.Pointer(&x)) + 4))))

	b := make([]byte, 8)
	l := make([]byte, 8)
	binary.BigEndian.PutUint64(b, x)
	binary.LittleEndian.PutUint64(l, x)
	fmt.Printf("BigEndian byte of x: %+v\n", b)
	fmt.Printf("LittleEndian byte of x: %+v\n", l)

}

func init() {
	TestFuncMap["test64bitdataendian"] = test64bitdataendian
}
