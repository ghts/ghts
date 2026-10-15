package trade

import (
	"testing"

	lb "github.com/ghts/ghts/lib"
)

// TestF와일더_이동_평균_입력_검증 : 무효한 윈도우_크기 또는 비어있는 입력에는 패닉이 발생해야 한다.
func TestF와일더_이동_평균_입력_검증(t *testing.T) {
	t.Parallel()

	값_모음 := []float64{1, 2, 3, 4, 5}

	lb.F테스트_패닉_발생(t, func() { _ = F와일더_이동_평균(값_모음, 0) })
	lb.F테스트_패닉_발생(t, func() { _ = F와일더_이동_평균(값_모음, -3) })
	lb.F테스트_패닉_발생(t, func() { _ = F와일더_이동_평균([]float64{}, 14) })

	결과_모음 := F와일더_이동_평균(값_모음, 3)
	lb.F테스트_다름(t, len(결과_모음), len(값_모음))
}
