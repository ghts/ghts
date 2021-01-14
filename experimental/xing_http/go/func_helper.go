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

package xing

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

var (
	F당일 = xt.F당일
	F전일 = xt.F전일
)

func f2에러(값 string) error {
	if strings.TrimSpace(값) == "" {
		return nil
	} else {
		return errors.New(값)
	}
}

func TR도우미(질의값 lib.I질의값, 결과값_포인터 interface{}) (에러 error) {
	return http질의_도우미(질의값.TR코드(), 질의값, 결과값_포인터)
}

func http질의_도우미(url string, 질의값, 결과값_포인터 interface{}) (에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 결과값_포인터 = nil }}.S실행()

	if !strings.HasPrefix(url, xt.F주소_C32_호출().URL()) {
		url = xt.F주소_C32_호출().URL() + "/" + url
	}

	바이트_모음_질의, 에러 := json.Marshal(질의값)
	lib.F확인(에러)

	http응답, 에러 := (&http.Client{Timeout: lib.P30초}).Post(url, "application/json", bytes.NewBuffer(바이트_모음_질의))
	lib.F확인(에러)

	바이트_모음_응답, 에러 := ioutil.ReadAll(http응답.Body)
	lib.F확인(에러)

	if lib.F종류(결과값_포인터) != reflect.Ptr {
		return lib.New에러with출력("포인터형이 아님. %T", 결과값_포인터)
	}

	//디버깅용 출력 문자열
	//if lib.F체크포인트(url); strings.Contains(url, "connected") {
	//	응답 := &xt.S응답_JSON{}
	//	lib.F체크포인트(url, 질의값)
	//	lib.F체크포인트(바이트_모음_응답)
	//	lib.F체크포인트(string(바이트_모음_응답))
	//	lib.F체크포인트(json.Unmarshal(바이트_모음_응답, 응답))
	//	lib.F체크포인트(응답)
	//	lib.F체크포인트(응답.E)
	//	lib.F체크포인트(응답.V)
	//
	//
	//	lib.F체크포인트(json.Unmarshal(바이트_모음_응답, 결과값_포인터))
	//	lib.F체크포인트(결과값_포인터)
	//}

	return json.Unmarshal(바이트_모음_응답, 결과값_포인터)
}

func F계좌번호_모음() (응답값 []string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌번호_모음 = nil }}.S실행()

	if len(계좌번호_모음) != 0 {
		return 계좌번호_모음, nil
	}

	s := struct {
		V []string
		E string
	}{nil, ""}

	lib.F확인(http질의_도우미("account_no_list", "", &s))

	if f2에러(s.E) == nil && len(s.V) > 0 {
		계좌번호_모음 = s.V
	}

	return s.V, f2에러(s.E)
}

func F계좌번호_존재함(계좌번호 string) bool {
	if 계좌번호_모음, 에러 := F계좌번호_모음(); 에러 != nil {
		lib.New에러with출력("계좌번호 모음 리스트 확보 실패.")
		return false
	} else {
		계좌번호 = strings.TrimSpace(계좌번호)

		for _, 값 := range 계좌번호_모음 {
			if 계좌번호 == 값 {
				return true
			}
		}
	}

	return false
}

func F계좌_상세명(계좌_번호 string) (계좌_상세명 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_상세명 = "" }}.S실행()

	s := struct {
		V string
		E error
	}{"", nil}

	lib.F확인(http질의_도우미("account_detail_name", 계좌_번호, &s))

	return s.V, s.E
}

func F접속됨() (접속됨 bool, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 접속됨 = false }, M출력_숨김: true}.S실행()

	if !lib.F포트_열림_확인(xt.F주소_C32_호출()) {
		return false, lib.New에러("TCP 포트 닫혀있음.")
	}

	질의값 := lib.New질의값_기본형(lib.TR접속됨, "")

	s := struct {
		V bool
		E string
	}{false, ""}

	lib.F확인(http질의_도우미("connected", 질의값, &s))

	return s.V, f2에러(s.E)
}

func F서버_구분() xt.T서버_구분 {
	return 서버_구분
}

func F계좌_번호(인덱스 int) (계좌_번호 string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 계좌_번호 = "" }}.S실행()

	if 계좌번호_모음, 에러 := F계좌번호_모음(); 에러 != nil {
		return "", 에러
	} else if 인덱스 >= len(계좌번호_모음) {
		return "", lib.New에러("잘못된 인덱스 %v 계좌번호 수량 %v", 인덱스, len(계좌번호_모음))
	} else {
		return 계좌번호_모음[인덱스], nil
	}
}

func F계좌_번호_단순형(인덱스 int) string {
	계좌_번호, 에러 := F계좌_번호(인덱스)
	lib.F확인(에러)

	return 계좌_번호
}
