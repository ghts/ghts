package dll32

import (
	"github.com/ghts/ghts/lib"
	nano "github.com/ghts/ghts/lib/nanomsg"
	xt "github.com/ghts/ghts/xing/base"
	"runtime"
	"sync"
	"syscall"
)

// 전역 변수는 항상 동시 액세스로 인한 오류의 위험이 있어서 한 군데 몰아서 관리함.

//goland:noinspection GoUnusedGlobalVariable
var (
	xing_api_dll syscall.Handle
	메시지_윈도우      uintptr

	// Xing API 함수 포인터
	etkConnect     uintptr
	etkIsConnected uintptr
	etkLogin       uintptr
	etkLogout      uintptr
	//etkDisconnect           uintptr	// syscall, cgo 방식 모두 에러 발생.
	etkRequest              uintptr
	etkAdviseRealData       uintptr
	etkUnadviseRealData     uintptr
	etkUnadviseWindow       uintptr
	etkGetAccountListCount  uintptr
	etkGetAccountList       uintptr
	etkGetAccountName       uintptr
	etkGetAccountDetailName uintptr // syscall 방식은 에러 발생. cgo 방식은 정상 작동.
	etkGetAccountNickName   uintptr
	etkGetServerName        uintptr
	etkGetLastError         uintptr
	etkGetErrorMessage      uintptr
	etkGetTRCountPerSec     uintptr
	etkGetTRCountBaseSec    uintptr
	etkGetTRCountLimit      uintptr
	etkGetTRCountRequest    uintptr
	etkReleaseRequestData   uintptr
	etkReleaseMessageData   uintptr
	etkDecompress           uintptr
)

// 다중 사용에 안전한 값들.
var (
	소켓REP_TR수신   = lib.F확인2(nano.NewNano소켓REP(xt.F주소_DLL32()))
	소켓PUB_실시간_정보 = lib.F확인2(nano.NewNano소켓PUB(xt.F주소_실시간()))

	소켓REQ_저장소 = lib.New소켓_저장소(20, func() (lib.I소켓_질의, error) {
		return nano.NewNano소켓REQ(xt.F주소_콜백(), lib.P30초)
	})

	접속_처리_잠금  sync.Mutex
	api_호출_잠금 sync.Mutex

	ch로그인 = make(chan bool, 1)
	Ch질의  = make(chan *lib.S채널_질의, 100)
	ch콜백  = make(chan lib.I콜백, 100)

	수신_도우미_수량 = lib.F최대값(runtime.NumCPU(), 2)
	콜백_도우미_수량 = lib.F최대값(runtime.NumCPU(), 2)

	Ch모니터링_루틴_종료   = make(chan lib.T신호, 1)
	Ch함수_호출_도우미_종료 = make(chan lib.T신호, 1)
	Ch수신_도우미_종료    = make(chan lib.T신호, 수신_도우미_수량)
	Ch콜백_도우미_종료    = make(chan lib.T신호, 콜백_도우미_수량)

	API_초기화_완료 = lib.New안전한_bool(false)
	API_초기화_잠금 = new(sync.Mutex)
)

// 초기화 이후에는 사실상 읽기 전용이어서, 다중 사용에 문제가 없는 값들.
var (
	계좌번호_모음 []string
	계좌_비밀번호 string
	서버_구분   xt.T서버_구분
)
