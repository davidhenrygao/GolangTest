package main

import (
	"fmt"
	"github.com/davidhenrygao/GolangTest/proto/example"
	"github.com/golang/protobuf/proto"
)

func protobufTest() {
	fmt.Println()
	//protobuf golang test

	test := &example.Test{
		Label: proto.String("fuck you!"),
		//Type:  proto.Int32(123),
		//Reps: []int64{3, 0, 6, 2, 4, 7, 0, 0},
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("too hard to use!"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Printf("protobuf Marshal error: %s!\n", err)
		return
	}
	resultTest := &example.Test{}
	err = proto.Unmarshal(data, resultTest)
	if err != nil {
		fmt.Printf("protobuf Unmarshal error: %s!\n", err)
		return
	}
	fmt.Printf("Decode result: \n")
	fmt.Printf("Label : %#v.\n", resultTest.GetLabel())
	fmt.Printf("Type : %#v.\n", resultTest.GetType())
	fmt.Printf("Reps : %#v.\n", resultTest.GetReps())
	fmt.Printf("Optionalgroup : %#v.\n", resultTest.GetOptionalgroup())
	fmt.Printf("resultTest.Type is nil: %#v.\n", resultTest.Type == nil)
}

func init() {
	TestFuncMap["protobufTest"] = protobufTest
}
