package lib

import (
	"context"
	"database/sql"
)

func SQL실행(db *sql.DB, sql문자열 string, 추가_인수 ...interface{}) (id int64, 에러 error) {
	var tx *sql.Tx

	defer S예외처리{M에러: &에러, M함수: func() {
		id = 0

		if tx != nil {
			tx.Rollback()
		}
	}}.S실행()

	txOpts := new(sql.TxOptions)
	txOpts.Isolation = sql.LevelDefault
	txOpts.ReadOnly = false

	if tx, 에러 = db.BeginTx(context.TODO(), txOpts); 에러 != nil {
		tx = nil
		return
	}

	stmt := F확인2(tx.Prepare(sql문자열))
	defer stmt.Close()

	결과 := F확인2(stmt.Exec(추가_인수...))
	id = F확인2(결과.LastInsertId())
	에러 = tx.Commit()

	return
}

func F정수값DB질의(db *sql.DB, sql문자열 string, 추가_인수 ...interface{}) (정수값 int64, 에러 error) {
	defer S예외처리{M에러: &에러}.S실행()

	stmt := F확인2(db.Prepare(sql문자열))
	defer stmt.Close()

	rows := F확인2(stmt.Query(추가_인수...))
	defer rows.Close()

	for rows.Next() {
		F확인1(rows.Scan(&정수값))
	}

	return 정수값, nil
}
