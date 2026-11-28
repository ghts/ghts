package th

import (
	"strings"
	"testing"
)

func TestF변수값_자료형_문자열(t *testing.T) {
	문자열 := f변수값_자료형_문자열("테스트_문자열", 1)

	F테스트_참임(t, strings.Contains(문자열, "테스트_문자열"))
	F테스트_참임(t, strings.Contains(문자열, "string"))
	F테스트_참임(t, strings.Contains(문자열, "1"))
	F테스트_참임(t, strings.Contains(문자열, "int"))
}

func TestF호출경로_포함(t *testing.T) {
	t.Parallel()

	문자열 := "github.com/ghts/sample.go.go:65:f샘플()\n\nFAIL	github.com/ghts/	23.231s"

	F테스트_참임(t, f호출경로_포함(문자열))
}
