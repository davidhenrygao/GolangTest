package main

import (
	"fmt"
)

type Embedded interface {
	foo()
	bar()
}

type InnerObj struct {
	val int
}

func (r InnerObj) foo() {
	fmt.Printf("foo: %+v.\n", r.val)
}

func (r InnerObj) bar() {
	fmt.Printf("bar: %+v.\n", r.val)
}

type Wrapper struct {
	val int
	Embedded
}

func (r Wrapper) bar() {
	fmt.Printf("bar wrapper: %+v.\n", r.val)
}

func embeddedStructTest() {
	fmt.Println()
	//struct embedded test
	iObj := InnerObj{1}
	var eInf Embedded
	eInf = iObj
	eInf.foo()
	eInf.bar()
	wObj := Wrapper{2, eInf}
	eInf = wObj
	eInf.foo()
	eInf.bar()
	//eInf.Embedded.bar() //compiled error!
	wObj.Embedded.bar()
	wObj2 := Wrapper{3, eInf}
	eInf = wObj2
	eInf.foo()
	eInf.bar()
	wObj2.Embedded.bar()
}

func init() {
	TestFuncMap["embeddedStructTest"] = embeddedStructTest
}
