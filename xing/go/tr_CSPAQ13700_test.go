package xing

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"testing"
	"time"
)

func TestCSPAQ13700_현물계좌_주문체결내역_질의값(t *testing.T) {
	t.Parallel()

	_, ok := interface{}(new(xt.CSPAQ13700_현물계좌_주문체결내역_질의값)).(lib.I질의값)

	lib.F테스트_참임(t, ok)
}

func TestCSPAQ13700_현물계좌_주문체결내역(t *testing.T) {
	t.Skip() // t.Parallel()

	계좌번호, 에러 := F계좌_번호(0)
	lib.F테스트_에러없음(t, 에러)

	for i := 0; i < 30; i++ {
		일자 := F당일().AddDate(0, 0, -1*i)
		testCSPAQ13700_현물계좌_주문체결내역_도우미(t, 계좌번호, 일자)
	}
}

func testCSPAQ13700_현물계좌_주문체결내역_도우미(t *testing.T, 계좌번호 string, 일자 time.Time) {
	값_모음, 에러 := TrCSPAQ13700_현물계좌_주문체결내역(
		계좌번호, 일자, lib.P매도_매수_전체, xt.CSPAQ13700_체결)
	lib.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lib.F테스트_참임(t, 값.M주문일.Before(F당일().AddDate(0, 0, 1)))
		lib.F테스트_다름(t, 값.M관리지점번호, "")
		lib.F테스트_다름(t, 값.M주문시장코드, "")
		lib.F테스트_참임(t, 값.M주문번호 >= 0)
		lib.F테스트_참임(t, 값.M원주문번호 >= 0)
		lib.F테스트_참임(t, F종목코드_존재함(값.M종목코드), 값.M종목코드)
		lib.F테스트_다름(t, 값.M종목명, "")
		lib.F테스트_같음(t, 값.M매도_매수_구분, lib.P매도, lib.P매수)
		lib.F테스트_같음(t, 값.M주문유형,
			xt.P주문유형_해당없음, xt.P주문_현금매도, xt.P주문_현금매수,
			xt.P주문_신용매도, xt.P주문_신용매수, xt.P주문_저축매도,
			xt.P주문_저축매수, xt.P주문_상품매도_대차, xt.P주문_상품매도,
			xt.P주문_상품매수, xt.P주문_현금매수_유가, xt.P주문_현금매수_정리, xt.P주문_장외매매) //xt.P주문_선물대용매도_일반, xt.P주문_선물대용매도_반대,
		lib.F테스트_같음(t, 값.M주문처리유형,
			xt.CSPAQ13700_정상처리, xt.CSPAQ13700_정정확인, xt.CSPAQ13700_정정거부_채권,
			xt.CSPAQ13700_취소확인, xt.CSPAQ13700_취소거부_채권)
		lib.F테스트_같음(t, 값.M정정취소구분, lib.P신규, lib.P정정, lib.P취소)
		lib.F테스트_참임(t, 값.M정정취소수량 >= 0)
		lib.F테스트_참임(t, 값.M정정취소가능수량 >= 0)
		lib.F테스트_참임(t, 값.M주문수량 >= 0)
		lib.F테스트_참임(t, 값.M주문가격 >= 0)
		lib.F테스트_참임(t, 값.M체결수량 >= 0)
		lib.F테스트_참임(t, 값.M체결가 >= 0)
		lib.F테스트_참임(t, 값.M체결처리시각.Hour() >= 9 && 값.M체결처리시각.Hour() <= 16, 값.M체결처리시각.Hour())
		lib.F테스트_참임(t, 값.M최종체결시각.Hour() >= 9 && 값.M최종체결시각.Hour() <= 16, 값.M최종체결시각.Hour())
		lib.F테스트_같음(t, 값.M호가유형,
			xt.P호가_지정가, xt.P호가_시장가, xt.P호가_조건부_지정가,
			xt.P호가_최유리_지정가, xt.P호가_최우선_지정가, xt.P호가_시장가_IOC,
			xt.P호가_최유리_지정가_IOC, xt.P호가_최유리_지정가_FOK, xt.P호가_지정가_전환,
			xt.P호가_지정가_IOC_전환, xt.P호가_지정가_FOK_전환, xt.P호가_부분충족_K_OTC,
			xt.P호가_전량충족_K_OTC, xt.P호가_장전_시간외, xt.P호가_장후_시간외,
			xt.P호가_시간외_단일가)
		lib.F테스트_같음(t, 값.M주문조건, lib.P주문조건_없음, lib.P주문조건_IOC, lib.P주문조건_FOK)
		lib.F테스트_참임(t, 값.M전체체결수량 >= 0)

		// 모의서버에서는 '50'이 수신되는 버그가 존재함. 게시판 질답에서 확인됨.
		lib.F테스트_같음(t, 값.M통신매체, xt.T통신매체구분(50),
			xt.P통신매체_아이폰, xt.P통신매체_안드로이드, xt.P통신매체_API, xt.P통신매체_HTS)
		lib.F테스트_다름(t, 값.M회원번호, "")
		lib.F테스트_같음(t, 값.M예약주문여부, xt.CSPAQ13700_예약주문_아님, xt.CSPAQ13700_예약주문)
		lib.F테스트_참임(t, 값.M대출일.Before(F당일().AddDate(0, 0, 1)))
		lib.F테스트_참임(t, 값.M주문시각.Hour() >= 9 && 값.M주문시각.Hour() <= 16, 값.M주문시각.Hour())
		lib.F테스트_다름(t, 값.M운용지시번호, "")
		lib.F테스트_다름(t, 값.M주문자ID, "")
	}
}
