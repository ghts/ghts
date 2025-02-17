package xing

import (
	"github.com/ghts/ghts/lib"
	mt "github.com/ghts/ghts/lib/market_time"
	"github.com/ghts/ghts/lib/nanomsg"
	xt "github.com/ghts/ghts/xing/base"

	"testing"
)

func TestF호가_잔량_실시간_정보(t *testing.T) {
	t.Parallel()

	if !mt.F한국증시_정규_거래_시간임() {
		t.SkipNow()
	}

	const 종목코드_코스피 = "005930" // 삼성전자
	const 종목코드_코스닥 = "091990" // 셀트리온 헬스케어
	const 종목코드_ETF = "069500" // KODEX 200

	종목_코스피, 에러 := F종목by코드(종목코드_코스피)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 종목_코스피.G시장구분(), lib.P시장구분_코스피)

	종목_코스닥, 에러 := F종목by코드(종목코드_코스닥)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 종목_코스닥.G시장구분(), lib.P시장구분_코스닥)

	종목_ETF, 에러 := F종목by코드(종목코드_ETF)
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_같음(t, 종목_ETF.G시장구분(), lib.P시장구분_ETF)

	소켓SUB_실시간 := lib.F확인2(nano.NewNano소켓SUB(xt.F주소_실시간()))
	lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_구독(종목코드_코스피))
	lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_구독(종목코드_코스닥))
	lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_구독(종목코드_ETF))

	defer func() {
		lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_코스피))
		lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_코스닥))
		lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_ETF))
	}()

	var 코스피_수신, 코스닥_수신, ETF_수신 bool

	// 실시간 정보 수신 확인
	for i := 0; i < 1000; i++ {
		바이트_변환_모음, 에러 := 소켓SUB_실시간.G수신()
		lib.F테스트_에러없음(t, 에러)

		i실시간_정보 := lib.F확인2(바이트_변환_모음.S해석기(xt.F바이트_변환값_해석).G해석값(0))

		값, ok := i실시간_정보.(*xt.S호가_잔량_실시간_정보)
		if !ok {
			continue
		}

		switch 값.M종목코드 {
		case 종목코드_코스피:
			코스피_수신 = true
			lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_코스피))
		case 종목코드_코스닥:
			코스닥_수신 = true
			lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_코스닥))
		case 종목코드_ETF:
			ETF_수신 = true
			lib.F테스트_에러없음(t, F호가_잔량_실시간_정보_해지(종목코드_ETF))
		}

		lib.F테스트_참임(t, 값.M중간가격 <= lib.F최소값(값.M매도호가_모음...))
		lib.F테스트_참임(t, 값.M중간가격 >= lib.F최대값(값.M매수호가_모음...))
		lib.F테스트_참임(t, 값.M매도중간가잔량합계수량 >= 0)
		lib.F테스트_참임(t, 값.M매수중간가잔량합계수량 >= 0)
		lib.F테스트_참임(t, 값.M중간가잔량합계수량 == 값.M매도중간가잔량합계수량+값.M매수중간가잔량합계수량)
		lib.F테스트_같음(t, 값.M중간가잔량구분, lib.P매도_매수_전체, lib.P매도, lib.P매수)

		if 코스피_수신 && 코스닥_수신 && ETF_수신 {
			return
		}
	}

	t.FailNow()
}
