package lib

import (
	"context"
	"database/sql"
	"math/big"
)

func SQL실행(db *sql.DB, sql문자열 string, 추가_인수 ...interface{}) (id int64, 에러 error) {
	var tx *sql.Tx
	var stmt *sql.Stmt

	defer S예외처리{M에러: &에러, M에러_실행: func() {
		id = 0
		if tx != nil {
			tx.Rollback()
		}
	}, M항상_실행: func() {
		if stmt != nil {
			stmt.Close()
		}
	}}.S실행()

	txOpts := new(sql.TxOptions)
	txOpts.Isolation = sql.LevelDefault
	txOpts.ReadOnly = false

	if tx, 에러 = db.BeginTx(context.TODO(), txOpts); 에러 != nil {
		tx = nil
		return
	}

	stmt = F확인2(tx.Prepare(sql문자열))
	결과 := F확인2(stmt.Exec(추가_인수...))
	id = F확인2(결과.LastInsertId())
	에러 = tx.Commit()

	return
}

// DB질의() : 단일 행, 단일 열의 질의 결과를 지정된 자료형으로 반환합니다.
// 지원 자료형 : 정수/실수(T숫자), big.Int, big.Rat (값/포인터), string, bool.
// big.Int/big.Rat 은 INTEGER 열(정수)에서 변환됩니다.
// (database/sql 의 내장 Scan 이 INTEGER(int64) 를 지원하지 않아 int64 로 스캔 후 변환합니다.)
// 포인터 자료형(*big.Int, *big.Rat)은 매 호출마다 새 인스턴스를 반환합니다.
func DB질의[T T숫자 | big.Int | *big.Int | big.Rat | *big.Rat | string | bool](db *sql.DB, sql문자열 string, 추가_인수 ...interface{}) (값 T, 에러 error) {
	defer S예외처리{M에러: &에러}.S실행()

	rows := F확인2(db.Query(sql문자열, 추가_인수...))
	defer rows.Close()

	for rows.Next() {
		switch 스캔_대상 := any(&값).(type) {
		case *big.Int: // 내장 Scan 이 int64 를 지원하지 않으므로 int64 로 스캔 후 변환.
			정수값 := new(int64)
			F확인1(rows.Scan(정수값))
			스캔_대상.SetInt64(*정수값)
		case *big.Rat:
			정수값 := new(int64)
			F확인1(rows.Scan(정수값))
			스캔_대상.SetInt64(*정수값)
		case **big.Int:
			정수값 := new(int64)
			F확인1(rows.Scan(정수값))
			*스캔_대상 = new(big.Int)
			(*스캔_대상).SetInt64(*정수값)
		case **big.Rat:
			정수값 := new(int64)
			F확인1(rows.Scan(정수값))
			*스캔_대상 = new(big.Rat)
			(*스캔_대상).SetInt64(*정수값)
		default:
			F확인1(rows.Scan(&값))
		}

		return 값, nil
	}

	에러 = New에러("DB질의() : 질의 결과 존재하지 않습니다.")

	return
}
