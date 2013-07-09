package nitro

//#cgo pkg-config: nitro
//#include "nitro.h"
//#include "wrapper.h"
import "C"
import "unsafe"
import "errors"

type NitroFrame *C.nitro_frame_t
type NitroSocket *C.nitro_socket_t
type NitroSockOpt *C.nitro_sockopt_t

func Start() {
	C.nitro_runtime_start()
}

func NitroError() error {
	e := C.nitro_error()
	msg := C.GoString(C.nitro_errmsg(e))
	return errors.New(msg)
}

func Bind(location string) (NitroSocket, error) {
	opt := C.nitro_sockopt_new()
	if opt == nil {
		return nil, NitroError()
	}
	return (NitroSocket)(C.nitro_socket_bind(C.CString(location), opt)), nil
}

func Connect(location string) (NitroSocket, error) {
	opt := C.nitro_sockopt_new()
	if opt == nil {
		return nil, NitroError()
	}
	return (NitroSocket)(C.nitro_socket_connect(C.CString(location), opt)), nil
}

// add fininalizer to destroy nitro frame
func BytesToFrame(b []byte) NitroFrame {
	unsafePtr := unsafe.Pointer(&b[0])
	size := C.uint32_t(len(b))
	return (NitroFrame)(C.nitro_frame_new_copy(unsafePtr, size))
}

// make a copy of the bytestring
func FrameToBytes(f NitroFrame) []byte {
	cframe := (*C.nitro_frame_t)(f)
	unsafePtr := C.nitro_frame_data(cframe)
	size := C.int(C.nitro_frame_size(cframe))
	return C.GoBytes(unsafePtr,size)
}

func Send(f NitroFrame, s NitroSocket, flags int) error {
	e := C.nitro_send_((*C.nitro_frame_t)(f), (*C.nitro_socket_t)(s), C.int(flags))
	if e < 0 {
		return NitroError()
	}
	return nil
}

// add finalizer to destroy nitro frame
func Recv(s NitroSocket, flags int) (NitroFrame, error) {
	f := (NitroFrame)(C.nitro_recv_((*C.nitro_socket_t)(s), C.int(flags)))
	if f == nil {
		return nil, NitroError()
	}
	return f, nil
}
