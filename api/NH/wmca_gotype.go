package NH

// #include "./wmca_type.h"
import "C"

import (
	공용 "github.com/ghts/ghts/common"

	"time"
	"unsafe"
)

//----------------------------------------------------------------------//
// WMCA_CONNECTED 로그인 구조체
//----------------------------------------------------------------------//
type S로그인_정보_블록 struct {
	TR구분번호  int
	M로그인_정보 S로그인_정보
}

func New로그인_정보_블록(c LoginBlock) S로그인_정보_블록 {
	return S로그인_정보_블록{
		TR구분번호:  int(c.TrIndex),
		M로그인_정보: New로그인_정보(*(c.LoginInfo))}
}

type S로그인_정보 struct {
	M접속시각  time.Time
	M접속서버  string
	M접속ID  string
	M계좌_목록 []S계좌_정보
}

func New로그인_정보(c LoginInfo) S로그인_정보 {
	시각, 에러 := time.Parse(공용.F2문자열(c.Date[:]), "값을 확인한 후 포맷 문자열 수정할 것.")
	공용.F에러_패닉(에러)

	계좌_수량, 에러 := 공용.F2정수(공용.F2문자열(c.AccountCount))
	공용.F에러_패닉(에러)

	계좌_목록 := make([]S계좌_정보, 계좌_수량)
	for i := 0; i < 계좌_수량; i++ {
		계좌_목록[i] = New계좌_정보(c.Accountlist[i])
	}

	return S로그인_정보{
		M접속시각:  시각,
		M접속서버:  공용.F2문자열(c.ServerName[:]),
		M접속ID:  공용.F2문자열(c.UserID[:]),
		M계좌_목록: 계좌_목록,
	}
}

type S계좌_정보 struct {
	M계좌_번호     string
	M계좌명       string
	M상품_코드     string
	M관리점_코드    string
	M위임_만기일    time.Time
	M일괄주문_허용계좌 bool // ('G': 허용)
	M주석        string
}

func New계좌_정보(c AccountInfo) S계좌_정보 {
	위임_만기일, 에러 := time.Parse(공용.F2문자열(c.ExpirationDate8[:]), "값을 확인한 후 포맷 문자열 수정할 것.")
	공용.F에러_패닉(에러)

	일괄주문_허용계좌 := false
	if 공용.F2문자열(c.Granted) == "G" {
		일괄주문_허용계좌 = true
	}

	return S계좌_정보{
		M계좌_번호:     공용.F2문자열(c.AccountNo[:]),
		M계좌명:       공용.F2문자열(c.AccountName[:]),
		M상품_코드:     공용.F2문자열(c.Act_pdt_cdz3[:]),
		M관리점_코드:    공용.F2문자열(c.Amn_tab_cdz4[:]),
		M위임_만기일:    위임_만기일,
		M일괄주문_허용계좌: 일괄주문_허용계좌,
		M주석:        공용.F2문자열(c.Filler[:]),
	}
}

//----------------------------------------------------------------------//
// WMCA 문자 message 구조체
//----------------------------------------------------------------------//
type S메시지 struct {
	M메시지_코드 string //00000:정상, 기타:비정상(코드값은 언제든지 변경될 수 있음.)
	M메시지_내용 string
}

func New메시지(c MsgHeader) S메시지 {
	return S메시지{
		M메시지_코드: 공용.F2문자열(c.MsgCode[:]),
		M메시지_내용: 공용.F2문자열(c.UsrMsg[:]),
	}
}

//----------------------------------------------------------------------//
// WMCA TR 응답 구조체
//----------------------------------------------------------------------//
type S수신_블록 struct {
	TR구분번호  int
	M수신_데이터 S수신_데이터
}

func New수신_블록(c OutDataBlock) S수신_블록 {
	return S수신_블록{
		TR구분번호:  int(c.TrIndex),
		M수신_데이터: New수신_데이터(*(c.DataStruct))}
}

type S수신_데이터 struct {
	M블록_이름 string
	M데이터   []byte
}

func New수신_데이터(c Received) S수신_데이터 {
	데이터 := C.GoBytes(unsafe.Pointer(c.DataString), C.int(c.Length))
	// 반대는 (*C.char)(unsafe.Pointer(&b[0]))

	return S수신_데이터{
		M블록_이름: 공용.F2문자열(c.BlockName),
		M데이터:   데이터}
}

//----------------------------------------------------------------------//
// 주식 현재가 조회 (c1101)
//----------------------------------------------------------------------//
func New주식_현재가_조회_질의(종목_코드 string) C.Tc1101InBlock {
	c := Tc1101InBlock{}
	c.Lang[0] = ([]byte("K"))[0]

	if len(종목_코드) > len(c.Code) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.Code[:]), len(종목_코드))

		panic(에러)
	}

	바이트_모음 := []byte(종목_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.Code[i] = 바이트_모음[i]
	}

	return *((*C.Tc1101InBlock)(unsafe.Pointer(&c)))
}

type S주식_현재가_조회_기본_자료 struct {
	M종목_코드             string
	M종목명               string // 첫자리는 KOSPI200은 ‘*’, 스타지수종목은 ‘#’
	M현재가               int64
	M등락부호              byte // 0x18 :상한, 0x1E :상승, 0x20 :보합, 0x19 :하한, 0x1F :하락
	M등락폭               int64
	M등락율               float64
	M매도_호가             int64
	M매수_호가             int64
	M거래량               int64
	M거래비율              float64
	M유동주_회전율           float64
	M거래대금              int64
	M상한가               int64
	M고가                int64
	M시가                int64
	M시가_대비_부호          byte
	M시가_대비_등락폭         int64
	M저가                int64
	M하한가               int64
	M시각                time.Time
	M매도_호가_최우선         int64
	M매도_호가_차선          int64
	M매도_호가_차차선         int64
	M매도_호가_4차선         int64
	M매도_호가_5차선         int64
	M매도_호가_6차선         int64
	M매도_호가_7차선         int64
	M매도_호가_8차선         int64
	M매도_호가_9차선         int64
	M매도_호가_10차선        int64
	M매수_호가_최우선         int64
	M매수_호가_차선          int64
	M매수_호가_차차선         int64
	M매수_호가_4차선         int64
	M매수_호가_5차선         int64
	M매수_호가_6차선         int64
	M매수_호가_7차선         int64
	M매수_호가_8차선         int64
	M매수_호가_9차선         int64
	M매수_호가_10차선        int64
	M매도_최우선_잔량         int64
	M매도_차선_잔량          int64
	M매도_차차선_잔량         int64
	M매도_4차선_잔량         int64
	M매도_5차선_잔량         int64
	M매도_6차선_잔량         int64
	M매도_7차선_잔량         int64
	M매도_8차선_잔량         int64
	M매도_9차선_잔량         int64
	M매도_10차선_잔량        int64
	M매수_최우선_잔량         int64
	M매수_차선_잔량          int64
	M매수_차차선_잔량         int64
	M매수_4차선_잔량         int64
	M매수_5차선_잔량         int64
	M매수_6차선_잔량         int64
	M매수_7차선_잔량         int64
	M매수_8차선_잔량         int64
	M매수_9차선_잔량         int64
	M매수_10차선_잔량        int64
	M매도_잔량_총합          int64
	M매수_잔량_총합          int64
	M시간외_매도_잔량         int64
	M시간외_매수_잔량         int64
	M피봇_2차_저항          int64  // 피봇가 + 전일 고가 – 전일 저가
	M피봇_1차_저항          int64  // (피봇가 * 2) – 전일 저가
	M피봇가               int64  // (전일 고가 + 전일 저가 + 전일 종가) / 3
	M피봇_1차_지지          int64  // (피봇가 * 2) – 전일 고가
	M피봇_2차_지지          int64  // 피봇가 – 전일고가 + 전일 저가
	M코스피_코스닥_구분        string // '코스피' , '코스닥'
	M업종명               string
	M자본금_규모            string
	M결산월               string
	M시장조치1             string
	M시장조치2             string
	M시장조치3             string
	M시장조치4             string
	M시장조치5             string
	M시장조치6             string
	M전환사채_구분           string
	M액면가               int64
	M전일종가_타이틀          string
	M전일종가              int64
	M대용가               int64 // 담보가치인 듯
	M공모가               int64
	M5일_고가             int64
	M5일_저가             int64
	M20일_고가            int64
	M20일_저가            int64
	M52주_고가            int64 // 거래일만 따지면 1년이 52주이었던 것으로 기억함.
	M52주_고가_일자         time.Time
	M52주_저가            int64
	M52주_저가_일자         time.Time
	M유동_주식수            int64
	M상장_주식_수량_1000주_단위 int64
	M시가_총액             int64
	M거래원_정보_수신_시간      time.Time
	M매도_거래원_1          string
	M매수_거래원_1          string
	M매도_거래량_1          int64
	M매수_거래량_1          int64
	M매도_거래원_2          string
	M매수_거래원_2          string
	M매도_거래량_2          int64
	M매수_거래량_2          int64
	M매도_거래원_3          string
	M매수_거래원_3          string
	M매도_거래량_3          int64
	M매수_거래량_3          int64
	M매도_거래원_4          string
	M매수_거래원_4          string
	M매도_거래량_4          int64
	M매수_거래량_4          int64
	M매도_거래원_5          string
	M매수_거래원_5          string
	M매도_거래량_5          int64
	M매수_거래량_5          int64
	M외국인_매도_거래량        int64
	M외국인_매수_거래량        int64
	M외국인_시간            time.Time
	M외국인_지분율           float64
	M결제일               time.Time
	M신용잔고_퍼센트          float64
	M유상_배정_기준일         time.Time
	M무상_배정_기준일         time.Time
	M유상_배정_비율          float64
	M외국인_변동주_수량        int64
	M무상_배정_비율          float64
	M당일_자사주_신청_여부      bool
	M상장일               time.Time
	M대주주_지분율           float64
	M대주주_지분율_정보_일자     time.Time
	M네잎클로버_종목_여부       bool
	M증거금_비율            float64
	M자본금               int64
	M전체_거래원_매도_합계      int64
	M전체_거래원_매수_합계      int64
	M종목명2              string
	M우회_상장_여부          bool
	M유동주_회전율_2         float64
	M코스피_구분_2          string    // 앞에 나온 '코스피/코스닥 구분'과 중복 아닌가?
	M공여율_기준일           time.Time // 공여율은 '신용거래 관련 비율'이라고 함.
	M공여율               float64   // 공여율(%)
	PER                float64   // PER
	M종목별_신용_한도         bool
	M가중_평균_가격          int64
	M상장_주식_수량          int64
	M추가_상장_주식_수량       int64
	M종목_코멘트            string
	M전일_거래량            int64
	M전일대비_등락부호         byte
	M전일대비_등락폭          int64
	M연중_최고가            int64     // 52주 최고가와 중복.
	M연중_최고가_일자         time.Time // 연중 최고가일
	M연중_최저가            int64
	M연중_최저가_일자         time.Time // 연중 최저가일
	M외국인_보유_주식수        int64     // 외국인 보유 주식수
	M외국인_지분_한도         float64   // % 단위
	M매매_수량_단위          int64     // 매매 수량 단위
	M대량_매매_방향          int8      // 0: 해당없음 1: 매도 2: 매수
	M대량_매매_존재          bool
}

func New주식_현재가_조회_기본_자료(c Tc1101OutBlock) S주식_현재가_조회_기본_자료 {
	return S주식_현재가_조회_기본_자료{
		M종목_코드:             공용.F2문자열(c.Code[:]),
		M종목명:               공용.F2문자열(c.Title[:]),
		M현재가:               공용.F2정수64_바이트(c.MarketPrice[:]),
		M등락부호:              c.DiffSign[0],
		M등락폭:               공용.F2정수64_바이트(c.Diff[:]),
		M등락율:               공용.F2실수_바이트(c.DiffRate[:]),
		M매도_호가:             공용.F2정수64_바이트(c.OfferPrice[:]),
		M매수_호가:             공용.F2정수64_바이트(c.BidPrice[:]),
		M거래량:               공용.F2정수64_바이트(c.Volume[:]),
		M거래비율:              공용.F2실수_바이트(c.TrVolRate[:]),
		M유동주_회전율:           공용.F2실수_바이트(c.FloatRate[:]),
		M거래대금:              공용.F2정수64_바이트(c.TrAmount[:]),
		M상한가:               공용.F2정수64_바이트(c.UpLmtPrice[:]),
		M고가:                공용.F2정수64_바이트(c.High[:]),
		M시가:                공용.F2정수64_바이트(c.Open[:]),
		M시가_대비_부호:          c.VsOpenSign[0],
		M시가_대비_등락폭:         공용.F2정수64_바이트(c.VsOpenDiff[:]),
		M저가:                공용.F2정수64_바이트(c.Low[:]),
		M하한가:               공용.F2정수64_바이트(c.LowLmtPrice[:]),
		M시각:                공용.F2시각_바이트(c.Time[:], "포맷 문자열은 내용 확인 후 수정할 것."),
		M매도_호가_최우선:         공용.F2정수64_바이트(c.OfferPrice1[:]),
		M매도_호가_차선:          공용.F2정수64_바이트(c.OfferPrice2[:]),
		M매도_호가_차차선:         공용.F2정수64_바이트(c.OfferPrice3[:]),
		M매도_호가_4차선:         공용.F2정수64_바이트(c.OfferPrice4[:]),
		M매도_호가_5차선:         공용.F2정수64_바이트(c.OfferPrice5[:]),
		M매도_호가_6차선:         공용.F2정수64_바이트(c.OfferPrice6[:]),
		M매도_호가_7차선:         공용.F2정수64_바이트(c.OfferPrice7[:]),
		M매도_호가_8차선:         공용.F2정수64_바이트(c.OfferPrice8[:]),
		M매도_호가_9차선:         공용.F2정수64_바이트(c.OfferPrice9[:]),
		M매도_호가_10차선:        공용.F2정수64_바이트(c.OfferPrice10[:]),
		M매수_호가_최우선:         공용.F2정수64_바이트(c.BidPrice1[:]),
		M매수_호가_차선:          공용.F2정수64_바이트(c.BidPrice2[:]),
		M매수_호가_차차선:         공용.F2정수64_바이트(c.BidPrice3[:]),
		M매수_호가_4차선:         공용.F2정수64_바이트(c.BidPrice4[:]),
		M매수_호가_5차선:         공용.F2정수64_바이트(c.BidPrice5[:]),
		M매수_호가_6차선:         공용.F2정수64_바이트(c.BidPrice6[:]),
		M매수_호가_7차선:         공용.F2정수64_바이트(c.BidPrice7[:]),
		M매수_호가_8차선:         공용.F2정수64_바이트(c.BidPrice8[:]),
		M매수_호가_9차선:         공용.F2정수64_바이트(c.BidPrice9[:]),
		M매수_호가_10차선:        공용.F2정수64_바이트(c.BidPrice10[:]),
		M매도_최우선_잔량:         공용.F2정수64_바이트(c.OfferVolume1[:]),
		M매도_차선_잔량:          공용.F2정수64_바이트(c.OfferVolume2[:]),
		M매도_차차선_잔량:         공용.F2정수64_바이트(c.OfferVolume3[:]),
		M매도_4차선_잔량:         공용.F2정수64_바이트(c.OfferVolume4[:]),
		M매도_5차선_잔량:         공용.F2정수64_바이트(c.OfferVolume5[:]),
		M매도_6차선_잔량:         공용.F2정수64_바이트(c.OfferVolume6[:]),
		M매도_7차선_잔량:         공용.F2정수64_바이트(c.OfferVolume7[:]),
		M매도_8차선_잔량:         공용.F2정수64_바이트(c.OfferVolume8[:]),
		M매도_9차선_잔량:         공용.F2정수64_바이트(c.OfferVolume9[:]),
		M매도_10차선_잔량:        공용.F2정수64_바이트(c.OfferVolume10[:]),
		M매수_최우선_잔량:         공용.F2정수64_바이트(c.BidVolume1[:]),
		M매수_차선_잔량:          공용.F2정수64_바이트(c.BidVolume2[:]),
		M매수_차차선_잔량:         공용.F2정수64_바이트(c.BidVolume3[:]),
		M매수_4차선_잔량:         공용.F2정수64_바이트(c.BidVolume4[:]),
		M매수_5차선_잔량:         공용.F2정수64_바이트(c.BidVolume5[:]),
		M매수_6차선_잔량:         공용.F2정수64_바이트(c.BidVolume6[:]),
		M매수_7차선_잔량:         공용.F2정수64_바이트(c.BidVolume7[:]),
		M매수_8차선_잔량:         공용.F2정수64_바이트(c.BidVolume8[:]),
		M매수_9차선_잔량:         공용.F2정수64_바이트(c.BidVolume9[:]),
		M매수_10차선_잔량:        공용.F2정수64_바이트(c.BidVolume10[:]),
		M매도_잔량_총합:          공용.F2정수64_바이트(c.OfferVolTot[:]),
		M매수_잔량_총합:          공용.F2정수64_바이트(c.BidVolTot[:]),
		M시간외_매도_잔량:         공용.F2정수64_바이트(c.OfferVolAfterHour[:]),
		M시간외_매수_잔량:         공용.F2정수64_바이트(c.BidVolAfterHour[:]),
		M피봇_2차_저항:          공용.F2정수64_바이트(c.PivotUp2[:]),
		M피봇_1차_저항:          공용.F2정수64_바이트(c.PivotUp1[:]),
		M피봇가:               공용.F2정수64_바이트(c.PivotPrice[:]),
		M피봇_1차_지지:          공용.F2정수64_바이트(c.PivotDown1[:]),
		M피봇_2차_지지:          공용.F2정수64_바이트(c.PivotDown2[:]),
		M코스피_코스닥_구분:        공용.F2문자열(c.Market[:]),
		M업종명:               공용.F2문자열(c.Sector[:]),
		M자본금_규모:            공용.F2문자열(c.CapSize[:]),
		M결산월:               공용.F2문자열(c.SettleMonth[:]),
		M시장조치1:             공용.F2문자열(c.MarketAction1[:]),
		M시장조치2:             공용.F2문자열(c.MarketAction2[:]),
		M시장조치3:             공용.F2문자열(c.MarketAction3[:]),
		M시장조치4:             공용.F2문자열(c.MarketAction4[:]),
		M시장조치5:             공용.F2문자열(c.MarketAction5[:]),
		M시장조치6:             공용.F2문자열(c.MarketAction6[:]),
		M전환사채_구분:           공용.F2문자열(c.ConvertBond[:]),
		M액면가:               공용.F2정수64_바이트(c.NominalPrice[:]),
		M전일종가_타이틀:          공용.F2문자열(c.PrevPriceTitle[:]),
		M전일종가:              공용.F2정수64_바이트(c.PrevPrice[:]),
		M대용가:               공용.F2정수64_바이트(c.MortgageValue[:]),
		M공모가:               공용.F2정수64_바이트(c.PublicOfferPrice[:]),
		M5일_고가:             공용.F2정수64_바이트(c.High5Day[:]),
		M5일_저가:             공용.F2정수64_바이트(c.Low5Day[:]),
		M20일_고가:            공용.F2정수64_바이트(c.High20Day[:]),
		M20일_저가:            공용.F2정수64_바이트(c.Low20Day[:]),
		M52주_고가:            공용.F2정수64_바이트(c.High1Year[:]),
		M52주_고가_일자:         공용.F2시각_바이트(c.High1YearDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M52주_저가:            공용.F2정수64_바이트(c.Low1Year[:]),
		M52주_저가_일자:         공용.F2시각_바이트(c.Low1YearDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M유동_주식수:            공용.F2정수64_바이트(c.FloatVolume[:]),
		M상장_주식_수량_1000주_단위: 공용.F2정수64_바이트(c.ListVolBy1000[:]),
		M시가_총액:             공용.F2정수64_바이트(c.MarketCapital[:]),
		M거래원_정보_수신_시간:      공용.F2시각_바이트(c.TraderInfoTime[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M매도_거래원_1:          공용.F2문자열(c.Seller1[:]),
		M매수_거래원_1:          공용.F2문자열(c.Buyer1[:]),
		M매도_거래량_1:          공용.F2정수64_바이트(c.Seller1Volume[:]),
		M매수_거래량_1:          공용.F2정수64_바이트(c.Buyer1Volume[:]),
		M매도_거래원_2:          공용.F2문자열(c.Seller2[:]),
		M매수_거래원_2:          공용.F2문자열(c.Buyer2[:]),
		M매도_거래량_2:          공용.F2정수64_바이트(c.Seller2Volume[:]),
		M매수_거래량_2:          공용.F2정수64_바이트(c.Buyer2Volume[:]),
		M매도_거래원_3:          공용.F2문자열(c.Seller3[:]),
		M매수_거래원_3:          공용.F2문자열(c.Buyer3[:]),
		M매도_거래량_3:          공용.F2정수64_바이트(c.Seller3Volume[:]),
		M매수_거래량_3:          공용.F2정수64_바이트(c.Buyer3Volume[:]),
		M매도_거래원_4:          공용.F2문자열(c.Seller4[:]),
		M매수_거래원_4:          공용.F2문자열(c.Buyer4[:]),
		M매도_거래량_4:          공용.F2정수64_바이트(c.Seller4Volume[:]),
		M매수_거래량_4:          공용.F2정수64_바이트(c.Buyer4Volume[:]),
		M매도_거래원_5:          공용.F2문자열(c.Seller5[:]),
		M매수_거래원_5:          공용.F2문자열(c.Buyer5[:]),
		M매도_거래량_5:          공용.F2정수64_바이트(c.Seller5Volume[:]),
		M매수_거래량_5:          공용.F2정수64_바이트(c.Buyer5Volume[:]),
		M외국인_매도_거래량:        공용.F2정수64_바이트(c.ForeignSellVolume[:]),
		M외국인_매수_거래량:        공용.F2정수64_바이트(c.ForeignBuyVolume[:]),
		M외국인_시간:            공용.F2시각_바이트(c.ForeignTime[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M외국인_지분율:           공용.F2실수_바이트(c.ForeignHoldingRate[:]),
		M결제일:               공용.F2시각_바이트(c.SettleDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M신용잔고_퍼센트:          공용.F2실수_바이트(c.DebtPercent[:]),
		M유상_배정_기준일:         공용.F2시각_바이트(c.RightsIssueDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M무상_배정_기준일:         공용.F2시각_바이트(c.BonusIssueDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M유상_배정_비율:          공용.F2실수_바이트(c.RightsIssueRate[:]),
		M무상_배정_비율:          공용.F2실수_바이트(c.BonusIssueRate[:]),
		M외국인_변동주_수량:        공용.F2정수64_바이트(c.ForeignFloatVol[:]),
		M당일_자사주_신청_여부:      공용.F2참거짓_바이트(c.TreasuryStock[:], "1", true),
		M상장일:               공용.F2시각_바이트(c.IpoDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M대주주_지분율:           공용.F2실수_바이트(c.MajorHoldRate[:]),
		M대주주_지분율_정보_일자:     공용.F2시각_바이트(c.MajorHoldInfoDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M네잎클로버_종목_여부:       공용.F2참거짓_바이트(c.FourLeafClover[:], "1", true),
		M증거금_비율:            공용.F2실수_바이트(c.MarginRate[:]),
		M자본금:               공용.F2정수64_바이트(c.Capital[:]),
		M전체_거래원_매도_합계:      공용.F2정수64_바이트(c.SellTotalSum[:]),
		M전체_거래원_매수_합계:      공용.F2정수64_바이트(c.BuyTotalSum[:]),
		M종목명2:              공용.F2문자열(c.Title2[:]),
		M우회_상장_여부:          공용.F2참거짓_바이트(c.BackdoorListing[:], "1", true),
		M유동주_회전율_2:         공용.F2실수_바이트(c.FloatRate2[:]),
		M코스피_구분_2:          공용.F2문자열(c.Market2[:]),
		M공여율_기준일:           공용.F2시각_바이트(c.DebtTrDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M공여율:               공용.F2실수_바이트(c.DebtTrPercent[:]),
		PER:                공용.F2실수_바이트(c.PER[:]),
		M종목별_신용_한도:   공용.F2참거짓_바이트(c.DebtLimit[:], "1", true),
		M가중_평균_가격:    공용.F2정수64_바이트(c.WeightAvgPrice[:]),
		M상장_주식_수량:    공용.F2정수64_바이트(c.ListedVolume[:]),
		M추가_상장_주식_수량: 공용.F2정수64_바이트(c.AddListing[:]),
		M종목_코멘트:      공용.F2문자열(c.Comment[:]),
		M전일_거래량:      공용.F2정수64_바이트(c.PrevVolume[:]),
		M전일대비_등락부호:   c.VsPrevSign[0],
		M전일대비_등락폭:    공용.F2정수64_바이트(c.VsPrevDiff[:]),
		M연중_최고가:      공용.F2정수64_바이트(c.High1Year2[:]),
		M연중_최고가_일자:   공용.F2시각_바이트(c.High1YearDate2[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M연중_최저가:      공용.F2정수64_바이트(c.Low1Year2[:]),
		M연중_최저가_일자:   공용.F2시각_바이트(c.Low1YearDate2[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M외국인_보유_주식수:  공용.F2정수64_바이트(c.ForeignHoldQty[:]),
		M외국인_지분_한도:   공용.F2실수_바이트(c.ForeignLmtPercent[:]),
		M매매_수량_단위:    공용.F2정수64_바이트(c.TrUnitVolume[:]),
		M대량_매매_방향:    int8(공용.F2정수64_바이트(c.DarkPoolOfferBid[:])), // 0: 해당없음 1: 매도 2: 매수
		M대량_매매_존재:    공용.F2참거짓_바이트(c.DarkPoolExist[:], "1", true)}
}

type S주식_현재가_조회_변동_거래량_자료 struct { // 변동거래량자료[반복]
	M시간     time.Time
	M현재가    int64
	M등락부호   byte
	M등락폭    int64
	M매도_호가  int64
	M매수_호가  int64
	M변동_거래량 int64
	M거래량    int64
}

func New주식_현재가_조회_변동_거래량_자료(c Tc1101OutBlock2) S주식_현재가_조회_변동_거래량_자료 {
	return S주식_현재가_조회_변동_거래량_자료{
		M시간:     공용.F2시각_바이트(c.Time[:], "내용 확인 후 포맷 문자열 작성할 것."),
		M현재가:    공용.F2정수64_바이트(c.MarketPrice[:]),
		M등락부호:   c.DiffSign[0],
		M등락폭:    공용.F2정수64_바이트(c.Diff[:]),
		M매도_호가:  공용.F2정수64_바이트(c.OfferPrice[:]),
		M매수_호가:  공용.F2정수64_바이트(c.BidPrice[:]),
		M변동_거래량: 공용.F2정수64_바이트(c.DiffVolume[:]),
		M거래량:    공용.F2정수64_바이트(c.Volume[:])}
}

type S주식_현재가_조회_종목_지표 struct { // 종목지표
	M동시_호가_구분       string // 0:동시호가 아님 1:동시호가 2:동시호가연장 3:시가범위연장 4:종가범위연장 5:배분개시 6:변동성 완화장치 발동
	M예상_체결가         int64
	M예상_체결부호        byte
	M예상_등락폭         int64
	M예상_등락율         float64
	M예상_체결수량        int64
	ECN정보_유무        bool // 우리나라에는 아직 ECN이 없는 것으로 알고 있음.
	ECN전일_종가        int64
	ECN등락_부호        byte
	ECN등락폭          int64
	ECN등락율          float64
	ECN체결_수량        int64
	ECN대비_예상_체결_부호  byte
	ECN대비_예상_체결_등락폭 int64
	ECN대비_예상_체결_등락율 float64
}

func New주식_현재가_조회_종목_지표(c Tc1101OutBlock3) S주식_현재가_조회_종목_지표 {
	return S주식_현재가_조회_종목_지표{
		M동시_호가_구분:       공용.F2문자열(c.SyncOfferBid[:]), // 0:동시호가 아님 1:동시호가 2:동시호가연장 3:시가범위연장 4:종가범위연장 5:배분개시 6:변동성 완화장치 발동
		M예상_체결가:         공용.F2정수64_바이트(c.EstmPrice[:]),
		M예상_체결부호:        c.EstmSign[0],
		M예상_등락폭:         공용.F2정수64_바이트(c.EstmDiff[:]),
		M예상_등락율:         공용.F2실수_바이트(c.EstmDiffRate[:]),
		M예상_체결수량:        공용.F2정수64_바이트(c.EstmVol[:]),
		ECN정보_유무:        공용.F2참거짓_바이트(c.ECN_InfoExist[:], "1", true), // 우리나라에는 아직 ECN이 없는 것으로 알고 있음.
		ECN전일_종가:        공용.F2정수64_바이트(c.ECN_PrevPrice[:]),
		ECN등락_부호:        c.ECN_DiffSign[0],
		ECN등락폭:          공용.F2정수64_바이트(c.ECN_Diff[:]),
		ECN등락율:          공용.F2실수_바이트(c.ECN_DiffRate[:]),
		ECN체결_수량:        공용.F2정수64_바이트(c.ECN_Volume[:]),
		ECN대비_예상_체결_부호:  c.VsECN_EstmSign[0],
		ECN대비_예상_체결_등락폭: 공용.F2정수64_바이트(c.VsECN_EstmDiff[:]),
		ECN대비_예상_체결_등락율: 공용.F2실수_바이트(c.VsECN_EstmDiffRate[:])}
}

//----------------------------------------------------------------------//
// ETF 현재가 조회 (c1151)
//----------------------------------------------------------------------//
func New_ETF_현재가_조회_질의(종목_코드 string) C.Tc1151InBlock {
	c := Tc1151InBlock{}
	c.Lang[0] = ([]byte("K"))[0]

	if len(종목_코드) > len(c.Code) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.Code[:]), len(종목_코드))

		panic(에러)
	}

	바이트_모음 := []byte(종목_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.Code[i] = 바이트_모음[i]
	}

	return *((*C.Tc1151InBlock)(unsafe.Pointer(&c)))
}

type S_ETF_현재가_조회_기본_자료 struct { // 종목마스타기본자료
	M종목_코드             string
	M종목명               string
	M현재가               int64
	M등락부호              byte
	M등락폭               int64
	M등락율               float64
	M매도_호가             int64
	M매수_호가             int64
	M거래량               int64
	M거래_비율             float64
	M유동주_회전율           float64
	M거래_대금             int64
	M상한가               int64
	M고가                int64
	M시가                int64
	M시가_대비_부호          byte
	M시가_대비_등락폭         int64
	M저가                int64
	M하한가               int64
	M시각                time.Time
	M매도_호가_최우선         int64
	M매도_호가_차선          int64
	M매도_호가_차차선         int64
	M매도_호가_4차선         int64
	M매도_호가_5차선         int64
	M매도_호가_6차선         int64
	M매도_호가_7차선         int64
	M매도_호가_8차선         int64
	M매도_호가_9차선         int64
	M매도_호가_10차선        int64
	M매수_호가_최우선         int64
	M매수_호가_차선          int64
	M매수_호가_차차선         int64
	M매수_호가_4차선         int64
	M매수_호가_5차선         int64
	M매수_호가_6차선         int64
	M매수_호가_7차선         int64
	M매수_호가_8차선         int64
	M매수_호가_9차선         int64
	M매수_호가_10차선        int64
	M매도_잔량_최우선         int64
	M매도_잔량_차선          int64
	M매도_잔량_차차선         int64
	M매도_잔량_4차선         int64
	M매도_잔량_5차선         int64
	M매도_잔량_6차선         int64
	M매도_잔량_7차선         int64
	M매도_잔량_8차선         int64
	M매도_잔량_9차선         int64
	M매도_잔량_10차선        int64
	M매수_잔량_최우선         int64
	M매수_잔량_차선          int64
	M매수_잔량_차차선         int64
	M매수_잔량_4차선         int64
	M매수_잔량_5차선         int64
	M매수_잔량_6차선         int64
	M매수_잔량_7차선         int64
	M매수_잔량_8차선         int64
	M매수_잔량_9차선         int64
	M매수_잔량_10차선        int64
	M매도_잔량_총합          int64
	M매수_잔량_총합          int64
	M시간외_매도_잔량         int64
	M시간외_매수_잔량         int64
	M피봇_2차_저항          int64
	M피봇_1차_저항          int64
	M피봇_가격             int64
	M피봇_1차_지지          int64
	M피봇_2차_지지          int64
	M코스피_코스닥_구분        string
	M업종명               string
	M자본금_규모            string
	M결산월               string
	M시장_조치_1           string
	M시장_조치_2           string
	M시장_조치_3           string
	M시장_조치_4           string
	M시장_조치_5           string
	M시장_조치_6           string
	M전환사채_구분           string
	M액면가               int64
	M전일_종가_타이틀         string // GUI화면에 쓸 자료이니 별 필요 없을 듯...
	M전일_종가             int64
	M대용가               int64
	M공모가               int64
	M5일_고가             int64
	M5일_저가             int64
	M20일_고가            int64
	M20일_저가            int64
	M52주_고가            int64
	M52주_고가_일자         time.Time
	M52주_저가            int64
	M52주_저가_일자         time.Time
	M유동_주식수            int64
	M상장_주식_수량_1000주_단위 int64
	M시가_총액             int64
	M거래원_정보_수신_시각      time.Time
	M매도_거래원_1          string
	M매수_거래원_1          string
	M매도_거래원_1_거래량      int64
	M매수_거래원_1_거래량      int64
	M매도_거래원_2          string
	M매수_거래원_2          string
	M매도_거래원_2_거래량      int64
	M매수_거래원_2_거래량      int64
	M매도_거래원_3          string
	M매수_거래원_3          string
	M매도_거래원_3_거래량      int64
	M매수_거래원_3_거래량      int64
	M매도_거래원_4          string
	M매수_거래원_4          string
	M매도_거래원_4_거래량      int64
	M매수_거래원_4_거래량      int64
	M매도_거래원_5          string
	M매수_거래원_5          string
	M매도_거래원_5_거래량      int64
	M매수_거래원_5_거래량      int64
	M외국인_매도_거래량        int64
	M외국인_매수_거래량        int64
	M외국인_시간            time.Time
	M외국인_지분율           float64
	M결제일               time.Time
	M신용잔고_퍼센트          float64
	M유상_배정_기준일         time.Time
	M무상_배정_기준일         time.Time
	M유상_배정_비율          float64
	M무상_배정_비율          float64
	M상장일               time.Time
	M상장_주식_수량          int64
	M전체_거래원_매도_합계      int64
	M전체_거래원_매수_합계      int64
}

func New_ETF_현재가_조회_기본_자료(c Tc1151OutBlock) S_ETF_현재가_조회_기본_자료 {
	return S_ETF_현재가_조회_기본_자료{
		M종목_코드:             공용.F2문자열(c.Code[:]),
		M종목명:               공용.F2문자열(c.Title[:]),
		M현재가:               공용.F2정수64_바이트(c.MarketPrice[:]),
		M등락부호:              c.DiffSign[0],
		M등락폭:               공용.F2정수64_바이트(c.Diff[:]),
		M등락율:               공용.F2실수_바이트(c.DiffRate[:]),
		M매도_호가:             공용.F2정수64_바이트(c.OfferPrice[:]),
		M매수_호가:             공용.F2정수64_바이트(c.BidPrice[:]),
		M거래량:               공용.F2정수64_바이트(c.Volume[:]),
		M거래_비율:             공용.F2실수_바이트(c.TrVolRate[:]),
		M유동주_회전율:           공용.F2실수_바이트(c.FloatVolRate[:]),
		M거래_대금:             공용.F2정수64_바이트(c.TrAmount[:]),
		M상한가:               공용.F2정수64_바이트(c.UpLmtPrice[:]),
		M고가:                공용.F2정수64_바이트(c.High[:]),
		M시가:                공용.F2정수64_바이트(c.Open[:]),
		M시가_대비_부호:          c.VsOpenSign[0],
		M시가_대비_등락폭:         공용.F2정수64_바이트(c.VsOpenDiff[:]),
		M저가:                공용.F2정수64_바이트(c.Low[:]),
		M하한가:               공용.F2정수64_바이트(c.LowLmtPrice[:]),
		M시각:                공용.F2시각_바이트(c.Time[:], "포맷 문자열은 내용 확인 후 작성할 것."),
		M매도_호가_최우선:         공용.F2정수64_바이트(c.OfferPrice1[:]),
		M매도_호가_차선:          공용.F2정수64_바이트(c.OfferPrice2[:]),
		M매도_호가_차차선:         공용.F2정수64_바이트(c.OfferPrice3[:]),
		M매도_호가_4차선:         공용.F2정수64_바이트(c.OfferPrice4[:]),
		M매도_호가_5차선:         공용.F2정수64_바이트(c.OfferPrice5[:]),
		M매도_호가_6차선:         공용.F2정수64_바이트(c.OfferPrice6[:]),
		M매도_호가_7차선:         공용.F2정수64_바이트(c.OfferPrice7[:]),
		M매도_호가_8차선:         공용.F2정수64_바이트(c.OfferPrice8[:]),
		M매도_호가_9차선:         공용.F2정수64_바이트(c.OfferPrice9[:]),
		M매도_호가_10차선:        공용.F2정수64_바이트(c.OfferPrice10[:]),
		M매수_호가_최우선:         공용.F2정수64_바이트(c.BidPrice1[:]),
		M매수_호가_차선:          공용.F2정수64_바이트(c.BidPrice2[:]),
		M매수_호가_차차선:         공용.F2정수64_바이트(c.BidPrice3[:]),
		M매수_호가_4차선:         공용.F2정수64_바이트(c.BidPrice4[:]),
		M매수_호가_5차선:         공용.F2정수64_바이트(c.BidPrice5[:]),
		M매수_호가_6차선:         공용.F2정수64_바이트(c.BidPrice6[:]),
		M매수_호가_7차선:         공용.F2정수64_바이트(c.BidPrice7[:]),
		M매수_호가_8차선:         공용.F2정수64_바이트(c.BidPrice8[:]),
		M매수_호가_9차선:         공용.F2정수64_바이트(c.BidPrice9[:]),
		M매수_호가_10차선:        공용.F2정수64_바이트(c.BidPrice10[:]),
		M매도_잔량_최우선:         공용.F2정수64_바이트(c.OfferVolume1[:]),
		M매도_잔량_차선:          공용.F2정수64_바이트(c.OfferVolume2[:]),
		M매도_잔량_차차선:         공용.F2정수64_바이트(c.OfferVolume3[:]),
		M매도_잔량_4차선:         공용.F2정수64_바이트(c.OfferVolume4[:]),
		M매도_잔량_5차선:         공용.F2정수64_바이트(c.OfferVolume5[:]),
		M매도_잔량_6차선:         공용.F2정수64_바이트(c.OfferVolume6[:]),
		M매도_잔량_7차선:         공용.F2정수64_바이트(c.OfferVolume7[:]),
		M매도_잔량_8차선:         공용.F2정수64_바이트(c.OfferVolume8[:]),
		M매도_잔량_9차선:         공용.F2정수64_바이트(c.OfferVolume9[:]),
		M매도_잔량_10차선:        공용.F2정수64_바이트(c.OfferVolume10[:]),
		M매수_잔량_최우선:         공용.F2정수64_바이트(c.BidVolume1[:]),
		M매수_잔량_차선:          공용.F2정수64_바이트(c.BidVolume2[:]),
		M매수_잔량_차차선:         공용.F2정수64_바이트(c.BidVolume3[:]),
		M매수_잔량_4차선:         공용.F2정수64_바이트(c.BidVolume4[:]),
		M매수_잔량_5차선:         공용.F2정수64_바이트(c.BidVolume5[:]),
		M매수_잔량_6차선:         공용.F2정수64_바이트(c.BidVolume6[:]),
		M매수_잔량_7차선:         공용.F2정수64_바이트(c.BidVolume7[:]),
		M매수_잔량_8차선:         공용.F2정수64_바이트(c.BidVolume8[:]),
		M매수_잔량_9차선:         공용.F2정수64_바이트(c.BidVolume9[:]),
		M매수_잔량_10차선:        공용.F2정수64_바이트(c.BidVolume10[:]),
		M매도_잔량_총합:          공용.F2정수64_바이트(c.OfferVolTot[:]),
		M매수_잔량_총합:          공용.F2정수64_바이트(c.BidVolTot[:]),
		M시간외_매도_잔량:         공용.F2정수64_바이트(c.OfferVolAfterHour[:]),
		M시간외_매수_잔량:         공용.F2정수64_바이트(c.BidVolAfterHour[:]),
		M피봇_2차_저항:          공용.F2정수64_바이트(c.PivotUp2[:]),
		M피봇_1차_저항:          공용.F2정수64_바이트(c.PivotUp1[:]),
		M피봇_가격:             공용.F2정수64_바이트(c.PivotPrice[:]),
		M피봇_1차_지지:          공용.F2정수64_바이트(c.PivotDown1[:]),
		M피봇_2차_지지:          공용.F2정수64_바이트(c.PivotDown2[:]),
		M코스피_코스닥_구분:        공용.F2문자열(c.Market[:]),
		M업종명:               공용.F2문자열(c.Sector[:]),
		M자본금_규모:            공용.F2문자열(c.CapSize[:]),
		M결산월:               공용.F2문자열(c.SettleMonth[:]),
		M시장_조치_1:           공용.F2문자열(c.MarketAction1[:]),
		M시장_조치_2:           공용.F2문자열(c.MarketAction2[:]),
		M시장_조치_3:           공용.F2문자열(c.MarketAction3[:]),
		M시장_조치_4:           공용.F2문자열(c.MarketAction4[:]),
		M시장_조치_5:           공용.F2문자열(c.MarketAction5[:]),
		M시장_조치_6:           공용.F2문자열(c.MarketAction6[:]),
		M전환사채_구분:           공용.F2문자열(c.ConvertBond[:]),
		M액면가:               공용.F2정수64_바이트(c.NominalPrice[:]),
		M전일_종가_타이틀:         공용.F2문자열(c.PrevPriceTitle[:]),
		M전일_종가:             공용.F2정수64_바이트(c.PrevPrice[:]),
		M대용가:               공용.F2정수64_바이트(c.MortgageValue[:]),
		M공모가:               공용.F2정수64_바이트(c.PublicOfferPrice[:]),
		M5일_고가:             공용.F2정수64_바이트(c.High5Day[:]),
		M5일_저가:             공용.F2정수64_바이트(c.Low5Day[:]),
		M20일_고가:            공용.F2정수64_바이트(c.High20Day[:]),
		M20일_저가:            공용.F2정수64_바이트(c.Low20Day[:]),
		M52주_고가:            공용.F2정수64_바이트(c.High1Year[:]),
		M52주_고가_일자:         공용.F2시각_바이트(c.High1YearDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M52주_저가:            공용.F2정수64_바이트(c.Low1Year[:]),
		M52주_저가_일자:         공용.F2시각_바이트(c.Low1YearDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M유동_주식수:            공용.F2정수64_바이트(c.FloatVolume[:]),
		M상장_주식_수량_1000주_단위: 공용.F2정수64_바이트(c.ListVolBy1000[:]),
		M시가_총액:             공용.F2정수64_바이트(c.MarketCapital[:]),
		M거래원_정보_수신_시각:      공용.F2시각_바이트(c.TraderInfoTime[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M매도_거래원_1:          공용.F2문자열(c.Seller1[:]),
		M매수_거래원_1:          공용.F2문자열(c.Buyer1[:]),
		M매도_거래원_1_거래량:      공용.F2정수64_바이트(c.Seller1Volume[:]),
		M매수_거래원_1_거래량:      공용.F2정수64_바이트(c.Buyer1Volume[:]),
		M매도_거래원_2:          공용.F2문자열(c.Seller2[:]),
		M매수_거래원_2:          공용.F2문자열(c.Buyer2[:]),
		M매도_거래원_2_거래량:      공용.F2정수64_바이트(c.Seller2Volume[:]),
		M매수_거래원_2_거래량:      공용.F2정수64_바이트(c.Buyer2Volume[:]),
		M매도_거래원_3:          공용.F2문자열(c.Seller3[:]),
		M매수_거래원_3:          공용.F2문자열(c.Buyer3[:]),
		M매도_거래원_3_거래량:      공용.F2정수64_바이트(c.Seller3Volume[:]),
		M매수_거래원_3_거래량:      공용.F2정수64_바이트(c.Buyer3Volume[:]),
		M매도_거래원_4:          공용.F2문자열(c.Seller4[:]),
		M매수_거래원_4:          공용.F2문자열(c.Buyer4[:]),
		M매도_거래원_4_거래량:      공용.F2정수64_바이트(c.Seller4Volume[:]),
		M매수_거래원_4_거래량:      공용.F2정수64_바이트(c.Buyer4Volume[:]),
		M매도_거래원_5:          공용.F2문자열(c.Seller5[:]),
		M매수_거래원_5:          공용.F2문자열(c.Buyer5[:]),
		M매도_거래원_5_거래량:      공용.F2정수64_바이트(c.Seller5Volume[:]),
		M매수_거래원_5_거래량:      공용.F2정수64_바이트(c.Buyer5Volume[:]),
		M외국인_매도_거래량:        공용.F2정수64_바이트(c.ForeignSellVolume[:]),
		M외국인_매수_거래량:        공용.F2정수64_바이트(c.ForeignBuyVolume[:]),
		M외국인_시간:            공용.F2시각_바이트(c.ForeignTime[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M외국인_지분율:           공용.F2실수_바이트(c.ForeignHoldingRate[:]),
		M결제일:               공용.F2시각_바이트(c.SettleDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M신용잔고_퍼센트:          공용.F2실수_바이트(c.DebtPercent[:]),
		M유상_배정_기준일:         공용.F2시각_바이트(c.RightsIssueDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M무상_배정_기준일:         공용.F2시각_바이트(c.BonusIssueDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M유상_배정_비율:          공용.F2실수_바이트(c.RightsIssueRate[:]),
		M무상_배정_비율:          공용.F2실수_바이트(c.BonusIssueRate[:]),
		M상장일:               공용.F2시각_바이트(c.IpoDate[:], "포맷 문자열은 내용 확인 후 결정할 것."),
		M상장_주식_수량:          공용.F2정수64_바이트(c.ListedVolume[:]),
		M전체_거래원_매도_합계:      공용.F2정수64_바이트(c.SellTotalSum[:]),
		M전체_거래원_매수_합계:      공용.F2정수64_바이트(c.BuyTotalSum[:])}
}

type S_ETF_현재가_조회_변동_거래량 struct {
	M시간     time.Time
	M현재가    int64
	M등락_부호  byte
	M등락폭    int64
	M매도_호가  int64
	M매수_호가  int64
	M변동_거래량 int64
	M거래량    int64
}

func New_ETF_현재가_조회_변동_거래량(c Tc1151OutBlock2) S_ETF_현재가_조회_변동_거래량 {
	return S_ETF_현재가_조회_변동_거래량{
		M시간:     공용.F2시각_바이트(c.Time[:], "포맷 문자열은 내용 확인 후 작성할 것."),
		M현재가:    공용.F2정수64_바이트(c.MarketPrice[:]),
		M등락_부호:  c.DiffSign[0],
		M등락폭:    공용.F2정수64_바이트(c.Diff[:]),
		M매도_호가:  공용.F2정수64_바이트(c.OfferPrice[:]),
		M매수_호가:  공용.F2정수64_바이트(c.BidPrice[:]),
		M변동_거래량: 공용.F2정수64_바이트(c.DiffVolume[:]),
		M거래량:    공용.F2정수64_바이트(c.Volume[:])}
}

type S_ETF_현재가_조회_예상_체결 struct {
	M동시_호가_구분  string
	M예상_체결가    int64
	M예상_체결_부호  byte
	M예상_체결_등락폭 int64
	M예상_체결_등락율 float64
	M예상_체결_수량  int64
}

func New_ETF_현재가_조회_예상_체결(c Tc1151OutBlock3) S_ETF_현재가_조회_예상_체결 {
	return S_ETF_현재가_조회_예상_체결{
		M동시_호가_구분:  공용.F2문자열(c.SyncOfferBid[:]),
		M예상_체결가:    공용.F2정수64_바이트(c.EstmPrice[:]),
		M예상_체결_부호:  c.EstmSign[0],
		M예상_체결_등락폭: 공용.F2정수64_바이트(c.EstmDiff[:]),
		M예상_체결_등락율: 공용.F2실수_바이트(c.EstmDiffRate[:]),
		M예상_체결_수량:  공용.F2정수64_바이트(c.EstmVolume[:])}
}

type S_ETF_현재가_조회_ETF자료 struct {
	ETF구분           string
	NAV             float64
	NAV_등락_부호       byte
	NAV_등락폭         float64
	M전일NAV          float64
	M괴리율            float64
	M괴리율_부호         byte
	M설정단위_당_현금_배당액  int64
	M구성_종목수         int64
	M순자산_총액_억원      int64
	M추적_오차율         float64
	LP_매도_최우선_잔량    int64
	LP_매도_차선_잔량     int64
	LP_매도_차차선_잔량    int64
	LP_매도_4차선_잔량    int64
	LP_매도_5차선_잔량    int64
	LP_매도_6차선_잔량    int64
	LP_매도_7차선_잔량    int64
	LP_매도_8차선_잔량    int64
	LP_매도_9차선_잔량    int64
	LP_매도_10차선_잔량   int64
	LP_매수_최우선_잔량    int64
	LP_매수_차선_잔량     int64
	LP_매수_차차선_잔량    int64
	LP_매수_4차선_잔량    int64
	LP_매수_5차선_잔량    int64
	LP_매수_6차선_잔량    int64
	LP_매수_7차선_잔량    int64
	LP_매수_8차선_잔량    int64
	LP_매수_9차선_잔량    int64
	LP_매수_10차선_잔량   int64
	ETF_복제_방법_구분_코드 string
	ETF_상품_유형_코드    string
}

func New_ETF_현재가_조회_ETF자료(c Tc1151OutBlock4) S_ETF_현재가_조회_ETF자료 {
	return S_ETF_현재가_조회_ETF자료{
		ETF구분:           공용.F2문자열(c.ETF[:]),
		NAV:             공용.F2실수_바이트(c.NAV[:]),
		NAV_등락_부호:       c.DiffSign[0],
		NAV_등락폭:         공용.F2실수_바이트(c.Diff[:]),
		M전일NAV:          공용.F2실수_바이트(c.PrevNAV[:]),
		M괴리율:            공용.F2실수_바이트(c.DivergeRate[:]),
		M괴리율_부호:         c.DivergeSign[0],
		M설정단위_당_현금_배당액:  공용.F2정수64_바이트(c.DividendPerCU[:]),
		M구성_종목수:         공용.F2정수64_바이트(c.ConstituentNo[:]),
		M순자산_총액_억원:      공용.F2정수64_바이트(c.NAVBy100Million[:]),
		M추적_오차율:         공용.F2실수_바이트(c.TrackingErrRate[:]),
		LP_매도_최우선_잔량:    공용.F2정수64_바이트(c.LP_OfferVolume1[:]),
		LP_매도_차선_잔량:     공용.F2정수64_바이트(c.LP_OfferVolume2[:]),
		LP_매도_차차선_잔량:    공용.F2정수64_바이트(c.LP_OfferVolume3[:]),
		LP_매도_4차선_잔량:    공용.F2정수64_바이트(c.LP_OfferVolume4[:]),
		LP_매도_5차선_잔량:    공용.F2정수64_바이트(c.LP_OfferVolume5[:]),
		LP_매도_6차선_잔량:    공용.F2정수64_바이트(c.LP_OfferVolume6[:]),
		LP_매도_7차선_잔량:    공용.F2정수64_바이트(c.LP_OfferVolume7[:]),
		LP_매도_8차선_잔량:    공용.F2정수64_바이트(c.LP_OfferVolume8[:]),
		LP_매도_9차선_잔량:    공용.F2정수64_바이트(c.LP_OfferVolume9[:]),
		LP_매도_10차선_잔량:   공용.F2정수64_바이트(c.LP_OfferVolume10[:]),
		LP_매수_최우선_잔량:    공용.F2정수64_바이트(c.LP_BidVolume1[:]),
		LP_매수_차선_잔량:     공용.F2정수64_바이트(c.LP_BidVolume2[:]),
		LP_매수_차차선_잔량:    공용.F2정수64_바이트(c.LP_BidVolume3[:]),
		LP_매수_4차선_잔량:    공용.F2정수64_바이트(c.LP_BidVolume4[:]),
		LP_매수_5차선_잔량:    공용.F2정수64_바이트(c.LP_BidVolume5[:]),
		LP_매수_6차선_잔량:    공용.F2정수64_바이트(c.LP_BidVolume6[:]),
		LP_매수_7차선_잔량:    공용.F2정수64_바이트(c.LP_BidVolume7[:]),
		LP_매수_8차선_잔량:    공용.F2정수64_바이트(c.LP_BidVolume8[:]),
		LP_매수_9차선_잔량:    공용.F2정수64_바이트(c.LP_BidVolume9[:]),
		LP_매수_10차선_잔량:   공용.F2정수64_바이트(c.LP_BidVolume10[:]),
		ETF_복제_방법_구분_코드: 공용.F2문자열(c.TrackingMethod[:]),
		ETF_상품_유형_코드:    공용.F2문자열(c.ETF_Type[:])}
}

type S_ETF_현재가_조회_기반_지수_자료 struct {
	M지수_코드       string
	M업종_코드       string
	M지수_이름       string
	M코스피200      float64
	M코스피200_등락부호 byte
	M코스피200_등락폭  float64
	M채권_지수       float64
	M채권_지수_등락부호  byte
	M채권_지수_등락폭   float64
	M해외_지수_코드    string
	M기타_업종_코드    string
	M채권_지수_코드    string
	M채권_지수_세부_코드 string
}

func New_ETF_현재가_조회_기반_지수_자료(c Tc1151OutBlock5) S_ETF_현재가_조회_기반_지수_자료 {
	return S_ETF_현재가_조회_기반_지수_자료{
		M지수_코드:       공용.F2문자열(c.IndexCode[:]),
		M업종_코드:       공용.F2문자열(c.SectorCode[:]),
		M지수_이름:       공용.F2문자열(c.IndexName[:]),
		M코스피200:      공용.F2실수_바이트(c.KP200Index[:]),
		M코스피200_등락부호: c.KP200Sign[0],
		M코스피200_등락폭:  공용.F2실수_바이트(c.KP200Diff[:]),
		M채권_지수:       공용.F2실수_바이트(c.BondIndex[:]),
		M채권_지수_등락부호:  c.BondDiff[0],
		M채권_지수_등락폭:   공용.F2실수_바이트(c.BondDiff[:]),
		M해외_지수_코드:    공용.F2문자열(c.ForeignIndexSymbol[:]),
		M기타_업종_코드:    공용.F2문자열(c.EtcSectorCode[:]),
		M채권_지수_코드:    공용.F2문자열(c.BondIndexCode[:]),
		M채권_지수_세부_코드: 공용.F2문자열(c.BondDetailCode[:])}
}

//----------------------------------------------------------------------//
// 코스피 호가 잔량 (h1)
//----------------------------------------------------------------------//
func New코스피_호가_잔량_질의(종목_코드 string) C.Th1InBlock {
	c := Th1InBlock{}

	if len(종목_코드) > len(c.Code) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.Code[:]), len(종목_코드))

		panic(에러)
	}

	바이트_모음 := []byte(종목_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.Code[i] = 바이트_모음[i]
	}

	return *((*C.Th1InBlock)(unsafe.Pointer(&c)))
}

type S코스피_호가_잔량 struct {
	M종목_코드         string
	M시각            time.Time
	M매도_호가         int64
	M매수_호가         int64
	M매도_호가_잔량      int64
	M매수_호가_잔량      int64
	M차선_매도_호가      int64
	M차선_매수_호가      int64
	M차선_매도_호가_잔량   int64
	M차선_매수_호가_잔량   int64
	M차차선_매도_호가     int64
	M차차선_매수_호가     int64
	M차차선_매도_호가_잔량  int64
	M차차선_매수_호가_잔량  int64
	M4차선_매도_호가     int64
	M4차선_매수_호가     int64
	M4차선_매도_호가_잔량  int64
	M4차선_매수_호가_잔량  int64
	M5차선_매도_호가     int64
	M5차선_매수_호가     int64
	M5차선_매도_호가_잔량  int64
	M5차선_매수_호가_잔량  int64
	M6차선_매도_호가     int64
	M6차선_매수_호가     int64
	M6차선_매도_호가_잔량  int64
	M6차선_매수_호가_잔량  int64
	M7차선_매도_호가     int64
	M7차선_매수_호가     int64
	M7차선_매도_호가_잔량  int64
	M7차선_매수_호가_잔량  int64
	M8차선_매도_호가     int64
	M8차선_매수_호가     int64
	M8차선_매도_호가_잔량  int64
	M8차선_매수_호가_잔량  int64
	M9차선_매도_호가     int64
	M9차선_매수_호가     int64
	M9차선_매도_호가_잔량  int64
	M9차선_매수_호가_잔량  int64
	M10차선_매도_호가    int64
	M10차선_매수_호가    int64
	M10차선_매도_호가_잔량 int64
	M10차선_매수_호가_잔량 int64
	M누적_거래량        int64
}

func New코스피_호가_잔량(c Th1OutBlock) S코스피_호가_잔량 {
	return S코스피_호가_잔량{
		M종목_코드:         공용.F2문자열(c.Code[:]),
		M시각:            공용.F2시각_바이트(c.Time[:], "포맷 문자열은 내용 확인 후 작성할 것."),
		M매도_호가:         공용.F2정수64_바이트(c.OfferPrice1[:]),
		M매수_호가:         공용.F2정수64_바이트(c.BidPrice1[:]),
		M매도_호가_잔량:      공용.F2정수64_바이트(c.OfferVolume1[:]),
		M매수_호가_잔량:      공용.F2정수64_바이트(c.BidVolume1[:]),
		M차선_매도_호가:      공용.F2정수64_바이트(c.OfferPrice2[:]),
		M차선_매수_호가:      공용.F2정수64_바이트(c.BidPrice2[:]),
		M차선_매도_호가_잔량:   공용.F2정수64_바이트(c.OfferVolume2[:]),
		M차선_매수_호가_잔량:   공용.F2정수64_바이트(c.BidVolume2[:]),
		M차차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice3[:]),
		M차차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice3[:]),
		M차차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume3[:]),
		M차차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume3[:]),
		M4차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice4[:]),
		M4차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice4[:]),
		M4차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume4[:]),
		M4차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume4[:]),
		M5차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice5[:]),
		M5차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice5[:]),
		M5차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume5[:]),
		M5차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume5[:]),
		M6차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice6[:]),
		M6차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice6[:]),
		M6차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume6[:]),
		M6차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume6[:]),
		M7차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice7[:]),
		M7차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice7[:]),
		M7차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume7[:]),
		M7차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume7[:]),
		M8차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice8[:]),
		M8차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice8[:]),
		M8차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume8[:]),
		M8차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume8[:]),
		M9차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice9[:]),
		M9차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice9[:]),
		M9차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume9[:]),
		M9차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume9[:]),
		M10차선_매도_호가:    공용.F2정수64_바이트(c.OfferPrice10[:]),
		M10차선_매수_호가:    공용.F2정수64_바이트(c.BidPrice10[:]),
		M10차선_매도_호가_잔량: 공용.F2정수64_바이트(c.OfferVolume10[:]),
		M10차선_매수_호가_잔량: 공용.F2정수64_바이트(c.BidVolume10[:]),
		M누적_거래량:        공용.F2정수64_바이트(c.Volume[:])}
}

//----------------------------------------------------------------------//
// 코스닥 호가 잔량 (k3)
//----------------------------------------------------------------------//
func New코스닥_호가_잔량_질의(종목_코드 string) C.Tk3InBlock {
	c := Tk3InBlock{}

	if len(종목_코드) > len(c.Code) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.Code[:]), len(종목_코드))

		panic(에러)
	}

	바이트_모음 := []byte(종목_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.Code[i] = 바이트_모음[i]
	}

	return *((*C.Tk3InBlock)(unsafe.Pointer(&c)))
}

type S코스닥_호가_잔량 struct {
	M종목_코드         string
	M시각            time.Time
	M매도_호가         int64
	M매수_호가         int64
	M매도_호가_잔량      int64
	M매수_호가_잔량      int64
	M차선_매도_호가      int64
	M차선_매수_호가      int64
	M차선_매도_호가_잔량   int64
	M차선_매수_호가_잔량   int64
	M차차선_매도_호가     int64
	M차차선_매수_호가     int64
	M차차선_매도_호가_잔량  int64
	M차차선_매수_호가_잔량  int64
	M4차선_매도_호가     int64
	M4차선_매수_호가     int64
	M4차선_매도_호가_잔량  int64
	M4차선_매수_호가_잔량  int64
	M5차선_매도_호가     int64
	M5차선_매수_호가     int64
	M5차선_매도_호가_잔량  int64
	M5차선_매수_호가_잔량  int64
	M6차선_매도_호가     int64
	M6차선_매수_호가     int64
	M6차선_매도_호가_잔량  int64
	M6차선_매수_호가_잔량  int64
	M7차선_매도_호가     int64
	M7차선_매수_호가     int64
	M7차선_매도_호가_잔량  int64
	M7차선_매수_호가_잔량  int64
	M8차선_매도_호가     int64
	M8차선_매수_호가     int64
	M8차선_매도_호가_잔량  int64
	M8차선_매수_호가_잔량  int64
	M9차선_매도_호가     int64
	M9차선_매수_호가     int64
	M9차선_매도_호가_잔량  int64
	M9차선_매수_호가_잔량  int64
	M10차선_매도_호가    int64
	M10차선_매수_호가    int64
	M10차선_매도_호가_잔량 int64
	M10차선_매수_호가_잔량 int64
	M누적_거래량        int64
}

func New코스닥_호가_잔량(c Tk3OutBlock) S코스닥_호가_잔량 {
	return S코스닥_호가_잔량{
		M종목_코드:         공용.F2문자열(c.Code[:]),
		M시각:            공용.F2시각_바이트(c.Time[:], "포맷 문자열은 내용 확인 후 작성할 것."),
		M매도_호가:         공용.F2정수64_바이트(c.OfferPrice1[:]),
		M매수_호가:         공용.F2정수64_바이트(c.BidPrice1[:]),
		M매도_호가_잔량:      공용.F2정수64_바이트(c.OfferVolume1[:]),
		M매수_호가_잔량:      공용.F2정수64_바이트(c.BidVolume1[:]),
		M차선_매도_호가:      공용.F2정수64_바이트(c.OfferPrice2[:]),
		M차선_매수_호가:      공용.F2정수64_바이트(c.BidPrice2[:]),
		M차선_매도_호가_잔량:   공용.F2정수64_바이트(c.OfferVolume2[:]),
		M차선_매수_호가_잔량:   공용.F2정수64_바이트(c.BidVolume2[:]),
		M차차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice3[:]),
		M차차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice3[:]),
		M차차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume3[:]),
		M차차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume3[:]),
		M4차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice4[:]),
		M4차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice4[:]),
		M4차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume4[:]),
		M4차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume4[:]),
		M5차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice5[:]),
		M5차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice5[:]),
		M5차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume5[:]),
		M5차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume5[:]),
		M6차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice6[:]),
		M6차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice6[:]),
		M6차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume6[:]),
		M6차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume6[:]),
		M7차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice7[:]),
		M7차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice7[:]),
		M7차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume7[:]),
		M7차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume7[:]),
		M8차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice8[:]),
		M8차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice8[:]),
		M8차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume8[:]),
		M8차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume8[:]),
		M9차선_매도_호가:     공용.F2정수64_바이트(c.OfferPrice9[:]),
		M9차선_매수_호가:     공용.F2정수64_바이트(c.BidPrice9[:]),
		M9차선_매도_호가_잔량:  공용.F2정수64_바이트(c.OfferVolume9[:]),
		M9차선_매수_호가_잔량:  공용.F2정수64_바이트(c.BidVolume9[:]),
		M10차선_매도_호가:    공용.F2정수64_바이트(c.OfferPrice10[:]),
		M10차선_매수_호가:    공용.F2정수64_바이트(c.BidPrice10[:]),
		M10차선_매도_호가_잔량: 공용.F2정수64_바이트(c.OfferVolume10[:]),
		M10차선_매수_호가_잔량: 공용.F2정수64_바이트(c.BidVolume10[:]),
		M누적_거래량:        공용.F2정수64_바이트(c.Volume[:])}
}

//----------------------------------------------------------------------//
// 코스피 시간외 호가 잔량 (h2)
//----------------------------------------------------------------------//

func New코스피_시간외_호가_잔량_질의(종목_코드 string) C.Th2InBlock {
	c := Th2InBlock{}

	if len(종목_코드) > len(c.Code) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.Code[:]), len(종목_코드))

		panic(에러)
	}

	바이트_모음 := []byte(종목_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.Code[i] = 바이트_모음[i]
	}

	return *((*C.Th2InBlock)(unsafe.Pointer(&c)))
}

type S코스피_시간외_호가_잔량 struct {
	M종목_코드    string
	M시각       time.Time
	M매도_호가_잔량 int64
	M매수_호가_잔량 int64
}

func New코스피_시간외_호가_잔량(c Th2OutBlock) S코스피_시간외_호가_잔량 {
	return S코스피_시간외_호가_잔량{
		M종목_코드:    공용.F2문자열(c.Code[:]),
		M시각:       공용.F2시각_바이트(c.Time[:], "포맷 문자열은 내용 확인 후 작성할 것."),
		M매도_호가_잔량: 공용.F2정수64_바이트(c.OfferVolume[:]),
		M매수_호가_잔량: 공용.F2정수64_바이트(c.BidVolume[:])}
}

//----------------------------------------------------------------------//
// 코스닥 시간외 호가 잔량 (k4)
//----------------------------------------------------------------------//

func New코스닥_시간외_호가_잔량_질의(종목_코드 string) C.Tk4InBlock {
	c := Tk4InBlock{}

	if len(종목_코드) > len(c.Code) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.Code[:]), len(종목_코드))

		panic(에러)
	}

	바이트_모음 := []byte(종목_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.Code[i] = 바이트_모음[i]
	}

	return *((*C.Tk4InBlock)(unsafe.Pointer(&c)))
}

type S코스닥_시간외_호가_잔량 struct {
	M종목_코드    string
	M시각       time.Time
	M매도_호가_잔량 int64
	M매수_호가_잔량 int64
}

func New코스닥_시간외_호가_잔량(c Tk4OutBlock) S코스닥_시간외_호가_잔량 {
	return S코스닥_시간외_호가_잔량{
		M종목_코드:    공용.F2문자열(c.Code[:]),
		M시각:       공용.F2시각_바이트(c.Time[:], "포맷 문자열은 내용 확인 후 작성할 것."),
		M매도_호가_잔량: 공용.F2정수64_바이트(c.OfferVolume[:]),
		M매수_호가_잔량: 공용.F2정수64_바이트(c.BidVolume[:])}
}

//----------------------------------------------------------------------//
// 코스피 예상 호가 잔량 (h3)
//----------------------------------------------------------------------//

func New코스피_예상_호가_잔량_질의(종목_코드 string) C.Th3InBlock {
	c := Th3InBlock{}

	if len(종목_코드) > len(c.Code) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.Code[:]), len(종목_코드))

		panic(에러)
	}

	바이트_모음 := []byte(종목_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.Code[i] = 바이트_모음[i]
	}

	return *((*C.Th3InBlock)(unsafe.Pointer(&c)))
}

type S코스피_예상_호가_잔량 struct {
	M종목_코드    string
	M시각       time.Time
	M동시_호가_구분 string
	M예상_체결가   int64
	M예상_등락_부호 byte
	M예상_등락폭   int64
	M예상_등락율   float64
	M예상_체결_수량 int64
	M매도_호가    int64
	M매수_호가    int64
	M매도_호가_잔량 int64
	M매수_호가_잔량 int64
}

func New코스피_예상_호가_잔량(c Th3OutBlock) S코스피_예상_호가_잔량 {
	return S코스피_예상_호가_잔량{
		M종목_코드:    공용.F2문자열(c.Code[:]),
		M시각:       공용.F2시각_바이트(c.Time[:], "포맷 문자열은 내용 확인 후 작성할 것."),
		M동시_호가_구분: 공용.F2문자열(c.SyncOfferBid[:]),
		M예상_체결가:   공용.F2정수64_바이트(c.EstmPrice[:]),
		M예상_등락_부호: c.EstmDiffSign[0],
		M예상_등락폭:   공용.F2정수64_바이트(c.EstmDiff[:]),
		M예상_등락율:   공용.F2실수_바이트(c.EstmDiffRate[:]),
		M예상_체결_수량: 공용.F2정수64_바이트(c.EstmVolume[:]),
		M매도_호가:    공용.F2정수64_바이트(c.OfferPrice[:]),
		M매수_호가:    공용.F2정수64_바이트(c.BidPrice[:]),
		M매도_호가_잔량: 공용.F2정수64_바이트(c.OfferVolume[:]),
		M매수_호가_잔량: 공용.F2정수64_바이트(c.BidVolume[:])}
}

//----------------------------------------------------------------------//
// 코스닥 예상 호가 잔량 (k5)
//----------------------------------------------------------------------//

func New코스닥_예상_호가_잔량_질의(종목_코드 string) C.Tk5InBlock {
	c := Tk5InBlock{}

	if len(종목_코드) > len(c.Code) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.Code[:]), len(종목_코드))

		panic(에러)
	}

	바이트_모음 := []byte(종목_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.Code[i] = 바이트_모음[i]
	}

	return *((*C.Tk5InBlock)(unsafe.Pointer(&c)))
}

type S코스닥_예상_호가_잔량 struct {
	M종목_코드    string
	M시각       time.Time
	M동시_호가_구분 string
	M예상_체결가   int64
	M예상_등락_부호 byte
	M예상_등락폭   int64
	M예상_등락율   float64
	M예상_체결_수량 int64
	M매도_호가    int64
	M매수_호가    int64
	M매도_호가_잔량 int64
	M매수_호가_잔량 int64
}

func New코스닥_예상_호가_잔량(c Tk5OutBlock) S코스닥_예상_호가_잔량 {
	return S코스닥_예상_호가_잔량{
		M종목_코드:    공용.F2문자열(c.Code[:]),
		M시각:       공용.F2시각_바이트(c.Time[:], "포맷 문자열"),
		M동시_호가_구분: 공용.F2문자열(c.SyncOfferBid[:]),
		M예상_체결가:   공용.F2정수64_바이트(c.EstmPrice[:]),
		M예상_등락_부호: c.EstmDiffSign[0],
		M예상_등락폭:   공용.F2정수64_바이트(c.EstmDiff[:]),
		M예상_등락율:   공용.F2실수_바이트(c.EstmDiffRate[:]),
		M예상_체결_수량: 공용.F2정수64_바이트(c.EstmVolume[:]),
		M매도_호가:    공용.F2정수64_바이트(c.OfferPrice[:]),
		M매수_호가:    공용.F2정수64_바이트(c.BidPrice[:]),
		M매도_호가_잔량: 공용.F2정수64_바이트(c.OfferVolume[:]),
		M매수_호가_잔량: 공용.F2정수64_바이트(c.BidVolume[:])}
}

//----------------------------------------------------------------------//
// 코스피 체결 (j8)
//----------------------------------------------------------------------//

func New코스피_체결_질의(종목_코드 string) C.Tj8InBlock {
	c := Tj8InBlock{}

	if len(종목_코드) > len(c.Code) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.Code[:]), len(종목_코드))

		panic(에러)
	}

	바이트_모음 := []byte(종목_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.Code[i] = 바이트_모음[i]
	}

	return *((*C.Tj8InBlock)(unsafe.Pointer(&c)))
}

type S코스피_체결 struct {
	M종목_코드        string
	M시각           time.Time
	M등락_부호        byte
	M등락폭          int64
	M현재가          int64
	M등락율          float64
	M고가           int64
	M저가           int64
	M매도_호가        int64
	M매수_호가        int64
	M누적_거래량       int64
	M전일_거래량_대비_비율 float64
	M변동_거래량       int64
	M거래_대금        int64
	M시가           int64
	M가중_평균_가격     int64
	M장구분          string
}

func New코스피_체결(c Tj8OutBlock) S코스피_체결 {
	return S코스피_체결{
		M종목_코드:        공용.F2문자열(c.Code[:]),
		M시각:           공용.F2시각_바이트(c.Time[:], "포맷 문자열은 내용 확인 후 작성할 것."),
		M등락_부호:        c.DiffSign[0],
		M등락폭:          공용.F2정수64_바이트(c.Diff[:]),
		M현재가:          공용.F2정수64_바이트(c.MarketPrice[:]),
		M등락율:          공용.F2실수_바이트(c.DiffRate[:]),
		M고가:           공용.F2정수64_바이트(c.High[:]),
		M저가:           공용.F2정수64_바이트(c.Low[:]),
		M매도_호가:        공용.F2정수64_바이트(c.OfferPrice[:]),
		M매수_호가:        공용.F2정수64_바이트(c.BidPrice[:]),
		M누적_거래량:       공용.F2정수64_바이트(c.Volume[:]),
		M전일_거래량_대비_비율: 공용.F2실수_바이트(c.VsPrevVolRate[:]),
		M변동_거래량:       공용.F2정수64_바이트(c.DiffVolume[:]),
		M거래_대금:        공용.F2정수64_바이트(c.TrAmount[:]),
		M시가:           공용.F2정수64_바이트(c.Open[:]),
		M가중_평균_가격:     공용.F2정수64_바이트(c.WeightAvgPrice[:]),
		M장구분:          공용.F2문자열(c.Market[:])}
}

//----------------------------------------------------------------------//
// 코스닥 체결 (k8)
//----------------------------------------------------------------------//

func New코스닥_체결_질의(종목_코드 string) C.Tk8InBlock {
	c := Tk8InBlock{}

	if len(종목_코드) > len(c.Code) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.Code[:]), len(종목_코드))

		panic(에러)
	}

	바이트_모음 := []byte(종목_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.Code[i] = 바이트_모음[i]
	}

	return *((*C.Tk8InBlock)(unsafe.Pointer(&c)))
}

type S코스닥_체결 struct {
	M종목_코드        string
	M시각           time.Time
	M등락_부호        byte
	M등락폭          int64
	M현재가          int64
	M등락율          float64
	M고가           int64
	M저가           int64
	M매도_호가        int64
	M매수_호가        int64
	M누적_거래량       int64
	M전일_거래량_대비_비율 float64
	M변동_거래량       int64
	M거래_대금        int64
	M시가           int64
	M가중_평균_가격     int64
	M장구분          string
}

func New코스닥_체결(c Tj8OutBlock) S코스닥_체결 {
	return S코스닥_체결{
		M종목_코드:        공용.F2문자열(c.Code[:]),
		M시각:           공용.F2시각_바이트(c.Time[:], "포맷 문자열은 내용 확인 후 작성할 것."),
		M등락_부호:        c.DiffSign[0],
		M등락폭:          공용.F2정수64_바이트(c.Diff[:]),
		M현재가:          공용.F2정수64_바이트(c.MarketPrice[:]),
		M등락율:          공용.F2실수_바이트(c.DiffRate[:]),
		M고가:           공용.F2정수64_바이트(c.High[:]),
		M저가:           공용.F2정수64_바이트(c.Low[:]),
		M매도_호가:        공용.F2정수64_바이트(c.OfferPrice[:]),
		M매수_호가:        공용.F2정수64_바이트(c.BidPrice[:]),
		M누적_거래량:       공용.F2정수64_바이트(c.Volume[:]),
		M전일_거래량_대비_비율: 공용.F2실수_바이트(c.VsPrevVolRate[:]),
		M변동_거래량:       공용.F2정수64_바이트(c.DiffVolume[:]),
		M거래_대금:        공용.F2정수64_바이트(c.TrAmount[:]),
		M시가:           공용.F2정수64_바이트(c.Open[:]),
		M가중_평균_가격:     공용.F2정수64_바이트(c.WeightAvgPrice[:]),
		M장구분:          공용.F2문자열(c.Market[:])}
}

//----------------------------------------------------------------------//
// 코스피 ETF NAV (j1) (예제코드가 없음. 패딩 필드가 필요한 지 추가 확인 필요함.)
//----------------------------------------------------------------------//

type S_코스피_ETF_NAV struct {
	M종목_코드    string
	M시각       time.Time
	M등락_부호    byte
	M등락폭      float64
	NAV_현재가   float64
	NAV_시가    float64
	NAV_고가    float64
	NAV_저가    float64
	M추적_오차_부호 byte
	M추적_오차    float64
	M괴리율_부호   byte
	M괴리율      float64
}

func New_코스피_ETF(c Tj1OutBlock) S_코스피_ETF_NAV {
	return S_코스피_ETF_NAV{
		M종목_코드:    공용.F2문자열(c.Code[:]),
		M시각:       공용.F2시각_바이트(c.Time[:], "포맷 문자열"),
		M등락_부호:    c.DiffSign[0],
		M등락폭:      공용.F2실수_바이트(c.Diff[:]),
		NAV_현재가:   공용.F2실수_바이트(c.Current[:]),
		NAV_시가:    공용.F2실수_바이트(c.Open[:]),
		NAV_고가:    공용.F2실수_바이트(c.High[:]),
		NAV_저가:    공용.F2실수_바이트(c.Low[:]),
		M추적_오차_부호: c.TrackErrSign[0],
		M추적_오차:    공용.F2실수_바이트(c.TrackingError[:]),
		M괴리율_부호:   c.DivergeSign[0],
		M괴리율:      공용.F2실수_바이트(c.DivergeRate[:])}
}

//----------------------------------------------------------------------//
// 코스닥 ETF NAV (j0) (예제코드가 없음. 패딩 필드가 필요한 지 추가 확인 필요함.)
//----------------------------------------------------------------------//

type S_코스닥_ETF_NAV struct {
	M종목_코드    string
	M시각       time.Time
	M등락_부호    byte
	M등락폭      int64
	NAV_현재가   float64
	NAV_시가    float64
	NAV_고가    float64
	NAV_저가    float64
	M추적_오차_부호 byte
	M추적_오차    float64
	M괴리율_부호   byte
	M괴리율      float64
}

func New_코스닥_ETF(c Tj0OutBlock) S_코스닥_ETF_NAV {
	return S_코스닥_ETF_NAV{
		M종목_코드:    공용.F2문자열(c.Code[:]),
		M시각:       공용.F2시각_바이트(c.Time[:], "포맷 문자열"),
		M등락_부호:    c.DiffSign[0],
		M등락폭:      공용.F2정수64_바이트(c.Diff[:]),
		NAV_현재가:   공용.F2실수_바이트(c.Current[:]),
		NAV_시가:    공용.F2실수_바이트(c.Open[:]),
		NAV_고가:    공용.F2실수_바이트(c.High[:]),
		NAV_저가:    공용.F2실수_바이트(c.Low[:]),
		M추적_오차_부호: c.TrackErrSign[0],
		M추적_오차:    공용.F2실수_바이트(c.TrackingError[:]),
		M괴리율_부호:   c.DivergeSign[0],
		M괴리율:      공용.F2실수_바이트(c.DivergeRate[:])}
}

/* 코스피/코스닥 업종코드 참고표
코스피 업종명			코스닥 업종명
00 	KRX 100			01 	코스닥지수
01 	코스피지수			03 	기타서비스
02 	대형주			04 	코스닥 IT
03 	중형주			06 	제조
04 	소형주			07 	건설
05 	음식료품			08 	유통
06 	섬유의복			10 	운송
07 	종이목재			11 	금융
08 	화학				12 	통신방송서비스
09 	의약품			13 	IT S/W & SVC
10 	비금속광물			14 IT H/W
11 	철강금속			15 	음식료담배
12 	기계				16 	섬유의류
13 	전기전자			17 	종이목재
14 	의료정밀			18 	출판매체복제
15 	운수장비			19 	화학
16 	유통업			20 	제약
17 	전기가스업			21 	비금속
18 	건설업			22 	금속
19 	운수창고			23 	기계장비
20 	통신업			24 	일반전기전자
21 	금융업			25 	의료정밀기기
22 	은행				26 	운송장비부품
24 	증권				27 	기타 제조
25 	보험				28 	통신서비스
26 	서비스업			29 	방송서비스
27 	제조업			30 	인터넷
28 	코스피 200		31 	디지털컨텐츠
29 	코스피 100		32 	소프트웨어
30 	코스피 50			33 	컴퓨터서비스
32 	코스피 배당		34 	통신장비
39 	KP200 건설기계		35 	정보기기
40 	KP200 조선운송		36 	반도체
41 	KP200 철강소재		37 	IT부품
42 	KP200 에너지화학	38 	KOSDAQ 100
43 	KP200 정보기술		39 	KOSDAQ MID 300
44 	KP200 금융		40 	KOSDAQ SMALL
45 	KP200 생활소비재	43 	코스닥 스타
46 	KP200 경기소비재	44 	오락문화
47 	동일가중 KP200		45 	프리미어
48 	동일가중 KP100		46 	우량기업부
49 	동일가중 KP50		47 	벤처기업부
　	　				48 	중견기업부
　	　				49 	기술성장기업부  */

//----------------------------------------------------------------------//
// 코스피 업종 지수 (u1)
//----------------------------------------------------------------------//

func New코스피_업종_지수_질의(업종_코드 string) C.Tu1InBlock {
	c := Tu1InBlock{}

	if len(업종_코드) > len(c.SectorCode) {
		에러 := 공용.F에러_생성(
			"업종 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.SectorCode[:]), len(업종_코드))

		panic(에러)
	}

	바이트_모음 := []byte(업종_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.SectorCode[i] = 바이트_모음[i]
	}

	return *((*C.Tu1InBlock)(unsafe.Pointer(&c)))
}

type S코스피_업종_지수 struct {
	M업종_코드  string
	M시각     time.Time
	M지수값    float64
	M등락_부호  byte
	M등락폭    float64
	M거래량    int64
	M거래_대금  int64
	M개장값    float64
	M최고값    float64
	M최고값_시각 time.Time
	M최저값    float64
	M최저값_시간 time.Time
	M지수_등락율 float64
	M거래_비중  float64
}

func New코스피_업종_지수(c Tu1OutBlock) S코스피_업종_지수 {
	return S코스피_업종_지수{
		M업종_코드:  공용.F2문자열(c.SectorCode[:]),
		M시각:     공용.F2시각_바이트(c.Time[:], "포맷 문자열"),
		M지수값:    공용.F2실수_바이트(c.IndexValue[:]),
		M등락_부호:  c.DiffSign[0],
		M등락폭:    공용.F2실수_바이트(c.Diff[:]),
		M거래량:    공용.F2정수64_바이트(c.Volume[:]),
		M거래_대금:  공용.F2정수64_바이트(c.TrAmount[:]),
		M개장값:    공용.F2실수_바이트(c.Open[:]),
		M최고값:    공용.F2실수_바이트(c.High[:]),
		M최고값_시각: 공용.F2시각_바이트(c.HighTime[:], "포맷 문자열"),
		M최저값:    공용.F2실수_바이트(c.Low[:]),
		M최저값_시간: 공용.F2시각_바이트(c.LowTime[:], "포맷 문자열"),
		M지수_등락율: 공용.F2실수_바이트(c.DiffRate[:]),
		M거래_비중:  공용.F2실수_바이트(c.TrVolRate[:])}
}

//----------------------------------------------------------------------//
// 코스닥 업종 지수 (k1)
//----------------------------------------------------------------------//

func New코스닥_업종_지수_질의(업종_코드 string) C.Tk1InBlock {
	c := Tk1InBlock{}

	if len(업종_코드) > len(c.SectorCode) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상 : %v, 실제 : %v.",
			len(c.SectorCode[:]), len(업종_코드))

		panic(에러)
	}

	바이트_모음 := []byte(업종_코드)

	for i := 0; i < len(바이트_모음); i++ {
		c.SectorCode[i] = 바이트_모음[i]
	}

	return *((*C.Tk1InBlock)(unsafe.Pointer(&c)))
}

type S코스닥_업종_지수 struct {
	M업종_코드  string
	M시각     time.Time
	M지수값    float64
	M등락_부호  byte
	M등락폭    float64
	M거래량    int64
	M거래_대금  int64
	M개장값    float64
	M최고값    float64
	M최고값_시각 time.Time
	M최저값    float64
	M최저값_시간 time.Time
	M지수_등락율 float64
	M거래_비중  float64
}

func New코스닥_업종_지수(c Tk1OutBlock) S코스닥_업종_지수 {
	return S코스닥_업종_지수{
		M업종_코드:  공용.F2문자열(c.SectorCode[:]),
		M시각:     공용.F2시각_바이트(c.Time[:], "포맷 문자열"),
		M지수값:    공용.F2실수_바이트(c.IndexValue[:]),
		M등락_부호:  c.DiffSign[0],
		M등락폭:    공용.F2실수_바이트(c.Diff[:]),
		M거래량:    공용.F2정수64_바이트(c.Volume[:]),
		M거래_대금:  공용.F2정수64_바이트(c.TrAmount[:]),
		M개장값:    공용.F2실수_바이트(c.Open[:]),
		M최고값:    공용.F2실수_바이트(c.High[:]),
		M최고값_시각: 공용.F2시각_바이트(c.HighTime[:], "포맷 문자열"),
		M최저값:    공용.F2실수_바이트(c.Low[:]),
		M최저값_시간: 공용.F2시각_바이트(c.LowTime[:], "포맷 문자열"),
		M지수_등락율: 공용.F2실수_바이트(c.DiffRate[:]),
		M거래_비중:  공용.F2실수_바이트(c.TrVolRate[:])}
}
