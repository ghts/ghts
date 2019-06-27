/* Copyright (C) 2015-2019 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package lib

import (
	"bytes"
	"container/list"
	"sync"
	"time"
)

type I_TR코드 interface {
	TR코드() string
}

type I질의값 interface {
	TR구분() TR구분
	TR코드() string
	G식별번호() int64 // NH Open API에서 필요함. Xing API에서는 사용하지 않음.
}

type I문자열 interface {
	G문자열() string
}

type I종목코드 interface {
	G종목코드() string
}

type I종목코드_모음 interface {
	G종목코드_모음() []string
	G전체_종목코드() string
}

func New질의값_기본형(TR구분 TR구분, TR코드 string) *S질의값_기본형 {
	s := new(S질의값_기본형)
	s.M구분 = TR구분
	s.M코드 = TR코드

	return s
}

type S질의값_기본형 struct {
	M구분   TR구분
	M코드   string
	M식별번호 int64
}

func (s S질의값_기본형) TR구분() TR구분   { return s.M구분 }
func (s S질의값_기본형) TR코드() string { return s.M코드 }
func (s S질의값_기본형) G식별번호() int64 { return s.M식별번호 }

func New질의값_정수(TR구분 TR구분, TR코드 string, 값 int) *S질의값_정수 {
	s := new(S질의값_정수)
	s.S질의값_기본형 = New질의값_기본형(TR구분, TR코드)
	s.M정수값 = 값

	return s
}

type S질의값_정수 struct {
	*S질의값_기본형
	M정수값 int
}

func New질의값_문자열(TR구분 TR구분, TR코드 string, 값 string) *S질의값_문자열 {
	s := new(S질의값_문자열)
	s.S질의값_기본형 = New질의값_기본형(TR구분, TR코드)
	s.M문자열 = 값

	return s
}

type S질의값_문자열 struct {
	*S질의값_기본형
	M문자열 string
}

func New질의값_문자열_모음(TR구분 TR구분, TR코드 string, 값 []string) *S질의값_문자열_모음 {
	s := new(S질의값_문자열_모음)
	s.S질의값_기본형 = New질의값_기본형(TR구분, TR코드)
	s.M문자열_모음 = 값

	return s
}

type S질의값_문자열_모음 struct {
	*S질의값_기본형
	M문자열_모음 []string
}

func New질의값_바이트_변환(TR구분 TR구분, TR코드 string, 값 interface{}) *S질의값_바이트_변환 {
	s := new(S질의값_바이트_변환)
	s.S질의값_기본형 = New질의값_기본형(TR구분, TR코드)
	s.M바이트_변환 = New바이트_변환_단순형(P변환형식_기본값, 값)

	return s
}

type S질의값_바이트_변환 struct {
	*S질의값_기본형
	M바이트_변환 *S바이트_변환
}

func New질의값_바이트_변환_모음(TR구분 TR구분, TR코드 string, 값_모음 ...interface{}) *S질의값_바이트_변환_모음 {
	s := new(S질의값_바이트_변환_모음)
	s.S질의값_기본형 = New질의값_기본형(TR구분, TR코드)
	s.M바이트_변환_모음 = New바이트_변환_모음_단순형(P변환형식_기본값, 값_모음...)

	return s
}

type S질의값_바이트_변환_모음 struct {
	*S질의값_기본형
	M바이트_변환_모음 *S바이트_변환_모음
}

type S질의값_단일_종목 struct {
	*S질의값_기본형
	M종목코드 string
}

func (s S질의값_단일_종목) G종목코드() string { return s.M종목코드 }

func New질의값_단일_종목() *S질의값_단일_종목 {
	s := new(S질의값_단일_종목)
	s.S질의값_기본형 = new(S질의값_기본형)

	return s
}

type S질의값_단일종목_연속키 struct {
	*S질의값_단일_종목
	M연속키 string
}

func New질의값_단일종목_연속키() *S질의값_단일종목_연속키 {
	s := new(S질의값_단일종목_연속키)
	s.S질의값_단일_종목 = New질의값_단일_종목()

	return s
}

type S질의값_복수_종목 struct {
	*S질의값_기본형
	M종목코드_모음 []string
}

func (s S질의값_복수_종목) G종목코드_모음() []string {
	return F슬라이스_복사(s.M종목코드_모음, nil).([]string)
}

func (s S질의값_복수_종목) G전체_종목코드() string {
	버퍼 := new(bytes.Buffer)

	for _, 종목코드 := range s.M종목코드_모음 {
		버퍼.WriteString(종목코드)
	}

	return 버퍼.String()
}

func New질의값_복수종목(TR구분 TR구분, TR코드 string, 종목코드_모음 []string) *S질의값_복수_종목 {
	if 종목코드_모음 == nil {
		종목코드_모음 = make([]string, 0)
	}

	s := new(S질의값_복수_종목)
	s.S질의값_기본형 = New질의값_기본형(TR구분, TR코드)
	s.M종목코드_모음 = 종목코드_모음

	return s
}

type S질의값_정상_주문 struct {
	*S질의값_단일_종목
	M증권사      T증권사
	M계좌번호     string
	M주문수량     int64
	M주문단가     int64
	M매도_매수_구분 T매도_매수_구분
	M호가유형     T호가유형
	M주문조건     T주문조건
}

func New질의값_정상_주문() *S질의값_정상_주문 {
	s := new(S질의값_정상_주문)
	s.S질의값_단일_종목 = New질의값_단일_종목()

	return s
}

type S질의값_정정_주문 struct {
	*S질의값_단일_종목
	M증권사   T증권사
	M계좌번호  string
	M원주문번호 int64
	M주문수량  int64
	M주문단가  int64
}

func New질의값_정정_주문() *S질의값_정정_주문 {
	s := new(S질의값_정정_주문)
	s.S질의값_단일_종목 = New질의값_단일_종목()

	return s
}

type S질의값_취소_주문 struct {
	*S질의값_단일_종목
	M증권사   T증권사
	M계좌번호  string
	M원주문번호 int64
	M주문수량  int64
}

func New질의값_취소_주문() *S질의값_취소_주문 {
	s := new(S질의값_취소_주문)
	s.S질의값_단일_종목 = New질의값_단일_종목()

	return s
}

// 전송 권한 관련
type I전송_권한 interface {
	I_TR코드
	G획득() I전송_권한
	S기록()
	G남은_수량() int
}

func New전송_권한(TR코드 string, 수량 int, 간격 time.Duration) I전송_권한 {
	s := new(s전송_권한)
	s.tr코드 = TR코드
	s.수량 = 수량
	s.간격 = 간격 + P100밀리초
	s.전송_기록_저장소 = list.New()

	switch {
	case 수량 > 10:
		s.수량 -= 1
	case 수량 > 100:
		s.수량 -= 2
	}

	return s
}

type s전송_권한 struct {
	sync.Mutex
	tr코드      string
	수량        int
	간격        time.Duration
	전송_기록_저장소 *list.List
}

func (s *s전송_권한) TR코드() string { return s.tr코드 }

func (s *s전송_권한) G획득() I전송_권한 {
	s.Lock()

	if s.G남은_수량() <= 0 {
		전송_시각 := s.전송_기록_저장소.Front().Value.(time.Time)

		if s.간격 > P10분 {
			지금 := F지금()
			F체크포인트(F2문자열("%v : %v초 대기 예정.",
				지금.Format("15:04:05.999"),
				s.간격-지금.Sub(전송_시각)/P1초))
		}

		F대기(s.간격 - F지금().Sub(전송_시각))

		if s.간격 > P10분 {
			F체크포인트(F2문자열("%v : %v초 대기 완료.", s.tr코드, F지금().Format("15:04:05.999")))
		}
	}

	return s
}

func (s *s전송_권한) S기록() {
	s.전송_기록_저장소.PushBack(F지금())
	s.Unlock()
}

func (s *s전송_권한) G남은_수량() int {
	s.s오래된_전송_기록_정리()

	return s.수량 - s.전송_기록_저장소.Len()
}

func (s *s전송_권한) s오래된_전송_기록_정리() {
	지금 := F지금()

	for {
		if s.전송_기록_저장소.Len() == 0 {
			return
		}

		전송_기록 := s.전송_기록_저장소.Front()
		전송_시각 := 전송_기록.Value.(time.Time)

		if 지금.Sub(전송_시각) > s.간격 {
			s.전송_기록_저장소.Remove(전송_기록)
		} else {
			return
		}
	}
}

type S문자열_모음 struct {
	M배열 []string
}

type S채널_질의_API struct {
	M질의값  I질의값
	Ch회신값 chan interface{}
	Ch에러  chan error
}
