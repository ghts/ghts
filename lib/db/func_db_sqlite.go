package db

import (
	"database/sql"
	"strings"

	lb "github.com/ghts/ghts/lib"
	_ "modernc.org/sqlite" // 실제로 사용하는 드라이버 — 항상 컴파일됨
)

// DSN_MySQL : MySQL DSN 문자열 생성.
// 드라이버에 의존하지 않는 순수 문자열 포매팅이라 mysql 태그와 무관하게 항상 사용 가능하다.
// 실제 연결(DB_MySQL)은 //go:build mysql인 func_db_mysql.go에서 진행한다.
func DSN_MySQL(address, username, password, dbname string) string {
	return lb.F2문자열("%v:%v@tcp(%v:3306)/%v?parseTime=true",
		username,
		password,
		address,
		dbname)
}

// DSN_SQLite : SQLite DSN 문자열(데이터베이스 파일 경로) 생성.
func DSN_SQLite(파일경로 string) string {
	pragma옵션 := "_journal_mode=WAL&_auto_vacuum=INCREMENTAL&_busy_timeout=5000"

	if strings.Contains(파일경로, "?") {
		return 파일경로 + "&" + pragma옵션
	}

	return 파일경로 + "?" + pragma옵션
}

// DB_SQLite : DSN(파일 경로)으로 SQLite 연결.
func DB_SQLite(DSN string) (*sql.DB, error) {
	db, 에러 := sql.Open("sqlite", DSN)
	if 에러 != nil {
		return nil, 에러
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(lb.P1시간)

	if 에러 = db.Ping(); 에러 != nil {
		return nil, 에러
	}

	return db, nil
}
