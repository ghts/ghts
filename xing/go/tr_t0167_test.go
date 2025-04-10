package xing

import (
	lb "github.com/ghts/ghts/lib"

	"testing"
	"time"
)

func TestT0167_시각_조회(t *testing.T) {
	t.Parallel()

	시각, 에러 := (<-TrT0167_시각_조회()).G값()
	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_같음(t, 시각.Year(), time.Now().Year())
	lb.F테스트_같음(t, 시각.Month(), time.Now().Month())
	lb.F테스트_같음(t, 시각.Day(), time.Now().Day())

	지금 := time.Now()
	차이 := 시각.Sub(지금)
	lb.F테스트_참임(t, 차이 > (-1*lb.P1시간) && 차이 < lb.P1시간, 시각, 지금)
}
