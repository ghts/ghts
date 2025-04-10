package mysql

import (
	"database/sql"
	lb "github.com/ghts/ghts/lib"
)
import _ "github.com/go-sql-driver/mysql"

func DSN_MySQL(address, username, password, dbname string) string {
	return lb.F2문자열("%v:%v@tcp(%v:3306)/%v?parseTime=true",
		username,
		password,
		address,
		dbname)
}

func DB_MySQL(DSN string) (*sql.DB, error) {
	db, _ := sql.Open("mysql", DSN)

	if 에러 := db.Ping(); 에러 != nil {
		return nil, 에러
	}

	return db, nil
}
