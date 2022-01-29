/* Copyright (C) 2015-2022 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2022년 UnHa Kim (unha.kim@ghts.org)

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
	"sync"
	"time"
)

type I송수신 interface {
	S송신(변환_형식 T변환, 값_모음 ...interface{}) error
	G수신() (*S바이트_변환_모음, error)
}

type I소켓 interface {
	I송수신
	S송신_단순형(변환_형식 T변환, 값_모음 ...interface{})
	G수신_단순형() *S바이트_변환_모음
	S타임아웃(타임아웃 time.Duration) I소켓
	S옵션(옵션_모음 ...interface{})
	Close() error
}

type I소켓with컨텍스트 interface {
	I소켓
	G컨텍스트() (I송수신, error)
}

type I소켓_질의 interface {
	I소켓
	G질의_응답(변환_형식 T변환, 값_모음 ...interface{}) (*S바이트_변환_모음, error)
	G질의_응답_검사(변환_형식 T변환, 값_모음 ...interface{}) *S바이트_변환_모음
}

func New소켓_저장소(수량 int, 생성함수 func() I소켓_질의) *s소켓_저장소 {
	s := new(s소켓_저장소)
	s.M저장소 = make(chan I소켓_질의, 수량)
	s.M생성함수 = 생성함수

	return s
}

type s소켓_저장소 struct {
	sync.Mutex
	M저장소  chan I소켓_질의
	M생성함수 func() I소켓_질의
}

func (s *s소켓_저장소) G소켓() I소켓_질의 {
	select {
	case <-Ch공통_종료():
		return nil
	case 소켓 := <-s.M저장소:
		return 소켓
	default:
		s.Lock()
		defer s.Unlock()

		return s.M생성함수()
	}
}

func (s *s소켓_저장소) S회수(소켓 I소켓_질의) {
	if 소켓 == nil {
		return
	}

	소켓.S타임아웃(P30초)

	select {
	case s.M저장소 <- 소켓:
		return
	default:
		소켓.Close()
	}
}

func (s *s소켓_저장소) S정리() {
	for i := 0; i < len(s.M저장소); i++ {
		소켓 := <-s.M저장소
		소켓.Close()
	}
}
