package btft

import (
	bfc "github.com/ghts/ghts/backtest/factor/common"
	"github.com/ghts/ghts/lib"
	"testing"
)

func TestF자료형_일치_여부(t *testing.T) {
	_, ok := interface{}(new(S데이터_처리기_백테스트)).(I데이터_처리기)
	lib.F테스트_참임(t, ok)

	_, ok = interface{}(new(S포트폴리오)).(I포트폴리오)
	lib.F테스트_참임(t, ok)

	_, ok = interface{}(new(S전략_실행기[*S팩터_세종, *bfc.S재무_세종])).(I전략_실행기)
	lib.F테스트_참임(t, ok)
}
