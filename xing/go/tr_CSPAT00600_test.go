package xing

import (
	lb "github.com/ghts/ghts/lib"
	krx "github.com/ghts/ghts/lib/market_time"
	"github.com/ghts/ghts/lib/nanomsg"
	"github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestCSPAT00600_현물_정상_주문_질의값(t *testing.T) {
	t.Parallel()

	_, ok := interface{}(new(xt.CSPAT00600_현물_정상_주문_질의값)).(lb.I질의값)
	lb.F테스트_참임(t, ok)
}

func TestCSPAT00600_현물_정상_주문(t *testing.T) {
	t.Parallel()

	if xt.F서버_구분() == xt.P서버_실거래 ||
		!F금일_한국증시_개장() ||
		!krx.F한국증시_정규_거래_시간임() {
		t.SkipNow()
	}

	소켓SUB_실시간 := lb.F확인2(nano.NewNano소켓SUB(xt.F주소_실시간()))
	lb.F대기(lb.P1초)

	lb.F테스트_에러없음(t, F주문_응답_실시간_정보_구독())

	const 반복_횟수 = 10
	const 수량 = 5 // 주문이 정상 작동하는 지만 확인하면 됨.
	const 호가_유형 = lb.P호가_시장가

	var 매수_주문_접수_확인_수량, 매도_주문_접수_확인_수량 int
	var 매수_체결_수량, 매도_체결_수량 int64
	var 매수_주문번호_모음, 매도_주문번호_모음 = make([]int64, 반복_횟수), make([]int64, 반복_횟수)
	var 종목 = lb.New종목("069500", "KODEX 200", lb.P시장구분_ETF)
	var 가격_정상주문 = int64(0)
	var p1분전 = time.Now().Add(-1 * lb.P1분)
	var p1분후 = time.Now().Add(lb.P1분)

	계좌번호, 에러 := F계좌_번호(0)
	lb.F테스트_에러없음(t, 에러)

	//계좌_상세명, 에러 := F계좌_상세명(계좌번호)
	//lb.F확인(에러)
	//lb.F테스트_거짓임(t, strings.Contains(계좌_상세명, "선물옵션")) // 현물 계좌이어야 함.

	질의값_매수 := xt.NewCSPAT00600_현물_정상_주문_질의값()
	질의값_매수.M계좌번호 = 계좌번호
	질의값_매수.M종목코드 = 종목.G코드()
	질의값_매수.M주문수량 = 수량
	질의값_매수.M주문단가 = 가격_정상주문 // 시장가 주문 시 가격은 무조건 '0'을 입력해야 함.
	질의값_매수.M매도_매수_구분 = lb.P매수
	질의값_매수.M호가유형 = 호가_유형
	질의값_매수.M주문조건 = lb.P주문조건_없음 // 모의투자에서는 IOC, FOK를 사용할 수 없음.
	질의값_매수.M신용거래_구분 = xt.P신용거래_해당없음
	질의값_매수.M대출일 = time.Time{}

	for i := 0; i < 반복_횟수; i++ {
		응답값, 에러 := TrCSPAT00600_현물_정상주문(질의값_매수)

		lb.F대기(lb.P100밀리초)

		lb.F테스트_에러없음(t, 에러)
		lb.F테스트_다름(t, 응답값.M응답1, nil)
		lb.F테스트_같음(t, 응답값.M응답1.M주문수량, 수량)
		lb.F테스트_같음(t, 응답값.M응답1.M종목코드, 종목.G코드())
		lb.F테스트_같음(t, 응답값.M응답1.M호가유형, 호가_유형)
		lb.F테스트_같음(t, 응답값.M응답1.M주문가격, 가격_정상주문)
		lb.F테스트_같음(t, 응답값.M응답1.M신용거래_구분, xt.P신용거래_해당없음)
		lb.F테스트_같음(t, 응답값.M응답1.M주문조건_구분, lb.P주문조건_없음)
		lb.F테스트_다름(t, 응답값.M응답2, nil)
		lb.F테스트_같음(t, 응답값.M응답2.M종목코드, 종목.G코드())
		lb.F테스트_참임(t, 응답값.M응답2.M주문번호 > 0)
		lb.F테스트_참임(t, 응답값.M응답2.M주문시각.After(p1분전), 응답값.M응답2.M주문시각, p1분전)
		lb.F테스트_참임(t, 응답값.M응답2.M주문시각.Before(p1분후), 응답값.M응답2.M주문시각, p1분후)

		매수_주문번호_모음[i] = 응답값.M응답2.M주문번호
	}

	// 매수 주문 확인
	for {
		바이트_변환_모음, 에러 := 소켓SUB_실시간.G수신()
		lb.F테스트_에러없음(t, 에러)

		실시간_정보, ok := lb.F확인2(바이트_변환_모음.S해석기(xt.F바이트_변환값_해석).G해석값(0)).(*xt.S현물_주문_응답_실시간_정보)

		switch {
		case !ok:
			continue
		case !f주문번호_포함(실시간_정보.M주문번호, 매수_주문번호_모음):
			continue
		}

		switch 실시간_정보.RT코드 {
		case xt.RT현물_주문_거부_SC4:
			lb.F문자열_출력("매수 주문 거부됨 : '%v'", 실시간_정보.M주문번호)
			t.FailNow()
		case xt.RT현물_주문_정정_SC2, xt.RT현물_주문_취소_SC3:
			lb.F문자열_출력("예상하지 못한 TR코드 : '%v'", 실시간_정보.RT코드)
			t.FailNow()
		case xt.RT현물_주문_접수_SC0:
			매수_주문_접수_확인_수량++
		case xt.RT현물_주문_체결_SC1:
			매수_체결_수량 = 매수_체결_수량 + 실시간_정보.M수량
		}

		if 매수_주문_접수_확인_수량 == 반복_횟수 &&
			매수_체결_수량 == 반복_횟수*수량 {
			break
		}
	}

	질의값_매도 := xt.NewCSPAT00600_현물_정상_주문_질의값()
	질의값_매도.M계좌번호 = 계좌번호
	질의값_매도.M종목코드 = 종목.G코드()
	질의값_매도.M주문수량 = 수량
	질의값_매도.M주문단가 = 가격_정상주문
	질의값_매도.M매도_매수_구분 = lb.P매도
	질의값_매도.M호가유형 = 호가_유형
	질의값_매도.M신용거래_구분 = xt.P신용거래_해당없음
	질의값_매도.M주문조건 = lb.P주문조건_없음
	질의값_매도.M대출일 = time.Time{}

	for i := 0; i < 반복_횟수; i++ {
		응답값, 에러 := TrCSPAT00600_현물_정상주문(질의값_매도)

		lb.F대기(lb.P100밀리초)

		lb.F테스트_에러없음(t, 에러)
		lb.F테스트_다름(t, 응답값.M응답1, nil)
		lb.F테스트_같음(t, 응답값.M응답1.M주문수량, 수량)
		lb.F테스트_같음(t, 응답값.M응답1.M종목코드, 종목.G코드())
		lb.F테스트_같음(t, 응답값.M응답1.M호가유형, 호가_유형)
		lb.F테스트_같음(t, 응답값.M응답1.M주문가격, 가격_정상주문)
		lb.F테스트_같음(t, 응답값.M응답1.M신용거래_구분, xt.P신용거래_해당없음)
		lb.F테스트_같음(t, 응답값.M응답1.M주문조건_구분, lb.P주문조건_없음)
		lb.F테스트_다름(t, 응답값.M응답2, nil)
		lb.F테스트_같음(t, 응답값.M응답2.M종목코드, 종목.G코드())
		lb.F테스트_참임(t, 응답값.M응답2.M주문번호 > 0)
		lb.F테스트_참임(t, 응답값.M응답2.M주문시각.After(p1분전), 응답값.M응답2.M주문시각, p1분전)
		lb.F테스트_참임(t, 응답값.M응답2.M주문시각.Before(p1분후), 응답값.M응답2.M주문시각, p1분후)

		매도_주문번호_모음[i] = 응답값.M응답2.M주문번호
	}

	// 매도 주문 확인
	for {
		바이트_변환_모음, 에러 := 소켓SUB_실시간.G수신()
		lb.F테스트_에러없음(t, 에러)

		실시간_정보, ok := lb.F확인2(바이트_변환_모음.S해석기(xt.F바이트_변환값_해석).G해석값(0)).(*xt.S현물_주문_응답_실시간_정보)

		switch {
		case !ok:
			continue
		case !f주문번호_포함(실시간_정보.M주문번호, 매도_주문번호_모음):
			continue
		}

		switch 실시간_정보.RT코드 {
		case xt.RT현물_주문_거부_SC4:
			lb.F문자열_출력("매도 주문 거부됨 : '%v'", 실시간_정보.M주문번호)
			t.FailNow()
		case xt.RT현물_주문_정정_SC2, xt.RT현물_주문_취소_SC3:
			lb.F문자열_출력("예상하지 못한 TR코드 : '%v'", 실시간_정보.RT코드)
			t.FailNow()
		case xt.RT현물_주문_접수_SC0:
			매도_주문_접수_확인_수량++
		case xt.RT현물_주문_체결_SC1:
			매도_체결_수량 = 매도_체결_수량 + 실시간_정보.M수량
		}

		if 매도_주문_접수_확인_수량 == 반복_횟수 &&
			매도_체결_수량 == 반복_횟수*수량 {
			break
		}
	}
}

func f주문번호_포함(주문번호 int64, 주문번호_모음 []int64) bool {
	for _, 주문번호2 := range 주문번호_모음 {
		if 주문번호 == 주문번호2 {
			return true
		}
	}

	return false
}
