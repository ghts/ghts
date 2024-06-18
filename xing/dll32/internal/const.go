package dll32

const (
	xing_dll = "xingAPI.dll"

	FALSE = 0
	TRUE  = 1

	WM_USER                    = 0x0400
	XM_INIT                    = WM_USER
	XM_DISCONNECT              = XM_INIT + 1
	XM_RECEIVE_DATA            = XM_INIT + 3
	XM_RECEIVE_REAL_DATA       = XM_INIT + 4
	XM_LOGIN                   = XM_INIT + 5
	XM_LOGOUT                  = XM_INIT + 6
	XM_TIMEOUT                 = XM_INIT + 7
	XM_RECEIVE_LINK_DATA       = XM_INIT + 8
	XM_RECEIVE_REAL_DATA_CHART = XM_INIT + 10

	RCV_TR_DATA      = 1
	RCV_MSG_DATA     = 2
	RCV_SYSTEM_ERROR = 3
	RCV_RELEASE      = 4
)
