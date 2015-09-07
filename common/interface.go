/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package common

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

// 안전한 일련 번호
type I안전한_일련_번호 interface {
	G값() int
}

func New안전한_일련_번호() I안전한_일련_번호 {
	return &s안전한_일련_번호{}
}

// 기본 메시지
type I메시지 interface {
	G구분() string
	G내용(인덱스 int) string
	G내용_전체() []string
	G내용_전체_가변형() []interface{} // zmq소켓 사용 시 편리함.
	G길이() int
	String() string
}

func New메시지(구분 string, 내용 ...interface{}) I메시지 {
	내용_모음 := make([]string, len(내용))

	for i := 0; i < len(내용); i++ {
		내용_모음[i] = F2문자열(내용[i])
	}

	return s기본_메시지{구분: 구분, 내용: 내용_모음}
}

// 질의
type I질의 interface {
	I메시지 // 질의 내용
	G검사(메시지_구분 string, 질의_길이 int) error
	G회신() I회신
	S질의(질의_채널 chan I질의) I질의
	S회신(에러 error, 내용 ...interface{}) error
}

func New질의(메시지_구분 string, 내용 ...interface{}) I질의 {
	switch 메시지_구분 {
	case P메시지_GET:
	case P메시지_SET:
	case P메시지_DEL:
	case P메시지_초기화:
	case P메시지_종료:
	default:
		에러 := F에러_생성("잘못된 질의 메시지 구분 %v", 메시지_구분)
		F에러_출력(에러)
		//panic(에러)
		return nil
	}

	회신_채널 := make(chan I회신, 100)
	메시지 := New메시지(메시지_구분, 내용...)

	return s질의_메시지{회신_채널: 회신_채널, s기본_메시지: 메시지.(s기본_메시지)}
}

func New질의_zmq메시지(zmq메시지 []string) I질의 {
	if zmq메시지 == nil || len(zmq메시지) == 0 {
		return nil
	}

	메시지_구분 := zmq메시지[0]

	if len(zmq메시지) == 1 {
		return New질의(메시지_구분)
	}

	질의 := New질의(메시지_구분, F2인터페이스_모음(zmq메시지[1:])...)

	return 질의
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
			내용 = make([]interface{}, len(문자열_모음))

			for 인덱스, 문자열 := range 문자열_모음 {
				내용[인덱스] = 문자열
			}
		}
	}

	return s회신_메시지{에러: 에러, s기본_메시지: New메시지(메시지_구분, 내용...).(s기본_메시지)}
}


// 기본 가변형 메시지
// I메시지는 immutable 자료형인 string을 사용해서 안전하지만,
// I메시지_가변형은 mutable 자료형을 주고 받게 되며,
// 참조형 데이터를 주고 받게 되면 위험성이 더 높아지니 주의 요망.
type I메시지_가변형 interface {
	G구분() string
	G내용(인덱스 int) interface{}
	G내용_전체() []interface{}
	G길이() int
	String() string
}

func New메시지_가변형(구분 string, 내용 ...interface{}) I메시지_가변형 {
	if 내용 == nil || len(내용) == 0 {
		내용 = make([]interface{}, 0)
	}
	
	return s기본_메시지_가변형{구분: 구분, 내용: 내용}
}

// 가변형 질의
type I질의_가변형 interface {
	I메시지_가변형 // 질의 내용
	G검사(메시지_구분 string, 질의_길이 int) error
	G회신() I회신_가변형
	S질의(질의_채널 chan I질의_가변형) I질의_가변형
	S회신(에러 error, 내용 ...interface{}) error
}

func New질의_가변형(메시지_구분 string, 내용 ...interface{}) I질의_가변형 {
	switch 메시지_구분 {
	case P메시지_GET:
	case P메시지_SET:
	case P메시지_DEL:
	case P메시지_초기화:
	case P메시지_종료:
	default:
		에러 := F에러_생성("잘못된 질의 메시지 구분 %v", 메시지_구분)
		F에러_출력(에러)
		//panic(에러)
		return nil
	}

	회신_채널 := make(chan I회신_가변형, 100)
	메시지 := New메시지_가변형(메시지_구분, 내용...).(s기본_메시지_가변형)

	return s질의_메시지_가변형{회신_채널: 회신_채널, s기본_메시지_가변형: 메시지}
}

// 회신
type I회신_가변형 interface {
	I메시지_가변형
	G에러() error
}

func New회신_가변형(에러 error, 내용 ...interface{}) I회신_가변형 {
	메시지_구분 := ""

	if 에러 == nil || F포맷된_문자열("%v", 에러) == "<nil>" {
		메시지_구분 = P메시지_OK
	} else {
		메시지_구분 = P메시지_에러
	}

	if len(내용) == 1 {
		가변형_모음, ok := 내용[0].([]interface{})

		if ok {
			내용 = make([]interface{}, 0)
			내용 = append(내용, 가변형_모음...)
		}
	}
	
	메시지 := New메시지_가변형(메시지_구분, 내용...).(s기본_메시지_가변형)

	return s회신_메시지_가변형{에러: 에러, s기본_메시지_가변형: 메시지}
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
	G단위() string
	G실수값() float64
	G정밀값() *dec.Decimal
	G문자열값() string
	G문자열값_고정소숫점(소숫점_이하_자릿수 int) string
	G비교(다른_통화 I통화) int
	G부호() int
	G복사본() I통화
	G변경불가() bool
	S동결()
	S더하기(값 float64) I통화
	S빼기(값 float64) I통화
	S곱하기(값 float64) I통화
	S나누기(값 float64) (I통화, error)
	S금액(값 float64) I통화
	String() string
}

// go-decimaldec.New()는 float64를 문자열로 바꾼 후 정밀값으로 변환하므로,
// 인수로 float64를 사용해도 큰 문제없다.

func New원화(금액 float64) I통화 { return New통화(KRW, 금액) }
func New달러(금액 float64) I통화 { return New통화(USD, 금액) }
func New유로(금액 float64) I통화 { return New통화(EUR, 금액) }
func New위안(금액 float64) I통화 { return New통화(CNY, 금액) }
func New통화(단위 string, 금액 float64) I통화 {
	에러 := F통화단위_검사(단위)
	if 에러 != nil {
		panic(에러)
		return nil
	}

	s := new(s통화)
	s.단위 = 단위
	s.금액 = dec.New(금액)
	s.변경불가 = false

	return s
}

func F통화단위_검사(통화단위 string) error {
	switch 통화단위 {
	case "KRW", "USD", "EUR", "CNY":
		return nil
	default:
		return F에러_생성("잘못된 통화단위 %v", 통화단위)
	}
}

// 가격정보
type I가격정보 interface {
	G종목코드() string
	G가격() I통화
	G시점() time.Time
}

func New가격정보(종목코드 string, 가격 I통화, 시점 time.Time) I가격정보 {
	s := s가격정보{종목코드: 종목코드, 가격: 가격.G복사본(), 시점: 시점}
	return &s
}

// 종목별 보유량
type I종목별_보유량 interface {
	G종목코드() string
	G롱포지션() int64
	G숏포지션() int64
	G순보유량() int64
	G총보유량() int64
	S더하기_롱포지션(수량 int64) error
	S더하기_숏포지션(수량 int64) error
}

func New종목별_보유량(종목코드 string, 롱포지션 int64, 숏포지션 int64) I종목별_보유량 {
	if 롱포지션 < 0 || 숏포지션 < 0 {
		에러 := F에러_생성("입력된 보유량 음수임. %v %v", 롱포지션, 숏포지션)
		F에러_출력(에러)
		panic(에러)
	}

	s := s종목별_보유량{종목코드: 종목코드, 롱포지션: 롱포지션, 숏포지션: 숏포지션}
	return &s
}

// 증권사
type I증권사 interface {
	G이름() string
}

func New증권사(이름 string) I증권사 {
	s := s증권사{이름: 이름}

	return &s
}

// 전송 권한 관련
type I전송_권한 interface {
	G전송_권한_획득()
	G남은_수량() int
}

func New전송_권한(수량 int, 재충전_주기 time.Duration) I전송_권한 {
	ch전송_권한 := make(chan S비어있는_구조체, 수량)

	s := s전송_권한{
		ch전송_권한:   ch전송_권한,
		재충전_주기:    재충전_주기,
		마지막_충전_시점: time.Now().Add(-2 * 재충전_주기)}

	// 첫 재충전은 무조건 실행되기 위해서 마지막 충전 시점을 과거 시점으로 설정했음.
	s.s전송_권한_재충전()

	return &s
}

type I코드별_전송_권한 interface {
	G코드() string
	I전송_권한
}

func New코드별_전송_권한(코드 string, 전송_권한_모음 []I전송_권한) I코드별_전송_권한 {
	s := s코드별_전송_권한{
		코드:       코드,
		전송_권한_모음: 전송_권한_모음}

	return &s
}

type I증권사별_전송_권한 interface {
	G증권사() I증권사
	G전송_권한_획득(코드 string)
	G남은_수량(코드 string) int
}

func New증권사별_전송_권한(
	증권사 I증권사,
	코드별_전송_권한_모음 []I코드별_전송_권한,
	기본_전송_권한_모음 []I전송_권한) I증권사별_전송_권한 {
	코드별_전송_권한_맵 := make(map[string]I코드별_전송_권한)

	for _, 코드별_전송_권한 := range 코드별_전송_권한_모음 {
		코드별_전송_권한_맵[코드별_전송_권한.G코드()] = 코드별_전송_권한
	}

	s := s증권사별_전송_권한{
		증권사:         증권사,
		코드별_전송_권한_맵: 코드별_전송_권한_맵,
		기본_전송_권한_모음: 기본_전송_권한_모음}

	return &s
}
