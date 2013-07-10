package main

import (
	"fmt"
	"time"
	"encoding/binary"
	"crypto/rand"
	"github.com/edahlgren/gonitro"
)

func main() {
	nitro.Start()

	location := "tcp://127.0.0.1:7723"
	go server(location)

	for {
		client(location)
		time.Sleep(1 * time.Second)
	}
}

func server(location string) {
	bind, e := nitro.Bind(location)
	if e != nil {
		fmt.Printf("server: %v\n", e.Error())
	}
	defer nitro.Close(bind)

	callback := func(frameRecvd nitro.NitroFrame, msg []byte) {
		f := nitro.BytesToFrame(msg)
		e := nitro.Reply(bind, frameRecvd, f)
		if e != nil {
			fmt.Printf("server: %v\n", e.Error())
		}
	}

	for {
		f, e := nitro.Recv(bind, nitro.WAIT)
		if e != nil {
			fmt.Printf("server: %v\n", e.Error())
		}
		msg := nitro.FrameToBytes(f)
		go callback(f, doSomething(msg))
	}
}

func doSomething(msg []byte) []byte {
	return []byte(fmt.Sprintf("%v", len(msg)))
}

func randByteArray() []byte {
	var n uint32
        binary.Read(rand.Reader, binary.LittleEndian, &n)
	return []byte(fmt.Sprintf("%v", n))
}

func client(location string) {
	connect, e := nitro.Connect(location)
	if e != nil {
		fmt.Printf("client: %v\n", e.Error())
	}
	defer nitro.Close(connect)

	msg := randByteArray()
	e = nitro.Send(connect, nitro.BytesToFrame(msg))
	if e != nil {
		fmt.Printf("client: %v\n", e.Error())
	}

	f, e := nitro.Recv(connect, nitro.WAIT)
	if e != nil {
		fmt.Printf("client: %v\n", e.Error())
	}
	msgBack := nitro.FrameToBytes(f)
	fmt.Printf("server things %v is %v bytes long\n", string(msg), string(msgBack))
}
