// +build windows
// COPIED AND MODIFIED source code at https://github.com/lxn/win

package w32

// Window message constants
const (
	WM_DESTROY = 2
	WM_QUIT    = 18
)

// Predefined window handles
const (
	HWND_MESSAGE = ^HWND(2) // -3
)

// Window message constants
const (
	WM_USER = 1024
)
