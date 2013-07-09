package main

import (
	"github.com/edahlgren/gonitro"
	"fmt"
)

func main() {
	nitro.Start()
	bind, e := nitro.Bind("tcp://127.0.0.1:7723")
	if e != nil {
		panic(e.Error())
	}
	connect, e := nitro.Connect("tcp://127.0.0.1:7723")
	if e != nil {
		panic(e.Error())
	}

	frame := nitro.BytesToFrame([]byte("What is your name?"))
	e = nitro.Send(frame, connect, 0)
	if e != nil {
		panic(e.Error())
	}

	frameBack, e := nitro.Recv(bind, 0)
	if e != nil {
		panic(e.Error())
	}
	msg := string(nitro.FrameToBytes(frameBack))
	fmt.Printf("%v\n", msg)
}


