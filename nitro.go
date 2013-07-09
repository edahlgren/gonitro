package nitro

//#cgo pkg-config: nitro
//#include "nitro.h"
//#include "wrapper.h"
import "C"
import "unsafe"

type NitroFrame *C.nitro_frame_t
type NitroSocket *C.nitro_socket_t
type NitroSockOpt *C.nitro_sockopt_t

func Start() {
	C.nitro_runtime_start()
}

func Bind(location string) NitroSocket {
	opt := C.nitro_sockopt_new()
	return (NitroSocket)(C.nitro_socket_bind(C.CString(location), opt))
}

func Connect(location string) NitroSocket {
	opt := C.nitro_sockopt_new()
	return (NitroSocket)(C.nitro_socket_connect(C.CString(location), opt))
}

func BytesToFrame(b []byte) NitroFrame {
	unsafePtr := unsafe.Pointer(&b[0])
	size := C.uint32_t(len(b))
	return (NitroFrame)(C.nitro_frame_new_copy(unsafePtr, size))
}

func FrameToBytes(f NitroFrame) []byte {
	cframe := (*C.nitro_frame_t)(f)
	unsafePtr := C.nitro_frame_data(cframe)
	size := C.int(C.nitro_frame_size(cframe))
	return C.GoBytes(unsafePtr,size)
}

func Send(f NitroFrame, s NitroSocket, flags int) {
	C.nitro_send_((*C.nitro_frame_t)(f), (*C.nitro_socket_t)(s), C.int(flags))
}

func Recv(s NitroSocket, flags int) NitroFrame {
	return (NitroFrame)(C.nitro_recv_((*C.nitro_socket_t)(s), C.int(flags)))
}
