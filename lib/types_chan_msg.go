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
	"time"
)

// Go채널을 통해서 주고 받는 메시지.
type I채널_메시지 interface {
	G에러() error
	G값(인덱스 int) interface{}
	G값_모음() []interface{}
	G길이() int
}

func New채널_메시지(값_모음 ...interface{}) I채널_메시지 {
	s := new(s채널_메시지)
	s.값_모음, s.에러 = F2안전한_전달값_모음(값_모음...)

	return s
}

func New채널_메시지_에러(에러 interface{}) I채널_메시지 {
	if 에러 == nil {
		return New채널_메시지_비어있음()
	}

	s := new(s채널_메시지)
	s.에러 = New에러(에러)
	s.값_모음 = nil

	return s
}

func New채널_메시지_비어있음() I채널_메시지 {
	s := new(s채널_메시지)
	s.에러 = nil
	s.값_모음 = nil

	return s
}

type s채널_메시지 struct {
	에러   error
	값_모음 []interface{}
}

func (s s채널_메시지) G에러() error {
	return s.에러
}
func (s s채널_메시지) G값(인덱스 int) interface{} {
	return s.값_모음[인덱스]
}
func (s s채널_메시지) G값_모음() []interface{} {
	return s.값_모음
}
func (s s채널_메시지) G길이() int {
	if s.값_모음 == nil {
		return 0
	}

	return len(s.값_모음)
}

// Go의 chan(채널)을 통해서 주고 받는 질의.
// 송수신을 위해서 대기하느라 실행이 멈추는 것을 최소화 했음.
type I채널_질의 interface {
	G값(인덱스 int) interface{}
	G값_모음() []interface{}
	G검사(값_모음_길이 int) error
	S질의(값_모음 ...interface{}) I채널_질의
	S응답(회신 I채널_메시지)
	G응답() I채널_메시지                   // 타임아웃이 될 때까지 회신 대기
	G응답_DONT_WAIT() (I채널_메시지, bool) // 회신이 도착하지 않았으면 바로 반환
	G응답_추가_수신() bool
}

func New채널_질의(ch질의 chan<- I채널_질의, 타임아웃 time.Duration, 응답_버퍼_길이 int) I채널_질의 {
	if 타임아웃 != P무기한 && int64(타임아웃) <= 0 {
		panic(New에러with출력("타임 아웃은 0보다 커야 합니다.\n%v", 타임아웃))
	}

	s := new(s채널_질의)
	s.잠금 = make(chan T신호, 1)
	s.잠금 <- P신호 // 질의 가능하도록 초기화
	s.ch질의 = ch질의
	s.타임아웃 = 타임아웃
	s.응답_버퍼_길이 = 응답_버퍼_길이

	return s
}

type s채널_질의 struct {
	잠금       chan T신호
	값_모음     []interface{} // 질의값
	타임아웃     time.Duration // 질의 후 타임 아웃까지의 시간
	응답_버퍼_길이 int
	ch질의     chan<- I채널_질의 // 질의 목적지 채널
	ch응답     chan I채널_메시지
	ch타임아웃   <-chan time.Time // 타임 아웃이 되면 알려주는 채널
}

func (s *s채널_질의) 질의_잠금() {
	<-s.잠금
}
func (s *s채널_질의) 질의_잠금_해제(ch응답 chan I채널_메시지) {
	if ch응답 == s.ch응답 {
		// 현재 질의와 응답된 질의가 같을 때만 잠금 해제.
		// 새로운 질의가 시작되었고, 수신한 응답이 현재 질의와 무관하면,
		// 잠금 해제하지 않음.
		select {
		case s.잠금 <- P신호:
		default:
		}
	}
}

func (s *s채널_질의) G값(인덱스 int) interface{} {
	return s.값_모음[인덱스]
}
func (s *s채널_질의) G값_모음() []interface{} {
	return s.값_모음
}
func (s *s채널_질의) G검사(길이 int) error {
	if len(s.값_모음) != 길이 {
		에러 := New에러with출력("길이가 예상과 다릅니다. 예상값 %v, 실제값 %v", 길이, len(s.값_모음))
		//s.ch응답 <- New응답_에러(에러)
		return 에러
	}

	return nil
}

func (s *s채널_질의) S질의(값_모음 ...interface{}) I채널_질의 {
	s.질의_잠금()

	// 응답
	s.값_모음 = F확인(F2안전한_전달값_모음(값_모음...)).([]interface{})
	s.ch응답 = make(chan I채널_메시지, s.응답_버퍼_길이)
	s.ch타임아웃 = time.After(s.타임아웃)

	ch질의 := s.ch질의
	ch타임아웃 := s.ch타임아웃

	select {
	case ch질의 <- s:
	case <-ch타임아웃:
	}

	return s
}

func (s *s채널_질의) S응답(응답 I채널_메시지) {
	ch타임아웃 := s.ch타임아웃
	ch종료 := F공통_종료_채널()

	select {
	case s.ch응답 <- 응답:
	case <-ch타임아웃:
	case <-ch종료:
	}
}

func (s *s채널_질의) G응답() I채널_메시지 {
	ch응답 := s.ch응답
	ch타임아웃 := s.ch타임아웃
	ch종료 := F공통_종료_채널()

	defer s.질의_잠금_해제(ch응답)

	if s.타임아웃 == P무기한 {
		// 타임아웃이 P무기한이면 'ch타임아웃'을 고려하지 않음.
		select {
		case 응답 := <-ch응답:
			return 응답
		case <-ch종료:
			return New채널_메시지_비어있음()
		}
	} else {
		select {
		case 응답 := <-ch응답:
			return 응답
		case <-ch타임아웃:
			return New채널_메시지_에러("타임 아웃.")
		case <-ch종료:
			return New채널_메시지_비어있음()
		}
	}
}

func (s *s채널_질의) G응답_DONT_WAIT() (I채널_메시지, bool) {
	ch응답 := s.ch응답
	ch타임아웃 := s.ch타임아웃
	ch종료 := F공통_종료_채널()

	select {
	case 응답 := <-ch응답:
		s.질의_잠금_해제(ch응답)
		return 응답, true
	case <-ch타임아웃:
		s.질의_잠금_해제(ch응답)
		return New채널_메시지_에러("타임 아웃"), true
	case <-ch종료:
		s.질의_잠금_해제(ch응답)
		return New채널_메시지_에러("종료"), true
	default:
		return nil, false
	}
}

func (s *s채널_질의) G응답_추가_수신() bool {
	if len(s.ch응답) > 0 {
		return true
	}

	return false
}
