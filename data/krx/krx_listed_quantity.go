package krx

import (
	"encoding/csv"
lb "github.com/ghts/ghts/lib"
"net/url"
"strings"
)

var map상장_주식_수량 map[string]int64

const 상장_주식_정보_최소_수량 = 2000

// F상장_주식_수량_맵 : HTTP 쿼리 빈도를 최소화 하기 위해서 로컬 캐시('map상장_주식_수량')를 사용.
// 참고자료 : https://statools.tistory.com/175
func F상장_주식_수량_맵() (상장주식수량_맵 map[string]int64, 에러 error) {
	if len(map상장_주식_수량) < 상장_주식_정보_최소_수량 {
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
	defer lb.S예외처리{M에러: &에러}.S실행()

	레코드_모음 := lb.F확인2(krx종목정보_불러오기())
	상장주식수량_맵 = make(map[string]int64)
	상장주식수량_인덱스 := len(레코드_모음[0]) - 1

	for _, 레코드 := range 레코드_모음 {
		종목코드 := 레코드[1]

		if len(종목코드) != 6 {
			panic(lb.New에러("종목코드 추출 오류. 잘못된 칼럼을 선택했습니다."))
		}

		if 상장주식수량, 에러 := lb.F2정수64(레코드[상장주식수량_인덱스]); 에러 == nil && 상장주식수량 > 0 {
			상장주식수량_맵[종목코드] = 상장주식수량
		}
	}

	return 상장주식수량_맵, nil
}

// krx종목정보_불러오기 : https://statools.tistory.com/175 참조. KRX 종목정보를 CSV형태로 다운로드.
func krx종목정보_불러오기() (레코드_모음 [][]string, 에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	const url_OTP = "https://data.krx.co.kr/comm/fileDn/GenerateOTP/generate.cmd"

	form데이터_OTP := url.Values{
		"mktId":       {"ALL"},
		"share":       {"1"},
		"csvxls_isNo": {"false"},
		"name":        {"fileDown"},
		"url":         {"dbms/MDC/STAT/standard/MDCSTAT01901"}}

	OTP := string(lb.F확인2(lb.HTTP_POST(url_OTP, form데이터_OTP)))

	//lb.F문자열_출력("OTP : '%v'", OTP)

	const url_CSV = "https://data.krx.co.kr/comm/fileDn/download_csv/download.cmd"

	form데이터_CSV := url.Values{"code": {OTP}}

	csv문자열 := lb.F2문자열_EUC_KR(lb.F확인2(lb.HTTP_POST(url_CSV, form데이터_CSV)))
	레코드_모음 = lb.F확인2(csv.NewReader(strings.NewReader(csv문자열)).ReadAll())

	return 레코드_모음[1:], nil
}
