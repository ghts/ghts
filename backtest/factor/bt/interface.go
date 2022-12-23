package bt

import (
	"github.com/ghts/ghts/lib"
	dpd "github.com/ghts/ghts/lib/daily_price_data"
)

type I데이터_처리기 interface {
	G영업일_모음() []uint32
	G현재가(string) float64
	G현재가_맵([]string) map[string]float64
	ETF_ETN_종목_여부(string) bool
	G제외_종목_여부(string) bool
	S상장_주식_수량_확인([]string) []string
}

// 백테스트에서만 사용됨. 실제 매매에서는 현재가/'현재가 맵'으로 충분함.
type I데이터_처리기_백테스트_전용 interface {
	S준비(uint32, T가격_구분)
	G일자() uint32
	G가격(uint32, T가격_구분, string) float64
	G가격_맵(uint32, T가격_구분, []string) map[string]float64
	G기준일_이전_전종목_가격_맵(uint32) map[string]*dpd.S종목별_일일_가격정보_모음
}

// 재무 데이터에서 생성된 '종목별 팩터 데이터'를 정렬/필터하는 자료형 생성
type I팩터_데이터_처리기[T팩터 T팩터_데이터, T재무 T재무_데이터] interface {
	G필터_정렬_처리기() *S필터_정렬_처리기[T팩터]
}

type I포트폴리오 interface {
	G보유_종목_코드_모음() []string
	G리밸런싱_필요(string, lib.T리밸런싱_주기) bool
	G리밸런싱_기준_가격(종목코드 string) float64
	G리밸런싱_기준_수량(종목코드 string) int64
	G보유_수량_맵() (map[string]int64, error)
	S리밸런싱_실행(종목코드_모음 []string) error
	S손절(종목코드 string)
	S전체_손절()
	S익절(종목코드 string)
	S부분_익절(종목코드 string, 비율 float64)
}

type I포트폴리오_백테스트_전용 interface {
	S준비(uint32, T가격_구분)
	S일일_결산()
	G투자_성과_계산() (float64, float64, float64) // CAGR/MDD/Sharpe
}

type I전략_식별_정보 interface {
	G전략명() string
	G계좌번호() string
	G전략_식별_문자열() string
}

type I전략_인수[T팩터 T팩터_데이터, T재무 T재무_데이터] interface {
	I전략_식별_정보
	G리밸런싱_주기() lib.T리밸런싱_주기
	G종목_수량() int
	G복합_등급_계산_함수() func(*S필터_정렬_처리기[T팩터])
	G급등_종목_제외() bool
	G급락_종목_제외() bool
	G버퍼_퍼센트() float64 // 버퍼가 '0%'이면 버퍼룰 적용 안 함.
	G데이터_처리기() I데이터_처리기
	G팩터_데이터_처리기() I팩터_데이터_처리기[T팩터, T재무]
	G포트폴리오() I포트폴리오
}

type I전략_실행기 interface {
	G실행() error
}
