package main

import (
	"github.com/edahlgren/gonitro"
	"fmt"
)

func main() {
	nitro.Start()
	bind := nitro.Bind("tcp://127.0.0.1:7723")
	connect := nitro.Connect("tcp://127.0.0.1:7723")

	frame := nitro.BytesToFrame([]byte("What is your name?"))
	nitro.Send(frame, connect, 0)

	frameBack := nitro.Recv(bind, 0)
	msg := string(nitro.FrameToBytes(frameBack))
	fmt.Printf("%v\n", msg)
}


