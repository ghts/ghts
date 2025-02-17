package dll32

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/dll"
	xt "github.com/ghts/ghts/xing/base"
	"runtime"
	"time"
	"unsafe"
)

// 단일 스레드에서 API를 호출.
// Win32 함수, 증권사 API 모두 Go언어와 같은 동시/병렬 처리에 대한 고려가 없던 시절에 만들어졌으므로,
// 가능한 단일 고정 스레드에서 호출하는 게 좋다.
func go함수_호출_도우미(ch초기화, ch종료 chan lib.T신호) {
	if lib.F공통_종료_채널_닫힘() {
		return
	}

	defer func() {
		recover()

		if lib.F공통_종료_채널_닫힘() {
			Ch함수_호출_도우미_종료 <- lib.P신호_종료
		} else {
			ch종료 <- lib.P신호_종료
		}
	}()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	f초기화_XingAPI() // 모든 API 액세스를 단일 스레드에서 하기 위해서 여기에서 API 초기화를 실행함.
	F메시지_윈도우_생성()

	ch공통_종료 := lib.Ch공통_종료()

	select {
	case ch초기화 <- lib.P신호_초기화:
	default:
	}

	for {
		select {
		case 질의 := <-Ch질의:
			f질의값_처리(질의)
		case <-ch공통_종료:
			return
		default:
			F윈도우_메시지_처리()
			time.Sleep(50)
		}
	}
}

func f질의값_처리(질의 *lib.S채널_질의) {
	var 에러 error

	defer lib.S예외처리{M에러: &에러, M함수: func() { 질의.Ch에러 <- 에러 }}.S실행()

	switch 질의.M값.TR구분() {
	case xt.TR조회, xt.TR주문:
		F조회_및_주문_질의_처리(질의)
	case xt.TR실시간_정보_구독, xt.TR실시간_정보_해지:
		F실시간_정보_구독_해지_처리(질의)
	case xt.TR실시간_정보_일괄_해지:
		F실시간_정보_일괄_해지(질의)
	case xt.TR접속_및_로그인:
		F접속_처리(질의)
	case xt.TR접속됨:
		F접속됨(질의)
	case xt.TR서버_이름:
		F서버_이름(질의)
	case xt.TR에러_코드:
		F에러_코드(질의)
	case xt.TR에러_메시지:
		F에러_메시지(질의)
	case xt.TR계좌_수량:
		F계좌_수량(질의)
	case xt.TR계좌번호_모음:
		F계좌번호_모음(질의)
	case xt.TR계좌_이름:
		F계좌_이름(질의)
	//case xt.TR계좌_상세명:
	//	F계좌_상세명(질의)
	case xt.TR계좌_별명:
		F계좌_별명(질의)
	case xt.TR코드별_전송_제한:
		TR코드별_전송_제한(질의)
	case xt.TR소켓_테스트:
		질의.Ch회신값 <- lib.P신호_OK
	case xt.TR서버_구분:
		질의.Ch회신값 <- int(서버_구분)
	case xt.TR종료:
		F종료_질의_처리(질의)
	//case xt.TR초기화:
	//	f초기화_XingAPI() // 모든 API 액세스를 단일 스레드에서 하기 위해서 여기에서 API 초기화를 실행함.
	//	F메시지_윈도우_생성()
	default:
		panic(lib.New에러("예상하지 못한 TR구분값 : '%v'", int(질의.M값.TR구분())))
	}
}

func F조회_및_주문_질의_처리(질의 *lib.S채널_질의) {
	var 에러 error
	defer lib.S예외처리{M에러: &에러, M함수: func() { 질의.Ch에러 <- 에러 }}.S실행()

	var c데이터 unsafe.Pointer

	var 길이 int
	연속_조회_여부 := false
	연속_조회_키 := ""
	질의값 := 질의.M값
	TR코드 := 질의값.(lib.I질의값).TR코드()

	switch TR코드 {
	//case xt.TR선물옵션_주문체결내역조회_CFOAQ00600:
	//	질의값_CFOAQ00600 := 질의값.(*xt.CFOAQ00600_선물옵션_주문체결내역_질의값)
	//	연속_조회_여부 = 질의값_CFOAQ00600.M연속조회_여부
	//	연속_조회_키 = 질의값_CFOAQ00600.M연속키
	//
	//	c데이터 = unsafe.Pointer(xt.NewCFOAQ00600InBlock1(질의값_CFOAQ00600, 계좌_비밀번호))
	//	길이 = xt.SizeCFOAQ00600InBlock1
	//case xt.TR선물옵션_정상주문_CFOAT00100:
	//	c데이터 = unsafe.Pointer(xt.NewCFOAT00100InBlock1(질의값.(*xt.CFOAT00100_선물옵션_정상주문_질의값), 계좌_비밀번호))
	//	길이 = xt.SizeCFOAT00100InBlock1
	//case xt.TR선물옵션_정정주문_CFOAT00200:
	//	c데이터 = unsafe.Pointer(xt.NewCFOAT00200InBlock1(질의값.(*xt.CFOAT00200_선물옵션_정정주문_질의값), 계좌_비밀번호))
	//	길이 = xt.SizeCFOAT00200InBlock1
	//case xt.TR선물옵션_취소주문_CFOAT00300:
	//	c데이터 = unsafe.Pointer(xt.NewCFOAT00300InBlock1(질의값.(*xt.CFOAT00300_선물옵션_취소주문_질의값), 계좌_비밀번호))
	//	길이 = xt.SizeCFOAT00300InBlock1
	//case xt.TR선물옵션_예탁금_증거금_조회_CFOBQ10500:
	//	질의값_CFOBQ10500 := 질의값.(*xt.CFOBQ10500_선물옵션_예탁금_증거금_조회_질의값)
	//	연속_조회_여부 = 질의값_CFOBQ10500.M연속조회_여부
	//	연속_조회_키 = 질의값_CFOBQ10500.M연속키
	//
	//	c데이터 = unsafe.Pointer(xt.NewCFOBQ105000InBlock1(질의값_CFOBQ10500, 계좌_비밀번호))
	//	길이 = xt.SizeCFOBQ10500InBlock1
	//case xt.TR선물옵션_미결제약정_현황_CFOFQ02400:
	//	질의값_CFOFQ02400 := 질의값.(*xt.CFOFQ02400_선물옵션_미결제약정_질의값)
	//	연속_조회_여부 = 질의값_CFOFQ02400.M연속조회_여부
	//	연속_조회_키 = 질의값_CFOFQ02400.M연속키
	//
	//	c데이터 = unsafe.Pointer(xt.NewCFOFQ02400InBlock1(질의값_CFOFQ02400, 계좌_비밀번호))
	//	길이 = xt.SizeCFOFQ02400InBlock1
	case xt.TR현물계좌_총평가_CSPAQ12200:
		계좌번호 := 질의값.(*lib.S질의값_문자열).M문자열

		c데이터 = unsafe.Pointer(xt.NewCSPAQ12200InBlock(계좌번호, f계좌_비밀번호()))
		길이 = xt.SizeCSPAQ12200InBlock1
	case xt.TR현물계좌_잔고내역_조회_CSPAQ12300:
		질의값_CSPAQ12300 := 질의값.(*xt.CSPAQ12300_현물계좌_잔고내역_질의값)
		연속_조회_여부 = 질의값_CSPAQ12300.M연속조회_여부
		연속_조회_키 = 질의값_CSPAQ12300.M연속키

		c데이터 = unsafe.Pointer(xt.NewCSPAQ12300InBlock(질의값.(*xt.CSPAQ12300_현물계좌_잔고내역_질의값), f계좌_비밀번호()))
		길이 = xt.SizeCSPAQ12300InBlock1
	case xt.TR현물계좌_주문체결내역_조회_CSPAQ13700:
		질의값_CSPAQ13700 := 질의값.(*xt.CSPAQ13700_현물계좌_주문체결내역_질의값)
		연속_조회_여부 = 질의값_CSPAQ13700.M연속조회_여부
		연속_조회_키 = 질의값_CSPAQ13700.M연속키

		c데이터 = unsafe.Pointer(xt.NewCSPAQ13700InBlock(질의값_CSPAQ13700, f계좌_비밀번호()))
		길이 = xt.SizeCSPAQ13700InBlock1
	case xt.TR현물계좌_예수금_주문가능금액_CSPAQ22200:
		계좌번호 := 질의값.(*lib.S질의값_문자열).M문자열

		c데이터 = unsafe.Pointer(xt.NewCSPAQ22200InBlock(계좌번호, f계좌_비밀번호()))
		길이 = xt.SizeCSPAQ22200InBlock1
	case xt.TR현물_정상_주문_CSPAT00600:
		c데이터 = unsafe.Pointer(xt.NewCSPAT00600InBlock(질의값.(*xt.CSPAT00600_현물_정상_주문_질의값), f계좌_비밀번호()))
		길이 = xt.SizeCSPAT00600InBlock1
	case xt.TR현물_정정_주문_CSPAT00700:
		c데이터 = unsafe.Pointer(xt.NewCSPAT00700InBlock(질의값.(*xt.CSPAT00700_현물_정정_주문_질의값), f계좌_비밀번호()))
		길이 = xt.SizeCSPAT00700InBlock1
	case xt.TR현물_취소_주문_CSPAT00800:
		c데이터 = unsafe.Pointer(xt.NewCSPAT00800InBlock(질의값.(*lib.S질의값_취소_주문), f계좌_비밀번호()))
		길이 = xt.SizeCSPAT00800InBlock1
	case xt.TR현물_당일_매매일지_t0150:
		c데이터 = unsafe.Pointer(xt.NewT0150InBlock(질의값.(*xt.T0150_현물_당일_매매일지_질의값)))
		길이 = xt.SizeT0150InBlock
	case xt.TR현물_일자별_매매일지_t0151:
		c데이터 = unsafe.Pointer(xt.NewT0151InBlock(질의값.(*xt.T0151_현물_일자별_매매일지_질의값)))
		길이 = xt.SizeT0151InBlock
	case xt.TR시간_조회_t0167:
		c데이터 = unsafe.Pointer(dll.F2ANSI문자열(""))
		길이 = 0
	case xt.TR현물_체결_미체결_조회_t0425:
		c데이터 = unsafe.Pointer(xt.NewT0425InBlock(질의값.(*xt.T0425_현물_체결_미체결_조회_질의값), f계좌_비밀번호()))
		길이 = xt.SizeT0425InBlock
	//case xt.TR선물옵션_체결_미체결_조회_t0434:
	//	c데이터 = unsafe.Pointer(xt.NewT0434InBlock(질의값.(*xt.T0434_선물옵션_체결_미체결_조회_질의값), 계좌_비밀번호))
	//	길이 = xt.SizeT0434InBlock
	case xt.TR현물_호가_조회_t1101:
		c데이터 = unsafe.Pointer(xt.NewT1101InBlock(질의값.(*lib.S질의값_단일_종목)))
		길이 = xt.SizeT1101InBlock
	case xt.TR현물_시세_조회_t1102:
		c데이터 = unsafe.Pointer(xt.NewT1102InBlock(질의값.(*xt.T1102_현물_시세_조회_질의값)))
		길이 = xt.SizeT1102InBlock
	case xt.TR현물_기간별_조회_t1305:
		연속키 := lib.F2문자열_공백_제거(질의값.(*xt.T1305_현물_기간별_조회_질의값).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT1305InBlock(질의값.(*xt.T1305_현물_기간별_조회_질의값)))
		길이 = xt.SizeT1305InBlock
	case xt.TR현물_당일_전일_분틱_조회_t1310:
		연속키 := lib.F2문자열_공백_제거(질의값.(*xt.T1310_현물_전일당일분틱조회_질의값).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT1310InBlock(질의값.(*xt.T1310_현물_전일당일분틱조회_질의값)))
		길이 = xt.SizeT1310InBlock
	case xt.TR관리_불성실_투자유의_조회_t1404:
		연속키 := lib.F2문자열_공백_제거(질의값.(*xt.T1404_관리종목_조회_질의값).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT1404InBlock(질의값.(*xt.T1404_관리종목_조회_질의값)))
		길이 = xt.SizeT1404InBlock
	case xt.TR투자경고_매매정지_정리매매_조회_t1405:
		연속키 := lib.F2문자열_공백_제거(질의값.(*xt.T1405_투자경고_조회_질의값).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT1405InBlock(질의값.(*xt.T1405_투자경고_조회_질의값)))
		길이 = xt.SizeT1405InBlock
	case xt.TR_ETF_시세_조회_t1901:
		c데이터 = unsafe.Pointer(xt.NewT1901InBlock(질의값.(*lib.S질의값_단일_종목)))
		길이 = xt.SizeT1901InBlock
	case xt.TR_ETF_시간별_추이_t1902:
		연속키 := lib.F2문자열_공백_제거(질의값.(*lib.S질의값_단일종목_연속키).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}
		c데이터 = unsafe.Pointer(xt.NewT1902InBlock(질의값.(*lib.S질의값_단일종목_연속키)))
		길이 = xt.SizeT1902InBlock
	case xt.TR_ETF_LP호가_조회_t1906:
		c데이터 = unsafe.Pointer(xt.NewT1906InBlock(질의값.(*lib.S질의값_단일_종목)))
		길이 = xt.SizeT1906InBlock
	case xt.TR기업정보_요약_t3320:
		c데이터 = unsafe.Pointer(xt.NewT3320InBlock(질의값.(*lib.S질의값_단일_종목)))
		길이 = xt.SizeT3320InBlock
	case xt.TR재무순위_종합_t3341:
		c데이터 = unsafe.Pointer(xt.NewT3341InBlock(질의값.(*xt.T3341_재무순위_질의값)))
		길이 = xt.SizeT3341InBlock
	case xt.TR현물_멀티_현재가_조회_t8407:
		c데이터 = unsafe.Pointer(xt.NewT8407InBlock(질의값.(*lib.S질의값_복수_종목)))
		길이 = xt.SizeT8407InBlock
	case xt.TR현물_차트_일주월년_t8410:
		연속키 := lib.F2문자열_공백_제거(질의값.(*xt.T8410_현물_차트_일주월년_질의값).M연속일자)

		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT8410InBlock(질의값.(*xt.T8410_현물_차트_일주월년_질의값)))
		길이 = xt.SizeT8410InBlock
	case xt.TR현물_차트_틱_t8411:
		연속키 := lib.F2문자열_공백_제거(질의값.(*xt.T8411_현물_차트_틱_질의값).M연속일자) +
			lib.F2문자열_공백_제거(질의값.(*xt.T8411_현물_차트_틱_질의값).M연속시간)

		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT8411InBlock(질의값.(*xt.T8411_현물_차트_틱_질의값)))
		길이 = xt.SizeT8411InBlock
	case xt.TR현물_차트_분_t8412:
		연속키 := lib.F2문자열_공백_제거(질의값.(*xt.T8412_현물_차트_분_질의값).M연속일자) +
			lib.F2문자열_공백_제거(질의값.(*xt.T8412_현물_차트_분_질의값).M연속시간)

		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT8412InBlock(질의값.(*xt.T8412_현물_차트_분_질의값)))
		길이 = xt.SizeT8412InBlock
	case xt.TR현물_차트_일주월_t8413:
		연속키 := lib.F2문자열_공백_제거(질의값.(*xt.T8413_현물_차트_일주월_질의값).M연속일자)

		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT8413InBlock(질의값.(*xt.T8413_현물_차트_일주월_질의값)))
		길이 = xt.SizeT8413InBlock
	case xt.TR증시_주변_자금_추이_t8428:
		연속키 := lib.F2문자열_공백_제거(질의값.(*xt.T8428_증시주변_자금추이_질의값).M연속키)
		if 연속키 != "" {
			연속_조회_여부 = true
			연속_조회_키 = 연속키
		}

		c데이터 = unsafe.Pointer(xt.NewT8428InBlock(질의값.(*xt.T8428_증시주변_자금추이_질의값)))
		길이 = xt.SizeT8428InBlock
	//case xt.TR지수선물_마스터_조회_t8432:
	//	c데이터 = unsafe.Pointer(xt.NewT8432InBlock(질의값.(*lib.S질의값_문자열)))
	//	길이 = xt.SizeT8428InBlock
	case xt.TR현물_종목_조회_t8436:
		c데이터 = unsafe.Pointer(xt.NewT8436InBlock(질의값.(*lib.S질의값_문자열)))
		길이 = xt.SizeT8436InBlock
	default:
		panic(lib.New에러("미구현 : '%v'", TR코드))
	}

	lib.F조건부_패닉(c데이터 == nil, "c데이터 설정 실패.")

	식별번호 := lib.F확인2(F질의(TR코드, c데이터, 길이, 연속_조회_여부, 연속_조회_키, lib.P30초))

	switch {
	case 식별번호 < 0:
		질의.Ch에러 <- lib.New에러("TR호출 실패. 반환된 식별번호가 음수임. '%v' '%v'", TR코드, 식별번호)
	default:
		질의.Ch회신값 <- 식별번호
	}
}

func F실시간_정보_구독_해지_처리(질의 *lib.S채널_질의) {
	var 함수 func(string, string, int) error
	var 전체_종목코드 string
	var 단위_길이 int

	질의값 := 질의.M값

	switch TR구분 := 질의값.TR구분(); TR구분 {
	case xt.TR실시간_정보_구독:
		함수 = F실시간_정보_구독
	case xt.TR실시간_정보_해지:
		함수 = F실시간_정보_해지
	default:
		panic(lib.F2문자열("예상하지 못한 경우 : '%v' '%v'", int(TR구분), TR구분.String()))
	}

	switch 변환값 := 질의값.(type) {
	case lib.I종목코드_모음:
		전체_종목코드 = 변환값.G전체_종목코드()
		단위_길이 = len(변환값.G종목코드_모음()[0])
	case lib.I종목코드:
		전체_종목코드 = 변환값.G종목코드()
		단위_길이 = len(변환값.G종목코드())
	default:
		전체_종목코드 = ""
		단위_길이 = 0
	}

	에러 := 함수(질의값.TR코드(), 전체_종목코드, 단위_길이)

	if 에러 == nil {
		질의.Ch회신값 <- lib.P신호_OK
	} else {
		질의.Ch에러 <- 에러
	}
}

func F접속_처리(질의 *lib.S채널_질의) {
	서버_구분 = xt.T서버_구분(질의.M값.(*lib.S질의값_정수).M정수값)

	접속_처리_잠금.Lock()
	defer 접속_처리_잠금.Unlock()

	if 에러_접속 := F접속(서버_구분); 에러_접속 != nil {
		질의.Ch에러 <- 에러_접속
	} else if 에러_로그인 := F로그인(); 에러_로그인 != nil {
		질의.Ch에러 <- 에러_로그인
	} else {
		질의.Ch회신값 <- lib.P신호_OK
	}
}

func F종료_질의_처리(질의 *lib.S채널_질의) {
	질의.Ch회신값 <- lib.P신호_종료

	lib.F대기(lib.P500밀리초) // 회신 전달 대기

	f종료()
}

func f종료() {
	f콜백_동기식(lib.New콜백_신호(lib.P신호_DLL32_종료))
	f실시간_정보_일괄_해지()
	F로그아웃()
	F소켓_정리()
	OnDestroy()
}
