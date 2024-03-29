/* Copyright (C) 2015-2023 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2023년 UnHa Kim (unha.kim@ghts.org)

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
	"testing"
)

type I모의_테스트 interface {
	S모의_테스트_리셋()
}

type S모의_테스트 struct {
	sync.Mutex
	*testing.T
	통과_여부 bool
}

func (s *S모의_테스트) G값() bool {
	s.Lock()
	defer s.Unlock()

	return s.통과_여부
}

func (s *S모의_테스트) S값(통과_여부 bool) {
	s.Lock()
	defer s.Unlock()

	s.통과_여부 = 통과_여부
}

func (s *S모의_테스트) Error(args ...interface{})                 { s.S값(false) }
func (s *S모의_테스트) Errorf(format string, args ...interface{}) { s.S값(false) }
func (s *S모의_테스트) Fail()                                     { s.S값(false) }
func (s *S모의_테스트) FailNow()                                  { s.S값(false) }
func (s *S모의_테스트) Failed() bool                              { return !s.G값() }
func (s *S모의_테스트) Fatal(args ...interface{})                 { s.S값(false) }
func (s *S모의_테스트) Fatalf(format string, args ...interface{}) { s.S값(false) }
func (s *S모의_테스트) Log(args ...interface{})                   {}
func (s *S모의_테스트) Logf(format string, args ...interface{})   {}
func (s *S모의_테스트) Skip(args ...interface{})                  {}
func (s *S모의_테스트) SkipNow()                                  {}
func (s *S모의_테스트) Skipf(format string, args ...interface{})  {}
func (s *S모의_테스트) Skipped() bool                             { return false }
func (s *S모의_테스트) S모의_테스트_리셋()                               { s.S값(true) }

// F에러패닉_처리()의 인수.
type S예외처리 struct {
	M에러    *error
	M함수    func()
	M함수_항상 func()
	M출력_숨김 bool
}

func (s S예외처리) S실행() {
	defer func() {
		if s.M함수_항상 != nil {
			s.M함수_항상()
		}
	}()

	패닉_복원값 := recover()

	var 에러 error
	에러_포인터 := s.M에러

	switch {
	case 패닉_복원값 != nil:
		if s.M출력_숨김 {
			에러 = New에러(패닉_복원값)
		} else {
			에러 = New에러with출력(패닉_복원값)
		}
	case 에러_포인터 != nil && *에러_포인터 != nil:
		if s.M출력_숨김 {
			에러 = New에러(*에러_포인터)
		} else {
			에러 = New에러with출력(*에러_포인터)
		}
	default: // 에러 및 패닉 없음.
		return
	}

	if 에러_포인터 != nil {
		*에러_포인터 = 에러
	}

	if s.M함수 != nil {
		s.M함수()
	}
}
