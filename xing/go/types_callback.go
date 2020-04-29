/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

package xg

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
	"sync"
	"time"
)

func newC32_콜백_대기_저장소() *c32_콜백_저장소 {
	s := new(c32_콜백_저장소)
	s.저장소 = make(map[int]*c32_콜백_대기_항목)
	s.최근_정리_시간 = time.Now()

	return s
}

type c32_콜백_대기_항목 struct {
	sync.Mutex
	식별번호   int
	ch회신   chan interface{}
	TR코드   string
	대기값    interface{}
	에러     error
	데이터_수신 bool
	메시지_수신 bool
	응답_완료  bool
	회신_완료  bool
	생성된_시각 time.Time
}

func (s *c32_콜백_대기_항목) G회신값() interface{} {
	switch 변환값 := s.대기값.(type) {
	case *xt.S이중_응답_일반형:
		return 변환값.G값(s.TR코드)
	case *xt.S헤더_반복값:
		return 변환값.G값(s.TR코드)
	default:
		return s.대기값
	}
}

func (s *c32_콜백_대기_항목) S회신() {
	if s.회신_완료 {
		return
	}

	if s.에러 != nil {
		select {
		case s.ch회신 <- s.에러:
		default:
			panic(lib.New에러with출력("채널 에러 회신 실패."))
		}
	} else {
		select {
		case s.ch회신 <- s.G회신값():
		default:
			panic(lib.New에러with출력("채널 회신 실패."))
		}
	}

	s.회신_완료 = true
}

//x32  응답을 기다리는 TR 저장.
type c32_콜백_저장소 struct {
	sync.RWMutex
	저장소      map[int]*c32_콜백_대기_항목
	최근_정리_시간 time.Time
}

func (s *c32_콜백_저장소) G값(식별번호 int) *c32_콜백_대기_항목 {
	s.s정리()

	s.RLock()
	값 := s.저장소[식별번호]
	s.RUnlock()

	return 값
}

func (s *c32_콜백_저장소) S추가(식별번호 int, TR코드 string) chan interface{} {
	s.s정리()

	대기_항목 := new(c32_콜백_대기_항목)
	대기_항목.식별번호 = 식별번호
	대기_항목.ch회신 = make(chan interface{}, 1)
	대기_항목.TR코드 = TR코드
	대기_항목.생성된_시각 = lib.F지금()

	s.Lock()
	s.저장소[식별번호] = 대기_항목
	s.Unlock()

	return 대기_항목.ch회신
}

func (s *c32_콜백_저장소) S회신(식별번호 int) {
	대기_항목 := s.G값(식별번호)
	대기_항목.S회신()

	s.Lock()
	delete(s.저장소, 식별번호)
	s.Unlock()
}

func (s *c32_콜백_저장소) s정리() {
	s.RLock()
	최근_정리_시간 := s.최근_정리_시간
	s.RUnlock()

	지금 := lib.F지금()

	if 지금.Sub(최근_정리_시간) < lib.P1분 {
		return // 정리한 지 얼마 안 되었음.
	}

	s.Lock()
	defer s.Unlock()

	for idx, 대기_항목 := range s.저장소 {
		if 지금.Sub(대기_항목.생성된_시각) > lib.P40초 {
			delete(s.저장소, idx)
		}
	}
}

func new대기_중_데이터_저장소() *s소켓_메시지_대기_저장소 {
	s := new(s소켓_메시지_대기_저장소)
	s.저장소 = make(map[*lib.S바이트_변환_모음]chan *lib.S바이트_변환_모음)

	return s
}
