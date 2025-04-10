package xing

import (
	lb "github.com/ghts/ghts/lib"
	"strings"
	"testing"
)

func TestF샘플_종목_모음(t *testing.T) {
	t.Parallel()

	샘플_종목_모음 := lb.F샘플_종목_모음_전체()

	lb.F테스트_참임(t, len(샘플_종목_모음) > 0)

	for _, 종목 := range 샘플_종목_모음 {
		lb.F테스트_참임(t, F종목코드_존재함(종목.G코드()), 종목.G식별_문자열())

		종목_비교값, 에러 := F종목by코드(종목.G코드())
		lb.F테스트_에러없음(t, 에러)

		종목명 := strings.ReplaceAll(종목.G코드(), " ", "")
		종목명2 := strings.ReplaceAll(종목_비교값.G코드(), " ", "")

		lb.F테스트_참임(t, strings.Contains(종목명, 종목명2), 종목.G이름(), 종목_비교값.G이름())
	}
}
