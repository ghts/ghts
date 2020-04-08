// COPIED AND MODIFIED source code at https://github.com/lxn/win

package w32

import "golang.org/x/sys/windows"

var (
	libuser32       = windows.NewLazySystemDLL("user32.dll")
	dispatchMessage = libuser32.NewProc("DispatchMessageW")
	peekMessage     = libuser32.NewProc("PeekMessageW")
	registerClassEx = libuser32.NewProc("RegisterClassExW")
	createWindowEx  = libuser32.NewProc("CreateWindowExW")
	postQuitMessage = libuser32.NewProc("PostQuitMessage")
	destroyWindow   = libuser32.NewProc("DestroyWindow")
	defWindowProc   = libuser32.NewProc("DefWindowProcW")
	sendMessage     = libuser32.NewProc("SendMessageW")
	postMessage     = libuser32.NewProc("PostMessageW")

	libkernel32     = windows.NewLazySystemDLL("kernel32.dll")
	getModuleHandle = libkernel32.NewProc("GetModuleHandleW")
)
