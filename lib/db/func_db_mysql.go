//go:build mysql

package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // 이 드라이버는 //go:build mysql 태그로만 컴파일되어 바이너리에 포함된다.
)

// DB_MySQL : DSN으로 MySQL 연결.
//
// mysql 패키지는 //go:build mysql 태그가 붙어야만 컴파일되므로,
// 일반 빌드(`go build ./...`, -tags 미지정)에서는 드라이버가 바이너리에 포함되지 않는다.
// 향후 MySQL을 실제 사용할 때만 `go build -tags mysql ...`로 빌드한다.
func DB_MySQL(DSN string) (*sql.DB, error) {
	db, _ := sql.Open("mysql", DSN)

	if 에러 := db.Ping(); 에러 != nil {
		return nil, 에러
	}

	return db, nil
}
