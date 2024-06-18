package xt

type TR_DATA struct {
	RequestID           int32
	DataLength          int32
	TotalDataBufferSize int32
	ElapsedTime         int32
	DataMode            int32
	TrCode              [10]byte
	X_TrCode            [1]byte
	Cont                [1]byte
	ContKey             [18]byte
	X_ContKey           [1]byte
	None                [31]byte
	BlockName           [16]byte
	X_BlockName         [1]byte
	Data                uintptr
}

type MSG_DATA struct {
	RequestID   int32
	SystemError int32
	MsgCode     [5]byte
	X_MsgCode   [1]byte
	MsgLength   int32
	MsgData     uintptr
}

type REALTIME_DATA struct {
	TrCode     [3]byte
	X_TrCode   [1]byte
	KeyLength  int32
	KeyData    [32]byte
	X_KeyData  [1]byte
	RegKey     [32]byte
	X_RegKey   [1]byte
	DataLength int32
	Data       uintptr
}

type TR코드별_전송_제한_정보 struct {
	TR코드         string
	M초당_전송_제한    int
	M초_베이스       int
	M10분당_전송_제한  int
	M10분간_전송한_수량 int
}

type TR코드별_전송_제한_정보_모음 struct {
	M맵 map[string]*TR코드별_전송_제한_정보
}

type S로그인_정보 struct {
	M로그인_ID  string
	M로그인_암호  string
	M인증서_암호  string
	M계좌_비밀번호 string
	M모의투자_암호 string
}
