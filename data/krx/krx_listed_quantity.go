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
//func F상장_주식_수량_맵() (상장주식수량_맵 map[string]int64, 에러 error) {
//	if len(map상장_주식_수량) == 0 {
//		for i := 0; i < 3; i++ { // map상장_주식_수량 초기화
//			if map상장_주식_수량, 에러 = f상장_주식_수량_맵(); 에러 == nil && len(map상장_주식_수량) > 1000 {
//				break
//			}
//		}
//	}
//
//	return f상장_주식_수량_맵_복사본(), 에러
//}

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
