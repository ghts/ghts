// COPIED AND MODIFIED source code at https://github.com/lxn/win

package w32

type HANDLE uintptr
type HWND uintptr
type HINSTANCE uintptr
type HMENU uintptr
type HICON uintptr
type HCURSOR uintptr
type HBRUSH uintptr
type ATOM uint16

type MSG struct {
	HWnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type POINT struct {
	X, Y int32
}

type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}
