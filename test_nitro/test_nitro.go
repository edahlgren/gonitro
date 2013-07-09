package main

import (
	"fmt"
	"github.com/edahlgren/gonitro"
)

func main() {
	nitro.Start()

	location := "tcp://127.0.0.1:7723"
	bind, e := nitro.Bind(location)
	if e != nil {
		panic(e.Error())
	}
	connect, e := nitro.Connect(location)
	if e != nil {
		panic(e.Error())
	}

	testSendReceive(bind, connect)
	fmt.Printf("ok: testSendReceive passed\n")
	testReply(bind, connect)
	fmt.Printf("ok: testReply passed\n")
	testRelay(bind, connect)
	fmt.Printf("ok: testRelay passed\n")
}

func testSendReceive(bind nitro.NitroSocket, connect nitro.NitroSocket) {
	msg := "What is your name?"
	msgBack, e := SendReceive(bind, connect, msg)
	if e != nil {
		panic(e.Error())
	}
	if msg != msgBack {
		panic("SendReceive: strings don't match")
	}
}

func SendReceive(bind nitro.NitroSocket, connect nitro.NitroSocket, msg string) (string, error) {
	frame := nitro.BytesToFrame([]byte(msg))
	e := nitro.Send(connect, frame)
	if e != nil {
		return "", e
	}

	frameBack, e := nitro.Recv(bind, nitro.WAIT)
	if e != nil {
		return "", e
	}
	return string(nitro.FrameToBytes(frameBack)), nil
}

func testReply(bind nitro.NitroSocket, connect nitro.NitroSocket) {
	msg := "Sir Lancelot of Camelot"
	msgBack, e := ReplyHelper(bind, connect, msg)
	if e != nil {
		panic(e.Error())
	}
	if msg != msgBack {
		panic("Reply: strings don't match")
	}
}

func ReplyHelper(bind nitro.NitroSocket, connect nitro.NitroSocket, msg string) (string, error) {
	frame := nitro.BytesToFrame([]byte("What is your name?"))
	e := nitro.Send(connect, frame)
	if e != nil {
		return "", e
	}

	frameBack, e := nitro.Recv(bind, nitro.WAIT)
	if e != nil {
		return "", e
	}

	e = nitro.Reply(bind, frameBack, nitro.BytesToFrame([]byte(msg)))
	if e != nil {
		return "", e
	}

	frameResp, e := nitro.Recv(connect, nitro.WAIT)
	if e != nil {
		return "", e
	}
	str := string(nitro.FrameToBytes(frameResp))
	return str, nil
}

func testRelay(bind nitro.NitroSocket, connect nitro.NitroSocket) {
	fwlocation := "tcp://127.0.0.1:7724"
	fwbind, e := nitro.Bind(fwlocation)
	if e != nil {
		panic(e.Error())
	}
	fwconnect, e := nitro.Connect(fwlocation)
	if e != nil {
		panic(e.Error())
	}

	msg := "Look, you stupid Bastard. You've got no arms left."
	msgFw, e := RelayHelper(bind, connect, fwbind, fwconnect, msg)
	if e != nil {
		panic(e.Error())
	}
	if msg != msgFw {
		panic("RelayFw: strings don't match")
	}

	nitro.Close(fwbind)
	nitro.Close(fwconnect)
}

func RelayHelper(bind nitro.NitroSocket, connect nitro.NitroSocket, fwbind nitro.NitroSocket, fwconnect nitro.NitroSocket, msg string) (string, error) {
	/**
	 original sender: connected to 7723
	 proxy:           bound to 7723
	                  connected to 7724
	 peer receiver:   bound to 7724
	 **/

	// original sender
	e := nitro.Send(connect, nitro.BytesToFrame([]byte(msg)))
	if e != nil {
		return "", e
	}

	// proxy
	frame, e := nitro.Recv(bind, nitro.WAIT)
	if e != nil {
		return "", e
	}
	e = nitro.RelayFw(fwconnect, frame, frame)
	if e != nil {
		return "", e
	}

	// peer receiver
	frameFw, e := nitro.Recv(fwbind, nitro.WAIT)
	if e != nil {
		return "", e
	}
	e = nitro.Reply(fwbind, frameFw, frameFw)
	if e != nil {
		return "", e
	}

	// proxy
	frameBack, e := nitro.Recv(fwconnect, nitro.WAIT)
	if e != nil {
		return "", e
	}
	e = nitro.RelayBk(bind, frameBack, frameBack)
	if e != nil {
		return "", e
	}

	// original sender
	frame, e = nitro.Recv(connect, nitro.WAIT)
	if e != nil {
		return "", e
	}
	str := string(nitro.FrameToBytes(frame))
	return str, nil
}



