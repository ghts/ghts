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

package xt

import (
	"github.com/ghts/ghts/lib"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func f속성값_초기화(질의값 interface{}) interface{} {
	값 := reflect.ValueOf(질의값).Elem()

	for i := 0; i < 값.NumField(); i++ {
		switch {
		case !strings.HasPrefix(값.Type().Field(i).Name, "X_"),
			값.Field(i).Kind() != reflect.Uint8,
			!값.Field(i).CanSet():
			continue
		}

		값.Field(i).SetUint(0x20)
	}

	return 값.Interface()
}

func F서버_구분() T서버_구분 {
	if 구분값, 에러 := lib.F2정수(os.Getenv(P서버_구분_환경변수명)); 에러 == nil && T서버_구분(구분값) == P서버_실거래 {
		return P서버_실거래
	} else {
		return P서버_모의투자
	}
}

func F서버_구분_설정(서버_구분 T서버_구분) {
	os.Setenv(P서버_구분_환경변수명, strconv.Itoa(int(서버_구분)))
}

func F주소_C32_호출() lib.T주소 {
	for {
		// 환경변수를 통하면 자동으로 자식 프로세스에 같은 값이 전달된다.
		if 주소, 에러 := lib.F2정수(os.Getenv(P주소_C32_호출_환경변수명)); 에러 != nil {
			F주소_설정()
		} else {
			return lib.T주소(주소)
		}
	}
}

func F주소_C32_콜백() lib.T주소 {
	for {
		// 환경변수를 통하면 자동으로 자식 프로세스에 같은 값이 전달된다.
		if 주소, 에러 := lib.F2정수(os.Getenv(P주소_C32_콜백_환경변수명)); 에러 != nil {
			F주소_설정()
		} else {
			return lib.T주소(주소)
		}
	}
}

func F주소_실시간() lib.T주소 {
	for {
		// 환경변수를 통하면 자동으로 자식 프로세스에 같은 값이 전달된다.
		if 주소, 에러 := lib.F2정수(os.Getenv(P주소_실시간_환경변수명)); 에러 != nil {
			F주소_설정()
		} else {
			return lib.T주소(주소)
		}
	}
}

func F주소_설정() {
	주소_설정_잠금.Lock()
	defer 주소_설정_잠금.Unlock()

	if 주소_설정_완료.G값() {
		return
	}

	for {
		임의_포트_번호 := lib.F임의_범위_이내_정수값(0, 30000)
		주소_C32_호출 := lib.T주소(임의_포트_번호)
		주소_C32_콜백 := lib.T주소(임의_포트_번호 + 1)
		주소_실시간 := lib.T주소(임의_포트_번호 + 2)

		if lib.F포트_닫힘_확인(주소_C32_호출) &&
			lib.F포트_닫힘_확인(주소_C32_콜백) &&
			lib.F포트_닫힘_확인(주소_실시간) {
			// 환경변수를 통하면 자동으로 자식 프로세스에 같은 값이 전달된다.
			os.Setenv(P주소_C32_호출_환경변수명, strconv.Itoa(int(주소_C32_호출)))
			os.Setenv(P주소_C32_콜백_환경변수명, strconv.Itoa(int(주소_C32_콜백)))
			os.Setenv(P주소_실시간_환경변수명, strconv.Itoa(int(주소_실시간)))
			주소_설정_완료.S값(true)

			return
		}
	}
}

func F주소_재설정() {
	주소_설정_완료.S값(false)
	F주소_설정()
}
