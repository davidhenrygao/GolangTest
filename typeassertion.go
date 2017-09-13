package main

import (
	"fmt"
	"io"
)

type MyPrinter interface {
	MyPrint()
}

type Counter struct {
	count int
}

func (r *Counter) Incs() {
	r.count++
}

func (r *Counter) Decs() {
	r.count--
}

func (r *Counter) MyPrint() {
	fmt.Printf("count number = %+v\n", r.count)
}

func typeAssertionTest() {
	fmt.Println()
	//type assertion test
	counter := new(Counter)
	counter.Incs()
	counter.Incs()
	counter.Decs()

	var pInf MyPrinter = counter
	pInf.MyPrint()
	//w := pInf.(io.Writer)    //will cause panic! fuck golang!
	w, ok := pInf.(io.Writer)
	if ok {
		fmt.Fprintf(w, "Fuck you!\n")
	} else {
		fmt.Println("Type assertion fails.")
	}
}

func init() {
	TestFuncMap["typeAssertionTest"] = typeAssertionTest
}
