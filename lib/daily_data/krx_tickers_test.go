package daily_data

import (
	"github.com/ghts/ghts/lib"
	"testing"
)

func TestF상장_법인_정보_맵(t *testing.T) {
	법인정보_맵, 에러 := F상장_법인_정보_맵()
	lib.F테스트_에러없음(t, 에러)
	lib.F테스트_참임(t, len(법인정보_맵) > 2_000)
}
