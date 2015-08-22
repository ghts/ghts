// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs D:\Go\src\github.com\ghts\ghts\api\NH\wmca_ctype_orig.go

package NH

type ErrBool struct {
	Value     bool
	Pad_cgo_0 [3]byte
	ErrorCode uint32
}

type AccountInfo struct {
	AccountNo       [11]int8
	AccountName     [40]int8
	Act_pdt_cdz3    [3]int8
	Amn_tab_cdz4    [4]int8
	ExpirationDate8 [8]int8
	Granted         int8
	Filler          [189]int8
}
type LoginInfo struct {
	Date         [14]int8
	ServerName   [15]int8
	UserID       [8]int8
	AccountCount [3]int8
	Accountlist  [999]AccountInfo
}
type LoginBlock struct {
	TrIndex    int32
	PLoginInfo *LoginInfo
}

type MsgHeader struct {
	MsgCode [5]int8
	UsrMsg  [80]int8
}

type Received struct {
	PBlockName *int8
	PData      *int8
	Length     int32
}
type OutDataBlock struct {
	TrIndex int32
	PData   *Received
}
