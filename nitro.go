package nitro

//#cgo pkg-config: nitro
//#include "nitro.h"
//#include "wrapper.h"
import "C"
import "unsafe"
import "errors"
import "runtime"

// NitroFrame is a go object so we can call SetFinalizer on it
type NitroFrame struct {
	ptr *C.nitro_frame_t
}

func (f *NitroFrame) Free() {
	C.nitro_frame_destroy_(f.ptr)
}

type NitroSocket *C.nitro_socket_t

func Close(s NitroSocket) {
	C.nitro_socket_close((*C.nitro_socket_t)(s))
}

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

// add a fininalizer to destroy nitro frame
func BytesToFrame(b []byte) NitroFrame {
	unsafePtr := unsafe.Pointer(&b[0])
	size := C.uint32_t(len(b))
	f := NitroFrame{(C.nitro_frame_new_copy(unsafePtr, size))}
	runtime.SetFinalizer(&f, (*NitroFrame).Free)
	return f
}

// make a copy of the bytestring
func FrameToBytes(f NitroFrame) []byte {
	fptr := (*C.nitro_frame_t)(f.ptr)
	unsafePtr := C.nitro_frame_data(fptr)
	size := C.int(C.nitro_frame_size(fptr))
	// GoBytes copies the original data:
	// https://code.google.com/p/go-wiki/wiki/cgo#Go_strings_and_C_strings
	return C.GoBytes(unsafePtr,size)
}

const (
	NOFLAG = iota
	REUSE
	NOWAIT
)
const WAIT = NOFLAG

// send with reuse so we can garabage collect it
// don't expose the flags api, because nowait doesn't make sense here and we should always use reuse
func Send(s NitroSocket, f NitroFrame) error {
	fptr := (*C.nitro_frame_t)(f.ptr)
	sptr := (*C.nitro_socket_t)(s)
	e := C.nitro_send_(fptr, sptr, C.int(REUSE))
	if e < 0 {
		return NitroError()
	}
	return nil
}

// add finalizer to destroy nitro frame
func Recv(s NitroSocket, flag int) (NitroFrame, error) {
	sptr := (*C.nitro_socket_t)(s)
	fptr := C.nitro_recv_(sptr, C.int(flag))
	f := NitroFrame{fptr}
	if fptr == nil {
		return f, NitroError()
	}
	runtime.SetFinalizer(&f, (*NitroFrame).Free)
	return f, nil
}

func Reply(s NitroSocket, recvd NitroFrame, resp NitroFrame) error {
	e := C.nitro_reply_(
		(*C.nitro_frame_t)(recvd.ptr),
		(*C.nitro_frame_t)(resp.ptr),
		(*C.nitro_socket_t)(s),
		C.int(REUSE))
	if e < 0 {
		return NitroError()
	}
	return nil
}

func RelayFw(s NitroSocket, recvd NitroFrame, resp NitroFrame) error {
	e := C.nitro_relay_fw_(
		(*C.nitro_frame_t)(recvd.ptr),
		(*C.nitro_frame_t)(resp.ptr),
		(*C.nitro_socket_t)(s),
		C.int(REUSE))
	if e < 0 {
		return NitroError()
	}
	return nil
}

func RelayBk(s NitroSocket, recvd NitroFrame, resp NitroFrame) error {
	e := C.nitro_relay_bk_(
		(*C.nitro_frame_t)(recvd.ptr),
		(*C.nitro_frame_t)(resp.ptr),
		(*C.nitro_socket_t)(s),
		C.int(REUSE))
	if e < 0 {
		return NitroError()
	}
	return nil
}

