package xing

import (
	lb "github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"testing"
)

func TestT8436_주식종목_조회(t *testing.T) {
	t.Parallel()

	시장_구분 := ([]lb.T시장구분{lb.P시장구분_전체, lb.P시장구분_코스피, lb.P시장구분_코스닥})[lb.F임의_범위_이내_정수값(0, 2)]

	값_모음, 에러 := TrT8436_주식종목_조회(시장_구분)
	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_참임(t, len(값_모음) > 0, len(값_모음))

	for _, 응답값 := range 값_모음 {
		종목, 에러 := F종목by코드(응답값.M종목코드)
		lb.F테스트_에러없음(t, 에러)

		인덱스 := lb.F최소값(len(응답값.M종목명), len(종목.G이름()))
		lb.F테스트_같음(t, 응답값.M종목명[:인덱스], 종목.G이름()[:인덱스])

		switch 응답값.M시장구분 {
		case lb.P시장구분_ETF:
			lb.F테스트_다름(t, 응답값.M증권그룹, xt.P증권그룹_주식)
			lb.F테스트_같음(t, 응답값.M증권그룹, xt.P증권그룹_상장지수펀드_ETF, xt.P증권그룹_해외ETF)
		case lb.P시장구분_ETN:
			lb.F테스트_다름(t, 응답값.M증권그룹, xt.P증권그룹_주식)
			lb.F테스트_같음(t, 응답값.M증권그룹, xt.P증권그룹_ETN)
		case lb.P시장구분_코스피, lb.P시장구분_코스닥:
			// 다음넷의 ETF종목정보가 불완전해서 주식만 테스트 함.
			lb.F테스트_같음(t, 응답값.M시장구분, 종목.G시장구분())
		default:
			panic(lb.F2문자열("예상하지 못한 경우 : '%v' '%v'", int(응답값.M시장구분), 응답값.M시장구분.String()))
		}

		if 응답값.M증권그룹 == xt.P증권그룹_주식 {
			lb.F테스트_같음(t, 응답값.M시장구분, lb.P시장구분_전체, lb.P시장구분_코스피, lb.P시장구분_코스닥)
		}

		//lb.F테스트_같음(t, 응답값.M주문수량단위, 1)

		// 상한가 예상값 계산에 예외가 너무 많아서 건너뜀.
		//호가단위, 에러 := lb.F최소_호가단위by시장구분_기준가(종목.G시장구분(), 응답값.M전일가)
		//lb.F테스트_에러없음(t, 에러)
		//예상값_상한가 := int64(float64(응답값.M전일가) * 1.3)
		//오차_상한가 := lb.F2절대값_정수64(응답값.M상한가 - 예상값_상한가)
		//오차율_상한가 := float64(오차_상한가) / float64(응답값.M상한가) * 100
		//lb.F테스트_참임(t,  오차_상한가 <= 호가단위 || 오차율_상한가 < 3,
		//	응답값.M종목코드, 응답값.M상한가, 예상값_상한가, 오차_상한가, 호가단위, 오차율_상한가)

		// 액면분할 하면 전일가가 상한가보다 높아짐.
		//lb.F테스트_참임(t, 응답값.M상한가 == 0 || 응답값.M상한가 > 응답값.M전일가, 응답값.M상한가, 응답값.M전일가)
		lb.F테스트_참임(t, 응답값.M상한가 == 0 || 응답값.M상한가 > 응답값.M하한가, 응답값.M상한가, 응답값.M하한가)
		lb.F테스트_참임(t, 응답값.M상한가 == 0 || 응답값.M상한가 > 응답값.M기준가, 응답값.M상한가, 응답값.M기준가)

		// 하한가 예상값 계산의 예외가 너무 많아서 건너뜀.
		//예상값_하한가 := int64(float64(응답값.M전일가) * 0.7)
		//오차_하한가 := lb.F2절대값_정수64(응답값.M하한가 - 예상값_하한가)
		//오차율_하한가 := float64(오차_하한가) / float64(응답값.M하한가) * 100
		//lb.F테스트_참임(t, 오차_하한가 <= 호가단위 || 오차율_하한가 < 3,
		//	응답값.M종목코드, 응답값.M하한가, 예상값_하한가, 오차_하한가, 호가단위, 오차율_하한가)

		// 액면분할 하면 전일가가 하한가보다 높아짐.
		//lb.F테스트_참임(t, 응답값.M전일가 == 0 || 응답값.M하한가 <= 응답값.M전일가, 응답값.M종목코드, 응답값.M하한가, 응답값.M전일가)
		lb.F테스트_참임(t, 응답값.M하한가 == 0 || 응답값.M하한가 <= 응답값.M기준가, 응답값.M종목코드, 응답값.M하한가, 응답값.M기준가)
		lb.F테스트_같음(t, 응답값.M증권그룹, xt.P증권그룹_주식, xt.P증권그룹_예탁증서,
			xt.P증권그룹_증권투자회사_뮤추얼펀드, xt.P증권그룹_Reits종목, xt.P증권그룹_상장지수펀드_ETF,
			xt.P증권그룹_선박투자회사, xt.P증권그룹_인프라투융자회사, xt.P증권그룹_해외ETF,
			xt.P증권그룹_해외원주, xt.P증권그룹_ETN)
	}
}
