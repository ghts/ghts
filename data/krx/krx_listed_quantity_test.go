package krx

import (
	lb "github.com/ghts/ghts/lib"
	"testing"
)

func TestF상장_주식_수량_맵(t *testing.T) {
	상장_주식_수량_맵, 에러 := F상장_주식_수량_맵()
	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_참임(t, len(상장_주식_수량_맵) > 2000)
}
