package main

import (
	"fmt"
	"strings"
)

func testbytesliceandstring() {
	fmt.Println()

	b := []byte{66, 67, 68, 0, 69, 70}
	s := "12345\000678"
	fmt.Printf("b = %+v\n", b)
	fmt.Printf("s = %+v\n", s)
	eb := []byte(s)
	es := string(b)
	fmt.Printf("eb = %+v\n", eb)
	fmt.Printf("es = %+v\n", es)

	fmt.Printf("len(es) = %+v\n", len(es))
	r := strings.NewReader(es)
	fmt.Printf("Reader(es) : r.Len() = %+v\n", r.Len())
}

func init() {
	TestFuncMap["testbytesliceandstring"] = testbytesliceandstring
}
