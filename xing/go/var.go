package xing

import (
	lb "github.com/ghts/ghts/lib"
	nano "github.com/ghts/ghts/lib/nanomsg"
	xt "github.com/ghts/ghts/xing/base"
	"runtime"
	"sync"
	"time"
)

var (
	소켓REP_TR콜백 lb.I소켓with컨텍스트
	소켓REQ_저장소  = lb.New소켓_저장소(20, func() (lb.I소켓_질의, error) {
		return nano.NewNano소켓REQ(xt.F주소_DLL32(), lb.P30초)
	})
	//소켓SUB_실시간_정보 lb.I소켓

	ch신호_DLL32_초기화 = make(chan lb.T신호_32비트_모듈, 1)
	ch신호_DLL32_로그인 = make(chan lb.T신호_32비트_모듈, 1)
	ch신호_DLL32_종료  = make(chan lb.T신호_32비트_모듈, 1)

	대기소_DLL32 = newDLL32_콜백_대기_저장소()

	질의_에러_연속_발생_횟수 = lb.New안전한_정수64(0)
	V콜백_도우미_수량     = lb.F최대값(runtime.NumCPU(), 2)

	Ch모니터링_루틴_종료 = make(chan lb.T신호, 1)
	Ch콜백_도우미_종료  = make(chan lb.T신호, V콜백_도우미_수량)
	Ch접속_끊김      = make(chan lb.T신호)

	전송_제한_정보_모음          *xt.TR코드별_전송_제한_정보_모음
	tr코드별_전송_제한_초당_1회_미만 = make(map[string]lb.I전송_권한)
	tr코드별_전송_제한_1초       = make(map[string]lb.I전송_권한)
	tr코드별_전송_제한_10분      = make(map[string]lb.I전송_권한)

	주문_응답_구독_중 = lb.New안전한_bool(false)

	종료_잠금 = sync.Mutex{}
	종료_시각 = lb.New안전한_시각(time.Time{})
)

// 종목 관련 저장소는 초기화 이후에는 사실상 읽기 전용. 다중 사용에 문제가 없음.
var (
	종목모음_설정_잠금   sync.Mutex
	종목모음_설정일     = lb.New안전한_시각(time.Time{})
	종목맵_전체       = make(map[string]*lb.S종목)
	종목모음_코스피     = make([]*lb.S종목, 0)
	종목맵_코스피      = make(map[string]*lb.S종목)
	종목모음_코스닥     = make([]*lb.S종목, 0)
	종목맵_코스닥      = make(map[string]*lb.S종목)
	종목모음_ETF     = make([]*lb.S종목, 0)
	종목모음_ETN     = make([]*lb.S종목, 0)
	종목모음_ETF_ETN = make([]*lb.S종목, 0)
	종목모음_전체      = make([]*lb.S종목, 0)
	특수_종목_맵      = make(map[string]*lb.S종목)
	기준가_맵        = make(map[string]int64)
	하한가_맵        = make(map[string]int64)
	계좌번호_모음      []string
	프로세스ID_DLL32 int
)

var (
	최근_조건부_지정가_주문_종목코드 = "" // 호가 적합성 오류가 발생할 경우를 대비한 기록.
	조건부_지정가_주문_불가_맵    = make(map[string]lb.S비어있음)
	조건부_지정가_주문_불가_맵_잠금 = new(sync.Mutex)
)
