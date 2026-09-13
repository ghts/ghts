package lib

import (
	"database/sql"
	"math/big"
	"testing"

	_ "modernc.org/sqlite"
)

func TestDB질의(t *testing.T) {
	t.Parallel()

	db := F확인2(sql.Open("sqlite", ":memory:"))
	// in-memory DB 는 연결당 독립이므로 단일 연결로 고정.
	db.SetMaxOpenConns(1)
	defer func() {
		F확인1(db.Close())
	}()

	F확인2(db.Exec("CREATE TABLE db질의_테스트 (a INTEGER, b TEXT, c INTEGER, d INTEGER)"))
	F확인2(db.Exec("INSERT INTO db질의_테스트 VALUES (42, '안녕', 1, 1700000000)"))
	F확인2(db.Exec("INSERT INTO db질의_테스트 VALUES (-7, '음수', 0, 1700000001)"))

	t.Run("정수", func(t *testing.T) {
		F테스트_같음(t, F확인2(DB질의[int64](db, "SELECT a FROM db질의_테스트")), int64(42))
		F테스트_같음(t, F확인2(DB질의[int](db, "SELECT a FROM db질의_테스트")), int(42))
		F테스트_같음(t, F확인2(DB질의[int8](db, "SELECT a FROM db질의_테스트")), int8(42))
		F테스트_같음(t, F확인2(DB질의[uint16](db, "SELECT a FROM db질의_테스트")), uint16(42))
		F테스트_같음(t, F확인2(DB질의[int64](db, "SELECT a FROM db질의_테스트 WHERE a = ?", -7)), int64(-7))
	})

	t.Run("실수", func(t *testing.T) {
		F테스트_같음(t, F확인2(DB질의[float32](db, "SELECT a FROM db질의_테스트")), float32(42))
		F테스트_같음(t, F확인2(DB질의[float64](db, "SELECT a FROM db질의_테스트")), float64(42))
	})

	t.Run("big_값", func(t *testing.T) {
		정수값 := F확인2(DB질의[big.Int](db, "SELECT a FROM db질의_테스트"))
		F테스트_같음(t, 정수값.String(), "42")
		분수값 := F확인2(DB질의[big.Rat](db, "SELECT a FROM db질의_테스트"))
		F테스트_같음(t, 분수값.String(), "42/1")
	})

	t.Run("big_포인터", func(t *testing.T) {
		// 포인터 반환: nil 이 아니어야 하고, 매 호출마다 새 인스턴스여야 함.
		결과_정수1, 에러_정수1 := DB질의[*big.Int](db, "SELECT a FROM db질의_테스트")
		F확인1(에러_정수1)
		F테스트_참임(t, 결과_정수1 != nil)
		F테스트_같음(t, 결과_정수1.String(), "42")

		결과_정수2, 에러_정수2 := DB질의[*big.Int](db, "SELECT a FROM db질의_테스트")
		F확인1(에러_정수2)
		F테스트_같음(t, 결과_정수2.String(), "42")
		F테스트_거짓임(t, 결과_정수2 == 결과_정수1)

		결과_분수1, 에러_분수1 := DB질의[*big.Rat](db, "SELECT a FROM db질의_테스트")
		F확인1(에러_분수1)
		F테스트_참임(t, 결과_분수1 != nil)
		F테스트_같음(t, 결과_분수1.String(), "42/1")

		결과_분수2, 에러_분수2 := DB질의[*big.Rat](db, "SELECT a FROM db질의_테스트")
		F확인1(에러_분수2)
		F테스트_거짓임(t, 결과_분수2 == 결과_분수1)
	})

	t.Run("문자열", func(t *testing.T) {
		F테스트_같음(t, F확인2(DB질의[string](db, "SELECT b FROM db질의_테스트")), "안녕")
	})

	t.Run("부울", func(t *testing.T) {
		F테스트_참임(t, F확인2(DB질의[bool](db, "SELECT c FROM db질의_테스트")))
		F테스트_거짓임(t, F확인2(DB질의[bool](db, "SELECT c FROM db질의_테스트 WHERE a = ?", -7)))
	})

	t.Run("결과_없음", func(t *testing.T) {
		값, 에러 := DB질의[int64](db, "SELECT a FROM db질의_테스트 WHERE a = ?", 999)
		F테스트_참임(t, 에러 != nil)
		F테스트_같음(t, 값, int64(0))

		값_문자열, 에러_문자열 := DB질의[string](db, "SELECT b FROM db질의_테스트 WHERE a = ?", 999)
		F테스트_참임(t, 에러_문자열 != nil)
		F테스트_같음(t, 값_문자열, "")
	})
}
