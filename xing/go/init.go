/* Copyright (C) 2015-2019 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"runtime"

	"fmt"
	"time"
)

func init() {
	ch신호_C32_모음 = make([]chan xt.T신호_C32, 2)

	for i := 0; i < len(ch신호_C32_모음); i++ {
		ch신호_C32_모음[i] = make(chan xt.T신호_C32, 1)
	}

	//TR현물_호가_조회_t1101 = "t1101"	// HTS 1101 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR현물_시세_조회_t1102 = "t1102"	// HTS 1101 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR현물_기간별_조회_t1305      = "t1305"	// HTS 1305 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR현물_당일_전일_분틱_조회_t1310 = "t1310"	// HTS 1310 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR관리_불성실_투자유의_조회_t1404 = "t1404"	// HTS 1404 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR투자경고_매매정지_정리매매_조회_t1405 = "t1405"	// HTS 1405 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR_ETF_시간별_추이_t1902    = "t1902"	// HTS 1902 화면,  DevCenter 소숫점 비교 확인 완료.
	//TR기업정보_요약_t3320        = "t3320"	// HTS 3301, 3302 화면 비교 확인 완료. (CPS,SPS, EBITDAS 확인 불가)
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

- 재무 정보 취득 방법 알아볼 것. (t3320의 데이터는 정확도가 아주 낮음.)
- 주식 분할, 합병, 소각, 유상/무상 증자등 주가에 영향을 미치는 이벤트 아는 방법 찾을 것.
- 업종별 조회 기능. (t8424, t4203, t1514, t8419)
- 테마별 조회 기능. (t8425, t1531)
- 뉴스 및 공시 정보. (t3102,t3202)
- TR 결과값을 출력해서 HTS와 대조 비교해 볼 것. (실시간 정보들)
`
	lib.F메모(메모)
}

func F초기화() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	f초기화_소켓()
	f초기화_Go루틴()
	lib.F확인(f초기화_xing_C32())
	lib.F조건부_패닉(!f초기화_작동_확인(), "초기화 작동 확인 실패.")
	lib.F확인(f초기화_TR전송_제한())
	lib.F확인(f종목모음_설정())
	lib.F확인(f전일_당일_설정())
	f접속유지_실행()

	fmt.Println("**     초기화 완료     **")

	return nil
}

func f초기화_소켓() {
	소켓REP_TR콜백 = lib.NewNano소켓REP_raw_단순형(lib.P주소_Xing_C함수_콜백)
	소켓SUB_실시간_정보 = lib.NewNano소켓SUB_단순형(lib.P주소_Xing_실시간).(lib.I소켓Raw)
}

func f초기화_xing_C32() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	xing_C32_실행_잠금.Lock()
	defer xing_C32_실행_잠금.Unlock()

	if !lib.F인터넷에_접속됨() {
		lib.F문자열_출력("인터넷을 확인하십시오.")
		return
	}

	switch runtime.GOOS {
	case "windows":
		if 프로세스ID := xing_C32_실행_중(); 프로세스ID >= 0 {
			lib.F문자열_출력("xing_C32 가 이미 실행 중입니다.")
			return nil
		}

		lib.F확인(lib.F외부_프로세스_실행(xing_C32_경로))
	default:
		lib.F문자열_출력("*********************************************\n"+
			"현재 OS(%v)에서는 'xing_C32'를 수동으로 실행해야 합니다.\n"+
			"*********************************************", runtime.GOOS)
	}

	return nil
}

func f초기화_Go루틴() {
	고루틴_모음 := []func(chan lib.T신호) error{go_TR호출, go_TR콜백_처리, go_RT_주문처리결과}
	ch초기화 := make(chan lib.T신호, len(고루틴_모음))

	for _, 고루틴 := range 고루틴_모음 {
		go 고루틴(ch초기화)
	}

	for range 고루틴_모음 {
		<-ch초기화
	}
}

func f초기화_작동_확인() (작동_여부 bool) {
	defer lib.S예외처리{M함수: func() { 작동_여부 = false }}.S실행()

	ch확인 := make(chan lib.T신호, 1)
	ch타임아웃 := time.After(lib.P10분)

	select {
	case <-ch신호_C32_모음[xt.P신호_C32_READY]: // 서버 접속된 상태임.
	case <-ch타임아웃:
		lib.F체크포인트("C32 초기화 타임아웃")
		return false
	}

	// C32 모듈의 소켓이 초기화 될 시간을 준다.
	// 이게 없으면 제대로 작동하지 않으며, 필수적인 부분임. 삭제하지 말 것.
	lib.F대기(lib.P10초)

	// 소켓REP_TR수신 동작 테스트
	go tr수신_소켓_동작_확인(ch확인)

	select {
	case <-ch확인:
	case <-ch타임아웃:
		lib.F체크포인트("F소켓REP_TR수신_동작_여부_확인() 타임아웃.")
		return false
	}

	// F접속됨() 테스트
	go f접속_확인(ch확인)

	select {
	case <-ch확인:
	case <-ch타임아웃:
		lib.F체크포인트("F접속됨_확인() 타임아웃.")
		return false
	}

	fmt.Println("**     C32 동작 확인 완료     **")

	return true
}

func tr수신_소켓_동작_확인(ch완료 chan lib.T신호) {
	defer func() { ch완료 <- lib.P신호_종료 }()

	for i := 0; i < 100; i++ {
		if 응답 := F질의(lib.New질의값_기본형(xt.TR소켓_테스트, ""), lib.P5초); 응답.G에러() == nil {
			return
		}
	}
}

func f접속_확인(ch완료 chan lib.T신호) {
	defer func() {
		if ch완료 != nil {
			ch완료 <- lib.P신호_종료
		}
	}()

	for i := 0; i < 10; i++ {
		if 접속됨, 에러 := F접속됨(); 에러 != nil {
			lib.F에러_출력(에러)
			lib.F대기(lib.P1초)
			continue
		} else if !접속됨 {
			panic(lib.New에러("이 시점에 접속되어 있어야 함."))
		}

		접속_여부.S값(true)

		return
	}
}

func tr동작_확인(ch완료 chan lib.T신호) {
	defer func() { ch완료 <- lib.P신호_종료 }()

	if len(tr코드별_초당_전송_제한) == 0 {
		tr코드별_초당_전송_제한[xt.TR시간_조회_t0167] = lib.New전송_권한(xt.TR시간_조회_t0167, 5, lib.P1초)
	}

	for i := 0; i < 100; i++ {
		시각, 에러 := (<-TrT0167_시각_조회()).G값()

		if 에러 != nil || 시각.Equal(time.Time{}) {
			continue
		} else if 차이 := time.Now().Sub(시각); 차이 < -1*lib.P10분 || 차이 > lib.P10분 {
			panic(lib.New에러("서버와 시스템 시각 불일치 : 차이 '%v'분", 차이.Minutes()))
		}

		return
	}
}

func f전일_당일_설정() (에러 error) {
	lib.S예외처리{M에러: &에러}.S실행()

	const 수량 = 30

	질의값_기간별_조회 := xt.NewT1305_현물_기간별_조회_질의값()
	질의값_기간별_조회.M구분 = xt.TR조회
	질의값_기간별_조회.M코드 = xt.TR현물_기간별_조회_t1305
	질의값_기간별_조회.M종목코드 = "069500"
	질의값_기간별_조회.M일주월_구분 = xt.P일주월_일
	질의값_기간별_조회.M연속키 = ""
	질의값_기간별_조회.M수량 = 수량

	i응답값, 에러 := F질의_단일TR(질의값_기간별_조회)
	lib.F확인(에러)

	switch 응답값 := i응답값.(type) {
	case *xt.T1305_현물_기간별_조회_응답:
		lib.F조건부_패닉(응답값.M헤더.M수량 != int64(수량), "예상하지 못한 수량 : '%v' '%v'", 응답값.M헤더.M수량, 수량)
		lib.F조건부_패닉(len(응답값.M반복값_모음.M배열) != 수량, "예상하지 못한 수량 : '%v' '%v'", len(응답값.M반복값_모음.M배열), 수량)
		lib.F조건부_패닉(응답값.M반복값_모음.M배열[0].M일자.Before(응답값.M반복값_모음.M배열[1].M일자), "예상하지 못한 순서")

		당일 = lib.New안전한_시각(응답값.M반복값_모음.M배열[0].M일자)
		전일 = lib.New안전한_시각(응답값.M반복값_모음.M배열[1].M일자)

		최근_영업일_모음 = make([]time.Time, 수량, 수량)

		for i, 값 := range 응답값.M반복값_모음.M배열 {
			최근_영업일_모음[i] = lib.F2일자(값.M일자)
		}

		xt.F전일_당일_설정(전일.G값(), 당일.G값())

		return nil
	default:
		panic(lib.New에러("예상하지 못한 자료형 : '%T'", i응답값))
	}
}

func C32_종료() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	defer 접속_여부.S값(false)

	// 종료 신호 전송
	func() {
		defer lib.S예외처리{M출력_숨김: true}.S실행()

		소켓REQ := 소켓REQ_저장소.G소켓()
		defer 소켓REQ_저장소.S회수(소켓REQ)

		소켓REQ.S옵션(lib.P1초)
		소켓REQ.S송신(lib.MsgPack, lib.New질의값_기본형(xt.TR종료, ""))
	}()

	select {
	case <-ch신호_C32_모음[xt.P신호_C32_종료]:
	case <-time.After(lib.P1초):
	}

	lib.F체크포인트()

	// 강제 종료
	for {
		if 프로세스ID := xing_C32_실행_중(); 프로세스ID < 0 {
			lib.F체크포인트()
			return
		} else {
			lib.F프로세스_종료by프로세스ID(프로세스ID)
			lib.F대기(lib.P3초)
		}
	}
}

func F리소스_정리() {
	C32_종료()
	lib.F공통_종료_채널_닫기()
	lib.F패닉억제_호출(소켓REP_TR콜백.Close)
}

func f초기화_TR전송_제한() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	TR코드_모음 := []string{
		xt.TR선물옵션_주문체결내역조회_CFOAQ00600,
		xt.TR선물옵션_정상주문_CFOAT00100,
		xt.TR선물옵션_정정주문_CFOAT00200,
		xt.TR선물옵션_취소주문_CFOAT00300,
		xt.TR선물옵션_예탁금_증거금_조회_CFOBQ10500,
		xt.TR선물옵션_미결제약정_현황_CFOFQ02400,
		xt.TR선물옵션_주문체결내역조회_CFOAQ00600,
		xt.TR선물옵션_정상주문_CFOAT00100,
		xt.TR선물옵션_정정주문_CFOAT00200,
		xt.TR선물옵션_취소주문_CFOAT00300,
		xt.TR선물옵션_예탁금_증거금_조회_CFOBQ10500,
		xt.TR선물옵션_미결제약정_현황_CFOFQ02400,
		xt.TR현물계좌_잔고내역_조회_CSPAQ12300,
		xt.TR현물계좌_주문체결내역_조회_CSPAQ13700,
		xt.TR현물_정상_주문_CSPAT00600,
		xt.TR현물_정정_주문_CSPAT00700,
		xt.TR현물_취소_주문_CSPAT00800,
		xt.TR현물_당일_매매일지_수수료_t0150,
		xt.TR현물_전일_매매일지_수수료_t0151,
		xt.TR시간_조회_t0167,
		xt.TR체결_미체결_조회_t0425,
		xt.TR현물_호가_조회_t1101,
		xt.TR현물_시세_조회_t1102,
		xt.TR현물_기간별_조회_t1305,
		xt.TR현물_당일_전일_분틱_조회_t1310,
		xt.TR관리_불성실_투자유의_조회_t1404,
		xt.TR투자경고_매매정지_정리매매_조회_t1405,
		xt.TR_ETF_시간별_추이_t1902,
		xt.TR기업정보_요약_t3320,
		xt.TR재무순위_종합_t3341,
		xt.TR현물_차트_틱_t8411,
		xt.TR현물_차트_분_t8412,
		xt.TR현물_차트_일주월_t8413,
		xt.TR증시_주변_자금_추이_t8428,
		xt.TR지수선물_마스터_조회_t8432,
		xt.TR현물_종목_조회_t8436}

	응답 := F질의(lib.New질의값_문자열_모음(xt.TR코드별_전송_제한, "", TR코드_모음), lib.P5초)
	lib.F확인(응답.G에러())

	전송_제한_정보_모음 = new(xt.TR코드별_전송_제한_정보_모음)
	lib.F확인(응답.G값(0, 전송_제한_정보_모음))
	lib.F조건부_패닉(len(TR코드_모음) != len(전송_제한_정보_모음.M배열),
		"서로 다른 길이 : '%v' '%v'", len(TR코드_모음), len(전송_제한_정보_모음.M배열))

	for _, 전송_제한_정보 := range 전송_제한_정보_모음.M배열 {
		TR코드 := 전송_제한_정보.TR코드

		if 전송_제한_정보.M초당_전송_제한 > 0 {
			tr코드별_초당_전송_제한[TR코드] = lib.New전송_권한(TR코드, 전송_제한_정보.M초당_전송_제한, lib.P1초)
		}

		if 전송_제한_정보.M10분당_전송_제한 > 0 {
			전송_권한 := lib.New전송_권한(TR코드, 전송_제한_정보.M10분당_전송_제한, lib.P10분)

			// 지난 10분간 이미 전송한 수량을 반영.
			for i := 0; i < 전송_제한_정보.M10분간_전송한_수량; i++ {
				전송_권한.G획득()
			}

			tr코드별_10분당_전송_제한[TR코드] = 전송_권한
		}

		if 전송_제한_정보.M초_베이스 > 1 {
			panic(lib.New에러("예상하지 못한 경우 : '%v' '%v' '%v' '%v'", TR코드,
				전송_제한_정보.M10분당_전송_제한,
				전송_제한_정보.M초당_전송_제한,
				전송_제한_정보.M초_베이스))
		}
	}

	return nil
}

//func f자료형_크기_비교_확인() (에러 error) {
//	lib.S예외처리{M에러: &에러}.S실행()
//
//	lib.F조건부_패닉(xt.Sizeof_C_TR_DATA != C.sizeof_TR_DATA, "C.TR_DATA 크기 불일치 ", xt.Sizeof_C_TR_DATA, C.sizeof_TR_DATA)
//	lib.F조건부_패닉(xt.Sizeof_C_MSG_DATA != C.sizeof_MSG_DATA, "C.MSG_DATA 크기 불일치 ")
//	lib.F조건부_패닉(xt.Sizeof_C_REALTIME_DATA != C.sizeof_REALTIME_DATA, "C.REALTIME_DATA 크기 불일치 ")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.TR_DATA{}) != unsafe.Sizeof(C.TR_DATA_UNPACKED{}), "TR_DATA 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.REALTIME_DATA{}) != unsafe.Sizeof(C.REALTIME_DATA_UNPACKED{}), "REALTIME_DATA_UNPACKED 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.MSG_DATA{}) != unsafe.Sizeof(C.MSG_DATA_UNPACKED{}), "MSG_DATA_UNPACKED 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00600InBlock1{}) != unsafe.Sizeof(C.CSPAT00600InBlock1{}), "CSPAT00600InBlock1 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00600OutBlock1{}) != unsafe.Sizeof(C.CSPAT00600OutBlock1{}), "CSPAT00600OutBlock1 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00600OutBlock2{}) != unsafe.Sizeof(C.CSPAT00600OutBlock2{}), "CSPAT00600OutBlock2 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00600OutBlock{}) != unsafe.Sizeof(C.CSPAT00600OutBlock{}), "CSPAT00600OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00700InBlock1{}) != unsafe.Sizeof(C.CSPAT00700InBlock1{}), "CSPAT00700InBlock1 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00700OutBlock1{}) != unsafe.Sizeof(C.CSPAT00700OutBlock1{}), "CSPAT00700OutBlock1 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00700OutBlock2{}) != unsafe.Sizeof(C.CSPAT00700OutBlock2{}), "CSPAT00700OutBlock2 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00700OutBlock{}) != unsafe.Sizeof(C.CSPAT00700OutBlock{}), "CSPAT00700OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00800InBlock1{}) != unsafe.Sizeof(C.CSPAT00800InBlock1{}), "CSPAT00800InBlock1 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00800OutBlock1{}) != unsafe.Sizeof(C.CSPAT00800OutBlock1{}), "CSPAT00800OutBlock1 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00800OutBlock2{}) != unsafe.Sizeof(C.CSPAT00800OutBlock2{}), "CSPAT00800OutBlock2 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.CSPAT00800OutBlock{}) != unsafe.Sizeof(C.CSPAT00800OutBlock{}), "CSPAT00800OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.SC0_OutBlock{}) != unsafe.Sizeof(C.SC0_OutBlock{}), "SC0_OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.SC1_OutBlock{}) != unsafe.Sizeof(C.SC1_OutBlock{}), "SC1_OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.SC2_OutBlock{}) != unsafe.Sizeof(C.SC2_OutBlock{}), "SC2_OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.SC3_OutBlock{}) != unsafe.Sizeof(C.SC3_OutBlock{}), "SC3_OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.SC4_OutBlock{}) != unsafe.Sizeof(C.SC4_OutBlock{}), "SC4_OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T0167OutBlock{}) != unsafe.Sizeof(C.T0167OutBlock{}), "T0167OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1101InBlock{}) != unsafe.Sizeof(C.T1101InBlock{}), "T1101InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1101OutBlock{}) != unsafe.Sizeof(C.T1101OutBlock{}), "T1101OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1102InBlock{}) != unsafe.Sizeof(C.T1102InBlock{}), "T1102InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1102OutBlock{}) != unsafe.Sizeof(C.T1102OutBlock{}), "T1102OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1305InBlock{}) != unsafe.Sizeof(C.T1305InBlock{}), "T1305InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1305OutBlock{}) != unsafe.Sizeof(C.T1305OutBlock{}), "T1305OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1305OutBlock1{}) != unsafe.Sizeof(C.T1305OutBlock1{}), "T1305OutBlock1 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1310InBlock{}) != unsafe.Sizeof(C.T1310InBlock{}), "T1310InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1310OutBlock{}) != unsafe.Sizeof(C.T1310OutBlock{}), "T1310OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1310OutBlock1{}) != unsafe.Sizeof(C.T1310OutBlock1{}), "T1310OutBlock1 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1901InBlock{}) != unsafe.Sizeof(C.T1901InBlock{}), "T1901InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1901OutBlock{}) != unsafe.Sizeof(C.T1901OutBlock{}), "T1901OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1902InBlock{}) != unsafe.Sizeof(C.T1902InBlock{}), "T1902InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1902OutBlock{}) != unsafe.Sizeof(C.T1902OutBlock{}), "T1902OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T1902OutBlock1{}) != unsafe.Sizeof(C.T1902OutBlock1{}), "T1902OutBlock1 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T3320InBlock{}) != unsafe.Sizeof(C.T3320InBlock{}), "T3320InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T3320OutBlock{}) != unsafe.Sizeof(C.T3320OutBlock{}), "T3320OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T3320OutBlock1{}) != unsafe.Sizeof(C.T3320OutBlock1{}), "T3320OutBlock1 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8411InBlock{}) != unsafe.Sizeof(C.T8411InBlock{}), "T8411InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8411OutBlock{}) != unsafe.Sizeof(C.T8411OutBlock{}), "T8411OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8411OutBlock1{}) != unsafe.Sizeof(C.T8411OutBlock1{}), "T8411OutBlock1 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8412InBlock{}) != unsafe.Sizeof(C.T8412InBlock{}), "T8412InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8412OutBlock{}) != unsafe.Sizeof(C.T8412OutBlock{}), "T8412OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8412OutBlock1{}) != unsafe.Sizeof(C.T8412OutBlock1{}), "T8412OutBlock1 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8413InBlock{}) != unsafe.Sizeof(C.T8413InBlock{}), "T8413InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8413OutBlock{}) != unsafe.Sizeof(C.T8413OutBlock{}), "T8413OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8413OutBlock1{}) != unsafe.Sizeof(C.T8413OutBlock1{}), "T8413OutBlock1 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8428InBlock{}) != unsafe.Sizeof(C.T8428InBlock{}), "T8428InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8428OutBlock{}) != unsafe.Sizeof(C.T8428OutBlock{}), "T8428OutBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8428OutBlock1{}) != unsafe.Sizeof(C.T8428OutBlock1{}), "T8428OutBlock1 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8436InBlock{}) != unsafe.Sizeof(C.T8436InBlock{}), "T8436InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.T8436OutBlock{}) != unsafe.Sizeof(C.T8436OutBlock{}), "T8436OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.H1_InBlock{}) != unsafe.Sizeof(C.H1_InBlock{}), "H1_InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.H1_OutBlock{}) != unsafe.Sizeof(C.H1_OutBlock{}), "H1_OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.H2_InBlock{}) != unsafe.Sizeof(C.H2_InBlock{}), "H2_InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.H2_OutBlock{}) != unsafe.Sizeof(C.H2_OutBlock{}), "H2_OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.S3_InBlock{}) != unsafe.Sizeof(C.S3_InBlock{}), "S3_InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.S3_OutBlock{}) != unsafe.Sizeof(C.S3_OutBlock{}), "S3_OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.YS3InBlock{}) != unsafe.Sizeof(C.YS3InBlock{}), "YS3InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.YS3OutBlock{}) != unsafe.Sizeof(C.YS3OutBlock{}), "YS3OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.I5_InBlock{}) != unsafe.Sizeof(C.I5_InBlock{}), "I5_InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.I5_OutBlock{}) != unsafe.Sizeof(C.I5_OutBlock{}), "I5_OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.VI_InBlock{}) != unsafe.Sizeof(C.VI_InBlock{}), "VI_InBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.VI_OutBlock{}) != unsafe.Sizeof(C.VI_OutBlock{}), "VI_OutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.DVIInBlock{}) != unsafe.Sizeof(C.DVIInBlock{}), "DVIInBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.DVIOutBlock{}) != unsafe.Sizeof(C.DVIOutBlock{}), "DVIOutBlock 크기 불일치")
//
//	lib.F조건부_패닉(unsafe.Sizeof(xt.JIFInBlock{}) != unsafe.Sizeof(C.JIFInBlock{}), "JIFInBlock 크기 불일치")
//	lib.F조건부_패닉(unsafe.Sizeof(xt.JIFOutBlock{}) != unsafe.Sizeof(C.JIFOutBlock{}), "JIFOutBlock 크기 불일치")
//
//	return nil
//}
