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

package xing

import (
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/xing/base"
)

func F실시간_정보_구독_단순TR(RT코드 string) (에러 error) {
	return F질의(lib.New질의값_기본형(xt.TR실시간_정보_구독, RT코드)).G에러()
}

func F실시간_정보_해지_단순TR(RT코드 string) (에러 error) {
	return F질의(lib.New질의값_기본형(xt.TR실시간_정보_해지, RT코드)).G에러()
}

func F실시간_정보_구독(RT코드 string, 종목코드_모음 []string) (에러 error) {
	return F질의(lib.New질의값_복수종목(xt.TR실시간_정보_구독, RT코드, 종목코드_모음)).G에러()
}

func F실시간_정보_해지(RT코드 string, 종목코드_모음 []string) (에러 error) {
	return F질의(lib.New질의값_복수종목(xt.TR실시간_정보_해지, RT코드, 종목코드_모음)).G에러()
}

func F실시간_정보_일괄_해지() (에러 error) {
	return F질의(lib.New질의값_기본형(xt.TR실시간_정보_일괄_해지, "")).G에러()
}

func F실시간_데이터_구독_ETF(종목코드_모음 []string) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F확인(F실시간_정보_구독(xt.RT코스피_호가_잔량_H1, 종목코드_모음))
	lib.F확인(F실시간_정보_구독(xt.RT코스피_체결_S3, 종목코드_모음))
	lib.F확인(F실시간_정보_구독(xt.RT코스피_ETF_NAV_I5, 종목코드_모음))
	lib.F확인(F실시간_정보_구독(xt.RT코스피_시간외_호가_잔량_H2, 종목코드_모음))
	lib.F확인(F실시간_정보_구독(xt.RT코스피_예상_체결_YS3, 종목코드_모음))

	return nil
}

func F실시간_데이터_해지_ETF(종목코드_모음 []string) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F확인(F실시간_정보_해지(xt.RT코스피_호가_잔량_H1, 종목코드_모음))
	lib.F확인(F실시간_정보_해지(xt.RT코스피_체결_S3, 종목코드_모음))
	lib.F확인(F실시간_정보_해지(xt.RT코스피_ETF_NAV_I5, 종목코드_모음))
	lib.F확인(F실시간_정보_해지(xt.RT코스피_시간외_호가_잔량_H2, 종목코드_모음))
	lib.F확인(F실시간_정보_해지(xt.RT코스피_예상_체결_YS3, 종목코드_모음))

	return nil
}

func F주문_응답_실시간_정보_구독() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	if 주문_응답_구독_중.G값() {
		return
	}

	lib.F확인(F실시간_정보_구독_단순TR(xt.RT현물_주문_접수_SC0))
	lib.F확인(F실시간_정보_구독_단순TR(xt.RT현물_주문_체결_SC1))
	lib.F확인(F실시간_정보_구독_단순TR(xt.RT현물_주문_정정_SC2))
	lib.F확인(F실시간_정보_구독_단순TR(xt.RT현물_주문_취소_SC3))
	lib.F확인(F실시간_정보_구독_단순TR(xt.RT현물_주문_거부_SC4))

	return nil
}

func F주문_응답_실시간_정보_해지() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	if !주문_응답_구독_중.G값() {
		return
	}

	defer 주문_응답_구독_중.S값(false)

	lib.F확인(F실시간_정보_해지_단순TR(xt.RT현물_주문_접수_SC0))
	lib.F확인(F실시간_정보_해지_단순TR(xt.RT현물_주문_체결_SC1))
	lib.F확인(F실시간_정보_해지_단순TR(xt.RT현물_주문_정정_SC2))
	lib.F확인(F실시간_정보_해지_단순TR(xt.RT현물_주문_취소_SC3))
	lib.F확인(F실시간_정보_해지_단순TR(xt.RT현물_주문_거부_SC4))

	return nil
}
