package xing

import (
	lb "github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestCSPAQ13700_현물계좌_주문체결내역_질의값(t *testing.T) {
	t.Parallel()

	_, ok := interface{}(new(xt.CSPAQ13700_현물계좌_주문체결내역_질의값)).(lb.I질의값)

	lb.F테스트_참임(t, ok)
}

func TestCSPAQ13700_현물계좌_주문체결내역(t *testing.T) {
	t.Skip() // t.Parallel()

	계좌번호, 에러 := F계좌_번호(0)
	lb.F테스트_에러없음(t, 에러)

	for i := 0; i < 30; i++ {
		일자 := F당일().AddDate(0, 0, -1*i)
		testCSPAQ13700_현물계좌_주문체결내역_도우미(t, 계좌번호, 일자)
	}
}

func testCSPAQ13700_현물계좌_주문체결내역_도우미(t *testing.T, 계좌번호 string, 일자 time.Time) {
	값_모음, 에러 := TrCSPAQ13700_현물계좌_주문체결내역(
		계좌번호, 일자, lb.P매도_매수_전체, xt.CSPAQ13700_체결)
	lb.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lb.F테스트_참임(t, 값.M주문일.Before(F당일().AddDate(0, 0, 1)))
		lb.F테스트_다름(t, 값.M관리지점번호, "")
		lb.F테스트_다름(t, 값.M주문시장코드, "")
		lb.F테스트_참임(t, 값.M주문번호 >= 0)
		lb.F테스트_참임(t, 값.M원주문번호 >= 0)
		lb.F테스트_참임(t, F종목코드_존재함(값.M종목코드), 값.M종목코드)
		lb.F테스트_다름(t, 값.M종목명, "")
		lb.F테스트_같음(t, 값.M매도_매수_구분, lb.P매도, lb.P매수)
		lb.F테스트_같음(t, 값.M주문유형,
			xt.P주문유형_해당없음, xt.P주문_현금매도, xt.P주문_현금매수,
			xt.P주문_신용매도, xt.P주문_신용매수, xt.P주문_저축매도,
			xt.P주문_저축매수, xt.P주문_상품매도_대차, xt.P주문_상품매도,
			xt.P주문_상품매수, xt.P주문_현금매수_유가, xt.P주문_현금매수_정리, xt.P주문_장외매매) //xt.P주문_선물대용매도_일반, xt.P주문_선물대용매도_반대,
		lb.F테스트_같음(t, 값.M주문처리유형,
			xt.CSPAQ13700_정상처리, xt.CSPAQ13700_정정확인, xt.CSPAQ13700_정정거부_채권,
			xt.CSPAQ13700_취소확인, xt.CSPAQ13700_취소거부_채권)
		lb.F테스트_같음(t, 값.M정정취소구분, lb.P신규, lb.P정정, lb.P취소)
		lb.F테스트_참임(t, 값.M정정취소수량 >= 0)
		lb.F테스트_참임(t, 값.M정정취소가능수량 >= 0)
		lb.F테스트_참임(t, 값.M주문수량 >= 0)
		lb.F테스트_참임(t, 값.M주문가격 >= 0)
		lb.F테스트_참임(t, 값.M체결수량 >= 0)
		lb.F테스트_참임(t, 값.M체결가 >= 0)
		lb.F테스트_참임(t, 값.M체결처리시각.Hour() >= 9 && 값.M체결처리시각.Hour() <= 16, 값.M체결처리시각.Hour())
		lb.F테스트_참임(t, 값.M최종체결시각.Hour() >= 9 && 값.M최종체결시각.Hour() <= 16, 값.M최종체결시각.Hour())
		lb.F테스트_같음(t, 값.M호가유형,
			lb.P호가_지정가, lb.P호가_시장가, lb.P호가_조건부_지정가,
			lb.P호가_최유리_지정가, lb.P호가_최우선_지정가, lb.P호가_중간가,
			lb.P호가_장전_시간외, lb.P호가_장후_시간외, lb.P호가_시간외_단일가)
		lb.F테스트_같음(t, 값.M주문조건, lb.P주문조건_없음, lb.P주문조건_IOC, lb.P주문조건_FOK)
		lb.F테스트_참임(t, 값.M전체체결수량 >= 0)

		// 모의서버에서는 '50'이 수신되는 버그가 존재함. 게시판 질답에서 확인됨.
		lb.F테스트_같음(t, 값.M통신매체, xt.T통신매체구분(50),
			xt.P통신매체_아이폰, xt.P통신매체_안드로이드, xt.P통신매체_API, xt.P통신매체_HTS)
		lb.F테스트_다름(t, 값.M회원번호, "")
		lb.F테스트_같음(t, 값.M예약주문여부, xt.CSPAQ13700_예약주문_아님, xt.CSPAQ13700_예약주문)
		lb.F테스트_참임(t, 값.M대출일.Before(F당일().AddDate(0, 0, 1)))
		lb.F테스트_참임(t, 값.M주문시각.Hour() >= 9 && 값.M주문시각.Hour() <= 16, 값.M주문시각.Hour())
		lb.F테스트_다름(t, 값.M운용지시번호, "")
		lb.F테스트_다름(t, 값.M주문자ID, "")
	}
}
