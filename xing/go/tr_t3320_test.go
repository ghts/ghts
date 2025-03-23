package xing

import (
	"github.com/ghts/ghts/lib"
	"math"
	"strings"
	"testing"
)

func TestT3320_F기업정보_요약(t *testing.T) {
	t.Parallel()

	//종목코드 := "051910" // LG전자
	종목코드 := "060380"

	값, 에러 := TrT3320_F기업정보_요약(종목코드)
	lib.F테스트_에러없음(t, 에러)

	값1 := 값.M응답1
	lib.F테스트_다름(t, strings.TrimSpace(값1.M업종구분명), "")
	lib.F테스트_다름(t, strings.TrimSpace(값1.M시장구분), "")
	lib.F테스트_다름(t, strings.TrimSpace(값1.M시장구분명), "")
	lib.F테스트_다름(t, strings.TrimSpace(값1.M한글기업명), "")
	lib.F테스트_다름(t, strings.TrimSpace(값1.M본사주소), "")
	lib.F테스트_다름(t, lib.F정규식_검색(값1.M본사전화번호, []string{`[1-9]+`}), "")
	lib.F테스트_다름(t, lib.F정규식_검색(값1.M최근결산년도, []string{`[1-9]+`}), "")
	lib.F테스트_다름(t, lib.F정규식_검색(값1.M결산월, []string{`[1-9]+`}), "")
	lib.F테스트_다름(t, lib.F정규식_검색(값1.M최근결산년월, []string{`[1-9]+`}), "")
	lib.F테스트_참임(t, 값1.M주당액면가 > 0)
	lib.F테스트_참임(t, 값1.M주식수 > 0)
	lib.F테스트_다름(t, strings.TrimSpace(값1.M홈페이지), "")
	//lib.F테스트_다름(t, strings.TrimSpace(값1.M그룹명), "")
	lib.F테스트_참임(t, 값1.M외국인_비중 >= 0)
	lib.F테스트_다름(t, lib.F정규식_검색(값1.M주담전화, []string{`[1-9]+`}), "")
	lib.F테스트_참임(t, 값1.M자본금_억 > 0)
	lib.F테스트_참임(t, 값1.M시가총액 > 0)
	lib.F테스트_참임(t, 값1.M배당금 >= 0)
	lib.F테스트_참임(t, 값1.M배당수익율 >= 0)
	lib.F테스트_참임(t, 값1.M현재가 > 0)
	lib.F테스트_참임(t, 값1.M전일종가 > 0)
	lib.F테스트_참임(t, 값1.M현재가 > int64(float64(값1.M전일종가)*0.4) || 값1.M현재가 < int64(float64(값1.M전일종가)*1.4))

	값2 := 값.M응답2
	lib.F테스트_같음(t, strings.TrimSpace(값2.M종목코드), 종목코드)
	lib.F테스트_다름(t, strings.TrimSpace(값2.M결산년월), "")
	lib.F테스트_다름(t, strings.TrimSpace(값2.M결산구분), "")
	//lib.F테스트_참임(t, math.Abs(값2.PER) < 500, 값2.PER)
	//값2.EPS      = lib.F확인2(lib.F2실수(g.Eps)
	lib.F테스트_참임(t, 값2.PBR > 0)
	lib.F테스트_참임(t, math.Abs(값2.ROA) < 100, 값2.ROA)
	lib.F테스트_참임(t, math.Abs(값2.ROE) < 100, 값2.ROE)
	//값2.EBITDA   = lib.F확인2(lib.F2실수(g.Ebitda)
	lib.F테스트_참임(t, math.Abs(값2.EVEBITDA) < 100, 값2.EVEBITDA)
	lib.F테스트_참임(t, 값2.M액면가 > 0)
	//값2.SPS      = lib.F확인2(lib.F2실수(g.Sps)
	//값2.CPS      = lib.F확인2(lib.F2실수(g.Cps)
	//값2.BPS      = lib.F확인2(lib.F2실수(g.Bps)
	//값2.T_PER    = lib.F확인2(lib.F2실수(g.Tper)
	//값2.T_EPS    = lib.F확인2(lib.F2실수(g.Teps)
	//값2.PEG      = lib.F확인2(lib.F2실수(g.Peg)
	//값2.T_PEG    = lib.F확인2(lib.F2실수(g.Tpeg)
	lib.F테스트_다름(t, lib.F정규식_검색(값2.M최근분기년도, []string{`[1-9]+`}), "")
}
