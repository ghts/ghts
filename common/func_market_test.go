package common

import (
	"testing"
	"time"
)

func TestF한국증시_최근_개장일(테스트 *testing.T) {
	지금 := time.Now()
	한달전 := 지금.Add(-30 * 24 * time.Hour)

	최근_개장일_함수_모음 := [](func() (time.Time, error)) {
		f최근_개장일_다음넷,
		f최근_개장일_네이버,
		f최근_개장일_네이트,
		f최근_개장일_야후,
		F한국증시_최근_개장일}
	
	for _, 도우미_함수 := range 최근_개장일_함수_모음 {
		최근_개장일, 에러 := 도우미_함수()
		
		F테스트_에러없음(테스트, 에러)
		F테스트_거짓임(테스트, 최근_개장일.IsZero())
		F테스트_참임(테스트, 최근_개장일.After(한달전), 최근_개장일)
		F테스트_참임(테스트, 최근_개장일.Before(지금), 최근_개장일)
		F테스트_같음(테스트, 최근_개장일.Location(), 지금.Location())
	}
}