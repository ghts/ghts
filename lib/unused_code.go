package lib

import (
	"bytes"
	"math/big"
	"strings"
)

const (
	KRW = T통화(byte('K'))
	USD = T통화(byte('U'))
	EUR = T통화(byte('E'))
	CNY = T통화(byte('C'))
)

type T통화 byte

func (v *T통화) String() string {
	switch *v {
	case KRW:
		return "KRW"
	case USD:
		return "USD"
	case EUR:
		return "EUR"
	case CNY:
		return "CNY"
	default:
		return F2문자열("예상하지 못한 값 : '%v'", byte(*v))
	}
}

func (v *T통화) Parse(값 string) {
	switch 값 {
	case "KRW":
		*v = KRW
	case "USD":
		*v = USD
	case "EUR":
		*v = EUR
	case "CNY":
		*v = CNY
	default:
		panic(New에러("예상하지 못한 값 : '%v'", 값))
	}
}

// S통화는 통화를 표현하는 편의 구조체입니다.
type S통화 struct {
	단위   T통화
	금액   *big.Float
	변경불가 bool
}

func (s *S통화) G단위() T통화 { return s.단위 }
func (s *S통화) G정수64() int64 {
	값, _ := s.금액.Int64()
	return 값
}
func (s *S통화) G실수64() float64 {
	실수64, _ := s.금액.Float64()
	return 실수64
}
func (s *S통화) G정밀값() *big.Float {
	return new(big.Float).Copy(s.금액)
}

func (s *S통화) G문자열() string { return s.금액.String() }
func (s *S통화) G문자열_고정소숫점(소숫점_이하_자릿수 int) string {
	return s.금액.Text('f', 소숫점_이하_자릿수)
}

func (s *S통화) G비교(다른_통화 *S통화) T비교 {
	switch {
	case s.단위 != 다른_통화.G단위():
		return P비교_불가
	default:
		return T비교(s.금액.Cmp(다른_통화.G정밀값()))
	}
}

func (s *S통화) G복사본() *S통화 {
	복사본 := new(S통화)
	복사본.단위 = s.G단위()
	복사본.금액 = s.G정밀값()
	복사본.변경불가 = false

	return 복사본
}

func (s *S통화) G변경불가() bool {
	return s.변경불가
}

func (s *S통화) S동결() { s.변경불가 = true }

func (s *S통화) S더하기(값 float64) *S통화 {
	if s.변경불가 {
		panic("변경불가능한 값입니다.")
		//return s
	}

	s.금액 = new(big.Float).Add(s.금액, big.NewFloat(값))

	return s
}

func (s *S통화) S빼기(값 float64) *S통화 {
	if s.변경불가 {
		panic(New에러with출력("변경불가능한 값입니다."))
		//return s
	}

	s.금액 = new(big.Float).Sub(s.금액, big.NewFloat(값))

	return s
}

func (s *S통화) S곱하기(값 float64) *S통화 {
	if s.변경불가 {
		panic(New에러("변경불가능한 값입니다."))
		//return s
	}

	s.금액 = new(big.Float).Mul(s.금액, big.NewFloat(값))

	return s
}

func (s *S통화) S나누기(값 float64) (*S통화, error) {
	switch {
	case s.변경불가:
		panic("변경불가능한 값입니다.")
		//return s, 에러
	case 값 == 0.0:
		return nil, New에러with출력("분모가 0인 나눗셈 불가.")
	default:
		s.금액 = new(big.Float).Quo(s.금액, big.NewFloat(값))
		return s, nil
	}
}

func (s *S통화) S금액(금액 float64) *S통화 {
	if s.변경불가 {
		panic("변경불가능한 값입니다.")
		//return s
	}

	s.금액 = big.NewFloat(금액)

	return s
}

func (s *S통화) String() string {
	return s.단위.String() + " " + s.금액.String()
}

func (s *S통화) MarshalBinary() ([]byte, error) {
	var 변경_불가 byte
	if !s.변경불가 {
		변경_불가 = byte(0)
	} else {
		변경_불가 = byte(1)
	}

	금액, 에러 := s.금액.MarshalText()
	if 에러 != nil {
		return nil, 에러
	}

	버퍼 := new(bytes.Buffer)
	버퍼.WriteByte(변경_불가)
	버퍼.WriteByte(byte(s.단위))
	버퍼.Write(금액)

	return 버퍼.Bytes(), nil
}

func (s *S통화) UnmarshalBinary(값 []byte) (에러 error) {
	defer func() {
		if 에러 != nil {
			s.단위 = T통화(byte(' '))
			s.금액 = big.NewFloat(0.0)
			s.변경불가 = true
		}
	}()

	const 헤더_길이 = 2

	switch {
	case len(값) == 0:
		return New에러with출력("비어있는 M값")
	case len(값) <= 헤더_길이:
		return New에러with출력("너무 짧은 M값. %v", len(값))
	}

	if 변경_불가 := int(값[0]); 변경_불가 == 0 {
		s.변경불가 = false
	} else {
		s.변경불가 = true
	}

	s.단위 = T통화(값[1])

	F확인1(s.금액.UnmarshalText(값[2:]))

	return
}

func (s *S통화) MarshalText() ([]byte, error) {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString("{")
	버퍼.WriteString(s.단위.String())
	버퍼.WriteString(",")
	금액 := F확인2(s.금액.MarshalText())
	버퍼.Write(금액)
	버퍼.WriteString(",")
	if s.변경불가 {
		버퍼.WriteString("T")
	} else {
		버퍼.WriteString("F")
	}
	버퍼.WriteString("}")

	return 버퍼.Bytes(), nil
}

func (s *S통화) UnmarshalText(값 []byte) (에러 error) {
	defer func() {
		if 에러 != nil {
			s.단위 = T통화(byte(' '))
			s.금액 = big.NewFloat(0.0)
			s.변경불가 = true
		}
	}()

	문자열 := string(값)
	문자열_모음 := strings.Split(문자열, ",")

	if len(문자열_모음) != 3 {
		return New에러with출력("예상하지 못한 경우. %v", len(문자열_모음))
	}

	s.단위.Parse(strings.TrimSpace(문자열_모음[0][1:]))

	if s.금액 == nil {
		s.금액 = new(big.Float)
	}
	금액_문자열 := strings.TrimSpace(문자열_모음[1])
	F확인1(s.금액.UnmarshalText([]byte(금액_문자열)))

	switch strings.TrimSpace(문자열_모음[2])[:1] {
	case "T", "t":
		s.변경불가 = true
	case "F", "f":
		s.변경불가 = false
	default:
		에러 = New에러with출력("예상하지 못한 문자열. '%v'", strings.TrimSpace(문자열_모음[2]))
	}

	return nil
}

func New원화(금액 float64) *S통화 { return New통화(KRW, 금액) }
func New달러(금액 float64) *S통화 { return New통화(USD, 금액) }
func New유로(금액 float64) *S통화 { return New통화(EUR, 금액) }
func New위안(금액 float64) *S통화 { return New통화(CNY, 금액) }
func New통화(단위 T통화, 금액 float64) *S통화 {
	F확인1(F통화단위_검사(단위))

	s := new(S통화)
	s.단위 = 단위
	s.금액 = big.NewFloat(금액)
	s.변경불가 = false

	return s
}

func F통화단위_검사(통화단위 T통화) error {
	switch 통화단위 {
	case KRW, USD, EUR, CNY:
		return nil
	default:
		return New에러with출력("잘못된 통화단위. '%v'", 통화단위)
	}
}
