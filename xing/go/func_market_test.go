package xing

import (
	"github.com/ghts/ghts/lib"
	"testing"
)

func TestF종목코드_존재함(t *testing.T) {
	t.Parallel()

	lib.F테스트_참임(t, F종목코드_존재함("069500"))
}

func TestF질의값_종목코드_검사(t *testing.T) {
	t.Parallel()

	질의값1 := lib.New질의값_단일_종목_단순형()
	질의값1.M종목코드 = "069500"

	질의값2 := lib.New질의값_복수_종목(lib.TR조회, "", []string{"069500"})

	lib.F테스트_에러없음(t, F질의값_종목코드_검사(질의값1))
	lib.F테스트_에러없음(t, F질의값_종목코드_검사(질의값2))
}

func TestETF_ETN_종목_여부(t *testing.T) {
	t.Parallel()

	lib.F테스트_거짓임(t, ETF_ETN_종목_여부("000020"))
	lib.F테스트_참임(t, ETF_ETN_종목_여부("069500"))
}
