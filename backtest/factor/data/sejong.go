package data

import (
	"bufio"
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	bfc "github.com/ghts/ghts/backtest/factor/common"
	"github.com/ghts/ghts/lib"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func F세종_데이터_저장() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	맵 := make(map[string]*bfc.S재무_세종)

	for _, 구분 := range []T세종_데이터_구분{P매출_세종, P영업이익_세종, P당기순이익_세종, P자산_세종, P자본_세종, P부채_세종} {
		파일명_모음 := CSV파일_검색_세종(path.Join(lib.F홈_디렉토리(), "Downloads"), 구분.G파일명_조각())
		lib.F조건부_패닉(len(파일명_모음) != 1, "%v 파일명_모음 수량이 1이 아님. : '%v'", 구분, len(파일명_모음))

		파일명 := 파일명_모음[0]
		레코드_모음 := lib.F확인2(csv파일_읽기_세종(파일명))
		lib.F조건부_패닉(len(레코드_모음) < 1000, "레코드_모음 수량 부족 : '%v'", len(레코드_모음))

		lib.F확인1(csv파일_해석_세종(맵, 구분, 레코드_모음))
	}

	for 키, 값 := range 맵 {
		if 값.M매출액 == 0 ||
			값.M영업이익 == 0 ||
			값.M당기순이익 == 0 ||
			값.M자산 == 0 ||
			값.M자본 == 0 ||
			값.M부채 == 0 {
			delete(맵, 키) // 누락 데이터 삭제
		}
	}

	db := lib.F확인2(sql.Open("sqlite", bfc.DB파일명()))
	defer db.Close()

	lib.F확인1(bfc.F세종_재무_테이블_삭제(db))
	lib.F확인1(bfc.F세종_재무_테이블_생성(db))
	lib.F확인1(db저장_세종(db, 맵))

	return nil
}

type T세종_데이터_구분 uint8

const (
	P매출_세종 T세종_데이터_구분 = iota
	P영업이익_세종
	P당기순이익_세종
	P자산_세종
	P자본_세종
	P부채_세종
)

func (v T세종_데이터_구분) String() string {
	switch v {
	case P매출_세종:
		return "매출"
	case P영업이익_세종:
		return "영업이익"
	case P당기순이익_세종:
		return "당기순이익"
	case P자산_세종:
		return "자산"
	case P자본_세종:
		return "자본"
	case P부채_세종:
		return "부채"
	default:
		panic(lib.New에러("예상하짐 못한 세종 구분값 : '%v'", int(v)))
	}
}

func (v T세종_데이터_구분) G파일명_조각() string {
	switch v {
	case P매출_세종:
		return "sales"
	case P영업이익_세종:
		return "bp"
	case P당기순이익_세종:
		return "np"
	case P자산_세종:
		return "asset"
	case P자본_세종:
		return "capital"
	case P부채_세종:
		return "liability"
	default:
		panic(lib.New에러("예상하짐 못한 세종 구분값 : '%v'", int(v)))
	}
}

type S세종_데이터_구분 struct {
	M구분     T세종_데이터_구분
	M파일명_조각 string
}

func F세종_데이터_구분_맵() map[T세종_데이터_구분]string {
	s := make(map[T세종_데이터_구분]string)
	s[P매출_세종] = P매출_세종.G파일명_조각()
	s[P영업이익_세종] = P영업이익_세종.G파일명_조각()
	s[P당기순이익_세종] = P당기순이익_세종.G파일명_조각()
	s[P자산_세종] = P자산_세종.G파일명_조각()
	s[P자본_세종] = P자본_세종.G파일명_조각()
	s[P부채_세종] = P부채_세종.G파일명_조각()

	return s
}

func CSV파일_검색_세종(검색_시작_디렉토리, 항목 string) []string {
	파일명_모음 := make([]string, 0)
	헤더 := lib.F2문자열("SejongData-%v-", 항목)

	filepath.WalkDir(검색_시작_디렉토리, func(s string, d fs.DirEntry, 에러 error) error {
		if 에러 != nil {
			return 에러
		}

		if strings.Contains(d.Name(), 헤더) &&
			strings.Contains(d.Name(), "csv") &&
			!strings.Contains(d.Name(), "~lock") {
			파일명_모음 = append(파일명_모음, s)
		}

		return nil
	})

	return 파일명_모음
}

func csv파일_읽기_세종(파일명 string) (레코드_모음 [][]string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 레코드_모음 = nil }}.S실행()

	파일 := lib.F확인2(os.Open(파일명))
	defer 파일.Close()

	스캐너 := bufio.NewScanner(파일)
	버퍼 := new(bytes.Buffer)

	for 스캐너.Scan() {
		문자열 := 스캐너.Text()

		if 문자열 == "" ||
			strings.HasPrefix(문자열, ",,") ||
			strings.HasPrefix(문자열, "종목번호,종목명") ||
			strings.HasPrefix(문자열, "기준,주재무제표") ||
			strings.HasPrefix(문자열, "단위,억원") ||
			strings.HasPrefix(문자열, "제공,세종기업데이터") ||
			strings.Contains(문자열, "www.sejongdata.com") {
			continue
		}

		문자열 = strings.Replace(문자열, `"=""`, `"`, -1)
		문자열 = strings.Replace(문자열, `"""`, `"`, -1)

		버퍼.WriteString(문자열)
		버퍼.WriteString("\n")
	}

	return csv.NewReader(버퍼).ReadAll()
}

func csv파일_해석_세종(맵 map[string]*bfc.S재무_세종, 구분 T세종_데이터_구분, 레코드_모음 [][]string) (에러 error) {
	for _, 레코드 := range 레코드_모음 {
		종목코드 := 레코드[0]

		for i := 2; i < len(레코드); i++ {
			연도 := uint16(1994 + i)
			값 := lib.F2실수_단순형_공백은_0(레코드[i])
			F세종_값_설정(맵, 종목코드, 연도, 0, 구분, 값)
		}
	}

	return nil
}

func F세종_값_설정(맵 map[string]*bfc.S재무_세종, 종목코드 string, 연도 uint16, 분기 uint8, 구분 T세종_데이터_구분, 값 float64) {
	키 := bfc.F재무_정보_키(종목코드, 연도, 분기)

	세종, 존재함 := 맵[키]
	if !존재함 || 세종 == nil {
		세종 = bfc.New재무_세종()
		세종.M종목코드 = 종목코드
		세종.M연도 = 연도
		세종.M분기 = 분기

		맵[키] = 세종
	}

	switch 구분 {
	case P매출_세종:
		세종.M매출액 = 값
	case P영업이익_세종:
		세종.M영업이익 = 값
	case P당기순이익_세종:
		세종.M당기순이익 = 값
	case P자산_세종:
		세종.M자산 = 값
	case P자본_세종:
		세종.M자본 = 값
	case P부채_세종:
		세종.M부채 = 값
	default:
		panic(lib.New에러("예상하지 못한 세종 구분값 : '%v'", int(구분)))
	}
}

func db저장_세종(db *sql.DB, 맵 map[string]*bfc.S재무_세종) (에러 error) {
	var tx *sql.Tx
	defer lib.S예외처리{M에러: &에러, M함수: func() {
		lib.F에러_출력(에러)

		if tx != nil {
			tx.Rollback()
		}
	}}.S실행()

	SQL := new(bytes.Buffer)
	SQL.WriteString("REPLACE INTO sejong (")
	SQL.WriteString("code,")
	SQL.WriteString("year,")
	SQL.WriteString("quarter,")
	SQL.WriteString("sales,")
	SQL.WriteString("operating_profit,")
	SQL.WriteString("net_profit,")
	SQL.WriteString("asset,")
	SQL.WriteString("capital,")
	SQL.WriteString("liability) VALUES (?,?,?,?,?,?,?,?,?)")

	txOpts := new(sql.TxOptions)
	txOpts.Isolation = sql.LevelDefault
	txOpts.ReadOnly = false

	tx = lib.F확인2(db.BeginTx(context.TODO(), txOpts))

	stmt := lib.F확인2(tx.Prepare(SQL.String()))
	defer stmt.Close()

	for _, 값 := range 맵 {
		stmt.Exec(
			값.M종목코드,
			값.M연도,
			값.M분기,
			값.M매출액,
			값.M영업이익,
			값.M당기순이익,
			값.M자산,
			값.M자본,
			값.M부채)
	}

	return tx.Commit()
}
