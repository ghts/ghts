package daily_price

import (
	"path/filepath"
	"testing"

	lb "github.com/ghts/ghts/lib"
	ldb "github.com/ghts/ghts/lib/db"
)

// TestSQLite_테이블_생성N저장N읽기 :
// modernc.org/sqlite 기준 테이블 생성 → 신규 INSERT → 기존 UPDATE → 읽기 전체 경로 검증.
func TestSQLite_테이블_생성N저장N읽기(t *testing.T) {
	dbSQLite, 에러 := ldb.DB_SQLite(ldb.DSN_SQLite(":memory:"))
	lb.F테스트_에러없음(t, 에러)
	defer dbSQLite.Close()

	lb.F테스트_에러없음(t, F일일_가격정보_테이블_생성(dbSQLite))

	모음, 에러 := New종목별_일일_가격정보_모음([]*S일일_가격정보{
		f테스트용_일일_가격정보("000000", -2, 1000),
	})
	lb.F테스트_에러없음(t, 에러)

	// 1) 신규 INSERT 경로 검증
	lb.F테스트_에러없음(t, 모음.DB저장(dbSQLite))

	// 2) 기존 UPDATE 경로 검증 (동일 데이터 재저장)
	lb.F테스트_에러없음(t, 모음.DB저장(dbSQLite))

	// 3) 신규 레코드 추가 후 다시 저장
	모음2, 에러 := New종목별_일일_가격정보_모음([]*S일일_가격정보{
		f테스트용_일일_가격정보("000000", -1, 1100),
	})
	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_에러없음(t, 모음2.DB저장(dbSQLite))

	// 4) 읽기 검증
	읽어온_모음, 에러 := New종목별_일일_가격정보_모음_DB읽기(dbSQLite, "000000")
	lb.F확인1(에러)

	lb.F테스트_참임(t, len(읽어온_모음.M저장소) == 2, "예상과 다른 레코드 수 : '%v'", len(읽어온_모음.M저장소))
}

// TestDSN_SQLite_WAL_기본_적용 : DSN_SQLite 기본값으로 파일 기반 DB가 WAL 모드인지 검증.
// (인메모리 DB는 WAL을 적용할 수 없어 파일 DB 사용)
func TestDSN_SQLite_WAL_기본_적용(t *testing.T) {
	파일경로 := filepath.Join(t.TempDir(), "테스트.db")
	dbSQLite, 에러 := ldb.DB_SQLite(ldb.DSN_SQLite(파일경로))
	lb.F테스트_에러없음(t, 에러)
	defer dbSQLite.Close()

	var 저널_모드 string
	lb.F확인1(dbSQLite.QueryRow("PRAGMA journal_mode").Scan(&저널_모드))
	lb.F테스트_참임(t, 저널_모드 == "wal", "예상과 다른 저널 모드 : '%v'", 저널_모드)
}

func f테스트용_일일_가격정보(종목코드 string, offset_일수 int, 종가 float64) *S일일_가격정보 {
	s := new(S일일_가격정보)
	s.M종목코드 = 종목코드
	s.M일자 = lb.F일자2정수(lb.F금일().AddDate(0, 0, offset_일수))
	s.M시가 = 종가
	s.M고가 = 종가 + 1
	s.M저가 = 종가 - 1
	s.M종가 = 종가
	s.M거래량 = 100

	return s
}
