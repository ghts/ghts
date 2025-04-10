package xing

import (
	"fmt"
	lb "github.com/ghts/ghts/lib"
	ep "github.com/ghts/ghts/lib/external_process"
	"github.com/ghts/ghts/lib/nanomsg"
	"github.com/ghts/ghts/xing/base"
	"github.com/mitchellh/go-ps"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

func init() {
	//TR현물_호가_조회_t1101 = "t1101"	// HTS 1101 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR현물_시세_조회_t1102 = "t1102"	// HTS 1101 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR현물_기간별_조회_t1305      = "t1305"	// HTS 1305 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR현물_당일_전일_분틱_조회_t1310 = "t1310"	// HTS 1310 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR관리_불성실_투자유의_조회_t1404 = "t1404"	// HTS 1404 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR투자경고_매매정지_정리매매_조회_t1405 = "t1405"	// HTS 1405 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR_ETF_시간별_추이_t1902    = "t1902"	// HTS 1902 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR재무순위_종합_t3341        = "t3341"	// HTS 3303 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR현물_차트_틱_t8411        = "t8411"	// HTS 4001 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR현물_차트_분_t8412        = "t8412"	// HTS 4001 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR현물_차트_일주월_t8413      = "t8413"		// HTS 1305 화면, DevCenter 소숫점 비교 확인 완료.
	//TR증시_주변_자금_추이_t8428    = "t8428"	// HTS 1503 화면, DevCenter 소숫점 비교 확인 완료.
	//TR현물_종목_조회_t8436       = "t8436"	// 종목 정보 대조 비교 완료.

	메모 := `
- 선물 주문 : CFOAT00100(선물옵션 정상주문) / CFOAT00200(선물옵션 정정주문) / CFOAT00300(선물옵션 취소주문)
- CFOAQ00600(선물옵션 계좌주문체결내역 조회)
- CFOFQ02400(선물옵션 계좌 미결제 약정현황(평균가)
- CFOBQ10500(선물옵션 계좌예탁금증거금조회)
- 실시간 TR : C01(선물주문체결) / O01(선물접수) / H01(선물주문정정취소)
- 데이터 수집 중 연결이 끊기는 경우가 자주 발생함. 재접속 하는 것 완성 및 디버깅 할 것.
- 매매일지 / 수수료 (t0150, t0151)
- 계좌 조회 기능 구현할 것. (CDPCQ04700, CSPAQ12200,CSPAQ12300, CSPAQ13700, FOCCQ33600)
- ETF 일별 추이 (t1903)
- 선물 가격 정보 (t8414, t8415, t8416, t9943)
- 해외 지수 (o3123, o3121, t3518, t3521) // o31xx는 해외선물 계좌 내지 모의투자 있어야 보임.
- 주식 분할, 합병, 소각, 유상/무상 증자등 주가에 영향을 미치는 이벤트 아는 방법 찾을 것.
- 업종별 조회 기능. (t8424, t4203, t1514, t8419)
- 테마별 조회 기능. (t8425, t1531)
- 뉴스 및 공시 정보. (t3102,t3202)
- TR 결과값을 출력해서 HTS와 대조 비교해 볼 것. (실시간 정보들)
`
	lb.F중복없는_문자열_출력(메모)
}

func F초기화(서버_구분 xt.T서버_구분, 로그인_정보 *xt.S로그인_정보) {
	// 자식 프로세스는 부모 프로세스의 환경 변수를 그대로 물려받음.
	// 로그인 정보는 환경 변수를 통해서 DLL32 모듈로 전달.
	xt.F서버_구분_설정(서버_구분)
	xt.F로그인_정보_환경_변수_설정(로그인_정보)
	F소켓_생성()
	F초기화_Go루틴()
	lb.F확인1(f초기화_DLL32())
	lb.F확인1(F접속_로그인())
	lb.F조건부_패닉(!f초기화_작동_확인(), "초기화 작동 확인 실패.")
	lb.F확인1(F초기화_TR전송_제한())
	lb.F확인1(F종목_정보_설정())
	lb.F확인1(F전일_당일_설정())

	fmt.Println("** Xing API 초기화 완료 **")
}

func F소켓_생성() {
	소켓REP_TR콜백 = lb.F확인2(nano.NewNano소켓REP(xt.F주소_콜백()))
}

func F초기화_Go루틴() {
	ch초기화 := make(chan lb.T신호, 1)
	go Go루틴_관리(ch초기화)
	<-ch초기화
}

func f초기화_DLL32() (에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수_항상: xt.F로그인_정보_환경_변수_삭제}.S실행()

	if !lb.F인터넷에_접속됨() {
		lb.F문자열_출력("인터넷을 확인하십시오.")
		return
	}

	switch runtime.GOOS {
	case "windows":
		GOARCH_원래값 := os.Getenv("GOARCH")
		os.Setenv("GOARCH", "386") // 64비트에서 컴파일을 막는 루틴을 회피하기 위해서 32비트인척 함.
		defer os.Setenv("GOARCH", GOARCH_원래값)

		if 에러 := DLL32_빌드(); 에러 != nil {
			panic(lb.New에러with출력("dll32_xing.exe 빌드 에러 발생.\n%v", 에러))
		} else if lb.F파일_없음(DLL32_실행_화일_경로()) {
			panic(lb.New에러with출력("빌드된 실행 화일 찾을 수 없음. '%v'", DLL32_실행_화일_경로()))
		}

		// 자식 프로세스는 부모 프로세스의 환경 변수를 그대로 물려받음.
		// 로그인 정보는 환경 변수를 통해서 전달.
		프로세스ID_DLL32 = lb.F확인2(ep.F외부_프로세스_실행(DLL32_실행_화일_경로()))
		<-ch신호_DLL32_초기화
	default:
		lb.F문자열_출력("*********************************************\n"+
			"현재 OS(%v)에서는 'dll32_xing.exe'를 수동으로 실행해야 합니다.\n"+
			"*********************************************", runtime.GOOS)
	}

	return nil
}

func DLL32_실행_화일_경로() string {
	if 현재_디렉토리, 에러 := os.Getwd(); 에러 != nil {
		return ""
	} else {
		return filepath.Join(현재_디렉토리, "dll32_xing.exe")
	}
}

func DLL32_소스_코드_화일_경로() string {
	return filepath.Join(os.Getenv("USERPROFILE"), `go\src\github.com\ghts\ghts\xing\dll32\dll32_xing.go`)
}

func DLL32_빌드() error {
	if lb.F파일_존재함(DLL32_실행_화일_경로()) {
		if lb.F파일_없음(DLL32_소스_코드_화일_경로()) {
			return nil // 컴파일 준비되어 있지 않으면 이미 존재하는 실행 화일 그대로 사용.
		} else {
			os.Remove(DLL32_실행_화일_경로()) // 컴파일 준비되어 있으면 삭제 후 최신 버전 재생성
		}
	}

	GOARCH_원래값 := os.Getenv("GOARCH")
	os.Setenv("GOARCH", "386") // 32비트 컴파일.
	defer os.Setenv("GOARCH", GOARCH_원래값)

	CGO_ENABLED_원래값 := os.Getenv("CGO_ENABLED")
	os.Setenv("CGO_ENABLED", "0") // cgo 비활성화
	defer os.Setenv("CGO_ENABLED", CGO_ENABLED_원래값)

	PATH_원래값 := os.Getenv("PATH")
	os.Setenv("PATH", lb.GOROOT()+`\bin;C:\msys64\mingw32\bin;C:\msys64\usr\bin`)
	defer os.Setenv("PATH", PATH_원래값)

	return exec.Command("go", "build", "-o", "dll32_xing.exe", "github.com/ghts/ghts/xing/dll32").Run()
}

func DLL32_삭제() (에러 error) {
	if lb.F파일_존재함(DLL32_실행_화일_경로()) {
		return os.Remove(DLL32_실행_화일_경로())
	}

	return nil
}

func F접속_로그인() (에러 error) {
	if !F확인_TR소켓() {
		return lb.New에러("DLL32 프로세스 REP소켓 접속 불가.")
	}

	질의값 := lb.New질의값_정수(lb.TR접속, "", int(xt.F서버_구분()))
	i응답값 := lb.F확인2(F질의(질의값).G해석값(0))

	if 응답값, ok := i응답값.(lb.T신호); !ok {
		return lb.New에러("F접속_로그인() 예상하지 못한 자료형 : '%T'", i응답값)
	} else if 응답값 != lb.P신호_OK {
		return lb.New에러("예상하지 못한 응답값 : '%v'", 응답값)
	}

	<-ch신호_DLL32_로그인

	return nil
}

func f초기화_작동_확인() (작동_여부 bool) {
	defer lb.S예외처리{M함수: func() { 작동_여부 = false }}.S실행()

	ch확인 := make(chan lb.T신호, 1)
	ch타임아웃 := time.After(lb.P1분)

	go f접속_확인(ch확인)
	go f시간_일치_확인(ch확인)

	for i := 0; i < 2; i++ {
		select {
		case <-ch확인:
		case <-ch타임아웃:
			lb.New에러with출력("f초기화_작동_확인() 타임아웃.")
			return false
		}
	}

	//fmt.Println("** dll32 동작 확인 완료**")

	return true
}

func F확인_TR소켓() bool {
	for i := 0; i < 100; i++ {
		if 응답 := F질의(lb.New질의값_기본형(xt.TR소켓_테스트, ""), lb.P5초); 응답.G에러() == nil {
			return true
		}

		lb.F대기(lb.P500밀리초)
	}

	return false
}

func f접속_확인(ch완료 chan lb.T신호) {
	defer func() {
		if ch완료 != nil {
			ch완료 <- lb.P신호_종료
		}
	}()

	for i := 0; i < 10; i++ {
		if 접속됨, 에러 := F접속됨(); 에러 == nil && 접속됨 {
			break
		} else if 에러 != nil {
			lb.F에러_출력(에러)
		}

		lb.F대기(lb.P1초)
	}

	if 접속됨, 에러 := F접속됨(); 에러 != nil || !접속됨 {
		panic(lb.New에러("이 시점에 접속되어 있어야 함."))
	}

	return
}

func f시간_일치_확인(ch완료 chan lb.T신호) {
	defer func() { ch완료 <- lb.P신호_종료 }()

	if len(tr코드별_전송_제한_1초) == 0 {
		tr코드별_전송_제한_1초[xt.TR시간_조회_t0167] = lb.New전송_권한(xt.TR시간_조회_t0167, 5, lb.P1초)
	}

	for i := 0; i < 100; i++ {
		시각, 에러 := (<-TrT0167_시각_조회()).G값()

		if 에러 != nil || 시각.Equal(time.Time{}) {
			continue
		} else if 차이 := time.Now().Sub(시각); 차이 < -1*lb.P10분 || 차이 > lb.P10분 {
			panic(lb.New에러("서버와 시스템 시각 불일치 : 차이 '%v'분", 차이.Minutes()))
		}

		return
	}
}

func F전일_당일_설정() (에러 error) {
	for i := 0; i < 3; i++ {
		if 에러 := f전일_당일_설정(); 에러 == nil {
			return nil
		}
	}

	return lb.New에러("전일/당일 설정 중 에러 발생.")
}

func f전일_당일_설정() (에러 error) {
	lb.S예외처리{M에러: &에러}.S실행()

	const 수량 = 3

	질의값_기간별_조회 := xt.NewT1305_현물_기간별_조회_질의값("069500", xt.P일주월_일, 수량, xt.P거래소_KRX, "")
	i응답값 := lb.F확인2(F질의_단일TR(질의값_기간별_조회))

	switch 응답값 := i응답값.(type) {
	case *xt.T1305_현물_기간별_조회_응답:
		lb.F조건부_패닉(응답값.M헤더.M수량 != int64(수량), "예상하지 못한 수량 : '%v' '%v'", 응답값.M헤더.M수량, 수량)
		lb.F조건부_패닉(len(응답값.M반복값_모음.M배열) != 수량, "예상하지 못한 수량 : '%v' '%v'", len(응답값.M반복값_모음.M배열), 수량)
		lb.F조건부_패닉(응답값.M반복값_모음.M배열[0].M일자.Before(응답값.M반복값_모음.M배열[1].M일자), "예상하지 못한 순서")

		당일 := 응답값.M반복값_모음.M배열[0].M일자
		전일 := 응답값.M반복값_모음.M배열[1].M일자
		xt.F전일_당일_설정(전일, 당일)

		return nil
	default:
		panic(lb.New에러with출력("F전일_당일_설정() 예상하지 못한 자료형 : '%T'", i응답값))
	}
}

func DLL32_종료됨() bool {
	프로세스, 에러 := ps.FindProcess(프로세스ID_DLL32)
	포트_닫힘_C함수_호출 := lb.F포트_닫힘_확인(xt.F주소_DLL32())

	return 프로세스 == nil && 에러 == nil && 포트_닫힘_C함수_호출
}

func DLL32_종료() (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	if !DLL32_종료됨() {
		소켓REQ := 소켓REQ_저장소.G소켓()

		소켓REQ.S송신(lb.P변환형식_기본값, lb.New질의값_기본형(lb.TR종료, ""))
		소켓REQ_저장소.S회수(소켓REQ)

		lb.F대기(lb.P20초)
	}

	ch타임아웃 := time.After(lb.P2분)

	select {
	case <-ch신호_DLL32_종료:
	case <-ch타임아웃:
		return lb.New에러with출력("DLL32 종료 타임아웃")
	}

	for i := 0; i < 10; i++ {
		if DLL32_종료됨() {
			break
		}

		ep.F프로세스_종료by프로세스ID(프로세스ID_DLL32)
		lb.F대기(lb.P1초)
	}

	return nil
}

func F종료() {
	종료_잠금.Lock()
	defer func() {
		종료_시각.S값(lb.F지금())
		종료_잠금.Unlock()
	}()

	if lb.F지금().Before(종료_시각.G값().Add(lb.P3분)) {
		return // 중복 실행 방지.
	}

	DLL32_종료()
	lb.F공통_종료_채널_닫기()
	F소켓_정리()

	타임_아웃 := time.After(lb.P1분)

	select {
	case <-Ch모니터링_루틴_종료:
		//lb.F문자열_출력("모니터링 루틴 종료.")
	case <-타임_아웃:
		//lb.F문자열_출력("종료 타임아웃.")
	}

	for i := 0; i < V콜백_도우미_수량; i++ {
		select {
		case <-Ch콜백_도우미_종료:
			//lb.F문자열_출력("콜백 루틴 %v/%v 종료.", i+1, V콜백_도우미_수량)
		case <-타임_아웃:
			//lb.F문자열_출력("종료 타임아웃.")
		}
	}

	os.Exit(0)
}

func F소켓_정리() {
	lb.F패닉억제_호출(소켓REP_TR콜백.Close)
	소켓REQ_저장소.S정리()
}

var TR전송_제한_초기화_잠금 sync.Mutex

func F초기화_TR전송_제한() (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	if f전체TR_전송_제한_초기화_완료() {
		return nil
	}

	TR코드_모음 := []string{
		//xt.TR선물옵션_주문체결내역조회_CFOAQ00600,
		//xt.TR선물옵션_정상주문_CFOAT00100,
		//xt.TR선물옵션_정정주문_CFOAT00200,
		//xt.TR선물옵션_취소주문_CFOAT00300,
		//xt.TR선물옵션_예탁금_증거금_조회_CFOBQ10500,
		//xt.TR선물옵션_미결제약정_현황_CFOFQ02400,
		//xt.TR선물옵션_주문체결내역조회_CFOAQ00600,
		//xt.TR선물옵션_정상주문_CFOAT00100,
		//xt.TR선물옵션_정정주문_CFOAT00200,
		//xt.TR선물옵션_취소주문_CFOAT00300,
		//xt.TR선물옵션_예탁금_증거금_조회_CFOBQ10500,
		//xt.TR선물옵션_미결제약정_현황_CFOFQ02400,
		xt.TR현물계좌_총평가_CSPAQ12200,
		xt.TR현물계좌_잔고내역_조회_CSPAQ12300,
		xt.TR현물계좌_주문체결내역_조회_CSPAQ13700,
		xt.TR현물계좌_예수금_주문가능금액_CSPAQ22200,
		xt.TR현물_정상_주문_CSPAT00600,
		xt.TR현물_정정_주문_CSPAT00700,
		xt.TR현물_취소_주문_CSPAT00800,
		xt.TR현물_당일_매매일지_t0150,
		xt.TR현물_일자별_매매일지_t0151,
		xt.TR시간_조회_t0167,
		xt.TR현물_체결_미체결_조회_t0425,
		//xt.TR선물옵션_체결_미체결_조회_t0434,
		xt.TR현물_호가_조회_t1101,
		xt.TR현물_시세_조회_t1102,
		xt.TR현물_기간별_조회_t1305,
		xt.TR현물_당일_전일_분틱_조회_t1310,
		xt.TR관리_불성실_투자유의_조회_t1404,
		xt.TR투자경고_매매정지_정리매매_조회_t1405,
		xt.TR_ETF_시세_조회_t1901,
		xt.TR_ETF_시간별_추이_t1902,
		xt.TR_ETF_LP호가_조회_t1906,
		xt.TR기업정보_요약_t3320,
		xt.TR재무순위_종합_t3341,
		xt.TR현물_멀티_현재가_조회_t8407,
		xt.TR현물_차트_일주월년_t8410,
		xt.TR현물_차트_틱_t8411,
		xt.TR현물_차트_분_t8412,
		xt.TR현물_차트_일주월_t8413,
		xt.TR증시_주변_자금_추이_t8428,
		//xt.TR지수선물_마스터_조회_t8432,
		xt.TR현물_종목_조회_t8436}

	return tr전송_제한_초기화(TR코드_모음)
}

func tr전송_제한_초기화(TR코드_모음 []string) (에러 error) {
	TR전송_제한_초기화_잠금.Lock()
	defer TR전송_제한_초기화_잠금.Unlock()

	if len(TR코드_모음) > 1 && f전체TR_전송_제한_초기화_완료() {
		return nil
	} else if len(TR코드_모음) == 1 && f단일TR_전송_제한_초기화_완료(TR코드_모음[0]) {
		return nil
	}

	// 중복 제거
	TR코드_맵 := make(map[string]lb.S비어있음)

	for _, TR코드 := range TR코드_모음 {
		TR코드_맵[TR코드] = lb.S비어있음{}
	}

	TR코드_모음 = make([]string, 0)

	for TR코드 := range TR코드_맵 {
		TR코드_모음 = append(TR코드_모음, TR코드)
	}

	for {
		응답 := F질의(lb.New질의값_문자열_모음(xt.TR코드별_전송_제한, "", TR코드_모음), lb.P5초)
		lb.F확인1(응답.G에러())

		전송_제한_정보_모음 = new(xt.TR코드별_전송_제한_정보_모음)
		lb.F확인1(응답.G값(0, 전송_제한_정보_모음))
		lb.F조건부_패닉(len(TR코드_모음) != len(전송_제한_정보_모음.M맵),
			"서로 다른 길이 : '%v' '%v'", len(TR코드_모음), len(전송_제한_정보_모음.M맵))

		정상 := false
		for _, 전송_제한_정보 := range 전송_제한_정보_모음.M맵 {
			if 전송_제한_정보.M초당_전송_제한 > 0 {
				정상 = true
				break
			}
		}

		if 정상 {
			break
		}

		lb.F대기(lb.P1초)
	}

	for TR코드, 전송_제한_정보 := range 전송_제한_정보_모음.M맵 {
		if 전송_제한_정보.M초_베이스 > 0 {
			if 전송_권한, 존재함 := tr코드별_전송_제한_1초[TR코드]; 존재함 {
				전송_권한.S수량_간격_변경(1, lb.P1초*time.Duration(전송_제한_정보.M초_베이스))
				tr코드별_전송_제한_1초[TR코드] = 전송_권한
			} else {
				tr코드별_전송_제한_1초[TR코드] = lb.New전송_권한(TR코드, 1, lb.P1초*time.Duration(전송_제한_정보.M초_베이스))
			}
		} else if 전송_제한_정보.M초당_전송_제한 > 0 {
			if 전송_권한, 존재함 := tr코드별_전송_제한_1초[TR코드]; 존재함 {
				전송_권한.S수량_간격_변경(전송_제한_정보.M초당_전송_제한, lb.P1초)
				tr코드별_전송_제한_1초[TR코드] = 전송_권한
			} else {
				tr코드별_전송_제한_1초[TR코드] = lb.New전송_권한(TR코드, 전송_제한_정보.M초당_전송_제한, lb.P1초)
			}
		}

		if 전송_제한_정보.M10분당_전송_제한 > 0 {
			if 전송_권한, 존재함 := tr코드별_전송_제한_10분[TR코드]; 존재함 {
				전송_권한.S수량_간격_변경(전송_제한_정보.M10분당_전송_제한, lb.P10분)
				tr코드별_전송_제한_10분[TR코드] = 전송_권한
			} else {
				전송_권한 = lb.New전송_권한(TR코드, 전송_제한_정보.M10분당_전송_제한, lb.P10분)

				// 지난 10분간 이미 전송한 수량을 반영.
				for i := 0; i < 전송_제한_정보.M10분간_전송한_수량; i++ {
					전송_권한.G획득()
				}

				tr코드별_전송_제한_10분[TR코드] = 전송_권한
			}
		}
	}

	return nil
}

func f전체TR_전송_제한_초기화_완료() bool {
	return len(tr코드별_전송_제한_1초) > 1 && len(tr코드별_전송_제한_10분) > 0
}

func f단일TR_전송_제한_초기화_완료(TR코드 string) bool {
	_, 존재함1 := tr코드별_전송_제한_1초[TR코드]
	_, 존재함2 := tr코드별_전송_제한_10분[TR코드]

	return 존재함1 || 존재함2
}
