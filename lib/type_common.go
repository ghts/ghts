package lib

import (
	"bytes"
	"encoding/binary"
	"strings"
	"time"
)

var 비어있는_일자값 = time.Time{}

func F비어있는_일자값() time.Time {
	return 비어있는_일자값
}

type S비어있음 struct{}

type T정수 interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type T실수 interface {
	float32 | float64
}

type T숫자 interface {
	T정수 | T실수
}

type S종목 struct {
	코드    string
	이름    string
	시장_구분 T시장구분
	전일_종가 int64
	상한가   int64
	하한가   int64
	기준가   int64
}

func (s *S종목) G코드() string     { return s.코드 }
func (s *S종목) G이름() string     { return s.이름 }
func (s *S종목) G식별_문자열() string { return F2문자열("%v[%v]", s.G이름(), s.G코드()) }
func (s *S종목) G시장구분() T시장구분    { return s.시장_구분 }
func (s *S종목) G전일_종가() int64   { return s.전일_종가 }
func (s *S종목) G상한가() int64     { return s.상한가 }
func (s *S종목) G하한가() int64     { return s.하한가 }
func (s *S종목) G기준가() int64     { return s.기준가 }

func (s *S종목) String() string {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString(s.이름)
	버퍼.WriteString("(")
	버퍼.WriteString(s.코드)
	버퍼.WriteString(" : ")
	버퍼.WriteString(s.시장_구분.String())
	버퍼.WriteString(")")
	return 버퍼.String()
}

func (s *S종목) G복제본() *S종목 {
	복제본 := new(S종목)
	복제본.코드 = s.코드
	복제본.이름 = s.이름
	복제본.시장_구분 = s.시장_구분
	복제본.전일_종가 = s.전일_종가
	복제본.상한가 = s.상한가
	복제본.하한가 = s.하한가
	복제본.기준가 = s.기준가

	return 복제본
}

func (s *S종목) MarshalBinary() ([]byte, error) {
	// 인텔/AMD CPU는 리틀 엔디언. ARM CPU는 OS에 따라서 엔디언 바뀜. 맥OS에서는 리틀 엔디언.
	// 현존하는 OS 대부분 리틀 엔디언.
	// 빅 엔디언은 네트워크 전송 및 일부 대형 컴퓨터에서만 사용된다고 함.

	속성 := make([]byte, 1)
	속성[0] = uint8(s.시장_구분)

	코드_길이 := make([]byte, 2)
	binary.LittleEndian.PutUint16(코드_길이, uint16(len(s.코드)))

	이름_길이 := make([]byte, 2)
	binary.LittleEndian.PutUint16(이름_길이, uint16(len(s.이름)))

	가격_정보_모음 := []int64{s.전일_종가, s.상한가, s.하한가, s.기준가}
	가격_정보_바이트 := make([]byte, 8*len(가격_정보_모음))

	for i, 가격 := range 가격_정보_모음 {
		binary.LittleEndian.PutUint64(가격_정보_바이트[i*8:(i+1)*8], uint64(가격))
	}

	값_모음 := [][]byte{속성, 코드_길이, 이름_길이, []byte(s.코드), []byte(s.이름), 가격_정보_바이트}
	버퍼 := new(bytes.Buffer)

	for _, 값 := range 값_모음 {
		if _, 에러 := 버퍼.Write(값); 에러 != nil {
			return nil, 에러
		}
	}

	return 버퍼.Bytes(), nil
}

func (s *S종목) UnmarshalBinary(값 []byte) (에러 error) {
	defer func() {
		if 에러 != nil {
			s.코드 = ""
			s.이름 = ""
			s.시장_구분 = P시장구분_전체
			s.전일_종가 = 0
			s.상한가 = 0
			s.하한가 = 0
			s.기준가 = 0
		}
	}()

	const 헤더_길이 = 5
	const 가격_정보_길이 = 32 // 전일_종가, 상한가, 하한가, 기준가 (4개 항목 x 8바이트)

	switch {
	case len(값) == 0:
		return New에러with출력("비어있는 M값")
	case len(값) < 헤더_길이:
		return New에러with출력("너무 짧은 M값. %v", len(값))
	}

	속성 := 값[:1]
	s.시장_구분 = T시장구분(속성[0])

	코드_길이 := int(binary.LittleEndian.Uint16(값[1:3])) // 인텔 및 AMD 계열 CPU는 리틀 엔디언
	이름_길이 := int(binary.LittleEndian.Uint16(값[3:5]))

	총_길이 := 헤더_길이 + 코드_길이 + 이름_길이 + 가격_정보_길이
	if len(값) != 총_길이 {
		return New에러with출력("무효한 M값. %v %v %v %v %v", len(값),
			헤더_길이, 코드_길이, 이름_길이, 가격_정보_길이)
	}

	시작점 := 헤더_길이
	s.코드 = string(값[시작점:(시작점 + 코드_길이)])

	시작점 = 시작점 + 코드_길이
	s.이름 = string(값[시작점:(시작점 + 이름_길이)])

	시작점 = 시작점 + 이름_길이
	s.전일_종가 = int64(binary.LittleEndian.Uint64(값[시작점:(시작점 + 8)]))
	s.상한가 = int64(binary.LittleEndian.Uint64(값[시작점+8 : (시작점 + 16)]))
	s.하한가 = int64(binary.LittleEndian.Uint64(값[시작점+16 : (시작점 + 24)]))
	s.기준가 = int64(binary.LittleEndian.Uint64(값[시작점+24 : 시작점+32]))

	return nil
}

func (s *S종목) MarshalText() ([]byte, error) {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString(`{"종목_코드": "`)
	버퍼.WriteString(s.코드)
	버퍼.WriteString(`", "종목_이름": "`)
	버퍼.WriteString(s.이름)
	버퍼.WriteString(`", "시장_구분": "`)
	버퍼.WriteString(s.시장_구분.String())
	버퍼.WriteString(`", "전일_종가": `)
	버퍼.WriteString(F2문자열(s.전일_종가))
	버퍼.WriteString(`, "상한가": `)
	버퍼.WriteString(F2문자열(s.상한가))
	버퍼.WriteString(`, "하한가": `)
	버퍼.WriteString(F2문자열(s.하한가))
	버퍼.WriteString(`, "기준가": `)
	버퍼.WriteString(F2문자열(s.기준가))
	버퍼.WriteString(`}`)

	return 버퍼.Bytes(), nil
}

func (s *S종목) UnmarshalText(값 []byte) error {
	문자열 := string(값)

	검색_결과 := F정규식_검색(문자열, []string{`{"종목_코드": ".+", "종목_이름": "`})
	코드, 에러 := f구간_추출(검색_결과, `{"종목_코드": "`, `", "종목_이름": "`)
	if 에러 != nil {
		return 에러
	}
	s.코드 = 코드

	검색_결과 = F정규식_검색(문자열, []string{`"종목_이름": ".+", "시장_구분": "`})
	이름, 에러 := f구간_추출(검색_결과, `"종목_이름": "`, `", "시장_구분": "`)
	if 에러 != nil {
		return 에러
	}
	s.이름 = 이름

	// 시장_구분은 더 이상 마지막 필드가 아니므로 다음 필드명을 종료 기준으로 사용.
	검색_결과 = F정규식_검색(문자열, []string{`"시장_구분": ".+", "전일_종가": `})
	시장_구분, 에러 := f구간_추출(검색_결과, `"시장_구분": "`, `", "전일_종가": `)
	if 에러 != nil {
		return 에러
	}
	if 에러 = s.시장_구분.Parse(시장_구분); 에러 != nil {
		return 에러
	}

	전일_종가, 에러 := f정수64_검색(문자열, "전일_종가")
	if 에러 != nil {
		return 에러
	}
	s.전일_종가 = 전일_종가

	상한가, 에러 := f정수64_검색(문자열, "상한가")
	if 에러 != nil {
		return 에러
	}
	s.상한가 = 상한가

	하한가, 에러 := f정수64_검색(문자열, "하한가")
	if 에러 != nil {
		return 에러
	}
	s.하한가 = 하한가

	기준가, 에러 := f정수64_검색(문자열, "기준가")
	if 에러 != nil {
		return 에러
	}
	s.기준가 = 기준가

	return nil
}

// f구간_추출은 검색_결과에서 시작_문자열과 종료_문자열 사이의 구간을 추출합니다.
// 구분자를 발견하지 못하면 인덱스 범위를 벗어나는 패닉 대신 에러를 반환합니다.
func f구간_추출(검색_결과, 시작_문자열, 종료_문자열 string) (string, error) {
	if !strings.HasPrefix(검색_결과, 시작_문자열) || !strings.HasSuffix(검색_결과, 종료_문자열) {
		return "", New에러with출력("무효한 M값. %v", 검색_결과)
	}

	값 := 검색_결과[len(시작_문자열):]
	값 = 값[:len(값)-len(종료_문자열)]

	return 값, nil
}

// f정수64_검색은 문자열에서 "필드명": 숫자 형태의 숫자 필드를 검색하여 int64로 변환합니다.
func f정수64_검색(문자열, 필드명 string) (int64, error) {
	접두사 := `"` + 필드명 + `": `
	검색_결과 := F정규식_검색(문자열, []string{`"` + 필드명 + `": -?[0-9]+`})

	if !strings.HasPrefix(검색_결과, 접두사) {
		return 0, New에러with출력("%v 필드 검색 실패. %v", 필드명, 문자열)
	}

	정수64값, 에러 := F2정수64(검색_결과[len(접두사):])
	if 에러 != nil {
		return 0, New에러with출력("%v 필드 변환 실패. %v", 필드명, 문자열)
	}

	return 정수64값, nil
}

// New종목은 S종목을 생성합니다.
func New종목(코드 string, 이름 string, 시장_구분 T시장구분) *S종목 {
	switch 시장_구분 {
	case P시장구분_코스피, P시장구분_코스닥, P시장구분_ETF, P시장구분_코넥스:
		if len(코드) != 6 {
			코드 = F정규식_검색(코드, []string{"[0-9]+K?"})
		}

		if len(코드) != 6 {
			panic(New에러with출력("잘못된 코드 '%v' '%v' '%v'", 코드, 이름, 시장_구분))
		}
	default:
		panic(F2문자열("예상하지 못한 경우 : '%v' '%v'", int(시장_구분), 시장_구분.String()))
	}

	s := new(S종목)
	s.코드 = 코드
	s.이름 = 이름
	s.시장_구분 = 시장_구분

	return s
}

func New종목with가격정보(코드 string, 이름 string, 시장_구분 T시장구분, 전일_종가, 상한가, 하한가, 기준가 int64) *S종목 {
	switch 시장_구분 {
	case P시장구분_코스피, P시장구분_코스닥, P시장구분_ETF, P시장구분_코넥스:
		if len(코드) != 6 {
			코드 = F정규식_검색(코드, []string{"[0-9]+K?"})
		}

		if len(코드) != 6 {
			panic(New에러with출력("잘못된 코드 '%v' '%v' '%v'", 코드, 이름, 시장_구분))
		}
	default:
		// PASS. 코드 검사 통과를 위해서 default문 추가함.
	}

	s := new(S종목)
	s.코드 = 코드
	s.이름 = 이름
	s.시장_구분 = 시장_구분
	s.전일_종가 = 전일_종가
	s.상한가 = 상한가
	s.하한가 = 하한가
	s.기준가 = 기준가

	return s
}

// 중복 없고 무작위 순서의 문자열 모음.
type I문자열_집합 interface {
	G슬라이스() []string
	G포함(값 string) bool
	G길이() int
	S추가(값 string)
	S삭제(값 string)
	String() string
}

func New문자열_집합() I문자열_집합 {
	s := new(s문자열_집합)
	s.맵 = make(map[string]S비어있음)
	return s
}

// 중복 없고 무작위 순서의 문자열 모음.
type s문자열_집합 struct {
	맵 map[string]S비어있음
}

func (s *s문자열_집합) G슬라이스() []string {
	길이 := len(s.맵)
	값 := make([]string, 길이)

	i := 0
	for 문자열 := range s.맵 {
		값[i] = 문자열
		i++
	}

	return 값
}

func (s *s문자열_집합) G길이() int { return len(s.맵) }

func (s *s문자열_집합) G포함(값 string) bool {
	for 문자열 := range s.맵 {
		if 값 == 문자열 {
			return true
		}
	}

	return false
}

func (s *s문자열_집합) S추가(값 string) { s.맵[값] = S비어있음{} }

func (s *s문자열_집합) S삭제(값 string) { delete(s.맵, 값) }

func (s *s문자열_집합) String() string {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString("[")

	마지막_인덱스 := len(s.맵) - 1
	i := 0
	for 문자열 := range s.맵 {
		버퍼.WriteString(문자열)

		if i != 마지막_인덱스 {
			버퍼.WriteString(", ")
		}
	}

	버퍼.WriteString("]")

	return 버퍼.String()
}
