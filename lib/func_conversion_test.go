package lib

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestF2실수_소숫점_추가(t *testing.T) {
	t.Parallel()

	실수값, 에러 := F2실수_소숫점_추가("1234567", 2)
	F테스트_에러없음(t, 에러)
	F테스트_같음(t, 실수값, 12345.67)
}

func TestF바이트2참거짓(t *testing.T) {
	t.Parallel()
	F테스트_참임(t, F문자열_비교([]byte("1"), "1", true))
	F테스트_거짓임(t, F문자열_비교([]byte("2"), "1", true))
}

func TestF문자열2바이트_복사(t *testing.T) {
	t.Parallel()
	바이트_배열 := make([]byte, 10)
	F바이트_복사_문자열(바이트_배열, "12345")

	F테스트_같음(t, string(바이트_배열), "12345     ")
}

func TestF정수2바이트_복사(t *testing.T) {
	t.Parallel()
	바이트_배열 := make([]byte, 10)
	F바이트_복사_정수(바이트_배열, 12345)

	F테스트_같음(t, string(바이트_배열), "0000012345")
}

func TestF2포맷된_시각(t *testing.T) {
	t.Parallel()

	시각, 에러 := F2포맷된_시각("2006-01-02 15:04:05", "2001-02-03 04:05:06")
	F테스트_에러없음(t, 에러)
	F테스트_같음(t, 시각.Year(), 2001)
	F테스트_같음(t, 시각.Month(), time.February)
	F테스트_같음(t, 시각.Day(), 3)
	F테스트_같음(t, 시각.Hour(), 4)
	F테스트_같음(t, 시각.Minute(), 5)
	F테스트_같음(t, 시각.Second(), 6)

	시각, 에러 = F2포맷된_시각("2006-01-02 15:04:05", "에러 발생 유발 문자열")
	F테스트_에러발생(t, 에러)
	F테스트_참임(t, 시각.IsZero(), 시각)
}

func TestF문자열_모음2인터페이스_모음(t *testing.T) {
	t.Parallel()

	문자열_모음 := []string{"테스트1", "2", "3.0"}
	인터페이스_모음 := F2인터페이스_모음(문자열_모음)

	F테스트_같음(t, len(인터페이스_모음), len(문자열_모음))

	for i := 0; i < len(문자열_모음); i++ {
		F테스트_같음(t, 인터페이스_모음[i].(string), 문자열_모음[i])
	}
}

func TestF인터페이스_모음2문자열_모음(t *testing.T) {
	t.Parallel()

	인터페이스_모음 := []interface{}{"테스트", 1, time.Now()}
	문자열_모음 := F2문자열_모음(인터페이스_모음)

	F테스트_같음(t, len(인터페이스_모음), len(문자열_모음))

	for i := 0; i < len(문자열_모음); i++ {
		F테스트_같음(t, F2문자열(인터페이스_모음[i]), 문자열_모음[i])
	}
}

func TestF2문자열_EUC_KR(t *testing.T) {
	t.Parallel()

	F테스트_같음(t, "아름다운 우리말", F2문자열_EUC_KR("\xbe\xc6\xb8\xa7\xb4\xd9\xbf\xee \xbf\xec\xb8\xae\xb8\xbb"))
	F테스트_같음(t, "똠방각하", F2문자열_EUC_KR("\x8cc\xb9\xe6\xb0\xa2\xc7\xcf"))
	F테스트_같음(t, "펩시콜라", F2문자열_EUC_KR("\xc6\xe9\xbd\xc3\xc4\xdd\xb6\xf3"))
}

//func TestF인코딩_디코딩(t *testing.T) {
//	t.Parallel()
//
//	변환형식_모음 := []T변환{GOB} //,JSON}
//
//	r := F임의값_생성기()
//
//	값_모음 := []interface{}{
//		r.Int(), r.Int63(), r.Float64(), r.Intn(1) == 0,
//		F임의_문자열(5, 100), []string{"test1", "test2"}, F임의_시각(),
//		[]int{r.Int(), r.Int(), r.Int()}}
//
//	for _, 변환형식 := range 변환형식_모음 {
//		for _, 값1 := range 값_모음 {
//			바이트_모음, 에러 := F인코딩(변환형식, 값1)
//			F테스트_에러없음(t, 에러)
//
//			값2 := reflect.New(reflect.TypeOf(값1)).Elem().Interface()
//			F디코딩(변환형식, 바이트_모음, &값2)
//
//			F테스트_같음(t, 값1, 값2, 변환형식)
//		}
//
//		// nil 대응
//		바이트_모음, 에러 := F인코딩(변환형식, nil)
//		F테스트_에러발생(t, 에러)
//
//		// 자료형 정보가 존재하면 구조체도 가능함. 그러나, interface{}로는 안 됨.
//		s1 := F샘플_구조체_1()
//		바이트_모음, 에러 = F인코딩(변환형식, s1)
//		F테스트_에러없음(t, 에러)
//
//		s1_복사본 := s샘플_구조체_1{}
//		F디코딩(변환형식, 바이트_모음, &s1_복사본)
//
//		F테스트_같음(t, s1, s1_복사본, 변환형식)
//	}
//}

type s슬라이스를_포함한_구조체 struct {
	M문자열  string
	M슬라이스 []string
}

func TestF인코딩_디코딩_슬라이스를_포함한_구조체(t *testing.T) {
	t.Parallel()

	r := F임의값_생성기()

	원래값 := new(s슬라이스를_포함한_구조체)
	원래값.M문자열 = F임의_문자열(4, 6)
	원래값.M슬라이스 = make([]string, 5+r.Intn(10))

	for i := 0; i < len(원래값.M슬라이스); i++ {
		원래값.M슬라이스[i] = F임의_문자열(4, 6)
	}

	for _, 변환_형식 := range f테스트용_변환형식_모음() {
		바이트_모음, 에러 := F인코딩(변환_형식, 원래값)
		F테스트_에러없음(t, 에러)

		복원값 := new(s슬라이스를_포함한_구조체)
		에러 = F디코딩(변환_형식, 바이트_모음, 복원값)
		F테스트_에러없음(t, 에러)

		F테스트_같음(t, 원래값.M문자열, 복원값.M문자열)
		F테스트_같음(t, len(원래값.M슬라이스), len(복원값.M슬라이스))

		for i, 원래_문자열 := range 원래값.M슬라이스 {
			F테스트_같음(t, 복원값.M슬라이스[i], 원래_문자열)
		}
	}
}

type s슬라이스를_포함한_구조체2 struct {
	M문자열  string
	M슬라이스 []*S종목 // '[]I종목'으로 하면 안 인코딩/디코딩이 안 됨.
}

func TestF인코딩_디코딩_슬라이스를_포함한_구조체2(t *testing.T) {
	t.Parallel()

	r := F임의값_생성기()

	원래값 := new(s슬라이스를_포함한_구조체2)
	원래값.M문자열 = F임의_문자열(4, 6)
	원래값.M슬라이스 = make([]*S종목, 5+r.Intn(10))

	for i := 0; i < len(원래값.M슬라이스); i++ {
		원래값.M슬라이스[i] = F임의_샘플_종목_코스피_주식()
	}

	for _, 변환_형식 := range f테스트용_변환형식_모음() {
		바이트_모음, 에러 := F인코딩(변환_형식, 원래값)
		F테스트_에러없음(t, 에러)

		복원값 := new(s슬라이스를_포함한_구조체2)
		복원값.M슬라이스 = make([]*S종목, 0)

		에러 = F디코딩(변환_형식, 바이트_모음, 복원값)
		F테스트_에러없음(t, 에러)

		F테스트_같음(t, 원래값.M문자열, 복원값.M문자열)
		F테스트_같음(t, len(원래값.M슬라이스), len(복원값.M슬라이스))

		for i, 종목 := range 원래값.M슬라이스 {
			F테스트_같음(t, 복원값.M슬라이스[i].G코드(), 종목.G코드())
			F테스트_같음(t, 복원값.M슬라이스[i].G이름(), 종목.G이름())
		}
	}
}

func TestF바이트_변환값_해석(t *testing.T) {
	t.Parallel()

	변환_형식_모음 := []T변환{JSON, GOB}

	원본값_모음 := []interface{}{
		new(S콜백_기본형), New콜백_정수값_기본형(), new(S콜백_문자열), new(S콜백_TR데이터), new(S콜백_메시지_및_에러)}

	for _, 변환_형식 := range 변환_형식_모음 {
		for _, 원본값 := range 원본값_모음 {
			매개체, 에러 := New바이트_변환(변환_형식, 원본값)
			F테스트_에러없음(t, 에러)

			해석값, 에러 := 매개체.S해석기(F바이트_변환값_해석).G해석값()
			F테스트_에러없음(t, 에러)

			자료형_문자열 := strings.Replace(f자료형_문자열(해석값), "*", "", -1)
			F테스트_같음(t, 자료형_문자열, 매개체.G자료형_문자열())
		}
	}
}

func f자료형_문자열(값 interface{}) string {
	if 값 == nil {
		return "nil"
	}

	자료형 := reflect.TypeOf(값).String()
	시작_인덱스 := strings.Index(자료형, ".") + 1

	return 자료형[시작_인덱스:]
}

func TestF2한국_시간(t *testing.T) {
	지금 := time.Now()
	한국_시간 := F2한국_시간(지금)
	UTC_시간 := 지금.UTC()
	UTC_시간_9 := UTC_시간.Add(9 * P1시간)

	F테스트_같음(t, 한국_시간.Format(P간략한_시간_형식), UTC_시간_9.Format(P간략한_시간_형식))
}
