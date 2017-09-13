//Note:
//windows 7
//go version 1.8.3
package main

import (
	"fmt"
)

type T struct {
	f int
}

func (rv T) mv() int {
	return rv.f
}
func (rp *T) mp(x int) {
	rp.f = x
}

type Tv struct {
	f1 int
}

func (r Tv) mv() int {
	return r.f1
}

type Tp struct {
	f1 int
}

func (r *Tp) mp() int {
	return r.f1
}

type I interface {
	mv() int
	mp(x int)
}
type TVwrap struct {
	T
	f2 int
}
type TIwrap struct {
	I
	f2 int
}

func methodSetTest() {
	fmt.Println()
	//method set test
	tv := T{1}
	tp := &tv
	fmt.Printf("tv: %#v\n", tv)
	fmt.Printf("tv.mv(): %#v\n", tv.mv())
	fmt.Printf("tp.mv(): %#v\n", tp.mv()) //syntax suger: compiler will transfer tp.mv to (*tp).mv
	tv.mp(2)                              //syntax suger: compiler will transfer tv.mp to (&tv).mp
	fmt.Printf("after tv.mp(2), tv.mv() value: %#v\n", tv.mv())
	tp.mp(3)
	fmt.Printf("after tp.mp(3), tp.mv() value: %#v\n", tp.mv())
	fmt.Printf("T.mv: %#v\n", T.mv)
	fmt.Printf("(*T).mv: %#v\n", (*T).mv)
	//fmt.Printf("T.mp: %#v\n", T.mp) //compile error!
	fmt.Printf("(*T).mp: %#v\n", (*T).mp)
	fmt.Printf("tv.mv: %#v\n", tv.mv)
	fmt.Printf("tp.mv: %#v\n", tp.mv)
	fmt.Printf("tv.mp: %#v\n", tv.mp)
	fmt.Printf("tp.mp: %#v\n", tp.mp)

	/*
		Given a struct type S and a type named T, promoted methods are included in the method set of the struct as follows:
			*If S contains an embedded field T, the method sets of S and *S both include promoted methods with receiver T.
			The method set of *S also includes promoted methods with receiver *T.
			*If S contains an embedded field *T, the method sets of S and *S both include promoted methods with receiver T or *T.
	*/
	tov := Tv{1}
	tovp := &tov
	top := Tp{1}
	topp := &top
	fmt.Printf("Tv.mv: %#v\n", Tv.mv)
	fmt.Printf("(*Tv).mv: %#v\n", (*Tv).mv)
	//fmt.Printf("Tp.mp: %#v\n", Tp.mp) //compile error!
	fmt.Printf("(*Tp).mp: %#v\n", (*Tp).mp)
	fmt.Printf("tov.mv: %#v\n", tov.mv)
	fmt.Printf("tovp.mv: %#v\n", tovp.mv)
	fmt.Printf("top.mp: %#v\n", top.mp)
	fmt.Printf("topp.mp: %#v\n", topp.mp)

	var inf I
	//inf = tv    //compile error!
	inf = tp
	fmt.Printf("inf.mv(): %#v\n", inf.mv())
	inf.mp(4)
	fmt.Printf("after inf.mp(4), inf.mv() value: %#v\n", inf.mv())

	tvw := TVwrap{tv, 5}
	fmt.Printf("tvw: %#v\n", tvw)
	fmt.Printf("tvw.mv(): %#v\n", tvw.mv())
	tvw.mp(6) //compile pass because syntax suger, compiler transfer tvw.mp to (&tvw).mp
	fmt.Printf("after tvw.mp(6), tvw.mv() value: %#v\n", tvw.mv())
	tpw := &tvw
	fmt.Printf("tpw.mv(): %#v\n", tpw.mv())
	tpw.mp(7)
	fmt.Printf("after tpw.mp(7), tpw.mv() value: %#v\n", tpw.mv())
	fmt.Printf("TVwrap.mv: %#v\n", TVwrap.mv)
	fmt.Printf("(*TVwrap).mv: %#v\n", (*TVwrap).mv)
	fmt.Printf("tvw.mv: %#v\n", tvw.mv)
	fmt.Printf("tpw.mv: %#v\n", tpw.mv) //print the same function address as tvw.mv. why?
	//Compare print result, tvw.mv == tpw.mv == tp.mv == tv.mv, Note! compiler transfer all to tv.mv!
	//fmt.Printf("TVwrap.mp: %#v\n", TVwrap.mp)    //compile error!
	fmt.Printf("(*TVwrap).mp: %#v\n", (*TVwrap).mp)
	fmt.Printf("tvw.mp: %#v\n", tvw.mp)
	fmt.Printf("tpw.mp: %#v\n", tpw.mp)
	//Compare print result, tvw.mp == tpw.mp == tv.mp == tp.mp, Note! compiler transfer all to tp.mp!

	//tviw := TIwrap{tvw, 8} //compile error!
	//tviw := TIwrap{tv, 8}   //compile error!
	//tviw := TIwrap{tp, 8} //ok
	tviw := TIwrap{tpw, 8}
	tpiw := &tviw
	fmt.Printf("tviw: %#v\n", tviw)
	fmt.Printf("tviw.mv(): %#v\n", tviw.mv())
	fmt.Printf("tpiw.mv(): %#v\n", tpiw.mv())
	tviw.mp(9)
	fmt.Printf("after tviw.mp(9), tviw.mv() value: %#v\n", tviw.mv())
	tpiw.mp(10)
	fmt.Printf("after tpiw.mp(10), tpiw.mv() value: %#v\n", tpiw.mv())
}

func init() {
	TestFuncMap["methodSetTest"] = methodSetTest
}
