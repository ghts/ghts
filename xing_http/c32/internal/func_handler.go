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

package x32_http

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"io/ioutil"
	"net/http"
)

func 계좌번호_리스트(w http.ResponseWriter, req *http.Request) {
	if 응답 := f질의_처리_도우미(w, lib.New질의값_기본형(xt.TR계좌번호_모음, "")); 응답.Error() != nil {
		F회신(w, 응답)
	} else if 계좌번호_모음, ok := 응답.V.([]string); !ok {
		F회신(w, xt.New응답(lib.New에러with출력("예상하지 못한 자료형 : '%v' '%T' ", 응답.V, 응답.V)))
	} else {
		F회신(w, xt.New응답(계좌번호_모음))
	}
}

//func 계좌_상세명(w http.ResponseWriter, req *http.Request) {
//	var 계좌번호 string
//
//	if 바이트_모음, 에러 := ioutil.ReadAll(req.Body); 에러 != nil {
//		F회신(w, xt.New응답(에러))
//	} else if 계좌번호 = strings.TrimSpace(string(바이트_모음)); 계좌번호 == "" {
//		F회신(w, xt.New응답(lib.New에러("비어있는 계좌번호 질의값.")))
//	} else if 응답 := f질의_처리_도우미(w, lib.New질의값_문자열(xt.TR계좌_상세명, "", 계좌번호)); 응답.E != nil {
//		F회신(w, 응답)
//	} else if 계좌_상세명, ok := 응답.V.(string); !ok {
//		F회신(w, xt.New응답(lib.New에러with출력("예상하지 못한 자료형 : '%v' '%T' ", 응답.V, 응답.V)))
//	} else {
//		F회신(w, xt.New응답(계좌_상세명))
//	}
//}

func CSPAQ12200(w http.ResponseWriter, req *http.Request) {
	if 바이트_모음, 에러 := ioutil.ReadAll(req.Body); 에러 != nil {
		F회신(w, xt.New응답(에러))
	} else if 계좌번호 := lib.F문자열_정리(string(바이트_모음)); 계좌번호 == "" {
		F회신(w, xt.New응답(lib.New에러("비어있는 계좌번호 질의값.")))
	} else if !F계좌번호_존재함(계좌번호) {
		F회신(w, xt.New응답(lib.New에러("존재하지 않는 계좌번호 : '%v'", 계좌번호)))
	} else {
		F질의_처리(w, lib.New질의값_문자열(xt.TR조회, xt.TR현물계좌_총평가_CSPAQ12200, 계좌번호))
	}
}

func CSPAQ12300(w http.ResponseWriter, req *http.Request) {
	질의값 := &xt.CSPAQ12300_현물계좌_잔고내역_질의값{}

	if 에러 := F질의값_추출(w, req, 질의값); 에러 != nil {
		return
	}

	질의값.M구분 = xt.TR조회
	질의값.M코드 = xt.TR현물계좌_잔고내역_조회_CSPAQ12300

	F질의_처리(w, 질의값)
}

func T0167(w http.ResponseWriter, req *http.Request) {
	F질의_처리(w, lib.New질의값_기본형(xt.TR조회, xt.TR시간_조회_t0167))
}

func T1305(w http.ResponseWriter, req *http.Request) {
	질의값 := &xt.T1305_현물_기간별_조회_질의값{}

	if 에러 := F질의값_추출(w, req, 질의값); 에러 != nil {
		return
	}

	질의값.M구분 = xt.TR조회
	질의값.M코드 = xt.TR현물_기간별_조회_t1305

	F질의_처리(w, 질의값)
}
