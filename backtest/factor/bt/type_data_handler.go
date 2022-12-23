package bt

import (
	"bytes"
	"database/sql"
	bfc "github.com/ghts/ghts/backtest/factor/common"
	"github.com/ghts/ghts/lib"
	dpd "github.com/ghts/ghts/lib/daily_price_data"
	xing "github.com/ghts/ghts/xing/go"
	_ "modernc.org/sqlite"
	"strings"
	"time"
)

// S데이터_처리기_백테스트 : 백테스트용 I데이터_처리기 구현체. 실제 매매의 경우에는 API를 통해서 증권사 서버에 질의하게 될 기능을 가상으로 구현.
type S데이터_처리기_백테스트 struct {
	M일자          uint32
	M가격구분        T가격_구분
	일일_가격정보_맵_원본 map[string]*dpd.S종목별_일일_가격정보_모음
}

func (s *S데이터_처리기_백테스트) G영업일_모음() []uint32 {
	const KODEX200 = "069500"
	return s.일일_가격정보_맵_원본[KODEX200].G일자_정수값_모음()
}

func (s *S데이터_처리기_백테스트) G현재가(종목코드 string) float64 {
	return s.G가격(s.M일자, s.M가격구분, 종목코드)
}

func (s *S데이터_처리기_백테스트) G현재가_맵(종목코드_모음 []string) (현재가_맵 map[string]float64) {
	현재가_맵 = make(map[string]float64)

	for _, 종목코드 := range 종목코드_모음 {
		현재가_맵[종목코드] = s.G가격(s.M일자, s.M가격구분, 종목코드)
	}

	return 현재가_맵
}

func (s *S데이터_처리기_백테스트) ETF_ETN_종목_여부(종목코드 string) bool {
	return xing.ETF_ETN_종목_여부(종목코드)
}

func (s *S데이터_처리기_백테스트) G제외_종목_여부(종목코드 string) bool {
	return !xing.ETF_ETN_종목_여부(종목코드) && // ETF/ETN 제외
		!xing.F특수_종목_여부(종목코드) && // SPAC, 리츠 제외
		!xing.F지주회사_종목_여부(종목코드) &&
		!xing.F금융사_종목_여부(종목코드) &&
		!strings.HasPrefix(종목코드, "9") // 한국 상장 해외 기업 제외. (좋은 해외 기업은 한국이 아니라 미국에 상장됨.)
}

func (s *S데이터_처리기_백테스트) S상장_주식_수량_확인(종목코드_모음 []string) []string {
	return 종목코드_모음 // 백테스트에서는 확인할 필요 없음.
}

func (s *S데이터_처리기_백테스트) S준비(일자_정수값 uint32, 가격구분 T가격_구분) {
	s.M일자 = 일자_정수값
	s.M가격구분 = 가격구분
}

func (s *S데이터_처리기_백테스트) G일자() time.Time {
	return lib.F확인2(lib.F정수2일자(s.M일자))
}

func (s *S데이터_처리기_백테스트) G가격(기준일 uint32, 가격구분 T가격_구분, 종목코드 string) float64 {
	가격정보, 에러 := s.일일_가격정보_맵_원본[종목코드].G값(기준일)

	if 에러 != nil {
		return 0.0
	}

	switch 가격구분 {
	case P시가:
		return 가격정보.M시가
	case P고가:
		return 가격정보.M고가
	case P저가:
		return 가격정보.M저가
	case P종가:
		return 가격정보.M종가
	default:
		panic(lib.New에러("예상하지 못한 가격구분 : '%v", int(s.M가격구분)))
	}
}

func (s *S데이터_처리기_백테스트) G가격_맵(기준일 uint32, 가격구분 T가격_구분, 종목코드_모음 []string) (가격_맵 map[string]float64) {
	가격_맵 = make(map[string]float64)

	for _, 종목코드 := range 종목코드_모음 {
		가격_맵[종목코드] = s.G가격(기준일, 가격구분, 종목코드)
	}

	return 가격_맵
}

func (s *S데이터_처리기_백테스트) G기준일_시가_종가_평균(기준일 uint32, 종목코드 string) float64 {
	if 가격정보, 에러 := s.일일_가격정보_맵_원본[종목코드].G값(기준일); 에러 != nil {
		return 0.0
	} else {
		return 가격정보.M시가 + 가격정보.M종가
	}
}

func (s *S데이터_처리기_백테스트) G기준일_이전_전종목_가격정보_맵(기준일 uint32) (추출본 map[string]*dpd.S종목별_일일_가격정보_모음) {
	추출본 = make(map[string]*dpd.S종목별_일일_가격정보_모음)

	for 키, 값 := range s.일일_가격정보_맵_원본 {
		추출본[키] = 값.G기준일_이전_정보_복사본(기준일)
	}

	return 추출본
}

var db = lib.F확인2(sql.Open("sqlite", bfc.DB파일명()))

func f상장주식수량(종목코드 string, 기준일 uint32) float64 {
	lib.F확인1(bfc.F시가총액_테이블_생성(db))

	SQL := new(bytes.Buffer)
	SQL.WriteString("SELECT listed_qty FROM marcap ")
	SQL.WriteString("WHERE code=?")
	SQL.WriteString(" AND date = ?")

	행_모음 := lib.F확인2(db.Query(SQL.String()))
	행_모음.Next()

	var 상장주식수량 float64

	if 에러 := 행_모음.Scan(&상장주식수량); 에러 == nil {
		return 상장주식수량
	}

	return 0
}
