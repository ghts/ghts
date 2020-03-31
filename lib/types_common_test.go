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

package lib

import (
	"strings"
	"testing"
)

func TestS통화(t *testing.T) {
	//t.Parallel()  // 문자열_출력_일시정지 모드로 인하여 병렬 실행 불가
	통화 := New통화(KRW, 100.01)
	F테스트_같음(t, 통화.G단위(), KRW)
	F테스트_같음(t, 통화.G실수64(), 100.01)
	F테스트_같음(t, 통화.G정밀값().String(), "100.01")
	F테스트_같음(t, 통화.G문자열(), "100.01")
	F테스트_같음(t, 통화.G문자열_고정소숫점(1), "100.0")
	F테스트_같음(t, 통화.G문자열_고정소숫점(2), "100.01")
	F테스트_같음(t, 통화.G문자열_고정소숫점(3), "100.010")

	F테스트_같음(t, 통화.G비교(New통화(KRW, 100.01)), P비교_같음)
	F테스트_같음(t, 통화.G비교(New통화(KRW, 100.00)), P비교_큼)
	F테스트_같음(t, 통화.G비교(New통화(KRW, 100.02)), P비교_작음)
	F테스트_같음(t, 통화.G비교(New통화(USD, 100.00)), P비교_불가)

	// 복사본에 변경을 가해도 원본값은 변경되지 않는 지 확인.
	통화 = New통화(KRW, 100.01)
	F테스트_같음(t, 통화.G복사본().G비교(New통화(KRW, 100.01)), P비교_같음)
	F테스트_같음(t, 통화.G복사본().S금액(10.00).G비교(New통화(KRW, 10.00)), P비교_같음)
	F테스트_같음(t, 통화.G비교(New통화(KRW, 100.01)), P비교_같음)
	F테스트_같음(t, 통화.G복사본().G비교(New통화(KRW, 100.01)), P비교_같음)

	통화 = New통화(KRW, 100.01)
	F테스트_거짓임(t, 통화.G변경불가())
	통화.S동결()
	F테스트_참임(t, 통화.G변경불가())

	F문자열_출력_일시정지_시작()
	F테스트_에러발생(t, F패닉2에러(통화.S더하기, 100.0))
	F테스트_에러발생(t, F패닉2에러(통화.S빼기, 100.0))
	F테스트_에러발생(t, F패닉2에러(통화.S곱하기, 100.0))
	F테스트_에러발생(t, F패닉2에러(통화.S나누기, 100.0))
	F테스트_에러발생(t, F패닉2에러(통화.S금액, 100.0))
	F문자열_출력_일시정지_해제()

	F테스트_같음(t, New통화(KRW, 100.00).S더하기(100).G비교(New통화(KRW, 200.00)), P비교_같음)
	F테스트_같음(t, New통화(KRW, 100.00).S빼기(100).G비교(New통화(KRW, 0.00)), P비교_같음)
	F테스트_같음(t, New통화(KRW, 100.00).S곱하기(100.0).G비교(New통화(KRW, 10000.00)), P비교_같음)

	통화, 에러 := New통화(KRW, 100.00).S나누기(100.0)
	F테스트_에러없음(t, 에러)
	F테스트_같음(t, 통화.G비교(New통화(KRW, 1.00)), P비교_같음)

	F테스트_같음(t, New통화(KRW, 100.00).String(), "KRW 100")

	F테스트_같음(t, New원화(100.00).G비교(New통화(KRW, 100.00)), P비교_같음)
	F테스트_같음(t, New달러(100.00).G비교(New통화(USD, 100.00)), P비교_같음)
	F테스트_같음(t, New유로(100.00).G비교(New통화(EUR, 100.00)), P비교_같음)
	F테스트_같음(t, New위안(100.00).G비교(New통화(CNY, 100.00)), P비교_같음)

	F문자열_출력_일시정지_시작()
	통화, 에러 = New통화(KRW, 100.00).S나누기(0.0)
	F문자열_출력_일시정지_해제()

	F테스트_에러발생(t, 에러)
	F테스트_같음(t, 통화, nil)

	통화 = New통화(KRW, 100.01)
	통화.S동결()

	바이트_모음, 에러 := 통화.MarshalBinary()
	F테스트_에러없음(t, 에러)

	복원값 := New통화(KRW, 0.0)
	에러 = 복원값.UnmarshalBinary(바이트_모음)
	F테스트_에러없음(t, 에러)

	F테스트_같음(t, 복원값.G단위(), 통화.G단위())
	F테스트_같음(t, 복원값.G실수64(), 통화.G실수64())
	F테스트_같음(t, 복원값.G변경불가(), 통화.G변경불가())

	바이트_모음, 에러 = 통화.MarshalText()
	F테스트_에러없음(t, 에러)

	복원값 = new(S통화)
	에러 = 복원값.UnmarshalText(바이트_모음)
	F테스트_에러없음(t, 에러)
	F테스트_같음(t, 복원값.G단위(), 통화.G단위())
	F테스트_같음(t, 복원값.G실수64(), 통화.G실수64())
	F테스트_같음(t, 복원값.G변경불가(), 통화.G변경불가())
}

func TestI문자열_집합(t *testing.T) {
	t.Parallel()

	문자열_집합 := New문자열_집합()

	문자열_집합.S추가("테스트1")
	문자열_집합.S추가("테스트2")
	문자열_집합.S추가("테스트3")
	문자열_집합.S추가("테스트4")
	문자열_집합.S추가("테스트1")
	문자열_집합.S추가("테스트1")

	F테스트_같음(t, 문자열_집합.G길이(), 4)
	F테스트_참임(t, 문자열_집합.G포함("테스트1"))
	F테스트_참임(t, 문자열_집합.G포함("테스트2"))
	F테스트_참임(t, 문자열_집합.G포함("테스트3"))
	F테스트_참임(t, 문자열_집합.G포함("테스트4"))
	F테스트_거짓임(t, 문자열_집합.G포함("포함되지_않은_문자열"))

	출력_문자열 := 문자열_집합.String()
	F테스트_참임(t, 출력_문자열[:1] == "[")
	F테스트_참임(t, 출력_문자열[len(출력_문자열)-1:] == "]")
	F테스트_참임(t, 출력_문자열[len(출력_문자열)-2:] != ",")
	F테스트_참임(t, 출력_문자열[len(출력_문자열)-2:] != " ")
	F테스트_참임(t, strings.Contains(출력_문자열, "테스트1"))
	F테스트_참임(t, strings.Contains(출력_문자열, "테스트2"))
	F테스트_참임(t, strings.Contains(출력_문자열, "테스트3"))
	F테스트_참임(t, strings.Contains(출력_문자열, "테스트4"))
	F테스트_거짓임(t, strings.Contains(출력_문자열, "포함되지_않은_문자열"))

	값_모음 := 문자열_집합.G슬라이스()
	값_확인 := true

	for _, 값 := range 값_모음 {
		switch 값 {
		case "테스트1", "테스트2", "테스트3", "테스트4":
			continue
		}

		값_확인 = false
		break
	}

	F테스트_참임(t, 값_확인, 문자열_집합)

	문자열_집합.S삭제("테스트2")
	F테스트_같음(t, 문자열_집합.G길이(), 3)
	F테스트_참임(t, 문자열_집합.G포함("테스트1"))
	F테스트_거짓임(t, 문자열_집합.G포함("테스트2"))
	F테스트_참임(t, 문자열_집합.G포함("테스트3"))
	F테스트_참임(t, 문자열_집합.G포함("테스트4"))
	F테스트_거짓임(t, 문자열_집합.G포함("포함되지_않은_문자열"))

	값_모음 = 문자열_집합.G슬라이스()
	값_확인 = true

	for _, 값 := range 값_모음 {
		switch 값 {
		case "테스트1", "테스트3", "테스트4":
			continue
		}

		값_확인 = false
		break
	}

	F테스트_참임(t, 값_확인, 문자열_집합)

	출력_문자열 = 문자열_집합.String()
	F테스트_참임(t, 출력_문자열[:1] == "[")
	F테스트_참임(t, 출력_문자열[len(출력_문자열)-1:] == "]")
	F테스트_참임(t, 출력_문자열[len(출력_문자열)-2:] != ",")
	F테스트_참임(t, 출력_문자열[len(출력_문자열)-2:] != " ")
	F테스트_참임(t, strings.Contains(출력_문자열, "테스트1"))
	F테스트_거짓임(t, strings.Contains(출력_문자열, "테스트2"))
	F테스트_참임(t, strings.Contains(출력_문자열, "테스트3"))
	F테스트_참임(t, strings.Contains(출력_문자열, "테스트4"))
	F테스트_거짓임(t, strings.Contains(출력_문자열, "포함되지_않은_문자열"))
}
