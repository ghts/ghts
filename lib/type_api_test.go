package lib

import (
	"testing"
)

func TestI질의값(t *testing.T) {
	t.Parallel()

	var 질의값 I질의값

	질의값 = new(S질의값_기본형)
	질의값.TR구분() // unused 에러 발생 방지용

	질의값 = New질의값_문자열(TR조회, "", "")
	질의값.TR구분() // unused 에러 발생 방지용

	질의값 = New질의값_단일_종목_단순형()
	질의값.TR구분() // unused 에러 발생 방지용

	질의값 = New질의값_단일종목_연속키()
	질의값.TR구분() // unused 에러 발생 방지용

	질의값 = New질의값_복수_종목(TR조회, "", []string{""})
	질의값.TR구분() // unused 에러 발생 방지용

	질의값 = New질의값_정상_주문()
	질의값.TR구분() // unused 에러 발생 방지용

	질의값 = New질의값_정정_주문()
	질의값.TR구분() // unused 에러 발생 방지용

	질의값 = New질의값_취소_주문()
	질의값.TR구분() // unused 에러 발생 방지용
}

func TestS질의값_문자열(t *testing.T) {
	t.Parallel()

	문자열 := F임의_문자열(5, 10)

	원래값 := New질의값_문자열(TR조회, F임의_문자열(2, 6), 문자열)

	바이트_변환값, 에러 := New바이트_변환(F임의_변환_형식(), 원래값)
	F테스트_에러없음(t, 에러)

	복원값 := new(S질의값_문자열)
	F테스트_에러없음(t, 바이트_변환값.G값(복원값))

	F테스트_같음(t, 복원값.M구분, 원래값.M구분)
	F테스트_같음(t, 복원값.M코드, 원래값.M코드)
	F테스트_같음(t, 복원값.M문자열, 문자열)
}

func TestS질의값_단일종목(t *testing.T) {
	t.Parallel()

	원래값 := New질의값_단일_종목_단순형()
	원래값.M구분 = TR조회
	원래값.M코드 = F임의_문자열(2, 6)
	원래값.M종목코드 = F임의_샘플_종목().G코드()

	바이트_변환값, 에러 := New바이트_변환(F임의_변환_형식(), 원래값)
	F테스트_에러없음(t, 에러)

	복원값 := new(S질의값_단일_종목)
	F테스트_에러없음(t, 바이트_변환값.G값(복원값))

	F테스트_같음(t, 복원값.M구분, 원래값.M구분)
	F테스트_같음(t, 복원값.M코드, 원래값.M코드)
	F테스트_같음(t, 복원값.M종목코드, 원래값.M종목코드)
}

func TestI콜백(t *testing.T) {
	값_모음 := []interface{}{
		New콜백_기본형(T콜백(0)),
		New콜백_정수값(T콜백(0), 0),
		New콜백_문자열(T콜백(0), ""),
		New콜백_TR데이터(0, nil, "", false, ""),
		New콜백_메시지("", ""),
		New콜백_에러("", "")}

	for _, 값 := range 값_모음 {
		f콜백_테스트_도우미(t, 값)
	}
}

func f콜백_테스트_도우미(t *testing.T, 값 interface{}) {
	switch 값.(type) {
	case I콜백:
		return
	}

	t.FailNow()
}
