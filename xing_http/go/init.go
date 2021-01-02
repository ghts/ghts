/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

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

package xing_http

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
)

func F전일_당일_설정() (에러 error) {
	lib.S예외처리{M에러: &에러}.S실행()

	질의값 := xt.NewT1305_현물_기간별_조회_질의값()
	질의값.M구분 = xt.TR조회
	질의값.M코드 = xt.TR현물_기간별_조회_t1305
	질의값.M종목코드 = "069500"
	질의값.M일주월_구분 = xt.P일주월_일
	질의값.M연속키 = ""
	질의값.M수량 = 5

	s := struct {
		V xt.T1305_현물_기간별_조회_응답
		E string
	}{xt.T1305_현물_기간별_조회_응답{}, ""}

	lib.F확인(HTTP질의_도우미(xt.TR현물_기간별_조회_t1305, 질의값, &s))
	lib.F조건부_패닉(s.E != "", s.E)
	값_모음 := s.V.M반복값_모음.M배열

	당일 := 값_모음[0].M일자
	전일 := 값_모음[1].M일자

	xt.F전일_당일_설정(당일, 전일)

	return nil
}
