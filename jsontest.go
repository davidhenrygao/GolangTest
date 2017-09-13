package main

import (
	"encoding/json"
	"fmt"
)

type Em struct {
	F float32 `json:"f"`
}

type Msg struct {
	Code uint64      `json:"code"`
	Body interface{} `json:"body"`
	Em
}

type Body struct {
	I int    `json:"i"`
	S string `json:"s"`
}

type Code struct {
	C uint64 `json:"code"`
}

type BodyBack struct {
	B Body `json:"body"`
}

func jsonTest() {
	fmt.Println()
	//json test

	b1 := Body{1, "fuck you!"}
	msg1 := Msg{1, b1, Em{1.2}}
	b, err := json.Marshal(msg1)
	if err != nil {
		fmt.Printf("json marshal error: %s.\n", err)
		return
	}
	fmt.Printf("msg1 json byte: %s.\n", b)
	fmt.Println()
	var msg2 Msg
	err = json.Unmarshal(b, &msg2)
	fmt.Printf("msg2 : %#v.\n", msg2)
	var code Code
	err = json.Unmarshal(b, &code)
	if err != nil {
		fmt.Printf("json unmarshal error: %s.\n", err)
		return
	}
	fmt.Printf("msg back : %#v.\n", code)
	var body BodyBack
	err = json.Unmarshal(b, &body)
	fmt.Printf("bodyback : %#v.\n", body.B)
}

func init() {
	TestFuncMap["jsonTest"] = jsonTest
}
