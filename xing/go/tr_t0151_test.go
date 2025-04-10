package xing

import (
	lb "github.com/ghts/ghts/lib"
	"testing"
)

func TestT0151_일자별_매매일지(t *testing.T) {
	t.Parallel()

	계좌번호, 에러 := F계좌_번호(0)
	lb.F테스트_에러없음(t, 에러)

	값_모음, 에러 := TrT0151_현물_일자별_매매일지(계좌번호, F전일())
	lb.F테스트_에러없음(t, 에러)

	for _, 값 := range 값_모음 {
		lb.F테스트_같음(t, 값.M매도_매수_구분, lb.P매도, lb.P매수)
		lb.F테스트_참임(t, F종목코드_존재함(값.M종목코드), 값.M종목코드)
		lb.F테스트_참임(t, 값.M수량 > 0)
		lb.F테스트_참임(t, 값.M단가 > 0)
		lb.F테스트_참임(t, 값.M약정금액 > 0)

		// 매도 매수 상관없이 수수료는 종목 소계에서 구해집니다.
		// 종목 소계의 수수료 항목 확인 부탁드립니다.

		switch 값.M매도_매수_구분 {
		case lb.P매수:
			lb.F테스트_같음(t, 값.M농특세, 0)
		case lb.P매도:
			if ETF_ETN_종목_여부(값.M종목코드) {
				lb.F테스트_같음(t, 값.M거래세, 0)
				lb.F테스트_같음(t, 값.M농특세, 0)
			} else {
				예상_금액 := 0.003 * float64(값.M수량) * float64(값.M단가)
				lb.F테스트_참임(t, 값.M거래세 > 0 && float64(값.M거래세) < 예상_금액*1.1,
					값.M종목코드, 값.M거래세, 예상_금액)
				lb.F테스트_참임(t, 값.M농특세 > 0 && float64(값.M농특세) < 예상_금액*1.1,
					값.M종목코드, 값.M거래세, 예상_금액)
			}
		default:
			panic(lb.F2문자열("예상하지 못한 경우 : '%v' '%v'", int(값.M매도_매수_구분), 값.M매도_매수_구분.String()))
		}

		예상_정산금액 := (값.M단가 * 값.M수량) - 값.M거래세 - 값.M농특세 //  - 값.M수수료
		lb.F테스트_같음(t, 값.M정산금액, 예상_정산금액)
	}
}
