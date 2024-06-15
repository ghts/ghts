/* Copyright (C) 2015-2024 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2024년 UnHa Kim (unha.kim@ghts.org)

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

package krx

import (
	"encoding/csv"
	"github.com/ghts/ghts/lib"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var map상장_주식_수량 map[string]int64

// F상장_주식_수량_맵 : HTTP 쿼리 빈도를 최소화 하기 위해서 로컬 캐시('map상장_주식_수량')를 사용.
// 참고자료 : https://statools.tistory.com/175
func F상장_주식_수량_맵() (상장주식수량_맵 map[string]int64, 에러 error) {
	if len(map상장_주식_수량) == 0 {
		for i := 0; i < 3; i++ { // map상장_주식_수량 초기화
			if map상장_주식_수량, 에러 = f상장_주식_수량_맵(); 에러 == nil && len(map상장_주식_수량) > 1000 {
				break
			}
		}
	}

	return f상장_주식_수량_맵_복사본(), 에러
}

func f상장_주식_수량_맵_복사본() (복사본 map[string]int64) {
	복사본 = make(map[string]int64)

	for 종목코드, 상장_주식_수량 := range map상장_주식_수량 {
		복사본[종목코드] = 상장_주식_수량
	}

	return 복사본
}

func f상장_주식_수량_맵() (상장주식수량_맵 map[string]int64, 에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	레코드_모음, 에러 := csv.NewReader(strings.NewReader(lib.F확인2(csv다운로드()))).ReadAll()
	상장주식수량_맵 = make(map[string]int64)
	상장주식수량_인덱스 := len(레코드_모음[0]) - 1

	for _, 레코드 := range 레코드_모음 {
		종목코드 := 레코드[1]

		if 상장주식수량, 에러 := lib.F2정수64(레코드[상장주식수량_인덱스]); 에러 == nil && 상장주식수량 > 0 {
			상장주식수량_맵[종목코드] = 상장주식수량
		}
	}

	return 상장주식수량_맵, nil
}

func csv다운로드() (CSV string, 에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	const url_CSV = "https://data.krx.co.kr/comm/fileDn/download_csv/download.cmd"

	폼_데이터 := url.Values{"code": {lib.F확인2(otp())}}

	응답 := lib.F확인2(http.PostForm(url_CSV, 폼_데이터))
	defer 응답.Body.Close()

	바이트_모음 := lib.F확인2(io.ReadAll(응답.Body))
	CSV = lib.F2문자열_EUC_KR(바이트_모음)

	return CSV, nil
}

func otp() (OTP string, 에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	const url_OTP = "https://data.krx.co.kr/comm/fileDn/GenerateOTP/generate.cmd"

	폼_데이터 := url.Values{
		"locale":      {"ko_KR"},
		"mktId":       {"ALL"},
		"share":       {"1"},
		"csvxls_isNo": {"false"},
		"name":        {"fileDown"},
		"url":         {"dbms/MDC/STAT/standard/MDCSTAT01901"}}

	응답 := lib.F확인2(http.PostForm(url_OTP, 폼_데이터))
	defer 응답.Body.Close()

	바이트_모음 := lib.F확인2(io.ReadAll(응답.Body))
	OTP = string(바이트_모음)

	return OTP, nil
}
