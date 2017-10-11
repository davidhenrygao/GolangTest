package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/davidhenrygao/GolangTest/proto/test"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
)

var SessionCnt int32 = 0

type MsgHead struct {
	Session  int32
	Cmd      int32
	DataLen  int16
	Checksum int32
}

type Message struct {
	MsgHead
	Data []byte
}

func Marshal(v interface{}) []byte {
	data, err := proto.Marshal(v.(proto.Message))
	if err != nil {
		fmt.Printf("Protobuf Marshal error: %s.\n", err)
		return nil
	}
	return data
}

func Unmarshal(b []byte, v interface{}) error {
	err := proto.Unmarshal(b, v.(proto.Message))
	return err
}

func binWriteMessage(w io.Writer, msg Message) error {
	err := binary.Write(w, binary.BigEndian, msg.Session)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, msg.Cmd)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, msg.DataLen)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, msg.Checksum)
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.BigEndian, msg.Data)
	if err != nil {
		return err
	}
	return nil
}

func binReadMessage(r io.Reader) (Message, error) {
	var msg Message
	err := binary.Read(r, binary.BigEndian, &msg.Session)
	if err != nil {
		return Message{}, err
	}
	err = binary.Read(r, binary.BigEndian, &msg.Cmd)
	if err != nil {
		return Message{}, err
	}
	err = binary.Read(r, binary.BigEndian, &msg.DataLen)
	if err != nil {
		return Message{}, err
	}
	err = binary.Read(r, binary.BigEndian, &msg.Checksum)
	if err != nil {
		return Message{}, err
	}
	msg.Data = make([]byte, msg.DataLen)
	err = binary.Read(r, binary.BigEndian, &msg.Data)
	if err != nil {
		return Message{}, err
	}
	return msg, nil
}

func sendReq(conn net.Conn, cmd test.CMD, data []byte) error {
	msghead := MsgHead{
		Session:  SessionCnt,
		Cmd:      int32(cmd),
		DataLen:  int16(len(data)),
		Checksum: 0,
	}
	msg := Message{
		MsgHead: msghead,
		Data:    data,
	}
	buf := new(bytes.Buffer)
	err := binWriteMessage(buf, msg)
	if err != nil {
		err = fmt.Errorf("Binary pack message error: %s.\n", err)
		return err
	}
	buflen := buf.Len()
	if buflen >= 2<<16 {
		err = fmt.Errorf("length(%d) exceeds %d", buflen, 2<<16)
		return err
	}
	ul := uint16(buflen)
	b := make([]byte, 2+ul)
	binary.BigEndian.PutUint16(b[:2], ul)
	copy(b[2:], buf.Bytes())

	_, err = conn.Write(b)
	if err != nil {
		err = fmt.Errorf("Connection write error: %s.\n", err)
		return err
	}

	SessionCnt++
	return nil
}

func RecvResp(conn net.Conn) (Message, error) {
	len := make([]byte, 2)
	_, err := io.ReadFull(conn, len)
	if err != nil {
		fmt.Printf("RecvResp read message length error: %s.\n", err)
		return Message{}, err
	}
	ul := binary.BigEndian.Uint16(len)
	b := make([]byte, ul)
	_, err = io.ReadFull(conn, b)
	if err != nil {
		fmt.Printf("RecvResp read message error: %s.\n", err)
		return Message{}, err
	}
	//fmt.Printf("resp origin stream: %#v.\n", b)
	buf := bytes.NewReader(b)
	msg, err := binReadMessage(buf)
	if err != nil {
		fmt.Printf("RecvResp binary unpack error: %s.\n", err)
		return Message{}, err
	}
	fmt.Printf("Recv msg head: { Session: %#v, Cmd: %#v, DataLen: %#v, Checksum: %#v }.\n",
		msg.Session, msg.Cmd, msg.DataLen, msg.Checksum)
	//fmt.Printf("Recv msg data: %#v.\n", msg.Data)
	return msg, nil
}

func doProcess(conn net.Conn, req interface{}) {
	var data []byte
	data = Marshal(req)
	if data == nil {
		return
	}
	fmt.Println("Send login request to server.")
	err := sendReq(conn, test.CMD_LOGIN, data)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	msg, err := RecvResp(conn)
	if err != nil {
		return
	}
	result := msg.Data
	resp := &test.Resp{}
	err = Unmarshal(result, resp)
	if err != nil {
		fmt.Printf("Protobuf Unmarshal error: %s.\n", err)
		return
	}
	fmt.Println("Receive login response from server.")
	if resp.GetRc() != 0 {
		//if resp.GetRc() != test.RetCode_SUCCESS {
		fmt.Printf("login error code: %#v.\n", resp.GetRc())
	} else {
		if resp.RespBody != nil {
			respbody := resp.GetRespBody()
			if respbody.Secrect != nil {
				fmt.Printf("login return secrect: %#v.\n", respbody.GetSecrect())
			}
			if respbody.MsgBack != nil {
				fmt.Printf("login return msg: %#v.\n", respbody.GetMsgBack())
			}
			fmt.Printf("login return options: %#v. \n", respbody.GetOptions())
		}
	}
}

func protobufclientTest() {
	fmt.Println()

	fmt.Println("Connect to server(192.168.2.188:10086).")
	conn, err := net.Dial("tcp", "192.168.2.188:10086")
	if err != nil {
		fmt.Printf("Connect to server error: %s.\n", err)
		return
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	req := &test.Req{
		Account:  proto.String("David"),
		Password: proto.String("fuck you"),
		Msg:      proto.String("Girl"),
	}
	doProcess(conn, req)
}

func init() {
	TestFuncMap["protobufclientTest"] = protobufclientTest
}
