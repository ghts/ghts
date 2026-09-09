# MySQL / SQLite 공용 저장 방안 (daily_price)

`data/daily_price` 패키지의 DB 접근 코드를 MySQL/SQLite 두 엔진에서 동일하게 동작하게 하는 방안.

핵심 원리: 일자 표현을 `time.Time` 대신 **`uint32`(YYYYMMDD)** 로 통일한다.
이렇게 하면 `time.Time`↔DB 변환이라는 엔진 간 가장 큰 분기점이 사라져,
DDL과 쿼리 한 세트로 두 엔진에서 동일하게 동작한다 

---

## 1. 현황 진단 (현재 코드의 SQLite 비호환 지점)

| # | 위치 | 문제 |
|---|------|------|
| 1 | `DB저장` | `INSERT IGNORE INTO`는 MySQL 전용 문법. SQLite는 `INSERT OR IGNORE INTO`를 요구 (서로 불허) |
| 2 | `DB저장` ↔ `DB읽기with시작일` | 쓰기는 `uint32`(20240101), 읽기는 `time.Time`으로 스캔. MySQL은 `DATE` 컬럼의 암묵 변환 덕분에 우연히 동작하나, SQLite는 `20240101`이 INTEGER로 저장되어 `time.Time` 스캔 시 `unsupported Scan` 에러 |
| 3 | `DB읽기with시작일` | `date>=?`에 `time.Time` 파라미터 바인딩. go-sqlite3는 `time.Time`을 `"2024-01-01 14:30:22+09:00"` 문자열로 변환 → SQLite의 문자열 사전순 비교로 시작일 당일 데이터가 제외될 수 있음 |

호환되는 부분 (변경 불필요):
`CREATE TABLE IF NOT EXISTS`, `PRIMARY KEY (code,date)`,
`?` 플레이스홀더,
`ORDER BY`, `sql.LevelDefault` 트랜잭션 옵션,
`CHAR(8)` / `DECIMAL(20,3)` / `BIGINT` 컬럼 타입
(SQLite는 `DECIMAL(20,3)`의 정밀도를 무시하고 REAL로 저장하나, 
가격은 Go 쪽에서 `float64`로 다루므로 기능상 문제없음).

---

## 2. 방안

### 2-1. DDL: `date DATE` → `date INTEGER`

```sql
CREATE TABLE IF NOT EXISTS daily_price (
    code   CHAR(8) NOT NULL,
    date   INTEGER NOT NULL,
    open   DECIMAL(20,3) NOT NULL,
    high   DECIMAL(20,3) NOT NULL,
    low    DECIMAL(20,3) NOT NULL,
    close  DECIMAL(20,3) NOT NULL,
    volume BIGINT NOT NULL,
    PRIMARY KEY (code, date)
)
```

- MySQL: `INTEGER`는 `INT`의 별칭이라 그대로 유효
- SQLite: 네이티브 INTEGER

### 2-2. I/O 통일: 전부 `uint32`(YYYYMMDD)

| 위치                                                        | 현재 | 변경 |
|-----------------------------------------------------------|------|------|
| `DB읽기with시작일` 시그니처                                        | `시작일 time.Time` | `시작일 uint32` |
| `rows.Scan`                                               | `&일자`(time.Time) → `lb.F일자2정수` 변환 | `&일일_가격정보.M일자`(uint32) 직접 스캔 |
| `WHERE date>=?`                                           | `time.Time` 파라미터 | `uint32` 파라미터 — 양쪽 모두 정수 비교 |
| `DB읽기()` (전체 조회)                                          | `time.Time{}` | `0` |
| 호출 4곳<br>(`New종목별_일일_가격정보_모음_{3년치,2년치,15개월치,13개월치}_DB읽기`) | `lb.F지금().Add(-3*365*lb.P1일)` 등 | `lb.F일자2정수(lb.F지금().Add(-3*365*lb.P1일))` 등 |

MySQL도 INT 컬럼은 드라이버가 `int64`로 돌려주므로 `uint32` 스캔이 정상 동작한다.

### 2-3. `INSERT IGNORE` 문제: 문법 분기 없이 해결

`INSERT IGNORE`(MySQL)와 `INSERT OR IGNORE`(SQLite)는 서로 불허하는 문법이다.
`INSERT ... ON DUPLICATE KEY UPDATE`(MySQL) / `ON CONFLICT DO UPDATE`(SQLite) 또한 서로 불허하는 문법이다.
MySQL과 SQLite 공용을 위해서 UPDATE를 먼저 실행한 후 `RowsAffected() == 0`(신규 레코드)이면 INSERT 하면 된다:

INSERT 문장은 실제 값을 담는다:
```sql
INSERT INTO daily_price (code, date, open, high, low, close, volume)
VALUES (?,?,?,?,?,?,?)
```

```go
for _, 값 := range s.M저장소 {
    // 1) 수정 시도
    rs := lb.F확인2(stmt수정.Exec(값.M시가, 값.M고가, 값.M저가, 값.M종가,
        값.M거래량, 값.M종목코드, 값.G일자()))

    // 2) 없던 레코드면 실제값으로 삽입
    if n, _ := rs.RowsAffected(); n == 0 {
        lb.F확인2(stmt생성.Exec(값.M종목코드, 값.G일자(),
            값.M시가, 값.M고가, 값.M저가, 값.M종가, 값.M거래량))
    }
}
```

### 2-4. MySQL 기존 데이터 마이그레이션 (일회성)

기존 MySQL 테이블이 `DATE` 타입이라면 단순 `ALTER`로는 값 변환이 안 되므로
데이터를 변환해야 한다 (SQLite 이전 작업과 일석이조):

테이블명은 `daily_price`를 유지하고, 원래 있던 테이블은 `daily_price_old`로 이름을 바꾼다:

```sql
-- 1) 원래 있던 테이블 이름 변경
RENAME TABLE daily_price TO daily_price_old;

-- 2) 새 스키마(2-1 참조)로 daily_price 재생성
CREATE TABLE daily_price (
    code   CHAR(8) NOT NULL,
    date   INTEGER NOT NULL,
    open   DECIMAL(20,3) NOT NULL,
    high   DECIMAL(20,3) NOT NULL,
    low    DECIMAL(20,3) NOT NULL,
    close  DECIMAL(20,3) NOT NULL,
    volume BIGINT NOT NULL,
    PRIMARY KEY (code, date)
);

-- 3) 일자 값을 YYYYMMDD 정수로 변환해서 이전
INSERT INTO daily_price (code, date, open, high, low, close, volume)
SELECT code,
       YEAR(date)*10000 + MONTH(date)*100 + DAY(date),
       open, high, low, close, volume
FROM daily_price_old;
```

---

## 3. 결과물

- DDL, SELECT, INSERT, UPDATE 전부 **엔진 무관한 단일 코드 경로**
- `time.Time`은 Go 구조체/비즈니스 로직에서만 사용하고,
  DB 경계에서는 `uint32`만 사용 → 드라이버별 로케일/타임존 함정 완전 제거
- `DECIMAL(20,3)`, `CHAR(8)`, `PRIMARY KEY (code,date)`는 그대로 유효해서 유지

## 4. 변경 대상 파일

- `type_daily_price.go`
  - `F일일_가격정보_테이블_생성` — `date DATE` → `date INTEGER` (완료)
  - `DB읽기` / `DB읽기with시작일` — 시그니처·스캔·파라미터를 `uint32`로 (완료)
  - `DB저장` — INSERT/UPDATE 순서 반전 (RowsAffected 기반) (완료)
  - `New종목별_일일_가격정보_모음_{3년치,2년치,15개월치,13개월치}_DB읽기` — 시작일 변환 (완료)
- `xing/util/daily_price_data.go`
  - `F당일_일일_가격정보_수집` — `한달전`을 `lb.F일자2정수(...)`로 변환 (완료)
