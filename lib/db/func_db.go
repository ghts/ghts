package db

import (
	"database/sql"
	"strings"

	lb "github.com/ghts/ghts/lib"
)

import (
	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
)

// DSN_MySQL : MySQL DSN 문자열 생성.
func DSN_MySQL(address, username, password, dbname string) string {
	return lb.F2문자열("%v:%v@tcp(%v:3306)/%v?parseTime=true",
		username,
		password,
		address,
		dbname)
}

// DB_MySQL : DSN으로 MySQL 연결.
func DB_MySQL(DSN string) (*sql.DB, error) {
	db, _ := sql.Open("mysql", DSN)

	if 에러 := db.Ping(); 에러 != nil {
		return nil, 에러
	}

	return db, nil
}

// DSN_SQLite : SQLite DSN 문자열(데이터베이스 파일 경로) 생성.
// 기본값으로 WAL(journal_mode)과 busy_timeout(5초)을 적용한다.
// (데이터 수집과 트레이딩 프로세스가 동시 실행될 때 "database is locked" 방지)
// 추가 옵션이 이미 포함되어 있으면(문자열에 '?' 존재) 뒤에 이어 붙인다.
func DSN_SQLite(파일경로 string) string {
	pragma옵션 := "_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)"

	if strings.Contains(파일경로, "?") {
		return 파일경로 + "&" + pragma옵션
	}

	return 파일경로 + "?" + pragma옵션
}

// DB_SQLite : DSN(파일 경로)으로 SQLite 연결.
func DB_SQLite(DSN string) (*sql.DB, error) {
	db, _ := sql.Open("sqlite", DSN)

	if 에러 := db.Ping(); 에러 != nil {
		return nil, 에러
	}

	return db, nil
}
