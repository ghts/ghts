package NH

// #include <stdlib.h>
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
	TR식별번호  uint32
	M로그인_정보 S로그인_정보
}

func New로그인_정보_블록(c *C.LOGINBLOCK) *S로그인_정보_블록 {
	// 포인터 형태 자료형은 모두 정리.
	defer func() {
		C.free(unsafe.Pointer(c.LoginInfo))
		C.free(unsafe.Pointer(c))
	}()
	
	s := new(S로그인_정보_블록)
	s.TR식별번호 = uint32(c.TrIndex)
	s.M로그인_정보 = new로그인_정보(c.LoginInfo)
	
	return s
}

type S로그인_정보 struct {
	M접속시각  time.Time
	M접속서버  string
	M접속ID  string
	M계좌_목록 []S계좌_정보
}

func new로그인_정보(c *C.LOGININFO) S로그인_정보 {
	g := (*LoginInfo)(unsafe.Pointer(c))
	
	시각, 에러 := time.Parse(공용.F2문자열(g.Date[:]), "값을 확인한 후 포맷 문자열 수정할 것.")
	공용.F에러_패닉(에러)

	계좌_수량, 에러 := 공용.F2정수(공용.F2문자열(g.AccountCount))
	공용.F에러_패닉(에러)

	계좌_목록 := make([]S계좌_정보, 계좌_수량)
	for i := 0; i < 계좌_수량; i++ {
		계좌_목록[i] = new계좌_정보(g.Accountlist[i])
	}

	s := new(S로그인_정보)
	s.M접속시각 = 시각
	s.M접속서버 = 공용.F2문자열(g.ServerName[:])
	s.M접속ID = 공용.F2문자열(g.UserID[:])
	s.M계좌_목록 = 계좌_목록
	
	return *s
}

type S계좌_정보 struct {
	M계좌_번호     string
	M계좌명       string
	M상품_코드     string
	M관리점_코드    string
	M위임_만기일    time.Time
	M일괄주문_허용계좌 bool // ('G' = 허용)
	M주석        string
}

func new계좌_정보(g AccountInfo) S계좌_정보 {
	위임_만기일, 에러 := time.Parse(공용.F2문자열(g.ExpirationDate8[:]), "값을 확인한 후 포맷 문자열 수정할 것.")
	공용.F에러_패닉(에러)

	일괄주문_허용계좌 := false
	if 공용.F2문자열(g.Granted) == "G" {
		일괄주문_허용계좌 = true
	}

	s := new(S계좌_정보)
	s.M계좌_번호 = 공용.F2문자열(g.AccountNo[:])
	s.M계좌명 = 공용.F2문자열(g.AccountName[:])
	s.M상품_코드 = 공용.F2문자열(g.Act_pdt_cdz3[:])
	s.M관리점_코드 = 공용.F2문자열(g.Amn_tab_cdz4[:])
	s.M위임_만기일 = 위임_만기일
	s.M일괄주문_허용계좌 = 일괄주문_허용계좌
	s.M주석 = 공용.F2문자열(g.Filler[:])

	return *s
}

//----------------------------------------------------------------------//
// WMCA 문자 message 구조체
//----------------------------------------------------------------------//
type S수신_메시지_블록 struct {
	TR식별번호  uint32
	M메시지_코드 string //00000:정상, 기타:비정상(코드값은 언제든지 변경될 수 있음.)
	M메시지_내용 string
}

func New수신_메시지_블록(c *C.OUTDATABLOCK) *S수신_메시지_블록 {
	defer func() {
		C.free(unsafe.Pointer(c.DataStruct))
		C.free(unsafe.Pointer(c))
	}()
	
	g := (*MsgHeader)(unsafe.Pointer(c.DataStruct))
	
	s := new(S수신_메시지_블록)
	s.TR식별번호 = uint32(c.TrIndex)
	s.M메시지_코드 = 공용.F2문자열(g.MsgCode[:])
	s.M메시지_내용 = 공용.F2문자열(g.UsrMsg[:])
	
	return s	
}

//----------------------------------------------------------------------//
// WMCA TR 응답 구조체
//----------------------------------------------------------------------//
type S수신_데이터_블록 struct {
	TR식별번호  uint32
	M수신_데이터 S수신_데이터
}

func New수신_데이터_블록(c *C.OUTDATABLOCK) *S수신_데이터_블록 {
	// 포인터로 표시된 자료형은 모두 정리.
	defer func() {
		C.free(unsafe.Pointer(c.DataStruct.BlockName))
		C.free(unsafe.Pointer(c.DataStruct.DataString))
		C.free(unsafe.Pointer(c.DataStruct))
		C.free(unsafe.Pointer(c))
	}()
	
	s := new(S수신_데이터_블록)
	s.TR식별번호 = uint32(c.TrIndex)
	s.M수신_데이터 = new수신_데이터(c.DataStruct)
	
	return s
}

type S수신_데이터 struct {
	M블록_이름 string
	M데이터   interface{}
}

func new수신_데이터(c *C.RECEIVED) S수신_데이터 {
	블록_이름, 데이터 := f_Go구조체로_변환(c)
	
	s := new(S수신_데이터)
	s.M블록_이름 = 블록_이름
	s.M데이터 = 데이터
	
	return *s
}

//----------------------------------------------------------------------//
// 주식 현재가 조회 (c1101)
//----------------------------------------------------------------------//
func New주식_현재가_조회_질의(종목_코드 string) C.Tc1101InBlock {
	// 포인터는 동시 실행에서 위험성이 있고,
	// 구조체의 데이터 크기가 얼마 되지 않으므로,
	// 포인터를 사용하지 않고 구조체를 그냥 사용함.		
	c := Tc1101InBlock{}
	c.Lang[0] = ([]byte("K"))[0]

	if len(종목_코드) > len(c.Code) {
		종목_코드 = 종목_코드[:len(c.Code)]
		
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 허용 한도를 초과함. 허용 한도  = %v, 종목_코드 길이  = %v.",
			len(c.Code[:]), len(종목_코드))
		공용.F에러_출력(에러)
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
	M대량_매매_방향          int8      // 0 = 해당없음 1 = 매도 2 = 매수
	M대량_매매_존재          bool
}

func New주식_현재가_조회_기본_자료(c *C.char) *S주식_현재가_조회_기본_자료 {
	g := (*Tc1101OutBlock)(unsafe.Pointer(c))
	
	s := new(S주식_현재가_조회_기본_자료)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M종목명 = 공용.F2문자열(g.Title[:])
	s.M현재가 = 공용.F2정수64_바이트(g.MarketPrice[:])
	s.M등락부호 = g.DiffSign[0]
	s.M등락폭 = 공용.F2정수64_바이트(g.Diff[:])
	s.M등락율 = 공용.F2실수_바이트(g.DiffRate[:])
	s.M매도_호가 = 공용.F2정수64_바이트(g.OfferPrice[:])
	s.M매수_호가 = 공용.F2정수64_바이트(g.BidPrice[:])
	s.M거래량 = 공용.F2정수64_바이트(g.Volume[:])
	s.M거래비율 = 공용.F2실수_바이트(g.TrVolRate[:])
	s.M유동주_회전율 = 공용.F2실수_바이트(g.FloatRate[:])
	s.M거래대금 = 공용.F2정수64_바이트(g.TrAmount[:])
	s.M상한가 = 공용.F2정수64_바이트(g.UpLmtPrice[:])
	s.M고가 = 공용.F2정수64_바이트(g.High[:])
	s.M시가 = 공용.F2정수64_바이트(g.Open[:])
	s.M시가_대비_부호 = g.VsOpenSign[0]
	s.M시가_대비_등락폭 = 공용.F2정수64_바이트(g.VsOpenDiff[:])
	s.M저가 = 공용.F2정수64_바이트(g.Low[:])
	s.M하한가 = 공용.F2정수64_바이트(g.LowLmtPrice[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열은 내용 확인 후 수정할 것.")
	s.M매도_호가_최우선 = 공용.F2정수64_바이트(g.OfferPrice1[:])
	s.M매도_호가_차선 = 공용.F2정수64_바이트(g.OfferPrice2[:])
	s.M매도_호가_차차선 = 공용.F2정수64_바이트(g.OfferPrice3[:])
	s.M매도_호가_4차선 = 공용.F2정수64_바이트(g.OfferPrice4[:])
	s.M매도_호가_5차선 = 공용.F2정수64_바이트(g.OfferPrice5[:])
	s.M매도_호가_6차선 = 공용.F2정수64_바이트(g.OfferPrice6[:])
	s.M매도_호가_7차선 = 공용.F2정수64_바이트(g.OfferPrice7[:])
	s.M매도_호가_8차선 = 공용.F2정수64_바이트(g.OfferPrice8[:])
	s.M매도_호가_9차선 = 공용.F2정수64_바이트(g.OfferPrice9[:])
	s.M매도_호가_10차선 = 공용.F2정수64_바이트(g.OfferPrice10[:])
	s.M매수_호가_최우선 = 공용.F2정수64_바이트(g.BidPrice1[:])
	s.M매수_호가_차선 = 공용.F2정수64_바이트(g.BidPrice2[:])
	s.M매수_호가_차차선 = 공용.F2정수64_바이트(g.BidPrice3[:])
	s.M매수_호가_4차선 = 공용.F2정수64_바이트(g.BidPrice4[:])
	s.M매수_호가_5차선 = 공용.F2정수64_바이트(g.BidPrice5[:])
	s.M매수_호가_6차선 = 공용.F2정수64_바이트(g.BidPrice6[:])
	s.M매수_호가_7차선 = 공용.F2정수64_바이트(g.BidPrice7[:])
	s.M매수_호가_8차선 = 공용.F2정수64_바이트(g.BidPrice8[:])
	s.M매수_호가_9차선 = 공용.F2정수64_바이트(g.BidPrice9[:])
	s.M매수_호가_10차선 = 공용.F2정수64_바이트(g.BidPrice10[:])
	s.M매도_최우선_잔량 = 공용.F2정수64_바이트(g.OfferVolume1[:])
	s.M매도_차선_잔량 = 공용.F2정수64_바이트(g.OfferVolume2[:])
	s.M매도_차차선_잔량 = 공용.F2정수64_바이트(g.OfferVolume3[:])
	s.M매도_4차선_잔량 = 공용.F2정수64_바이트(g.OfferVolume4[:])
	s.M매도_5차선_잔량 = 공용.F2정수64_바이트(g.OfferVolume5[:])
	s.M매도_6차선_잔량 = 공용.F2정수64_바이트(g.OfferVolume6[:])
	s.M매도_7차선_잔량 = 공용.F2정수64_바이트(g.OfferVolume7[:])
	s.M매도_8차선_잔량 = 공용.F2정수64_바이트(g.OfferVolume8[:])
	s.M매도_9차선_잔량 = 공용.F2정수64_바이트(g.OfferVolume9[:])
	s.M매도_10차선_잔량 = 공용.F2정수64_바이트(g.OfferVolume10[:])
	s.M매수_최우선_잔량 = 공용.F2정수64_바이트(g.BidVolume1[:])
	s.M매수_차선_잔량 = 공용.F2정수64_바이트(g.BidVolume2[:])
	s.M매수_차차선_잔량 = 공용.F2정수64_바이트(g.BidVolume3[:])
	s.M매수_4차선_잔량 = 공용.F2정수64_바이트(g.BidVolume4[:])
	s.M매수_5차선_잔량 = 공용.F2정수64_바이트(g.BidVolume5[:])
	s.M매수_6차선_잔량 = 공용.F2정수64_바이트(g.BidVolume6[:])
	s.M매수_7차선_잔량 = 공용.F2정수64_바이트(g.BidVolume7[:])
	s.M매수_8차선_잔량 = 공용.F2정수64_바이트(g.BidVolume8[:])
	s.M매수_9차선_잔량 = 공용.F2정수64_바이트(g.BidVolume9[:])
	s.M매수_10차선_잔량 = 공용.F2정수64_바이트(g.BidVolume10[:])
	s.M매도_잔량_총합 = 공용.F2정수64_바이트(g.OfferVolTot[:])
	s.M매수_잔량_총합 = 공용.F2정수64_바이트(g.BidVolTot[:])
	s.M시간외_매도_잔량 = 공용.F2정수64_바이트(g.OfferVolAfterHour[:])
	s.M시간외_매수_잔량 = 공용.F2정수64_바이트(g.BidVolAfterHour[:])
	s.M피봇_2차_저항 = 공용.F2정수64_바이트(g.PivotUp2[:])
	s.M피봇_1차_저항 = 공용.F2정수64_바이트(g.PivotUp1[:])
	s.M피봇가 = 공용.F2정수64_바이트(g.PivotPrice[:])
	s.M피봇_1차_지지 = 공용.F2정수64_바이트(g.PivotDown1[:])
	s.M피봇_2차_지지 = 공용.F2정수64_바이트(g.PivotDown2[:])
	s.M코스피_코스닥_구분 = 공용.F2문자열(g.Market[:])
	s.M업종명 = 공용.F2문자열(g.Sector[:])
	s.M자본금_규모 = 공용.F2문자열(g.CapSize[:])
	s.M결산월 = 공용.F2문자열(g.SettleMonth[:])
	s.M시장조치1 = 공용.F2문자열(g.MarketAction1[:])
	s.M시장조치2 = 공용.F2문자열(g.MarketAction2[:])
	s.M시장조치3 = 공용.F2문자열(g.MarketAction3[:])
	s.M시장조치4 = 공용.F2문자열(g.MarketAction4[:])
	s.M시장조치5 = 공용.F2문자열(g.MarketAction5[:])
	s.M시장조치6 = 공용.F2문자열(g.MarketAction6[:])
	s.M전환사채_구분 = 공용.F2문자열(g.ConvertBond[:])
	s.M액면가 = 공용.F2정수64_바이트(g.NominalPrice[:])
	s.M전일종가_타이틀 = 공용.F2문자열(g.PrevPriceTitle[:])
	s.M전일종가 = 공용.F2정수64_바이트(g.PrevPrice[:])
	s.M대용가 = 공용.F2정수64_바이트(g.MortgageValue[:])
	s.M공모가 = 공용.F2정수64_바이트(g.PublicOfferPrice[:])
	s.M5일_고가 = 공용.F2정수64_바이트(g.High5Day[:])
	s.M5일_저가 = 공용.F2정수64_바이트(g.Low5Day[:])
	s.M20일_고가 = 공용.F2정수64_바이트(g.High20Day[:])
	s.M20일_저가 = 공용.F2정수64_바이트(g.Low20Day[:])
	s.M52주_고가 = 공용.F2정수64_바이트(g.High1Year[:])
	s.M52주_고가_일자 = 공용.F2시각_바이트(g.High1YearDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M52주_저가 = 공용.F2정수64_바이트(g.Low1Year[:])
	s.M52주_저가_일자 = 공용.F2시각_바이트(g.Low1YearDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M유동_주식수 = 공용.F2정수64_바이트(g.FloatVolume[:])
	s.M상장_주식_수량_1000주_단위 = 공용.F2정수64_바이트(g.ListVolBy1000[:])
	s.M시가_총액 = 공용.F2정수64_바이트(g.MarketCapital[:])
	s.M거래원_정보_수신_시간 = 공용.F2시각_바이트(g.TraderInfoTime[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M매도_거래원_1 = 공용.F2문자열(g.Seller1[:])
	s.M매수_거래원_1 = 공용.F2문자열(g.Buyer1[:])
	s.M매도_거래량_1 = 공용.F2정수64_바이트(g.Seller1Volume[:])
	s.M매수_거래량_1 = 공용.F2정수64_바이트(g.Buyer1Volume[:])
	s.M매도_거래원_2 = 공용.F2문자열(g.Seller2[:])
	s.M매수_거래원_2 = 공용.F2문자열(g.Buyer2[:])
	s.M매도_거래량_2 = 공용.F2정수64_바이트(g.Seller2Volume[:])
	s.M매수_거래량_2 = 공용.F2정수64_바이트(g.Buyer2Volume[:])
	s.M매도_거래원_3 = 공용.F2문자열(g.Seller3[:])
	s.M매수_거래원_3 = 공용.F2문자열(g.Buyer3[:])
	s.M매도_거래량_3 = 공용.F2정수64_바이트(g.Seller3Volume[:])
	s.M매수_거래량_3 = 공용.F2정수64_바이트(g.Buyer3Volume[:])
	s.M매도_거래원_4 = 공용.F2문자열(g.Seller4[:])
	s.M매수_거래원_4 = 공용.F2문자열(g.Buyer4[:])
	s.M매도_거래량_4 = 공용.F2정수64_바이트(g.Seller4Volume[:])
	s.M매수_거래량_4 = 공용.F2정수64_바이트(g.Buyer4Volume[:])
	s.M매도_거래원_5 = 공용.F2문자열(g.Seller5[:])
	s.M매수_거래원_5 = 공용.F2문자열(g.Buyer5[:])
	s.M매도_거래량_5 = 공용.F2정수64_바이트(g.Seller5Volume[:])
	s.M매수_거래량_5 = 공용.F2정수64_바이트(g.Buyer5Volume[:])
	s.M외국인_매도_거래량 = 공용.F2정수64_바이트(g.ForeignSellVolume[:])
	s.M외국인_매수_거래량 = 공용.F2정수64_바이트(g.ForeignBuyVolume[:])
	s.M외국인_시간 = 공용.F2시각_바이트(g.ForeignTime[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M외국인_지분율 = 공용.F2실수_바이트(g.ForeignHoldingRate[:])
	s.M결제일 = 공용.F2시각_바이트(g.SettleDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M신용잔고_퍼센트 = 공용.F2실수_바이트(g.DebtPercent[:])
	s.M유상_배정_기준일 = 공용.F2시각_바이트(g.RightsIssueDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M무상_배정_기준일 = 공용.F2시각_바이트(g.BonusIssueDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M유상_배정_비율 = 공용.F2실수_바이트(g.RightsIssueRate[:])
	s.M무상_배정_비율 = 공용.F2실수_바이트(g.BonusIssueRate[:])
	s.M외국인_변동주_수량 = 공용.F2정수64_바이트(g.ForeignFloatVol[:])
	s.M당일_자사주_신청_여부 = 공용.F2참거짓_바이트(g.TreasuryStock[:], "1", true)
	s.M상장일 = 공용.F2시각_바이트(g.IpoDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M대주주_지분율 = 공용.F2실수_바이트(g.MajorHoldRate[:])
	s.M대주주_지분율_정보_일자 = 공용.F2시각_바이트(g.MajorHoldInfoDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M네잎클로버_종목_여부 = 공용.F2참거짓_바이트(g.FourLeafClover[:], "1", true)
	s.M증거금_비율 = 공용.F2실수_바이트(g.MarginRate[:])
	s.M자본금 = 공용.F2정수64_바이트(g.Capital[:])
	s.M전체_거래원_매도_합계 = 공용.F2정수64_바이트(g.SellTotalSum[:])
	s.M전체_거래원_매수_합계 = 공용.F2정수64_바이트(g.BuyTotalSum[:])
	s.M종목명2 = 공용.F2문자열(g.Title2[:])
	s.M우회_상장_여부 = 공용.F2참거짓_바이트(g.BackdoorListing[:], "1", true)
	s.M유동주_회전율_2 = 공용.F2실수_바이트(g.FloatRate2[:])
	s.M코스피_구분_2 = 공용.F2문자열(g.Market2[:])
	s.M공여율_기준일 = 공용.F2시각_바이트(g.DebtTrDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M공여율 = 공용.F2실수_바이트(g.DebtTrPercent[:])
	s.PER = 공용.F2실수_바이트(g.PER[:])
	s.M종목별_신용_한도 = 공용.F2참거짓_바이트(g.DebtLimit[:], "1", true)
	s.M가중_평균_가격 = 공용.F2정수64_바이트(g.WeightAvgPrice[:])
	s.M상장_주식_수량 = 공용.F2정수64_바이트(g.ListedVolume[:])
	s.M추가_상장_주식_수량 = 공용.F2정수64_바이트(g.AddListing[:])
	s.M종목_코멘트 = 공용.F2문자열(g.Comment[:])
	s.M전일_거래량 = 공용.F2정수64_바이트(g.PrevVolume[:])
	s.M전일대비_등락부호 = g.VsPrevSign[0]
	s.M전일대비_등락폭 = 공용.F2정수64_바이트(g.VsPrevDiff[:])
	s.M연중_최고가 = 공용.F2정수64_바이트(g.High1Year2[:])
	s.M연중_최고가_일자 = 공용.F2시각_바이트(g.High1YearDate2[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M연중_최저가 = 공용.F2정수64_바이트(g.Low1Year2[:])
	s.M연중_최저가_일자 = 공용.F2시각_바이트(g.Low1YearDate2[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M외국인_보유_주식수 = 공용.F2정수64_바이트(g.ForeignHoldQty[:])
	s.M외국인_지분_한도 = 공용.F2실수_바이트(g.ForeignLmtPercent[:])
	s.M매매_수량_단위 = 공용.F2정수64_바이트(g.TrUnitVolume[:])
	s.M대량_매매_방향 = int8(공용.F2정수64_바이트(g.DarkPoolOfferBid[:])) // 0 = 해당없음 1 = 매도 2 = 매수
	s.M대량_매매_존재 = 공용.F2참거짓_바이트(g.DarkPoolExist[:], "1", true)
	
	return s
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

func New주식_현재가_조회_변동_거래량_자료(c *C.Tc1101OutBlock2) *S주식_현재가_조회_변동_거래량_자료 {
	g := (*Tc1101OutBlock2)(unsafe.Pointer(c))
	
	s := new(S주식_현재가_조회_변동_거래량_자료)
	s.M시간 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열")
	s.M현재가 = 공용.F2정수64_바이트(g.MarketPrice[:])
	s.M등락부호 = g.DiffSign[0]
	s.M등락폭 = 공용.F2정수64_바이트(g.Diff[:])
	s.M매도_호가 = 공용.F2정수64_바이트(g.OfferPrice[:])
	s.M매수_호가 = 공용.F2정수64_바이트(g.BidPrice[:])
	s.M변동_거래량 = 공용.F2정수64_바이트(g.DiffVolume[:])
	s.M거래량 = 공용.F2정수64_바이트(g.Volume[:])
	
	return s
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

func New주식_현재가_조회_종목_지표(c *C.char) *S주식_현재가_조회_종목_지표 {
	g := (*Tc1101OutBlock3)(unsafe.Pointer(c))
	
	s := new(S주식_현재가_조회_종목_지표)
	s.M동시_호가_구분 = 공용.F2문자열(g.SyncOfferBid[:]) // 0:동시호가 아님 1:동시호가 2:동시호가연장 3:시가범위연장 4:종가범위연장 5:배분개시 6:변동성 완화장치 발동
	s.M예상_체결가 = 공용.F2정수64_바이트(g.EstmPrice[:])
	s.M예상_체결부호 = g.EstmSign[0]
	s.M예상_등락폭 = 공용.F2정수64_바이트(g.EstmDiff[:])
	s.M예상_등락율 = 공용.F2실수_바이트(g.EstmDiffRate[:])
	s.M예상_체결수량 = 공용.F2정수64_바이트(g.EstmVol[:])
	s.ECN정보_유무 = 공용.F2참거짓_바이트(g.ECN_InfoExist[:], "1", true) // 우리나라에는 아직 ECN이 없는 것으로 알고 있음.
	s.ECN전일_종가 = 공용.F2정수64_바이트(g.ECN_PrevPrice[:])
	s.ECN등락_부호 = g.ECN_DiffSign[0]
	s.ECN등락폭 = 공용.F2정수64_바이트(g.ECN_Diff[:])
	s.ECN등락율 = 공용.F2실수_바이트(g.ECN_DiffRate[:])
	s.ECN체결_수량 = 공용.F2정수64_바이트(g.ECN_Volume[:])
	s.ECN대비_예상_체결_부호 = g.VsECN_EstmSign[0]
	s.ECN대비_예상_체결_등락폭 = 공용.F2정수64_바이트(g.VsECN_EstmDiff[:])
	s.ECN대비_예상_체결_등락율 = 공용.F2실수_바이트(g.VsECN_EstmDiffRate[:])
		
	return s
}

//----------------------------------------------------------------------//
// ETF 현재가 조회 (c1151)
//----------------------------------------------------------------------//
func New_ETF_현재가_조회_질의(종목_코드 string) C.Tc1151InBlock {
	c := Tc1151InBlock{}
	c.Lang[0] = ([]byte("K"))[0]

	if len(종목_코드) > len(c.Code) {
		종목_코드 = 종목_코드[:len(c.Code)]
		
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 허용 한도를 초과함. 허용 한도  = %v, 종목_코드 길이  = %v.",
			len(c.Code[:]), len(종목_코드))
		공용.F에러_출력(에러)
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

func New_ETF_현재가_조회_기본_자료(c *C.char) *S_ETF_현재가_조회_기본_자료 {
	g := (*Tc1151OutBlock)(unsafe.Pointer(c))
	
	s := new(S_ETF_현재가_조회_기본_자료)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M종목명 = 공용.F2문자열(g.Title[:])
	s.M현재가 = 공용.F2정수64_바이트(g.MarketPrice[:])
	s.M등락부호 = g.DiffSign[0]
	s.M등락폭 = 공용.F2정수64_바이트(g.Diff[:])
	s.M등락율 = 공용.F2실수_바이트(g.DiffRate[:])
	s.M매도_호가 = 공용.F2정수64_바이트(g.OfferPrice[:])
	s.M매수_호가 = 공용.F2정수64_바이트(g.BidPrice[:])
	s.M거래량 = 공용.F2정수64_바이트(g.Volume[:])
	s.M거래_비율 = 공용.F2실수_바이트(g.TrVolRate[:])
	s.M유동주_회전율 = 공용.F2실수_바이트(g.FloatVolRate[:])
	s.M거래_대금 = 공용.F2정수64_바이트(g.TrAmount[:])
	s.M상한가 = 공용.F2정수64_바이트(g.UpLmtPrice[:])
	s.M고가 = 공용.F2정수64_바이트(g.High[:])
	s.M시가 = 공용.F2정수64_바이트(g.Open[:])
	s.M시가_대비_부호 = g.VsOpenSign[0]
	s.M시가_대비_등락폭 = 공용.F2정수64_바이트(g.VsOpenDiff[:])
	s.M저가 = 공용.F2정수64_바이트(g.Low[:])
	s.M하한가 = 공용.F2정수64_바이트(g.LowLmtPrice[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열은 내용 확인 후 작성할 것.")
	s.M매도_호가_최우선 = 공용.F2정수64_바이트(g.OfferPrice1[:])
	s.M매도_호가_차선 = 공용.F2정수64_바이트(g.OfferPrice2[:])
	s.M매도_호가_차차선 = 공용.F2정수64_바이트(g.OfferPrice3[:])
	s.M매도_호가_4차선 = 공용.F2정수64_바이트(g.OfferPrice4[:])
	s.M매도_호가_5차선 = 공용.F2정수64_바이트(g.OfferPrice5[:])
	s.M매도_호가_6차선 = 공용.F2정수64_바이트(g.OfferPrice6[:])
	s.M매도_호가_7차선 = 공용.F2정수64_바이트(g.OfferPrice7[:])
	s.M매도_호가_8차선 = 공용.F2정수64_바이트(g.OfferPrice8[:])
	s.M매도_호가_9차선 = 공용.F2정수64_바이트(g.OfferPrice9[:])
	s.M매도_호가_10차선 = 공용.F2정수64_바이트(g.OfferPrice10[:])
	s.M매수_호가_최우선 = 공용.F2정수64_바이트(g.BidPrice1[:])
	s.M매수_호가_차선 = 공용.F2정수64_바이트(g.BidPrice2[:])
	s.M매수_호가_차차선 = 공용.F2정수64_바이트(g.BidPrice3[:])
	s.M매수_호가_4차선 = 공용.F2정수64_바이트(g.BidPrice4[:])
	s.M매수_호가_5차선 = 공용.F2정수64_바이트(g.BidPrice5[:])
	s.M매수_호가_6차선 = 공용.F2정수64_바이트(g.BidPrice6[:])
	s.M매수_호가_7차선 = 공용.F2정수64_바이트(g.BidPrice7[:])
	s.M매수_호가_8차선 = 공용.F2정수64_바이트(g.BidPrice8[:])
	s.M매수_호가_9차선 = 공용.F2정수64_바이트(g.BidPrice9[:])
	s.M매수_호가_10차선 = 공용.F2정수64_바이트(g.BidPrice10[:])
	s.M매도_잔량_최우선 = 공용.F2정수64_바이트(g.OfferVolume1[:])
	s.M매도_잔량_차선 = 공용.F2정수64_바이트(g.OfferVolume2[:])
	s.M매도_잔량_차차선 = 공용.F2정수64_바이트(g.OfferVolume3[:])
	s.M매도_잔량_4차선 = 공용.F2정수64_바이트(g.OfferVolume4[:])
	s.M매도_잔량_5차선 = 공용.F2정수64_바이트(g.OfferVolume5[:])
	s.M매도_잔량_6차선 = 공용.F2정수64_바이트(g.OfferVolume6[:])
	s.M매도_잔량_7차선 = 공용.F2정수64_바이트(g.OfferVolume7[:])
	s.M매도_잔량_8차선 = 공용.F2정수64_바이트(g.OfferVolume8[:])
	s.M매도_잔량_9차선 = 공용.F2정수64_바이트(g.OfferVolume9[:])
	s.M매도_잔량_10차선 = 공용.F2정수64_바이트(g.OfferVolume10[:])
	s.M매수_잔량_최우선 = 공용.F2정수64_바이트(g.BidVolume1[:])
	s.M매수_잔량_차선 = 공용.F2정수64_바이트(g.BidVolume2[:])
	s.M매수_잔량_차차선 = 공용.F2정수64_바이트(g.BidVolume3[:])
	s.M매수_잔량_4차선 = 공용.F2정수64_바이트(g.BidVolume4[:])
	s.M매수_잔량_5차선 = 공용.F2정수64_바이트(g.BidVolume5[:])
	s.M매수_잔량_6차선 = 공용.F2정수64_바이트(g.BidVolume6[:])
	s.M매수_잔량_7차선 = 공용.F2정수64_바이트(g.BidVolume7[:])
	s.M매수_잔량_8차선 = 공용.F2정수64_바이트(g.BidVolume8[:])
	s.M매수_잔량_9차선 = 공용.F2정수64_바이트(g.BidVolume9[:])
	s.M매수_잔량_10차선 = 공용.F2정수64_바이트(g.BidVolume10[:])
	s.M매도_잔량_총합 = 공용.F2정수64_바이트(g.OfferVolTot[:])
	s.M매수_잔량_총합 = 공용.F2정수64_바이트(g.BidVolTot[:])
	s.M시간외_매도_잔량 = 공용.F2정수64_바이트(g.OfferVolAfterHour[:])
	s.M시간외_매수_잔량 = 공용.F2정수64_바이트(g.BidVolAfterHour[:])
	s.M피봇_2차_저항 = 공용.F2정수64_바이트(g.PivotUp2[:])
	s.M피봇_1차_저항 = 공용.F2정수64_바이트(g.PivotUp1[:])
	s.M피봇_가격 = 공용.F2정수64_바이트(g.PivotPrice[:])
	s.M피봇_1차_지지 = 공용.F2정수64_바이트(g.PivotDown1[:])
	s.M피봇_2차_지지 = 공용.F2정수64_바이트(g.PivotDown2[:])
	s.M코스피_코스닥_구분 = 공용.F2문자열(g.Market[:])
	s.M업종명 = 공용.F2문자열(g.Sector[:])
	s.M자본금_규모 = 공용.F2문자열(g.CapSize[:])
	s.M결산월 = 공용.F2문자열(g.SettleMonth[:])
	s.M시장_조치_1 = 공용.F2문자열(g.MarketAction1[:])
	s.M시장_조치_2 = 공용.F2문자열(g.MarketAction2[:])
	s.M시장_조치_3 = 공용.F2문자열(g.MarketAction3[:])
	s.M시장_조치_4 = 공용.F2문자열(g.MarketAction4[:])
	s.M시장_조치_5 = 공용.F2문자열(g.MarketAction5[:])
	s.M시장_조치_6 = 공용.F2문자열(g.MarketAction6[:])
	s.M전환사채_구분 = 공용.F2문자열(g.ConvertBond[:])
	s.M액면가 = 공용.F2정수64_바이트(g.NominalPrice[:])
	s.M전일_종가_타이틀 = 공용.F2문자열(g.PrevPriceTitle[:])
	s.M전일_종가 = 공용.F2정수64_바이트(g.PrevPrice[:])
	s.M대용가 = 공용.F2정수64_바이트(g.MortgageValue[:])
	s.M공모가 = 공용.F2정수64_바이트(g.PublicOfferPrice[:])
	s.M5일_고가 = 공용.F2정수64_바이트(g.High5Day[:])
	s.M5일_저가 = 공용.F2정수64_바이트(g.Low5Day[:])
	s.M20일_고가 = 공용.F2정수64_바이트(g.High20Day[:])
	s.M20일_저가 = 공용.F2정수64_바이트(g.Low20Day[:])
	s.M52주_고가 = 공용.F2정수64_바이트(g.High1Year[:])
	s.M52주_고가_일자 = 공용.F2시각_바이트(g.High1YearDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M52주_저가 = 공용.F2정수64_바이트(g.Low1Year[:])
	s.M52주_저가_일자 = 공용.F2시각_바이트(g.Low1YearDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M유동_주식수 = 공용.F2정수64_바이트(g.FloatVolume[:])
	s.M상장_주식_수량_1000주_단위 = 공용.F2정수64_바이트(g.ListVolBy1000[:])
	s.M시가_총액 = 공용.F2정수64_바이트(g.MarketCapital[:])
	s.M거래원_정보_수신_시각 = 공용.F2시각_바이트(g.TraderInfoTime[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M매도_거래원_1 = 공용.F2문자열(g.Seller1[:])
	s.M매수_거래원_1 = 공용.F2문자열(g.Buyer1[:])
	s.M매도_거래원_1_거래량 = 공용.F2정수64_바이트(g.Seller1Volume[:])
	s.M매수_거래원_1_거래량 = 공용.F2정수64_바이트(g.Buyer1Volume[:])
	s.M매도_거래원_2 = 공용.F2문자열(g.Seller2[:])
	s.M매수_거래원_2 = 공용.F2문자열(g.Buyer2[:])
	s.M매도_거래원_2_거래량 = 공용.F2정수64_바이트(g.Seller2Volume[:])
	s.M매수_거래원_2_거래량 = 공용.F2정수64_바이트(g.Buyer2Volume[:])
	s.M매도_거래원_3 = 공용.F2문자열(g.Seller3[:])
	s.M매수_거래원_3 = 공용.F2문자열(g.Buyer3[:])
	s.M매도_거래원_3_거래량 = 공용.F2정수64_바이트(g.Seller3Volume[:])
	s.M매수_거래원_3_거래량 = 공용.F2정수64_바이트(g.Buyer3Volume[:])
	s.M매도_거래원_4 = 공용.F2문자열(g.Seller4[:])
	s.M매수_거래원_4 = 공용.F2문자열(g.Buyer4[:])
	s.M매도_거래원_4_거래량 = 공용.F2정수64_바이트(g.Seller4Volume[:])
	s.M매수_거래원_4_거래량 = 공용.F2정수64_바이트(g.Buyer4Volume[:])
	s.M매도_거래원_5 = 공용.F2문자열(g.Seller5[:])
	s.M매수_거래원_5 = 공용.F2문자열(g.Buyer5[:])
	s.M매도_거래원_5_거래량 = 공용.F2정수64_바이트(g.Seller5Volume[:])
	s.M매수_거래원_5_거래량 = 공용.F2정수64_바이트(g.Buyer5Volume[:])
	s.M외국인_매도_거래량 = 공용.F2정수64_바이트(g.ForeignSellVolume[:])
	s.M외국인_매수_거래량 = 공용.F2정수64_바이트(g.ForeignBuyVolume[:])
	s.M외국인_시간 = 공용.F2시각_바이트(g.ForeignTime[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M외국인_지분율 = 공용.F2실수_바이트(g.ForeignHoldingRate[:])
	s.M결제일 = 공용.F2시각_바이트(g.SettleDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M신용잔고_퍼센트 = 공용.F2실수_바이트(g.DebtPercent[:])
	s.M유상_배정_기준일 = 공용.F2시각_바이트(g.RightsIssueDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M무상_배정_기준일 = 공용.F2시각_바이트(g.BonusIssueDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M유상_배정_비율 = 공용.F2실수_바이트(g.RightsIssueRate[:])
	s.M무상_배정_비율 = 공용.F2실수_바이트(g.BonusIssueRate[:])
	s.M상장일 = 공용.F2시각_바이트(g.IpoDate[:], "포맷 문자열은 내용 확인 후 결정할 것.")
	s.M상장_주식_수량 = 공용.F2정수64_바이트(g.ListedVolume[:])
	s.M전체_거래원_매도_합계 = 공용.F2정수64_바이트(g.SellTotalSum[:])
	s.M전체_거래원_매수_합계 = 공용.F2정수64_바이트(g.BuyTotalSum[:])
	
	return s
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

func New_ETF_현재가_조회_변동_거래량(c *C.char) *S_ETF_현재가_조회_변동_거래량 {
	g := (*Tc1151OutBlock2)(unsafe.Pointer(c))
	
	s := new(S_ETF_현재가_조회_변동_거래량)
	s.M시간 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열은 내용 확인 후 작성할 것.")
	s.M현재가 = 공용.F2정수64_바이트(g.MarketPrice[:])
	s.M등락_부호 = g.DiffSign[0]
	s.M등락폭 = 공용.F2정수64_바이트(g.Diff[:])
	s.M매도_호가 = 공용.F2정수64_바이트(g.OfferPrice[:])
	s.M매수_호가 = 공용.F2정수64_바이트(g.BidPrice[:])
	s.M변동_거래량 = 공용.F2정수64_바이트(g.DiffVolume[:])
	s.M거래량 = 공용.F2정수64_바이트(g.Volume[:])
	
	return s
}

type S_ETF_현재가_조회_예상_체결 struct {
	M동시_호가_구분  string
	M예상_체결가    int64
	M예상_체결_부호  byte
	M예상_체결_등락폭 int64
	M예상_체결_등락율 float64
	M예상_체결_수량  int64
}

func New_ETF_현재가_조회_예상_체결(c *C.char) *S_ETF_현재가_조회_예상_체결 {
	g := (*Tc1151OutBlock3)(unsafe.Pointer(c))
	
	s := new(S_ETF_현재가_조회_예상_체결)
	s.M동시_호가_구분 = 공용.F2문자열(g.SyncOfferBid[:])
	s.M예상_체결가 = 공용.F2정수64_바이트(g.EstmPrice[:])
	s.M예상_체결_부호 = g.EstmSign[0]
	s.M예상_체결_등락폭 = 공용.F2정수64_바이트(g.EstmDiff[:])
	s.M예상_체결_등락율 = 공용.F2실수_바이트(g.EstmDiffRate[:])
	s.M예상_체결_수량 = 공용.F2정수64_바이트(g.EstmVolume[:])
	
	return s
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

func New_ETF_현재가_조회_ETF자료(c *C.char) *S_ETF_현재가_조회_ETF자료 {
	g := (*Tc1151OutBlock4)(unsafe.Pointer(c))
	
	s := new(S_ETF_현재가_조회_ETF자료)
	s.ETF구분 = 공용.F2문자열(g.ETF[:])
	s.NAV = 공용.F2실수_바이트(g.NAV[:])
	s.NAV_등락_부호 = g.DiffSign[0]
	s.NAV_등락폭 = 공용.F2실수_바이트(g.Diff[:])
	s.M전일NAV = 공용.F2실수_바이트(g.PrevNAV[:])
	s.M괴리율 = 공용.F2실수_바이트(g.DivergeRate[:])
	s.M괴리율_부호 = g.DivergeSign[0]
	s.M설정단위_당_현금_배당액 = 공용.F2정수64_바이트(g.DividendPerCU[:])
	s.M구성_종목수 = 공용.F2정수64_바이트(g.ConstituentNo[:])
	s.M순자산_총액_억원 = 공용.F2정수64_바이트(g.NAVBy100Million[:])
	s.M추적_오차율 = 공용.F2실수_바이트(g.TrackingErrRate[:])
	s.LP_매도_최우선_잔량 = 공용.F2정수64_바이트(g.LP_OfferVolume1[:])
	s.LP_매도_차선_잔량 = 공용.F2정수64_바이트(g.LP_OfferVolume2[:])
	s.LP_매도_차차선_잔량 = 공용.F2정수64_바이트(g.LP_OfferVolume3[:])
	s.LP_매도_4차선_잔량 = 공용.F2정수64_바이트(g.LP_OfferVolume4[:])
	s.LP_매도_5차선_잔량 = 공용.F2정수64_바이트(g.LP_OfferVolume5[:])
	s.LP_매도_6차선_잔량 = 공용.F2정수64_바이트(g.LP_OfferVolume6[:])
	s.LP_매도_7차선_잔량 = 공용.F2정수64_바이트(g.LP_OfferVolume7[:])
	s.LP_매도_8차선_잔량 = 공용.F2정수64_바이트(g.LP_OfferVolume8[:])
	s.LP_매도_9차선_잔량 = 공용.F2정수64_바이트(g.LP_OfferVolume9[:])
	s.LP_매도_10차선_잔량 = 공용.F2정수64_바이트(g.LP_OfferVolume10[:])
	s.LP_매수_최우선_잔량 = 공용.F2정수64_바이트(g.LP_BidVolume1[:])
	s.LP_매수_차선_잔량 = 공용.F2정수64_바이트(g.LP_BidVolume2[:])
	s.LP_매수_차차선_잔량 = 공용.F2정수64_바이트(g.LP_BidVolume3[:])
	s.LP_매수_4차선_잔량 = 공용.F2정수64_바이트(g.LP_BidVolume4[:])
	s.LP_매수_5차선_잔량 = 공용.F2정수64_바이트(g.LP_BidVolume5[:])
	s.LP_매수_6차선_잔량 = 공용.F2정수64_바이트(g.LP_BidVolume6[:])
	s.LP_매수_7차선_잔량 = 공용.F2정수64_바이트(g.LP_BidVolume7[:])
	s.LP_매수_8차선_잔량 = 공용.F2정수64_바이트(g.LP_BidVolume8[:])
	s.LP_매수_9차선_잔량 = 공용.F2정수64_바이트(g.LP_BidVolume9[:])
	s.LP_매수_10차선_잔량 = 공용.F2정수64_바이트(g.LP_BidVolume10[:])
	s.ETF_복제_방법_구분_코드 = 공용.F2문자열(g.TrackingMethod[:])
	s.ETF_상품_유형_코드 = 공용.F2문자열(g.ETF_Type[:])
	
	return s
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

func New_ETF_현재가_조회_기반_지수_자료(c *C.char) *S_ETF_현재가_조회_기반_지수_자료 {
	g := (*Tc1151OutBlock5)(unsafe.Pointer(c))
	
	s := new(S_ETF_현재가_조회_기반_지수_자료)
	s.M지수_코드 = 공용.F2문자열(g.IndexCode[:])
	s.M업종_코드 = 공용.F2문자열(g.SectorCode[:])
	s.M지수_이름 = 공용.F2문자열(g.IndexName[:])
	s.M코스피200 = 공용.F2실수_바이트(g.KP200Index[:])
	s.M코스피200_등락부호 = g.KP200Sign[0]
	s.M코스피200_등락폭 = 공용.F2실수_바이트(g.KP200Diff[:])
	s.M채권_지수 = 공용.F2실수_바이트(g.BondIndex[:])
	s.M채권_지수_등락부호 = g.BondDiff[0]
	s.M채권_지수_등락폭 = 공용.F2실수_바이트(g.BondDiff[:])
	s.M해외_지수_코드 = 공용.F2문자열(g.ForeignIndexSymbol[:])
	s.M기타_업종_코드 = 공용.F2문자열(g.EtcSectorCode[:])
	s.M채권_지수_코드 = 공용.F2문자열(g.BondIndexCode[:])
	s.M채권_지수_세부_코드 = 공용.F2문자열(g.BondDetailCode[:])
	
	return s
}

//----------------------------------------------------------------------//
// 코스피 호가 잔량 (h1)
//----------------------------------------------------------------------//
func New코스피_호가_잔량_질의(종목_코드 string) C.Th1InBlock {
	c := Th1InBlock{}

	if len(종목_코드) > len(c.Code) {
		종목_코드 = 종목_코드[:len(c.Code)]
		
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 허용 한도를 초과함. 허용 한도  = %v, 종목_코드 길이  = %v.",
			len(c.Code[:]), len(종목_코드))
		공용.F에러_출력(에러)
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

func New코스피_호가_잔량(c *C.char) *S코스피_호가_잔량 {
	g := (*Th1OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스피_호가_잔량)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열은 내용 확인 후 작성할 것.")
	s.M매도_호가 = 공용.F2정수64_바이트(g.OfferPrice1[:])
	s.M매수_호가 = 공용.F2정수64_바이트(g.BidPrice1[:])
	s.M매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume1[:])
	s.M매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume1[:])
	s.M차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice2[:])
	s.M차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice2[:])
	s.M차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume2[:])
	s.M차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume2[:])
	s.M차차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice3[:])
	s.M차차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice3[:])
	s.M차차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume3[:])
	s.M차차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume3[:])
	s.M4차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice4[:])
	s.M4차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice4[:])
	s.M4차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume4[:])
	s.M4차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume4[:])
	s.M5차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice5[:])
	s.M5차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice5[:])
	s.M5차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume5[:])
	s.M5차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume5[:])
	s.M6차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice6[:])
	s.M6차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice6[:])
	s.M6차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume6[:])
	s.M6차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume6[:])
	s.M7차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice7[:])
	s.M7차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice7[:])
	s.M7차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume7[:])
	s.M7차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume7[:])
	s.M8차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice8[:])
	s.M8차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice8[:])
	s.M8차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume8[:])
	s.M8차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume8[:])
	s.M9차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice9[:])
	s.M9차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice9[:])
	s.M9차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume9[:])
	s.M9차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume9[:])
	s.M10차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice10[:])
	s.M10차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice10[:])
	s.M10차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume10[:])
	s.M10차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume10[:])
	s.M누적_거래량 = 공용.F2정수64_바이트(g.Volume[:])
	
	return s
}

//----------------------------------------------------------------------//
// 코스닥 호가 잔량 (k3)
//----------------------------------------------------------------------//
func New코스닥_호가_잔량_질의(종목_코드 string) C.Tk3InBlock {
	c := Tk3InBlock{}

	if len(종목_코드) > len(c.Code) {
		종목_코드 = 종목_코드[:len(c.Code)]
		
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 허용 한도를 초과함. 허용 한도  = %v, 종목_코드 길이  = %v.",
			len(c.Code[:]), len(종목_코드))
		공용.F에러_출력(에러)
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

func New코스닥_호가_잔량(c *C.char) *S코스닥_호가_잔량 {
	g := (*Tk3OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스닥_호가_잔량)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열은 내용 확인 후 작성할 것.")
	s.M매도_호가 = 공용.F2정수64_바이트(g.OfferPrice1[:])
	s.M매수_호가 = 공용.F2정수64_바이트(g.BidPrice1[:])
	s.M매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume1[:])
	s.M매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume1[:])
	s.M차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice2[:])
	s.M차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice2[:])
	s.M차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume2[:])
	s.M차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume2[:])
	s.M차차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice3[:])
	s.M차차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice3[:])
	s.M차차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume3[:])
	s.M차차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume3[:])
	s.M4차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice4[:])
	s.M4차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice4[:])
	s.M4차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume4[:])
	s.M4차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume4[:])
	s.M5차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice5[:])
	s.M5차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice5[:])
	s.M5차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume5[:])
	s.M5차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume5[:])
	s.M6차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice6[:])
	s.M6차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice6[:])
	s.M6차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume6[:])
	s.M6차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume6[:])
	s.M7차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice7[:])
	s.M7차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice7[:])
	s.M7차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume7[:])
	s.M7차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume7[:])
	s.M8차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice8[:])
	s.M8차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice8[:])
	s.M8차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume8[:])
	s.M8차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume8[:])
	s.M9차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice9[:])
	s.M9차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice9[:])
	s.M9차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume9[:])
	s.M9차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume9[:])
	s.M10차선_매도_호가 = 공용.F2정수64_바이트(g.OfferPrice10[:])
	s.M10차선_매수_호가 = 공용.F2정수64_바이트(g.BidPrice10[:])
	s.M10차선_매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume10[:])
	s.M10차선_매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume10[:])
	s.M누적_거래량 = 공용.F2정수64_바이트(g.Volume[:])
	
	return s
}

//----------------------------------------------------------------------//
// 코스피 시간외 호가 잔량 (h2)
//----------------------------------------------------------------------//

func New코스피_시간외_호가_잔량_질의(종목_코드 string) C.Th2InBlock {
	c := Th2InBlock{}

	if len(종목_코드) > len(c.Code) {
		종목_코드 = 종목_코드[:len(c.Code)]
		
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 허용 한도를 초과함. 허용 한도  = %v, 종목_코드 길이  = %v.",
			len(c.Code[:]), len(종목_코드))
		공용.F에러_출력(에러)
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

func New코스피_시간외_호가_잔량(c *C.char) *S코스피_시간외_호가_잔량 {
	g := (*Th2OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스피_시간외_호가_잔량)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열")
	s.M매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume[:])
	s.M매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume[:])
	
	return s
}

//----------------------------------------------------------------------//
// 코스닥 시간외 호가 잔량 (k4)
//----------------------------------------------------------------------//

func New코스닥_시간외_호가_잔량_질의(종목_코드 string) C.Tk4InBlock {
	c := Tk4InBlock{}

	if len(종목_코드) > len(c.Code) {
		종목_코드 = 종목_코드[:len(c.Code)]
		
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 허용 한도를 초과함. 허용 한도  = %v, 종목_코드 길이  = %v.",
			len(c.Code[:]), len(종목_코드))
		공용.F에러_출력(에러)
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

func New코스닥_시간외_호가_잔량(c *C.char) *S코스닥_시간외_호가_잔량 {
	g := (*Tk4OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스닥_시간외_호가_잔량)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열은 내용 확인 후 작성할 것.")
	s.M매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume[:])
	s.M매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume[:])
	
	return s
}

//----------------------------------------------------------------------//
// 코스피 예상 호가 잔량 (h3)
//----------------------------------------------------------------------//

func New코스피_예상_호가_잔량_질의(종목_코드 string) C.Th3InBlock {
	c := Th3InBlock{}

	if len(종목_코드) > len(c.Code) {
		종목_코드 = 종목_코드[:len(c.Code)]
		
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 허용 한도를 초과함. 허용 한도  = %v, 종목_코드 길이  = %v.",
			len(c.Code[:]), len(종목_코드))
		공용.F에러_출력(에러)
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

func New코스피_예상_호가_잔량(c *C.char) *S코스피_예상_호가_잔량 {
	g := (*Th3OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스피_예상_호가_잔량)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열")
	s.M동시_호가_구분 = 공용.F2문자열(g.SyncOfferBid[:])
	s.M예상_체결가 = 공용.F2정수64_바이트(g.EstmPrice[:])
	s.M예상_등락_부호 = g.EstmDiffSign[0]
	s.M예상_등락폭 = 공용.F2정수64_바이트(g.EstmDiff[:])
	s.M예상_등락율 = 공용.F2실수_바이트(g.EstmDiffRate[:])
	s.M예상_체결_수량 = 공용.F2정수64_바이트(g.EstmVolume[:])
	s.M매도_호가 = 공용.F2정수64_바이트(g.OfferPrice[:])
	s.M매수_호가 = 공용.F2정수64_바이트(g.BidPrice[:])
	s.M매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume[:])
	s.M매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume[:])
	
	return s
}

//----------------------------------------------------------------------//
// 코스닥 예상 호가 잔량 (k5)
//----------------------------------------------------------------------//

func New코스닥_예상_호가_잔량_질의(종목_코드 string) C.Tk5InBlock {
	c := Tk5InBlock{}

	if len(종목_코드) > len(c.Code) {
		종목_코드 = 종목_코드[:len(c.Code)]
		
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 허용 한도를 초과함. 허용 한도  = %v, 종목_코드 길이  = %v.",
			len(c.Code[:]), len(종목_코드))
		공용.F에러_출력(에러)
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

func New코스닥_예상_호가_잔량(c *C.char) *S코스닥_예상_호가_잔량 {
	g := (*Tk5OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스닥_예상_호가_잔량)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열")
	s.M동시_호가_구분 = 공용.F2문자열(g.SyncOfferBid[:])
	s.M예상_체결가 = 공용.F2정수64_바이트(g.EstmPrice[:])
	s.M예상_등락_부호 = g.EstmDiffSign[0]
	s.M예상_등락폭 = 공용.F2정수64_바이트(g.EstmDiff[:])
	s.M예상_등락율 = 공용.F2실수_바이트(g.EstmDiffRate[:])
	s.M예상_체결_수량 = 공용.F2정수64_바이트(g.EstmVolume[:])
	s.M매도_호가 = 공용.F2정수64_바이트(g.OfferPrice[:])
	s.M매수_호가 = 공용.F2정수64_바이트(g.BidPrice[:])
	s.M매도_호가_잔량 = 공용.F2정수64_바이트(g.OfferVolume[:])
	s.M매수_호가_잔량 = 공용.F2정수64_바이트(g.BidVolume[:])
	
	return s
}

//----------------------------------------------------------------------//
// 코스피 체결 (j8)
//----------------------------------------------------------------------//

func New코스피_체결_질의(종목_코드 string) C.Tj8InBlock {
	c := Tj8InBlock{}

	if len(종목_코드) > len(c.Code) {
		종목_코드 = 종목_코드[:len(c.Code)]
		
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 허용 한도를 초과함. 허용 한도  = %v, 종목_코드 길이  = %v.",
			len(c.Code[:]), len(종목_코드))
		공용.F에러_출력(에러)
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

func New코스피_체결(c *C.char) *S코스피_체결 {
	g := (*Tj8OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스피_체결)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열은 내용 확인 후 작성할 것.")
	s.M등락_부호 = g.DiffSign[0]
	s.M등락폭 = 공용.F2정수64_바이트(g.Diff[:])
	s.M현재가 = 공용.F2정수64_바이트(g.MarketPrice[:])
	s.M등락율 = 공용.F2실수_바이트(g.DiffRate[:])
	s.M고가 = 공용.F2정수64_바이트(g.High[:])
	s.M저가 = 공용.F2정수64_바이트(g.Low[:])
	s.M매도_호가 = 공용.F2정수64_바이트(g.OfferPrice[:])
	s.M매수_호가 = 공용.F2정수64_바이트(g.BidPrice[:])
	s.M누적_거래량 = 공용.F2정수64_바이트(g.Volume[:])
	s.M전일_거래량_대비_비율 = 공용.F2실수_바이트(g.VsPrevVolRate[:])
	s.M변동_거래량 = 공용.F2정수64_바이트(g.DiffVolume[:])
	s.M거래_대금 = 공용.F2정수64_바이트(g.TrAmount[:])
	s.M시가 = 공용.F2정수64_바이트(g.Open[:])
	s.M가중_평균_가격 = 공용.F2정수64_바이트(g.WeightAvgPrice[:])
	s.M장구분 = 공용.F2문자열(g.Market[:])
	
	return s
}

//----------------------------------------------------------------------//
// 코스닥 체결 (k8)
//----------------------------------------------------------------------//

func New코스닥_체결_질의(종목_코드 string) C.Tk8InBlock {
	c := Tk8InBlock{}

	if len(종목_코드) > len(c.Code) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상  = %v, 실제  = %v.",
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

func New코스닥_체결(c *C.char) *S코스닥_체결 {
	g := (*Tj8OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스닥_체결)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열은 내용 확인 후 작성할 것.")
	s.M등락_부호 = g.DiffSign[0]
	s.M등락폭 = 공용.F2정수64_바이트(g.Diff[:])
	s.M현재가 = 공용.F2정수64_바이트(g.MarketPrice[:])
	s.M등락율 = 공용.F2실수_바이트(g.DiffRate[:])
	s.M고가 = 공용.F2정수64_바이트(g.High[:])
	s.M저가 = 공용.F2정수64_바이트(g.Low[:])
	s.M매도_호가 = 공용.F2정수64_바이트(g.OfferPrice[:])
	s.M매수_호가 = 공용.F2정수64_바이트(g.BidPrice[:])
	s.M누적_거래량 = 공용.F2정수64_바이트(g.Volume[:])
	s.M전일_거래량_대비_비율 = 공용.F2실수_바이트(g.VsPrevVolRate[:])
	s.M변동_거래량 = 공용.F2정수64_바이트(g.DiffVolume[:])
	s.M거래_대금 = 공용.F2정수64_바이트(g.TrAmount[:])
	s.M시가 = 공용.F2정수64_바이트(g.Open[:])
	s.M가중_평균_가격 = 공용.F2정수64_바이트(g.WeightAvgPrice[:])
	s.M장구분 = 공용.F2문자열(g.Market[:])
	
	return s
}

//----------------------------------------------------------------------//
// 코스피 ETF NAV (j1) (예제코드가 없음. 패딩 필드가 필요한 지 추가 확인 필요함.)
//----------------------------------------------------------------------//

type S코스피_ETF_NAV struct {
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

func New코스피_ETF(c *C.char) *S코스피_ETF_NAV {
	g := (*Tj1OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스피_ETF_NAV)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열")
	s.M등락_부호 = g.DiffSign[0]
	s.M등락폭 = 공용.F2실수_바이트(g.Diff[:])
	s.NAV_현재가 = 공용.F2실수_바이트(g.Current[:])
	s.NAV_시가 = 공용.F2실수_바이트(g.Open[:])
	s.NAV_고가 = 공용.F2실수_바이트(g.High[:])
	s.NAV_저가 = 공용.F2실수_바이트(g.Low[:])
	s.M추적_오차_부호 = g.TrackErrSign[0]
	s.M추적_오차 = 공용.F2실수_바이트(g.TrackingError[:])
	s.M괴리율_부호 = g.DivergeSign[0]
	s.M괴리율 = 공용.F2실수_바이트(g.DivergeRate[:])
	
	return s
}

//----------------------------------------------------------------------//
// 코스닥 ETF NAV (j0) (예제코드가 없음. 패딩 필드가 필요한 지 추가 확인 필요함.)
//----------------------------------------------------------------------//

type S코스닥_ETF_NAV struct {
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

func New코스닥_ETF(c *C.char) *S코스닥_ETF_NAV {
	g := (*Tj0OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스닥_ETF_NAV)
	s.M종목_코드 = 공용.F2문자열(g.Code[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열")
	s.M등락_부호 = g.DiffSign[0]
	s.M등락폭 = 공용.F2실수_바이트(g.Diff[:])
	s.NAV_현재가 = 공용.F2실수_바이트(g.Current[:])
	s.NAV_시가 = 공용.F2실수_바이트(g.Open[:])
	s.NAV_고가 = 공용.F2실수_바이트(g.High[:])
	s.NAV_저가 = 공용.F2실수_바이트(g.Low[:])
	s.M추적_오차_부호 = g.TrackErrSign[0]
	s.M추적_오차 = 공용.F2실수_바이트(g.TrackingError[:])
	s.M괴리율_부호 = g.DivergeSign[0]
	s.M괴리율 = 공용.F2실수_바이트(g.DivergeRate[:])
	
	return s
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
			"업종 코드 길이가 예상보다 긺. 예상  = %v, 실제  = %v.",
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

func New코스피_업종_지수(c *C.char) *S코스피_업종_지수 {
	g := (*Tu1OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스피_업종_지수)
	s.M업종_코드 = 공용.F2문자열(g.SectorCode[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열")
	s.M지수값 = 공용.F2실수_바이트(g.IndexValue[:])
	s.M등락_부호 = g.DiffSign[0]
	s.M등락폭 = 공용.F2실수_바이트(g.Diff[:])
	s.M거래량 = 공용.F2정수64_바이트(g.Volume[:])
	s.M거래_대금 = 공용.F2정수64_바이트(g.TrAmount[:])
	s.M개장값 = 공용.F2실수_바이트(g.Open[:])
	s.M최고값 = 공용.F2실수_바이트(g.High[:])
	s.M최고값_시각 = 공용.F2시각_바이트(g.HighTime[:], "포맷 문자열")
	s.M최저값 = 공용.F2실수_바이트(g.Low[:])
	s.M최저값_시간 = 공용.F2시각_바이트(g.LowTime[:], "포맷 문자열")
	s.M지수_등락율 = 공용.F2실수_바이트(g.DiffRate[:])
	s.M거래_비중 = 공용.F2실수_바이트(g.TrVolRate[:])
	
	return s
}

//----------------------------------------------------------------------//
// 코스닥 업종 지수 (k1)
//----------------------------------------------------------------------//

func New코스닥_업종_지수_질의(업종_코드 string) C.Tk1InBlock {
	c := Tk1InBlock{}

	if len(업종_코드) > len(c.SectorCode) {
		에러 := 공용.F에러_생성(
			"종목 코드 길이가 예상보다 긺. 예상  = %v, 실제  = %v.",
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

func New코스닥_업종_지수(c *C.char) *S코스닥_업종_지수 {
	g := (*Tk1OutBlock)(unsafe.Pointer(c))
	
	s := new(S코스닥_업종_지수)
	s.M업종_코드 = 공용.F2문자열(g.SectorCode[:])
	s.M시각 = 공용.F2시각_바이트(g.Time[:], "포맷 문자열")
	s.M지수값 = 공용.F2실수_바이트(g.IndexValue[:])
	s.M등락_부호 = g.DiffSign[0]
	s.M등락폭 = 공용.F2실수_바이트(g.Diff[:])
	s.M거래량 = 공용.F2정수64_바이트(g.Volume[:])
	s.M거래_대금 = 공용.F2정수64_바이트(g.TrAmount[:])
	s.M개장값 = 공용.F2실수_바이트(g.Open[:])
	s.M최고값 = 공용.F2실수_바이트(g.High[:])
	s.M최고값_시각 = 공용.F2시각_바이트(g.HighTime[:], "포맷 문자열")
	s.M최저값 = 공용.F2실수_바이트(g.Low[:])
	s.M최저값_시간 = 공용.F2시각_바이트(g.LowTime[:], "포맷 문자열")
	s.M지수_등락율 = 공용.F2실수_바이트(g.DiffRate[:])
	s.M거래_비중 = 공용.F2실수_바이트(g.TrVolRate[:])
	
	return s
}