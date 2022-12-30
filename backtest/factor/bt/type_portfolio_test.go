package btft

import (
	"github.com/ghts/ghts/lib"
	"math"
	"testing"
)

// 복리 수익율 계산 공식 확인
func TestF복리수익율(t *testing.T) {
	const 초기_투자금 = 100_000_000 // 1억
	const 수익율 float64 = 15.0 / 100.0
	const 투자_기간 = 10

	투자_결과 := 초기_투자금 * math.Pow(1+수익율, 투자_기간)
	산출_수익율 := math.Pow(math.E, math.Log(투자_결과/초기_투자금)/투자_기간) - 1.0

	lib.F테스트_참임(t, lib.F절대값((수익율-산출_수익율)/수익율) < 0.01)
}
