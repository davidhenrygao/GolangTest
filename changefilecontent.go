//Note:
//windows 7
//go version 1.8.3
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	//"runtime"
	"strings"
)

var LINE string = "\n"

const (
	PACKAGE           = "package "
	OPTION_GO_PACKAGE = "option go_package = "
	GO_PACKAGE        = "\"github.com/davidhenrygao/LoginTest/proto/login\";"
)

/*
func judgeOS() {
	fmt.Println("Current OS: ", runtime.GOOS)
	if runtime.GOOS == "windows" {
		LINE = "\r\n"
	} else {
		LINE = "\n"
	}
}
*/

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Printf("Close file(%s) err = %+v\n", f.Name(), err)
	}
}

func changefilecontent() {
	fmt.Println()

	//delete content
	origfilename := "proto/common/extrainfo.proto"
	origF, err := os.Open(origfilename)
	if err != nil {
		fmt.Printf("Open file err = %+v\n", err)
		return
	}
	fileinfo, err := origF.Stat()
	if err != nil {
		fmt.Printf("Get origin file stat err = %+v\n", err)
		return
	}
	fmt.Println("Original file base name: ", fileinfo.Name())
	defer closeFile(origF)

	csF, err := os.Create("exchangekey.cs.proto")
	if err != nil {
		fmt.Printf("Create charp proto file err = %+v\n", err)
		return
	}
	defer closeFile(csF)

	goF, err := os.Create("exchangekey.go.proto")
	if err != nil {
		fmt.Printf("Create charp proto file err = %+v\n", err)
		return
	}
	defer closeFile(goF)

	reader := bufio.NewReader(origF)
	var lBuf bytes.Buffer
	var csBuf bytes.Buffer
	var goBuf bytes.Buffer
	var flag bool = true
	var packagelineflag bool = true
	var line []byte
	for flag {
		line, err = reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				flag = false
			} else {
				fmt.Printf("Original file readline err = %+v\n", err)
				return
			}
		}
		_, err = lBuf.Write(line)
		if err != nil {
			fmt.Printf("Buffer Write line content err = %+v\n", err)
			return
		}
		str := lBuf.String()
		if packagelineflag {
			str = strings.TrimLeft(str, " \t")
			if len(str) > len(PACKAGE) {
				r := strings.Compare(PACKAGE, str[:len(PACKAGE)])
				if r == 0 {
					packagelineflag = false
					_, err = goBuf.Write(lBuf.Bytes())
					if err != nil {
						fmt.Printf("Buffer Write go file buffer package line err = %+v\n", err)
						return
					}
					_, err = goBuf.Write([]byte(LINE + OPTION_GO_PACKAGE + GO_PACKAGE + LINE))
					if err != nil {
						fmt.Printf("Buffer Write go file buffer option line err = %+v\n", err)
						return
					}
					lBuf.Truncate(0)
				}
			}
		}
		_, err = goBuf.Write(lBuf.Bytes())
		if err != nil {
			fmt.Printf("Buffer Write go file buffer err = %+v\n", err)
			return
		}
		_, err = csBuf.Write(lBuf.Bytes())
		if err != nil {
			fmt.Printf("Buffer Write csharp file buffer err = %+v\n", err)
			return
		}
		lBuf.Truncate(0)
	}

	_, err = goF.Write(goBuf.Bytes())
	if err != nil {
		fmt.Printf("Write go file error = %+v\n", err)
		return
	}
	err = goF.Sync()
	if err != nil {
		fmt.Printf("Sync go file error = %+v\n", err)
		return
	}
	_, err = csF.Write(csBuf.Bytes())
	if err != nil {
		fmt.Printf("Write csharp file error = %+v\n", err)
		return
	}
	err = csF.Sync()
	if err != nil {
		fmt.Printf("Sync csharp file error = %+v\n", err)
		return
	}

	fmt.Println("Finish changing.")

	return
}

func init() {
	TestFuncMap["changefilecontent"] = changefilecontent
	//judgeOS()
}
