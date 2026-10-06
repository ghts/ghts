package daily_price

import (
	"path/filepath"
	"testing"

	lb "github.com/ghts/ghts/lib"
	ldb "github.com/ghts/ghts/lib/db"
)

func TestSQLite_동일값_재저장(t *testing.T) {
	dbSQLite, 에러 := ldb.DB_SQLite(ldb.DSN_SQLite(":memory:"))
	lb.F테스트_에러없음(t, 에러)
	defer dbSQLite.Close()

	lb.F테스트_에러없음(t, F일일_가격정보_테이블_생성(dbSQLite))

	모음, 에러 := New종목별_일일_가격정보_모음([]*S일일_가격정보{
		f테스트용_일일_가격정보("000000", -2, 1000),
		f테스트용_일일_가격정보("000000", -1, 1100),
	})
	lb.F테스트_에러없음(t, 에러)

	// 1) 최초 저장 (신규 INSERT 경로)
	lb.F테스트_에러없음(t, 모음.DB저장(dbSQLite))

	// 2) 수집 흐름 시뮬레이션: DB에서 로드한 뒤 값 그대로 재저장
	로드_모음, 에러 := New종목별_일일_가격정보_모음_DB읽기(dbSQLite, "000000")
	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_에러없음(t, 로드_모음.DB저장(dbSQLite))

	// 3) 존재 레코드는 UPDATE로 처리되어야 해서 레코드 수 유지 (중복 INSERT 없음)
	확인_모음, 에러 := New종목별_일일_가격정보_모음_DB읽기(dbSQLite, "000000")
	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_같음(t, 확인_모음.Len(), 2)
}

// TestSQLite_변경값_저장 : 존재하는 레코드의 값을 변경해 저장하면 UPDATE가 실제로 반영되어야 한다.
func TestSQLite_변경값_저장(t *testing.T) {
	dbSQLite, 에러 := ldb.DB_SQLite(ldb.DSN_SQLite(":memory:"))
	lb.F테스트_에러없음(t, 에러)
	defer dbSQLite.Close()

	lb.F테스트_에러없음(t, F일일_가격정보_테이블_생성(dbSQLite))

	모음, 에러 := New종목별_일일_가격정보_모음([]*S일일_가격정보{
		f테스트용_일일_가격정보("000000", -2, 1000),
	})
	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_에러없음(t, 모음.DB저장(dbSQLite))

	// 동일 (종목코드, 일자)에 값이 바뀐 레코드로 재저장
	변경_모음, 에러 := New종목별_일일_가격정보_모음([]*S일일_가격정보{
		f테스트용_일일_가격정보("000000", -2, 1500),
	})
	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_에러없음(t, 변경_모음.DB저장(dbSQLite))

	확인_모음, 에러 := New종목별_일일_가격정보_모음_DB읽기(dbSQLite, "000000")
	lb.F테스트_에러없음(t, 에러)
	lb.F테스트_같음(t, 확인_모음.Len(), 1)
	lb.F테스트_같음(t, 확인_모음.G종가(), 1500.0)
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
	return New일일_가격정보(종목코드, lb.F금일().AddDate(0, 0, offset_일수),
		int64(종가), int64(종가+1), int64(종가-1), int64(종가), 100)
}
