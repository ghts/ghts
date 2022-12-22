package data

import (
	bfc "github.com/ghts/ghts/backtest/factor/common"
	"github.com/ghts/ghts/lib"
	"path"
	"sort"
	"testing"
)

func TestF세종_데이터_저장(t *testing.T) {
	F세종_데이터_저장()
}

func TestCSV파일_읽기_세종(t *testing.T) {
	맵 := make(map[string]*bfc.S재무_세종)

	for _, 구분 := range []T세종_데이터_구분{P매출_세종, P영업이익_세종, P당기순이익_세종, P자산_세종, P자본_세종, P부채_세종} {
		파일명_모음 := CSV파일_검색_세종(path.Join(lib.F홈_디렉토리(), "Downloads"), 구분.G파일명_조각())
		lib.F테스트_참임(t, len(파일명_모음) == 1, len(파일명_모음))

		파일명 := 파일명_모음[0]
		레코드_모음, 에러 := csv파일_읽기_세종(파일명)
		lib.F테스트_에러없음(t, 에러)
		lib.F테스트_참임(t, len(레코드_모음) > 1000, len(레코드_모음))

		에러 = csv파일_해석_세종(맵, 구분, 레코드_모음)
		lib.F테스트_에러없음(t, 에러)
	}

	for 키, 값 := range 맵 {
		if 값.M매출액 == 0 ||
			값.M영업이익 == 0 ||
			값.M당기순이익 == 0 ||
			값.M자산 == 0 ||
			값.M자본 == 0 ||
			값.M부채 == 0 {
			delete(맵, 키)
		}
	}

	// 데이터 누락 종목 확인.
	종목코드_맵 := make(map[string]lib.S비어있음)

	for _, 값 := range 맵 {
		if 값.M매출액 == 0.0 ||
			값.M영업이익 == 0 ||
			값.M당기순이익 == 0 ||
			값.M자산 == 0 ||
			값.M자본 == 0 ||
			값.M부채 == 0 {
			종목코드_맵[값.M종목코드] = lib.S비어있음{}
		}
	}

	종목코드_모음 := make([]string, len(종목코드_맵))

	i := 0
	for 종목코드 := range 종목코드_맵 {
		종목코드_모음[i] = 종목코드
		i++
	}

	sort.Strings(종목코드_모음)

	lib.F체크포인트(len(종목코드_모음))

	for _, 종목코드 := range 종목코드_모음 {
		println(종목코드)
	}
}
