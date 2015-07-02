package shared

import (
	"strings"
	"testing"
)

func TestT주소String(테스트 *testing.T) {
	F테스트_참임(테스트, strings.HasPrefix(P주소_주소정보.String(), "tcp://127.0.0.1:"))
}

func TestT통화단위String(테스트 *testing.T) {
	F테스트_같음(테스트, KRW.String(), "KRW")
}
