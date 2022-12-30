package ftfw

import (
	bfc "github.com/ghts/ghts/backtest/factor/common"
	"github.com/ghts/ghts/lib"
	dpd "github.com/ghts/ghts/lib/daily_price_data"
	xing "github.com/ghts/ghts/xing/go"
)

type S팩터_세종 struct {
	M종목명 string
	*bfc.S재무_세종
	M전년_재무_정보 *bfc.S재무_세종
	// --- 비재무 정보 ---
	M상장주식수량 float64
	M전일_거래량 float64
	M기준가    float64
	// --- 사이즈 팩터 ---
	M시가총액    float64
	M시가총액_등급 float64
	// --- 가치 팩터 ---
	PSR   float64
	PSR등급 float64
	POR   float64
	POR등급 float64
	PER   float64
	PER등급 float64
	PBR   float64
	PBR등급 float64
	PEG   float64
	PEG등급 float64
	// --- 성장성 팩터 ---
	M매출액_성장율      float64
	M매출액_성장율_등급   float64
	M영업이익_성장율     float64
	M영업이익_성장율_등급  float64
	M당기순이익_성장율    float64
	M당기순이익_성장율_등급 float64
	// --- 수익성 팩터 --
	M자산회전율    float64 // 자산회전율 = 매출액 / 자산
	M자산회전율_등급 float64
	M자본회전율    float64 // 자본회전율 = 매출액 / 자본
	M자본회전율_등급 float64
	OPA       float64 // 영업이익 / 자산
	OPA등급     float64
	OPE       float64 // 영업이익 / 자본
	OPE등급     float64
	OPM       float64 // 영업이익 / 매출액
	OPM등급     float64
	ROA       float64 // 순이익 / 자산
	ROA등급     float64
	ROE       float64 // 순이익 / 자본
	ROE등급     float64
	NPM       float64 // 순이익 / 매출액
	NPM등급     float64
	// --- 추세 팩터 ---
	M3개월_모멘텀           float64
	M3개월_모멘텀_등급        float64
	M6개월_모멘텀           float64
	M6개월_모멘텀_등급        float64
	M9개월_모멘텀           float64
	M9개월_모멘텀_등급        float64
	M12개월_모멘텀          float64
	M12개월_모멘텀_등급       float64
	M12개월_고점_대비_하락율    float64
	M12개월_고점_대비_하락율_등급 float64
	// --- 퀄리티 팩터 ---
	M부채_비율        float64
	M부채_비율_등급     float64
	M부채_증가율       float64
	M부채_증가율_등급    float64
	M자본_증가율       float64
	M자본_증가율_등급    float64
	M자산_증가율       float64
	M자산_증가율_등급    float64
	M거래_금액_중간값_억  float64 // 1개월 거래 금액 중간값
	M거래_금액_중간값_등급 float64
	M가격_변동성       float64 // 6개월 일일 수익율 표준 편차
	M가격_변동성_등급    float64
	M거래금액_변동성     float64 // 6개월 거래 금액 변동율 표준 편차.
	M거래금액_변동성_등급  float64
	// --- 급등락 ---
	M최근_급등  bool
	M최근_급등일 uint32
	M최근_급락  bool
	M최근_급락일 uint32
	// --- 팩터 카테고리별 점수 ---
	M가치_점수  float64
	M가치_등급  float64
	M성장성_점수 float64
	M성장성_등급 float64
	M수익성_점수 float64
	M수익성_등급 float64
	M추세_점수  float64
	M추세_등급  float64
	M퀄리티_점수 float64
	M퀄리티_등급 float64
	// --- 최종 종합 점수 ---
	M복합_점수 float64
	M복합_등급 float64
}

func (s S팩터_세종) G최근_급등() bool    { return s.M최근_급등 }
func (s S팩터_세종) G최근_급락() bool    { return s.M최근_급락 }
func (s S팩터_세종) G복합_등급() float64 { return s.M복합_등급 }

func New팩터_세종(종목코드 string, 기준일 uint32,
	일일_가격정보_모음 *dpd.S종목별_일일_가격정보_모음,
	재무_정보_저장소 *S재무_정보_저장소[*bfc.S재무_세종]) (s *S팩터_세종) {
	var 에러 error
	defer lib.S예외처리{M에러: &에러, M함수: func() { s = nil }}.S실행()

	최신_재무_정보 := 재무_정보_저장소.G기준일_최신_연도_정보(종목코드, 기준일)
	전년_재무_정보 := 재무_정보_저장소.G기준일_차최신_연도_정보(종목코드, 기준일)

	if 일일_가격정보_모음 == nil ||
		일일_가격정보_모음.Len() < 12*21+10 ||
		최신_재무_정보 == nil ||
		전년_재무_정보 == nil ||
		일일_가격정보_모음.G종가() <= 0 {
		return nil
	}

	_1개월전_종가 := 일일_가격정보_모음.G이전_종가(1 * 21)
	_3개월전_종가 := 일일_가격정보_모음.G이전_종가(3 * 21)
	_6개월전_종가 := 일일_가격정보_모음.G이전_종가(6 * 21)
	_9개월전_종가 := 일일_가격정보_모음.G이전_종가(9 * 21)
	_12개월전_종가 := 일일_가격정보_모음.G이전_종가(12 * 21)
	_12개월_고점_종가 := 일일_가격정보_모음.G기간_고점_종가(0, 12*21)

	s = new(S팩터_세종)
	s.M종목명 = lib.F확인2(xing.F종목명by코드(종목코드))
	s.S재무_세종 = 최신_재무_정보
	s.M전년_재무_정보 = 전년_재무_정보
	s.M상장주식수량 = f상장주식수량(s.M종목코드, 기준일)
	s.M전일_거래량 = 일일_가격정보_모음.M저장소[len(일일_가격정보_모음.M저장소)-1].M거래량
	s.M기준가 = 일일_가격정보_모음.G종가() // 실제 매매에서는 '당일 시가'로 할 것.
	s.M시가총액 = s.M기준가 * s.M상장주식수량
	s.PSR = s.M시가총액 / s.M매출액
	s.POR = s.M시가총액 / s.M영업이익
	s.PER = s.M시가총액 / s.M당기순이익
	s.PBR = s.M시가총액 / s.M자본
	s.M매출액_성장율 = (s.M매출액 - s.M전년_재무_정보.M매출액) / s.M시가총액
	s.M영업이익_성장율 = (s.M영업이익 - s.M전년_재무_정보.M영업이익) / s.M시가총액
	s.M당기순이익_성장율 = (s.M당기순이익 - s.M전년_재무_정보.M당기순이익) / s.M시가총액
	s.PEG = s.PER / s.M당기순이익_성장율
	s.M자산회전율 = s.M매출액 / s.M자산
	s.M자본회전율 = s.M매출액 / s.M자본
	s.OPA = s.M영업이익 / s.M자산
	s.OPE = s.M영업이익 / s.M자본
	s.OPM = s.M영업이익 / s.M매출액
	s.ROA = s.M당기순이익 / s.M자산
	s.ROE = s.M당기순이익 / s.M자본
	s.NPM = s.M당기순이익 / s.M매출액
	s.M3개월_모멘텀 = (_1개월전_종가 - _3개월전_종가) / _3개월전_종가
	s.M6개월_모멘텀 = (_1개월전_종가 - _6개월전_종가) / _6개월전_종가
	s.M9개월_모멘텀 = (_1개월전_종가 - _9개월전_종가) / _9개월전_종가
	s.M12개월_모멘텀 = (_1개월전_종가 - _12개월전_종가) / _12개월전_종가
	s.M12개월_고점_대비_하락율 = (_12개월_고점_종가 - s.M기준가) / _12개월_고점_종가
	s.M부채_비율 = s.M부채 / s.M자본 * 100
	s.M부채_증가율 = (s.M부채 - s.M전년_재무_정보.M부채) / s.M시가총액
	s.M자본_증가율 = (s.M자본 - s.M전년_재무_정보.M자본) / s.M시가총액
	s.M자산_증가율 = (s.M자산 - s.M전년_재무_정보.M자산) / s.M시가총액
	s.M거래_금액_중간값_억 = 일일_가격정보_모음.G거래_금액_중간값(21)
	s.M가격_변동성 = 일일_가격정보_모음.G일_변동성(21)
	s.M거래금액_변동성 = 일일_가격정보_모음.G거래_금액_변동성(21)
	s.M최근_급등, s.M최근_급등일 = f최근_급등(일일_가격정보_모음)
	s.M최근_급락, s.M최근_급락일 = f최근_급락(일일_가격정보_모음)

	return s
}
