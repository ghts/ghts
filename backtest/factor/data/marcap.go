package data

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	bfc "github.com/ghts/ghts/backtest/factor/common"
	"github.com/ghts/ghts/lib"
	_ "modernc.org/sqlite"
	"path"
)

func F시총_데이터_저장() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	레코드_모음 := make([][]string, 0)

	for 연도:=1995;연도<=2022;연도++ {
		파일명 := lib.F2문자열("marcap-%v.csv", 연도)
		연도별_레코드_모음 := lib.F확인2(csv파일_읽기(파일명))
		레코드_모음 = append(레코드_모음, 연도별_레코드_모음...)
		fmt.Printf("'%v', %v, %v\n", 파일명, lib.F정수_쉼표_추가(len(연도별_레코드_모음)), lib.F정수_쉼표_추가(len(레코드_모음)))
	}

	db := lib.F확인2(sql.Open("sqlite", bfc.DB파일명()))
	//lib.F확인1(F시가총액_테이블_삭제(db))
	lib.F확인1(F시가총액_테이블_생성(db))
	lib.F확인1(db저장(db, 레코드_모음))

	return nil
}

func csv파일_읽기(파일명 string) (레코드_모음 [][]string, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 레코드_모음 = nil }}.S실행()

	파일경로 := lib.F확인2(lib.F파일_검색(path.Join(lib.F홈_디렉토리(), "Downloads"), 파일명))
	레코드_모음 = lib.F확인2(lib.CSV읽기(파일경로, ",", nil))

	return 레코드_모음[1:], nil
}

func db저장(db *sql.DB, 레코드_모음 [][]string) (에러 error) {
	var tx *sql.Tx
	defer lib.S예외처리{M에러: &에러, M함수: func() {
		lib.F에러_출력(에러)

		if tx != nil {
			tx.Rollback()
		}
	}}.S실행()

	SQL := new(bytes.Buffer)
	SQL.WriteString("REPLACE INTO marcap (")
	SQL.WriteString("code,")
	SQL.WriteString("date,")
	SQL.WriteString("name,")
	SQL.WriteString("market,")
	SQL.WriteString("open,")
	SQL.WriteString("high,")
	SQL.WriteString("low,")
	SQL.WriteString("close,")
	SQL.WriteString("volume,")
	SQL.WriteString("amount,")
	SQL.WriteString("listed_qty,")
	SQL.WriteString("cap_mil,")
	SQL.WriteString("cap_rank) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)")

	txOpts := new(sql.TxOptions)
	txOpts.Isolation = sql.LevelDefault
	txOpts.ReadOnly = false

	tx = lib.F확인2(db.BeginTx(context.TODO(), txOpts))

	stmt := lib.F확인2(tx.Prepare(SQL.String()))
	defer stmt.Close()

	for i, 레코드 := range 레코드_모음 {
		종목코드 := 레코드[0]
		종목명 := 레코드[1]
		시장구분 := 레코드[2]
		//Dept := 레코드[3]	// 사용되지 않는 듯.
		종가 := lib.F확인2(lib.F2실수(레코드[4]))
		//ChangeCode := 레코드[5]
		//Changes := 레코드[6]
		//ChangeRatio := 레코드[7]
		시가 := lib.F확인2(lib.F2실수(레코드[8]))
		고가 := lib.F확인2(lib.F2실수(레코드[9]))
		저가 := lib.F확인2(lib.F2실수(레코드[10]))
		거래량 := lib.F확인2(lib.F2실수(레코드[11]))
		거래금액 := lib.F확인2(lib.F2실수(레코드[12]))
		시가총액_백만 := lib.F확인2(lib.F2실수(레코드[13]))
		상장주식수량 := lib.F확인2(lib.F2정수64(레코드[14]))
		//MarketId := 레코드[15]
		시총순위 := lib.F확인2(lib.F2정수(레코드[16]))
		일자 := 레코드[17]

		switch 시장구분 {
		case "KOSPI":
			시장구분 = "KSP"
		case "KOSDAQ", "KOSDAQ GLOBAL":
			시장구분 = "KSD"
		case "KONEX":
			시장구분 = "KNX"
		default:
			panic(lib.F2문자열("예상하지 못한 시장구분값 : '%v'", 시장구분))
		}

		lib.F확인2(stmt.Exec(종목코드,
			일자,
			종목명,
			시장구분,
			시가,
			고가,
			저가,
			종가,
			거래량,
			거래금액,
			상장주식수량,
			시가총액_백만,
			시총순위))

		if i>0 && i%100000 == 0 {
			lib.F체크포인트(lib.F정수_쉼표_추가(i))
		}
	}

	return tx.Commit()
}



func F시가총액_테이블_삭제(db *sql.DB) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	_, 에러 = db.Exec("DROP TABLE IF EXISTS marcap")

	return 에러
}

func F시가총액_테이블_생성(db *sql.DB) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	SQL := new(bytes.Buffer)
	SQL.WriteString("CREATE TABLE IF NOT EXISTS marcap (")
	SQL.WriteString("code CHAR(8) NOT NULL,")
	SQL.WriteString("date DATE NOT NULL,")
	SQL.WriteString("name VARCHAR(255) NOT NULL,")
	SQL.WriteString("market VARCHAR(255) NOT NULL,")
	SQL.WriteString("open DECIMAL(20,3) NOT NULL,")
	SQL.WriteString("high DECIMAL(20,3) NOT NULL,")
	SQL.WriteString("low DECIMAL(20,3) NOT NULL,")
	SQL.WriteString("close DECIMAL(20,3) NOT NULL,")
	SQL.WriteString("volume BIGINT NOT NULL,")
	SQL.WriteString("amount BIGINT NOT NULL,")
	SQL.WriteString("listed_qty BIGINT NOT NULL,")
	SQL.WriteString("cap_mil BIGINT NOT NULL,")
	SQL.WriteString("cap_rank INT NOT NULL,")
	SQL.WriteString("CONSTRAINT pk_marcap PRIMARY KEY (code,date)")
	SQL.WriteString(")")

	_, 에러 = db.Exec(SQL.String())

	return 에러
}