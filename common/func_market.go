package common

import (
	"time"
)

var 한국증시_최근_개장일 = New안전한_시각(time.Time{})
var 한국증시_최근_개장일_갱신_시간 = New안전한_시각(time.Time{})

func F한국증시_장중() bool {
	지금 := time.Now()
	최근_개장일, 에러 := F한국증시_최근_개장일()
	F에러_패닉(에러)
	개장_시각 := time.Date(최근_개장일.Year(), 최근_개장일.Month(),
		최근_개장일.Day(), 9, 0, 0, 0, 지금.Location())
	폐장_시각 := time.Date(최근_개장일.Year(), 최근_개장일.Month(),
		최근_개장일.Day(), 15, 0, 0, 0, 지금.Location())

	if 지금.Year() == 최근_개장일.Year() &&
		지금.Month() == 최근_개장일.Month() &&
		지금.Day() == 최근_개장일.Day() &&
		지금.After(개장_시각) &&
		지금.Before(폐장_시각) {
		return true
	}

	return false
}

func F한국증시_최근_개장일() (time.Time, error) {
	지금 := time.Now()
	금일_0시_5분 := time.Date(지금.Year(), 지금.Month(), 지금.Day(),
		0, 5, 0, 0, 지금.Location())
	일분전 := time.Now().Add(-1 * time.Minute)
	삼십일전 := 지금.Add(-30 * 24 * time.Hour)
	최근_갱신_완료 := false

	if 한국증시_최근_개장일_갱신_시간.G값().After(금일_0시_5분) ||
		한국증시_최근_개장일_갱신_시간.G값().After(일분전) {
		최근_갱신_완료 = true
	}

	// 최근 갱신했고, 임시 보관 중인 값이 정상적일 경우, 갱신하지 않고 바로 반환.
	if 최근_갱신_완료 &&
		한국증시_최근_개장일.G값().After(삼십일전) {
		return 한국증시_최근_개장일.G값(), nil
	}

	// 개장일 데이터 갱신 반환
	도우미_함수_모음 := [](func() (time.Time, error)){
		f최근_개장일_다음넷,
		f최근_개장일_네이버, 
		f최근_개장일_네이트,
		f최근_개장일_야후}

	최근_개장일_모음 := make([]time.Time, 0)

	for _, 도우미_함수 := range 도우미_함수_모음 {
		최근_개장일, 에러 := 도우미_함수()

		if 에러 != nil {
			continue
		}

		최근_개장일_모음 = append(최근_개장일_모음, 최근_개장일)
	}

	switch len(최근_개장일_모음) {
	case 0:
		에러 := F에러("개장일 데이터를 하나도 수집하지 못함.")
		return time.Time{}, 에러
	case 1:
		return 최근_개장일_모음[0], nil
	}

	for i := 0; i < len(최근_개장일_모음); i++ {
		for j := i; j < len(최근_개장일_모음); j++ {
			개장일1 := 최근_개장일_모음[i]
			개장일2 := 최근_개장일_모음[j]

			if 개장일1.Year() == 개장일2.Year() &&
				개장일1.Month() == 개장일2.Month() &&
				개장일1.Day() == 개장일2.Day() {
				최근_개장일_반환값 := time.Date(
					개장일1.Year(), 개장일1.Month(), 개장일1.Day(),
					0, 0, 0, 0, time.Now().Location())

				한국증시_최근_개장일_갱신_시간.S값(지금)
				한국증시_최근_개장일.S값(최근_개장일_반환값)

				return 최근_개장일_반환값, nil
			}
		}
	}

	에러 := F에러("개장일 데이터가 서로 일치하지 않음.")
	return time.Time{}, 에러
}

func f최근_개장일_다음넷() (time.Time, error) {
	본문, 에러 := F_HTTP회신_본문("http://finance.daum.net")

	if 에러 != nil {
		return time.Time{}, 에러
	}

	정규_표현식_모음_장중 := []string{
		`<span class='time'>[0-9]{1,2}:[0-9]{1,2}</span>(\t| ){0,5}<span class='txt_standard'>실시간</span>`}
	장중_문자열 := F문자열_검색_복수_정규식(본문, 정규_표현식_모음_장중)
	
	정규_표현식_모음_장외 := []string{
		`<span class='date'>[0-9]{1,2}\.[0-9]{1,2}</span>`,
		`[0-9]{1,2}\.[0-9]{1,2}`}
	장외_문자열 := F문자열_검색_복수_정규식(본문, 정규_표현식_모음_장외)
	
	switch {
	case 장중_문자열 == "" && 장외_문자열 == "":
		return time.Time{}, F에러("검색 결과 없음")
	case 장중_문자열 != "" && 장외_문자열 == "":
		지금 := time.Now()
		최근_개장일 := time.Date(지금.Year(), 지금.Month(), 지금.Day(),
			0, 0, 0, 0, 지금.Location())
		
		return 최근_개장일, nil
	case 장중_문자열 == "" && 장외_문자열 != "":
		연도 := F2문자열(time.Now().Year())
		일자_문자열 := 연도 + "." + 장외_문자열

		return F2포맷된_시각("2006.01.02", 일자_문자열)
	default:
		return time.Time{}, F에러("예상치 못한 경우")
	}
}

func f최근_개장일_네이버() (time.Time, error) {
	본문, 에러 := F_HTTP회신_본문("http://finance.naver.com")

	if 에러 != nil {
		return time.Time{}, 에러
	}

	정규_표현식_모음 := []string{
		`<span id="time">(.|\r|\n)+` +
			`[0-9]{4}\.[0-9]{1,2}\.[0-9]{1,2}` +
			`(.|\r|\n)+<span>`,
		`[0-9]{4}\.[0-9]{1,2}\.[0-9]{1,2}`}

	일자_문자열 := F문자열_검색_복수_정규식(본문, 정규_표현식_모음)

	if 일자_문자열 == "" {
		return time.Time{}, F에러("검색 결과 없음")
	}

	return F2포맷된_시각("2006.01.02", 일자_문자열)
}

func f최근_개장일_네이트() (time.Time, error) {
	본문, 에러 := F_HTTP회신_본문("http://stock.nate.com")

	if 에러 != nil {
		return time.Time{}, 에러
	}

	정규_표현식_모음 := []string{
		`[0-9]{4}/[0-9]{1,2}/[0-9]{1,2} [0-9]{1,2}:[0-9]{1,2}`}

	일자_문자열 := F문자열_검색_복수_정규식(본문, 정규_표현식_모음)
	
	if 일자_문자열 == "" {
		return time.Time{}, F에러("검색 결과 없음")
	}

	시각, 에러 := F2포맷된_시각("2006/01/02 15:04", 일자_문자열)
	
	if 에러 != nil {
		F에러_출력(에러)
		return time.Time{}, 에러
	}
	
	최근_개장일 := time.Date(시각.Year(), 시각.Month(), 시각.Day(), 
		0, 0, 0, 0, 시각.Location())

	return 최근_개장일, nil
}

func f최근_개장일_야후() (time.Time, error) {
	url := `https://finance.yahoo.com/q?uhb=uh3_finance_vert&fr=&type=2button&s=000030.KS`

	본문, 에러 := F_HTTP회신_본문(url)

	if 에러 != nil {
		return time.Time{}, 에러
	}
	
	정규_표현식_모음_장중 := []string{
		`<span class="time_rtq">.{0,100}</span>`,
		`[0-9]{1,2}:[0-9]{2}(AM|PM) [A-Z]{3}`}
	장중_문자열 := F문자열_검색_복수_정규식(본문, 정규_표현식_모음_장중)

	정규_표현식_모음_장외 := []string{
		`<span class="time_rtq">.{0,100}</span>`,
		`[a-zA-Z]{3} [0-9]{1,2}, [0-9]{1,2}:[0-9]{2}(AM|PM) [A-Z]{3}`}
	장외_문자열 := F문자열_검색_복수_정규식(본문, 정규_표현식_모음_장외)
	
	switch {
	case 장중_문자열 == "" && 장외_문자열 == "":
		return time.Time{}, F에러("검색 결과 없음")
	case 장중_문자열 != "" && 장외_문자열 == "":
		지금 := time.Now()
		최근_개장일 := time.Date(지금.Year(), 지금.Month(), 지금.Day(),
			0, 0, 0, 0, 지금.Location())
		return 최근_개장일, nil
	case 장중_문자열 == "" && 장외_문자열 != "":
		연도 := F2문자열(time.Now().Year())
		일자_문자열 := 연도 + " " + 장외_문자열
		일자, 에러 := F2포맷된_시각("3:04PM MST", 일자_문자열)
	
		if 에러 != nil {
			return time.Time{}, 에러
		}
	
		일자 = time.Date(일자.Year(), 일자.Month(), 일자.Day(),
				0, 0, 0, 0, 일자.Location())
	
		return 일자, nil
	default:
		return time.Time{}, F에러("예상치 못한 경우")
	}	
}
