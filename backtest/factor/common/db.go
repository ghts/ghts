package bfc //bt_factor_common

import (
	"bytes"
	"database/sql"
	"github.com/ghts/ghts/lib"
	"path"
)

func DB파일명() string {
	return path.Join(lib.F홈_디렉토리(), "backtest_factor.dat")
}

func F시가총액_테이블_생성(db *sql.DB) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	SQL := new(bytes.Buffer)
	SQL.WriteString("CREATE TABLE IF NOT EXISTS marcap (")
	SQL.WriteString("code CHAR(8) NOT NULL,")
	SQL.WriteString("date DATE NOT NULL,")
	SQL.WriteString("name VARCHAR(255) NOT NULL,")
	SQL.WriteString("market CHAR(3) NOT NULL,")
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

func F시가총액_테이블_삭제(db *sql.DB) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	_, 에러 = db.Exec("DROP TABLE IF EXISTS marcap")

	return 에러
}

func F세종_재무_테이블_생성(db *sql.DB) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	SQL := new(bytes.Buffer)
	SQL.WriteString("CREATE TABLE IF NOT EXISTS sejong (")
	SQL.WriteString("code CHAR(8) NOT NULL,")
	SQL.WriteString("year INT NOT NULL,")
	SQL.WriteString("quarter TINYINT NOT NULL,")
	SQL.WriteString("sales DOUBLE NOT NULL,")
	SQL.WriteString("operating_profit DOUBLE NOT NULL,")
	SQL.WriteString("net_profit DOUBLE NOT NULL,")
	SQL.WriteString("asset DOUBLE NOT NULL,")
	SQL.WriteString("capital DOUBLE NOT NULL,")
	SQL.WriteString("liability DOUBLE NOT NULL,")
	SQL.WriteString("CONSTRAINT pk_sejong PRIMARY KEY (code,year,quarter)")
	SQL.WriteString(")")

	_, 에러 = db.Exec(SQL.String())

	return 에러
}

func F세종_재무_테이블_삭제(db *sql.DB) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	_, 에러 = db.Exec("DROP TABLE IF EXISTS sejong")

	return 에러
}
