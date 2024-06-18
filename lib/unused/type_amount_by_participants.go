package unused

import (
	"bytes"
	"context"
	"database/sql"
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/trade"
	"time"
)

func New종목별_매매주체_동향(종목코드 string, 일자 time.Time, 기관, 외국인, 개인 float64) *S종목별_매매주체_동향 {
	if len(종목코드) != 6 {
		panic(lib.New에러("예상과 다른 종목코드 길이 : '%v' '%v'", 종목코드, len(종목코드)))
	} else if 일자 = lib.F2일자(일자); 일자.Before(lib.F지금().AddDate(-40, 0, 0)) {
		panic(lib.New에러("너무 오래된 일자 : '%v'", 일자.Format(lib.P일자_형식)))
	}

	return &S종목별_매매주체_동향{
		M종목코드:     종목코드,
		M일자:       lib.F일자2정수(일자),
		M기관_순매수액:  기관,
		M외국인_순매수액: 외국인,
		M개인_순매수액:  개인}
}

type S종목별_매매주체_동향 struct {
	M종목코드     string
	M일자       uint32
	M기관_순매수액  float64
	M외국인_순매수액 float64
	M개인_순매수액  float64
}

func (s *S종목별_매매주체_동향) G합계() float64 {
	return s.M기관_순매수액 + s.M외국인_순매수액 + s.M개인_순매수액
}

func F종목별_매매주체_동향_DB읽기(db *sql.DB, 종목코드 string, 시작일 time.Time) (값_모음 []*S종목별_매매주체_동향, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { 값_모음 = nil }}.S실행()

	종목코드 = trade.F종목코드_보정(종목코드)
	lib.F확인1(F매매주체_동향_정보_테이블_생성(db))

	SQL := new(bytes.Buffer)
	SQL.WriteString("SELECT")
	SQL.WriteString(" code,")
	SQL.WriteString(" date,")
	SQL.WriteString(" institution,")
	SQL.WriteString(" foreigner,")
	SQL.WriteString(" individual ")
	SQL.WriteString("FROM amount_by_participants ")
	SQL.WriteString("WHERE code=?")
	SQL.WriteString(" AND date>=? ")
	SQL.WriteString(" AND (institution!=0 OR foreigner!=0 OR individual!=0) ")
	SQL.WriteString("ORDER BY date")

	stmt := lib.F확인2(db.Prepare(SQL.String()))
	defer stmt.Close()

	rows := lib.F확인2(stmt.Query(종목코드, 시작일))
	defer rows.Close()

	값_모음 = make([]*S종목별_매매주체_동향, 0)

	주말 := lib.F금일().Weekday() == time.Saturday || lib.F금일().Weekday() == time.Sunday
	금일 := lib.F일자2정수(lib.F금일())

	var 일자 time.Time

	for rows.Next() {
		값 := new(S종목별_매매주체_동향)

		lib.F확인1(rows.Scan(
			&값.M종목코드,
			&일자,
			&값.M기관_순매수액,
			&값.M외국인_순매수액,
			&값.M개인_순매수액))

		값.M일자 = lib.F일자2정수(일자)

		if 주말 && 값.M일자 == 금일 {
			continue // 주말 데이터 수집 중 발생할 수 있는 오류 건너뜀.
		} else if 값.M기관_순매수액 == 0 && 값.M외국인_순매수액 == 0 && 값.M개인_순매수액 == 0 {
			continue // 값이 비어있는 데이터 제외.
		}

		값_모음 = append(값_모음, 값)
	}

	return 값_모음, nil
}

func F종목별_매매주체_동향_모음_DB저장(db *sql.DB, 값_모음 []*S종목별_매매주체_동향) (에러 error) {
	if len(값_모음) == 0 {
		return nil
	}

	var tx *sql.Tx
	defer lib.S예외처리{M에러: &에러, M함수: func() {
		lib.F에러_출력(에러)

		if tx != nil {
			tx.Rollback()
		}
	}}.S실행()

	lib.F확인1(F매매주체_동향_정보_테이블_생성(db))

	SQL생성 := new(bytes.Buffer)
	SQL생성.WriteString("INSERT IGNORE INTO amount_by_participants (")
	SQL생성.WriteString("  code,")
	SQL생성.WriteString("  date,")
	SQL생성.WriteString("  institution,")
	SQL생성.WriteString("  foreigner,")
	SQL생성.WriteString("  individual")
	SQL생성.WriteString(") VALUES (?,?,0,0,0)")

	SQL수정 := new(bytes.Buffer)
	SQL수정.WriteString("UPDATE amount_by_participants SET")
	SQL수정.WriteString("  institution=?,")
	SQL수정.WriteString("  foreigner=?,")
	SQL수정.WriteString("  individual=? ")
	SQL수정.WriteString("WHERE code=? AND date=?")

	txOpts := new(sql.TxOptions)
	txOpts.Isolation = sql.LevelDefault
	txOpts.ReadOnly = false

	tx = lib.F확인2(db.BeginTx(context.TODO(), txOpts))

	stmt생성 := lib.F확인2(tx.Prepare(SQL생성.String()))
	defer stmt생성.Close()

	stmt수정 := lib.F확인2(tx.Prepare(SQL수정.String()))
	defer stmt수정.Close()

	for _, 값 := range 값_모음 {
		if 값 == nil || (값.M기관_순매수액 == 0 && 값.M외국인_순매수액 == 0 && 값.M개인_순매수액 == 0) {
			continue // 오류 발생한 데이터 무시.
		}

		lib.F확인2(stmt생성.Exec(값.M종목코드, 값.M일자))
		lib.F확인2(stmt수정.Exec(값.M기관_순매수액, 값.M외국인_순매수액, 값.M개인_순매수액, 값.M종목코드, 값.M일자))
	}

	return tx.Commit()
}

func F매매주체_동향_정보_테이블_생성(db *sql.DB) error {
	SQL := new(bytes.Buffer)
	SQL.WriteString("CREATE TABLE IF NOT EXISTS amount_by_participants (")
	SQL.WriteString("code CHAR(8) NOT NULL,")
	SQL.WriteString("date DATE NOT NULL,")
	SQL.WriteString("institution DOUBLE  NOT NULL,")
	SQL.WriteString("foreigner DOUBLE  NOT NULL,")
	SQL.WriteString("individual DOUBLE  NOT NULL,")
	SQL.WriteString("PRIMARY KEY (code,date)")
	SQL.WriteString(")")

	_, 에러 := db.Exec(SQL.String())

	return 에러
}
