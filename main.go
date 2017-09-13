package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

var TestFuncMap = make(map[string]interface{})

func main() {
	var fname string
	var fInf interface{}
	var fValue reflect.Value
	fmt.Println(`Usage: type 'exit' to exit the program;
	type 'list' to list all the test.`)
	for {
		fmt.Fscanln(os.Stdin, &fname)
		if strings.ToLower(fname) == "exit" {
			break
		}
		if strings.ToLower(fname) == "list" {
			fmt.Printf("TestFuncMap = %#v\n", TestFuncMap)
			continue
		}
		fInf = TestFuncMap[fname]
		if fInf == nil {
			fmt.Println("Test function not exist: ", fname)
		} else {
			fValue = reflect.ValueOf(fInf)
			switch fValue.Kind() {
			case reflect.Func:
				fValue.Call(nil)
			default:
				fmt.Println("Error kind: ", fValue.Kind())
			}
		}
	}
}
