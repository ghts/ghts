package krx

import (
	"github.com/PuerkitoBio/goquery"
lb "github.com/ghts/ghts/lib"
"io"
"net/http"
"strings"
"time"
)
type S상장_법인_정보 struct {
	M회사명  string
	M종목코드 string
	M업종   string
	M주요제품 string
	M상장일  time.Time
	M결산월  time.Month
}

var map상장_법인_정보 map[string]*S상장_법인_정보

// F상장_법인_정보_맵 : HTTP 쿼리 빈도를 최소화 하기 위해서 로컬 캐시('map상장_법인_정보')를 사용.
func F상장_법인_정보_맵() (법인정보_맵 map[string]*S상장_법인_정보, 에러 error) {
	if len(map상장_법인_정보) == 0 {
		for i := 0; i < 3; i++ { // map상장_법인_정보 초기화
			if map상장_법인_정보, 에러 = f상장_법인_정보_맵(); 에러 == nil && len(법인정보_맵) > 1000 {
				break
			}
		}
	}

	return f상장_법인_정보_맵_복사본(), nil
}

func f상장_법인_정보_맵_복사본() (복사본 map[string]*S상장_법인_정보) {
	복사본 = make(map[string]*S상장_법인_정보)

	for 키, 값 := range map상장_법인_정보 {
		복사값 := new(S상장_법인_정보)
		복사값.M회사명 = 값.M회사명
		복사값.M종목코드 = 값.M종목코드
		복사값.M업종 = 값.M업종
		복사값.M주요제품 = 값.M주요제품
		복사값.M상장일 = f복사_Time(값.M상장일)
		복사값.M결산월 = 값.M결산월

		복사본[키] = 복사값
	}

	return 복사본
}

func f복사_Time(값 time.Time) time.Time {
	return time.Date(값.Year(), 값.Month(), 값.Day(), 값.Hour(), 값.Minute(), 값.Second(), 값.Nanosecond(), 값.Location())
}

func f상장_법인_정보_맵() (법인정보_맵 map[string]*S상장_법인_정보, 에러 error) {
	defer lb.S예외처리{M에러: &에러, M함수: func() { 법인정보_맵 = nil }}.S실행()

	url := `https://kind.krx.co.kr/corpgeneral/corpList.do?method=download&searchType=13`
	응답 := lb.F확인2(http.Get(url))
	defer 응답.Body.Close()

	응답값 := lb.F확인2(io.ReadAll(응답.Body))
	문서 := lb.F확인2(goquery.NewDocumentFromReader(strings.NewReader(lb.F2문자열_EUC_KR(응답값))))
	법인정보_맵 = make(map[string]*S상장_법인_정보)

	문서.Find("body > table > tbody > tr").Each(func(행 int, s *goquery.Selection) {
		법인_정보 := new(S상장_법인_정보)

		s.Find("td").Each(func(열 int, s *goquery.Selection) {
			문자열 := lb.F2문자열_공백_제거(s.Text())

			switch 열 {
			case 0:
				법인_정보.M회사명 = 문자열
			case 1:
				법인_정보.M종목코드 = 문자열
			case 2:
				법인_정보.M업종 = 문자열
			case 3:
				법인_정보.M주요제품 = 문자열
			case 4:
				if 상장일, 에러 := lb.F2포맷된_일자(lb.P일자_형식, 문자열); 에러 == nil {
					법인_정보.M상장일 = 상장일
				} else {
					lb.F문자열_출력("상장일 에러 : %v '%v'", 법인_정보.M종목코드, 문자열)
				}
			case 5:
				if 월_정수, 에러 := lb.F2정수(lb.F정규식_검색(문자열, []string{`[0-9]+`})); 에러 == nil {
					월_모음 := []time.Month{
						time.January,
						time.February,
						time.March,
						time.April,
						time.May,
						time.June,
						time.July,
						time.August,
						time.September,
						time.October,
						time.November,
						time.December}

					for _, 월 := range 월_모음 {
						if 월_정수 == int(월) {
							법인_정보.M결산월 = 월
							break
						}
					}
				} else {
					lb.F문자열_출력("결산월 에러 : %v '%v'", 법인_정보.M종목코드, 문자열)
				}
			default:
				// PASS
			}
		})

		if 법인_정보.M종목코드 != "" {
			법인정보_맵[법인_정보.M종목코드] = 법인_정보
		}
	})

	return 법인정보_맵, nil
}
