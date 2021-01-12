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
MERCHANTABILITY or FITNESS FOR A PAxt.RTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package x32

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"net/http"
)

func 접속됨(w http.ResponseWriter, req *http.Request) {
	if !로그인_완료.G값() {
		F회신(w, xt.New응답(false))
	} else if 응답 := f질의_처리_도우미(w, lib.New질의값_기본형(xt.TR접속됨, "")); 응답.Error() != nil {
		F회신(w, 응답)
	} else if 접속_여부, ok := 응답.V.(bool); !ok {
		F회신(w, xt.New응답(lib.New에러with출력("예상하지 못한 자료형 : '%v' '%T' ", 응답.V, 응답.V)))
	} else {
		F회신(w, xt.New응답(접속_여부))
	}
}

func 종료(w http.ResponseWriter, req *http.Request) {
	if 응답 := f질의_처리_도우미(w, lib.New질의값_기본형(xt.TR종료, "")); 응답.Error() != nil {
		F회신(w, 응답)
	} else {
		F회신(w, xt.New응답(lib.P신호_종료))
	}
}

func 계좌번호_리스트(w http.ResponseWriter, req *http.Request) {
	if 응답 := f질의_처리_도우미(w, lib.New질의값_기본형(xt.TR계좌번호_모음, "")); 응답.Error() != nil {
		F회신(w, 응답)
	} else if 계좌번호_모음, ok := 응답.V.([]string); !ok {
		F회신(w, xt.New응답(lib.New에러with출력("예상하지 못한 자료형 : '%v' '%T' ", 응답.V, 응답.V)))
	} else {
		F회신(w, xt.New응답(계좌번호_모음))
	}
}

func 계좌_상세명(w http.ResponseWriter, req *http.Request) {
	if 계좌번호, 에러 := F문자열_추출(req); 에러 != nil {
		F회신(w, xt.New응답(에러))
	} else if 응답 := f질의_처리_도우미(w, lib.New질의값_문자열(xt.TR계좌_상세명, "", 계좌번호)); 응답.Error() != nil {
		F회신(w, 응답)
	} else if 계좌_상세명, ok := 응답.V.(string); !ok {
		F회신(w, xt.New응답(lib.New에러with출력("예상하지 못한 자료형 : '%v' '%T' ", 응답.V, 응답.V)))
	} else {
		F회신(w, xt.New응답(계좌_상세명))
	}
}

// 미리 질의값 자료형이 고정되어 있으므로 1개로 충분.
func 실시간_정보_구독_및_해지(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_RT처리(w, req)
}

func CSPAQ12200_현물계좌_총평가(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물계좌_총평가_CSPAQ12200,
		&lib.S질의값_문자열{})
}

func CSPAQ12300_현물계좌_잔고내역_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물계좌_잔고내역_조회_CSPAQ12300,
		&xt.CSPAQ12300_현물계좌_잔고내역_질의값{})
}

func CSPAQ13700_현물계좌_주문체결내역_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물계좌_주문체결내역_조회_CSPAQ13700,
		&xt.CSPAQ13700_현물계좌_주문체결내역_질의값{})
}

func CSPAQ22200_현물계좌_예수금_주문가능금액(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물계좌_예수금_주문가능금액_CSPAQ22200,
		&lib.S질의값_문자열{})
}

func CSPAT00600_현물_정상_주문(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR주문, xt.TR현물_정상_주문_CSPAT00600,
		&xt.CSPAT00600_현물_정상_주문_질의값{})
}

func CSPAT00700_현물_정정_주문(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR주문, xt.TR현물_정정_주문_CSPAT00700,
		&xt.CSPAT00700_현물_정정_주문_질의값{})
}

func CSPAT00800_현물_취소_주문(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR주문, xt.TR현물_취소_주문_CSPAT00800,
		&lib.S질의값_취소_주문{})
}

func T0150_현물_당일_매매일지(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_당일_매매일지_t0150,
		&xt.T0150_현물_당일_매매일지_질의값{})
}

func T0151_현물_일자별_매매일지(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_일자별_매매일지_t0151,
		&xt.T0151_현물_일자별_매매일지_질의값{})
}

func T0167_시간_조회(w http.ResponseWriter, req *http.Request) {
	F질의_처리(w, lib.New질의값_기본형(xt.TR조회, xt.TR시간_조회_t0167))
}

func T0425_현물_체결_미체결_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_체결_미체결_조회_t0425,
		&xt.T0425_현물_체결_미체결_조회_질의값{})
}

func T1101_현물_호가_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_호가_조회_t1101,
		&lib.S질의값_단일_종목{})
}

func T1102_현물_시세_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_시세_조회_t1102,
		&lib.S질의값_단일_종목{})
}

func T1305_현물_기간별_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_기간별_조회_t1305,
		&xt.T1305_현물_기간별_조회_질의값{})
}

func T1310_현물_당일_전일_분틱_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_당일_전일_분틱_조회_t1310,
		&xt.T1310_현물_전일당일분틱조회_질의값{})
}

func T1404_관리_불성실_투자유의_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR관리_불성실_투자유의_조회_t1404,
		&xt.T1404_관리종목_조회_질의값{})
}

func T1405_투자경고_매매정지_정리매매_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR투자경고_매매정지_정리매매_조회_t1405,
		&xt.T1405_투자경고_조회_질의값{})
}

func T1902_ETF_시간별_추이(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR_ETF_시간별_추이_t1902,
		lib.New질의값_단일종목_연속키())
}

func T1906_ETF_LP호가_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR_ETF_LP호가_조회_t1906,
		lib.New질의값_단일_종목())
}

func T3341_재무순위_종합(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR재무순위_종합_t3341,
		&xt.T3341_재무순위_질의값{})
}

func T8407_현물_멀티_현재가_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_멀티_현재가_조회_t8407,
		&lib.S질의값_복수_종목{})
}

func T8411_현물_차트_틱(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_차트_틱_t8411,
		&xt.T8411_현물_차트_틱_질의값{})
}

func T8412_현물_차트_분(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_차트_분_t8412,
		&xt.T8412_현물_차트_분_질의값{})
}

func T8413_현물_차트_일주월(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_차트_일주월_t8413,
		&xt.T8413_현물_차트_일주월_질의값{})
}

func T8428_증시_주변_자금_추이(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR증시_주변_자금_추이_t8428,
		&xt.T8428_증시주변_자금추이_질의값{})
}

func T8436_현물_종목_조회(w http.ResponseWriter, req *http.Request) {
	F질의값_추출_TR처리(w, req, xt.TR조회, xt.TR현물_종목_조회_t8436,
		&lib.S질의값_문자열{})
}
