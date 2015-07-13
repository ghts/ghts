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
	"strings"
	"testing"
	"time"
)

func TestI안전한_bool(테스트 *testing.T) {
	테스트.Parallel()

	안전한_bool := New안전한_bool(false)

	F테스트_거짓임(테스트, 안전한_bool.G값())
	F테스트_에러발생(테스트, 안전한_bool.S값(false))
	F테스트_거짓임(테스트, 안전한_bool.G값())

	F테스트_에러없음(테스트, 안전한_bool.S값(true))
	F테스트_참임(테스트, 안전한_bool.G값())
}

func TestI안전한_string(테스트 *testing.T) {
	테스트.Parallel()

	안전한_string := New안전한_string("테스트")

	F테스트_같음(테스트, 안전한_string.G값(), "테스트")
	안전한_string.S값("테스트 2")
	F테스트_같음(테스트, 안전한_string.G값(), "테스트 2")
}

func TestI메시지(테스트 *testing.T) {
	테스트.Parallel()

	메시지 := New메시지(P메시지_GET, "테스트")
	F테스트_같음(테스트, 메시지.G구분(), P메시지_GET)
	F테스트_같음(테스트, 메시지.G내용(0), "테스트")
	F테스트_같음(테스트, 메시지.G길이(), 1)
	F테스트_같음(테스트, len(메시지.G내용_전체()), 메시지.G길이())
	F테스트_같음(테스트, 메시지.G내용_전체()[0], 메시지.G내용(0))

	문자열 := 메시지.String()
	F테스트_참임(테스트, strings.Contains(문자열, 메시지.G구분()))
	F테스트_참임(테스트, strings.Contains(문자열, 메시지.G내용(0)))

	메시지 = New메시지(P메시지_OK, "테스트", 1)
	F테스트_같음(테스트, 메시지.G구분(), P메시지_OK)
	F테스트_같음(테스트, 메시지.G내용(0), "테스트")
	F테스트_같음(테스트, 메시지.G내용(1), "1")
	F테스트_같음(테스트, 메시지.G길이(), 2)
	F테스트_같음(테스트, len(메시지.G내용_전체()), 메시지.G길이())
	F테스트_같음(테스트, 메시지.G내용_전체()[0], 메시지.G내용(0))
	F테스트_같음(테스트, 메시지.G내용_전체()[1], 메시지.G내용(1))

	문자열 = 메시지.String()
	F테스트_참임(테스트, strings.Contains(문자열, 메시지.G구분()))
	F테스트_참임(테스트, strings.Contains(문자열, 메시지.G내용(0)))
	F테스트_참임(테스트, strings.Contains(문자열, 메시지.G내용(1)))
}

func TestI질의_I회신(테스트 *testing.T) {
	테스트.Parallel()

	질의 := New질의(P메시지_GET, "질의", 1)

	_, ok := 질의.(I질의)
	F테스트_참임(테스트, ok)

	F테스트_같음(테스트, 질의.G구분(), P메시지_GET)
	F테스트_같음(테스트, 질의.G길이(), 2)
	F테스트_같음(테스트, 질의.G내용(0), "질의")
	F테스트_같음(테스트, 질의.G내용(1), "1")

	질의.G회신_채널() <- New회신(F에러_생성("에러"))

	회신 := <-질의.G회신_채널()

	_, ok = 회신.(I회신)
	F테스트_참임(테스트, ok)
	F테스트_같음(테스트, 회신.G에러().Error(), "에러")
	F테스트_같음(테스트, 회신.G구분(), P메시지_에러)
	F테스트_같음(테스트, 회신.G길이(), 0)

	질의.G회신_채널() <- New회신(nil, "회신", 10)

	회신 = <-질의.G회신_채널()
	F테스트_같음(테스트, 회신.G에러(), nil)
	F테스트_같음(테스트, 회신.G구분(), P메시지_OK)
	F테스트_같음(테스트, 회신.G길이(), 2)
	F테스트_같음(테스트, 회신.G내용(0), "회신")
	F테스트_같음(테스트, 회신.G내용(1), "10")
}

func TestI종목(테스트 *testing.T) {
	테스트.Parallel()

	종목 := New종목("코드", "이름")
	F테스트_같음(테스트, 종목.G코드(), "코드")
	F테스트_같음(테스트, 종목.G이름(), "이름")
}

func TestI통화(테스트 *testing.T) {
	테스트.Parallel()

	통화 := New통화(KRW, "100.01")
	F테스트_같음(테스트, 통화.G단위(), KRW)
	F테스트_같음(테스트, 통화.G실수값(), 100.01)
	F테스트_같음(테스트, 통화.G정밀값().Float(), 100.01)
	F테스트_같음(테스트, 통화.G실수_문자열(2), "100.01")
	F테스트_같음(테스트, 통화.G비교(New통화(KRW, "100.01")), P같음)
	F테스트_같음(테스트, 통화.G비교(New통화(KRW, "100.02")), P큼)
	F테스트_같음(테스트, 통화.G비교(New통화(KRW, "100.00")), P작음)
	F테스트_같음(테스트, 통화.G비교(New통화(USD, "100.00")), P비교불가)
	F테스트_같음(테스트, New통화(KRW, "100.00").G부호(), P양수)
	F테스트_같음(테스트, New통화(KRW, "-100.00").G부호(), P음수)
	F테스트_같음(테스트, New통화(KRW, "0.0").G부호(), P영)

	F테스트_같음(테스트, 통화.G복사본().G비교(New통화(KRW, "100.01")), P같음)
	F테스트_같음(테스트, 통화.G복사본().S금액("10.00").G비교(New통화(KRW, "10.00")), P같음)
	F테스트_같음(테스트, 통화.G비교(New통화(KRW, "100.01")), P같음)

	F테스트_거짓임(테스트, 통화.G변경불가())
	통화.S동결()
	F테스트_참임(테스트, 통화.G변경불가())
	F테스트_패닉발생(테스트, 통화.S더하기, New통화(KRW, "100.00"))
	F테스트_패닉발생(테스트, 통화.S빼기, New통화(KRW, "100.00"))
	F테스트_패닉발생(테스트, 통화.S곱하기, New통화(KRW, "100.00"))
	F테스트_패닉발생(테스트, 통화.S나누기, New통화(KRW, "100.00"))
	F테스트_패닉발생(테스트, 통화.S금액, "10.00")

	F테스트_같음(테스트, New통화(KRW, "100.00").S더하기(New통화(KRW, "100.00")).G비교(New통화(KRW, "200.00")), P같음)
	F테스트_같음(테스트, New통화(KRW, "100.00").S빼기(New통화(KRW, "100.00")).G비교(New통화(KRW, "0.00")), P같음)
	F테스트_같음(테스트, New통화(KRW, "100.00").S곱하기(New통화(KRW, "100.00")).G비교(New통화(KRW, "10000.00")), P같음)
	F테스트_같음(테스트, New통화(KRW, "100.00").S나누기(New통화(KRW, "100.00")).G비교(New통화(KRW, "1.00")), P같음)
	F테스트_같음(테스트, New통화(KRW, "100.00").String(), "KRW 100.00")

	F테스트_같음(테스트, New원화("100.00").G비교(New통화(KRW, "100.00")), P같음)
	F테스트_같음(테스트, New달러("100.00").G비교(New통화(USD, "100.00")), P같음)
	F테스트_같음(테스트, New유로("100.00").G비교(New통화(EUR, "100.00")), P같음)
	F테스트_같음(테스트, New위안("100.00").G비교(New통화(CNY, "100.00")), P같음)

	F문자열_출력_일시정지_시작()
	defer F문자열_출력_일시정지_해제()

	F테스트_같음(테스트, New통화(KRW, "Not_a_number"), nil)
	F테스트_같음(테스트, New원화("100").S금액("Not_a_number").G정밀값(), nil)
	F테스트_같음(테스트, New통화(KRW, "100.00").S더하기(New통화(USD, "100.00")).G정밀값(), nil)
	F테스트_같음(테스트, New통화(KRW, "100.00").S더하기(New통화(KRW, "100.00").S금액("Invalid_value")).G정밀값(), nil)
	F테스트_같음(테스트, New통화(KRW, "100.00").S금액("Invalid_value").S더하기(New통화(KRW, "100.00")).G정밀값(), nil)

	F테스트_같음(테스트, New통화(KRW, "100.00").S빼기(New통화(USD, "100.00")).G정밀값(), nil)
	F테스트_같음(테스트, New통화(KRW, "100.00").S빼기(New통화(KRW, "100.00").S금액("Invalid_value")).G정밀값(), nil)
	F테스트_같음(테스트, New통화(KRW, "100.00").S금액("Invalid_value").S빼기(New통화(KRW, "100.00")).G정밀값(), nil)

	F테스트_같음(테스트, New통화(KRW, "100.00").S곱하기(New통화(USD, "100.00")).G정밀값(), nil)
	F테스트_같음(테스트, New통화(KRW, "100.00").S곱하기(New통화(KRW, "100.00").S금액("Invalid_value")).G정밀값(), nil)
	F테스트_같음(테스트, New통화(KRW, "100.00").S금액("Invalid_value").S곱하기(New통화(KRW, "100.00")).G정밀값(), nil)

	F테스트_같음(테스트, New통화(KRW, "100.00").S나누기(New통화(USD, "100.00")).G정밀값(), nil)
	F테스트_같음(테스트, New통화(KRW, "100.00").S나누기(New통화(KRW, "100.00").S금액("Invalid_value")).G정밀값(), nil)
	F테스트_같음(테스트, New통화(KRW, "100.00").S금액("Invalid_value").S나누기(New통화(KRW, "100.00")).G정밀값(), nil)
	F테스트_같음(테스트, New통화(KRW, "100.00").S나누기(New통화(KRW, "0.00")).G정밀값(), nil)
}

func TestI가격정보(테스트 *testing.T) {
	테스트.Parallel()

	시점1 := time.Now()
	가격정보 := New가격정보(New종목("종목코드", "종목이름"), New원화("100.00"))
	시점2 := time.Now()

	F테스트_같음(테스트, 가격정보.G종목().G코드(), "종목코드")
	F테스트_같음(테스트, 가격정보.G종목().G이름(), "종목이름")
	F테스트_같음(테스트, 가격정보.G가격().G단위(), KRW)
	F테스트_같음(테스트, 가격정보.G가격().G실수값(), 100.0)

	F테스트_참임(테스트, 가격정보.G시점().Equal(시점1) || 가격정보.G시점().After(시점1))
	F테스트_참임(테스트, 가격정보.G시점().Equal(시점2) || 가격정보.G시점().Before(시점2))
}
