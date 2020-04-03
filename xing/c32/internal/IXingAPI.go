package x32

// #include "IXingAPI.h"
import "C"
import "syscall"

func runIXingAPITest() {
	xing_api_dll = syscall.Handle(uintptr(C.GetSafeHandle()))
	//C.SetSafeHandle(unsafe.Pointer(xing_api_dll))
	C.CheckAccountFunctions()
}