/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package shared

import (
	dec "github.com/wayn3h0/go-decimal"

	"time"
)

// 안전한 bool
type I안전한_bool interface {
	G값() bool
	S값(값 bool) error
}

func New안전한_bool(값 bool) I안전한_bool {
	return &s안전한_bool{값: 값}
}

// 안전한 string
type I안전한_string interface {
	G값() string
	S값(값 string)
}

func New안전한_string(값 string) I안전한_string {
	return &s안전한_string{값: 값}
}

// 기본 메시지
type I메시지 interface {
	G구분() string
	G내용(인덱스 int) string
	G내용_전체() []string
	G길이() int
	String() string
}

func New메시지(구분 string, 내용 ...interface{}) I메시지 {
	내용_모음 := make([]string, len(내용))

	for i := 0; i < len(내용); i++ {
		내용_모음[i] = F포맷된_문자열("%v", 내용[i])
	}

	return s기본_메시지{구분: 구분, 내용: 내용_모음}
}

// 질의
type I질의 interface {
	I메시지 // 질의 내용
	G회신(채널 chan I질의) I회신
	G회신_채널() chan I회신
	G검사(타이틀 string, 질의_길이 int) error
}

// New질의(...).G회신() 혹은 F질의(...) 둘 다 똑같다.
func New질의(메시지_구분 string, 내용 ...interface{}) I질의 {
	switch 메시지_구분 {
	case P메시지_GET:
	case P메시지_SET:
	case P메시지_DEL:
	case P메시지_종료:
	default:
		에러 := F에러_생성("잘못된 질의 메시지 구분 %v", 메시지_구분)
		F에러_출력(에러.Error())
		panic(에러)
	}

	회신_채널 := make(chan I회신, 1)
	메시지 := New메시지(메시지_구분, 내용...)

	return s질의_메시지{회신_채널: 회신_채널, s기본_메시지: 메시지.(s기본_메시지)}
}

func F질의(질의_채널 chan I질의, 메시지_구분 string, 내용 ...interface{}) I회신 {
	질의 := New질의(메시지_구분, 내용...)
	질의_채널 <- 질의

	return <-질의.G회신_채널()
}

// 회신
type I회신 interface {
	I메시지
	G에러() error
}

func New회신(에러 error, 내용 ...interface{}) I회신 {
	메시지_구분 := ""

	if 에러 == nil || F포맷된_문자열("%v", 에러) == "<nil>" {
		메시지_구분 = P메시지_OK
	} else {
		메시지_구분 = P메시지_에러
	}

	if len(내용) == 1 {
		문자열_모음, ok := 내용[0].([]string)

		if ok {
			내용_원본 := 내용
			내용 = make([]interface{}, len(내용_원본))

			for 인덱스, 문자열 := range 문자열_모음 {
				내용[인덱스] = 문자열
			}
		}
	}

	return s회신_메시지{에러: 에러, s기본_메시지: New메시지(메시지_구분, 내용...).(s기본_메시지)}
}

// 종목
type I종목 interface {
	G코드() string
	G이름() string
	String() string
}

func New종목(코드 string, 이름 string) I종목 {
	종목 := new(s종목)
	종목.코드 = 코드
	종목.이름 = 이름

	return 종목
}

// 통화
type I통화 interface {
	G단위() T통화단위
	G실수값() float64
	G정밀값() *dec.Decimal
	G실수_문자열(소숫점_이하_자릿수 int) string
	G비교(다른_통화 I통화) T비교결과
	G부호() T부호
	G복사본() I통화
	G변경불가() bool
	S동결()
	S더하기(다른_통화 I통화) I통화
	S빼기(다른_통화 I통화) I통화
	S곱하기(다른_통화 I통화) I통화
	S나누기(다른_통화 I통화) I통화
	S금액(금액 string) I통화
	String() string
}

func New원화(금액 string) I통화 { return New통화(KRW, 금액) }
func New달러(금액 string) I통화 { return New통화(USD, 금액) }
func New유로(금액 string) I통화 { return New통화(EUR, 금액) }
func New위안(금액 string) I통화 { return New통화(CNY, 금액) }
func New통화(단위 T통화단위, 금액 string) I통화 {
	정밀값, 에러 := dec.Parse(금액)

	if 에러 != nil {
		F에러_출력(에러.Error())
		return nil
	}

	s := new(s통화)
	s.단위 = 단위
	s.금액 = 정밀값
	s.변경불가 = false

	return s
}

// 가격정보
type I가격정보 interface {
	G종목() I종목
	G가격() I통화
	G시점() time.Time
}

func New가격정보(종목 I종목, 가격 I통화) I가격정보 {
	s := s가격정보{종목: 종목, 가격: 가격.G복사본(), 시점: time.Now()}
	return &s
}
