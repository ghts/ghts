package xing

import (
	lb "github.com/ghts/ghts/lib"
	"testing"
	"time"
)

func TestF접속됨(t *testing.T) {
	t.Parallel()

	접속됨, 에러 := F접속됨()
	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_참임(t, 접속됨)
}

func TestF계좌번호_모음(t *testing.T) {
	t.Parallel()

	계좌번호_모음, 에러 := F계좌번호_모음()

	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_참임(t, len(계좌번호_모음) > 0)

	for _, 계좌번호 := range 계좌번호_모음 {
		lb.F테스트_참임(t, len(계좌번호) > 0)
	}
}

func TestF영업일_기준_전일_당일(t *testing.T) {
	t.Parallel()

	전일 := F전일()
	당일 := F당일()

	lb.F테스트_다름(t, 전일, time.Time{})
	lb.F테스트_다름(t, 당일, time.Time{})
	lb.F테스트_참임(t, 전일.After(time.Now().AddDate(-1, 0, 0)))
	lb.F테스트_참임(t, 당일.After(전일))
	lb.F테스트_참임(t, 당일.Before(time.Now().AddDate(0, 0, 1)))
	lb.F테스트_같음(t, 전일.Hour(), 0)
	lb.F테스트_같음(t, 전일.Minute(), 0)
	lb.F테스트_같음(t, 전일.Second(), 0)
	lb.F테스트_같음(t, 전일.Nanosecond(), 0)
	lb.F테스트_같음(t, 당일.Hour(), 0)
	lb.F테스트_같음(t, 당일.Minute(), 0)
	lb.F테스트_같음(t, 당일.Second(), 0)
	lb.F테스트_같음(t, 당일.Nanosecond(), 0)
}
