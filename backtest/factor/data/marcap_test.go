package data

import (
	"github.com/ghts/ghts/lib"
	"testing"
)

// 시가총액 데이터셋 불러들이기
// https://github.com/FinanceData/marcap
func TestCSV파일_읽기(t *testing.T) {
	for year := 1995; year <= 2022; year++ {
		파일명 := lib.F2문자열("marcap-%v.csv", year)
		레코드_모음, 에러 := csv파일_읽기(파일명)
		lib.F테스트_에러없음(t, 에러)
		lib.F테스트_참임(t, len(레코드_모음) > 1000)
		lib.F테스트_참임(t, len(레코드_모음[1]) == 18)
	}
}

func TestF시총_데이터_저장(t *testing.T) {
	lib.F테스트_에러없음(t, F시총_데이터_저장())
}
