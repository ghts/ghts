package val

import (
	"bytes"
	"context"
	"database/sql"
	lb "github.com/ghts/ghts/lib"
	"strconv"
	"time"
)

func New내재가치_정보_모음() *S내재가치_정보_모음 {
	s := new(S내재가치_정보_모음)
	s.M저장소 = make(map[string]*S내재가치_정보)

	return s
}

type S내재가치_정보_모음 struct {
	M저장소 map[string]*S내재가치_정보
}

func (s *S내재가치_정보_모음) G종목별_최신_정보(종목코드 string) *S내재가치_정보 {
	시작_연도 := lb.F지금().Year()
	완료_연도 := 시작_연도 - 2

	for 연도 := 시작_연도; 연도 >= 완료_연도; 연도-- {
		for 월 := 12; 월 > 0; 월-- {
			키 := 종목코드 + lb.F2문자열(연도*100+월)
			if 값, 존재함 := s.M저장소[키]; 존재함 {
				return 값
			}
		}
	}

	return nil
}

func (s *S내재가치_정보_모음) G종목별_차최신_정보(종목코드 string) *S내재가치_정보 {
	최신_정보 := s.G종목별_최신_정보(종목코드)
	if 최신_정보 == nil {
		return nil
	}

	연도 := 최신_정보.S내재가치_식별정보.G일자2().Year() - 1

	for 월 := 12; 월 > 0; 월-- {
		키 := 종목코드 + lb.F2문자열(연도*100+월)
		if 값, 존재함 := s.M저장소[키]; 존재함 {
			return 값
		}
	}

	return nil
}

func (s *S내재가치_정보_모음) G종목별_차차최신_정보(종목코드 string) *S내재가치_정보 {
	최신_정보 := s.G종목별_최신_정보(종목코드)
	if 최신_정보 == nil {
		return nil
	}

	연도 := 최신_정보.S내재가치_식별정보.G일자2().Year() - 2

	for 월 := 12; 월 > 0; 월-- {
		키 := 종목코드 + lb.F2문자열(연도*100+월)
		if 값, 존재함 := s.M저장소[키]; 존재함 {
			return 값
		}
	}

	return nil
}

func (s *S내재가치_정보_모음) S상장주식수_업데이트(db *sql.DB, 종목코드 string, 수량 int64) error {
	if 값 := s.G종목별_최신_정보(종목코드); 값 != nil {
		값.M상장주식수 = float64(수량)

		값_맵 := make(map[string]*S내재가치_정보)
		값_맵[종목코드] = 값

		return F내재가치_정보_모음_DB저장(db, 값_맵)
	}

	return nil
}

func (s *S내재가치_정보_모음) S파일_읽기(파일명 string) error {
	return lb.JSON_파일_읽기(파일명, s)
}

func (s *S내재가치_정보_모음) S파일_저장(파일명 string) error {
	return lb.JSON_파일_저장(s, 파일명)
}

func (s *S내재가치_정보_모음) DB읽기(db *sql.DB) (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	lb.F확인1(F내재가치_정보_테이블_생성(db))

	SQL := new(bytes.Buffer)
	SQL.WriteString("SELECT code, date, json ")
	SQL.WriteString("FROM fundamental_data ")
	SQL.WriteString("ORDER BY code, date")

	stmt := lb.F확인2(db.Prepare(SQL.String()))
	defer stmt.Close()

	rows := lb.F확인2(stmt.Query())
	defer rows.Close()

	for rows.Next() {
		var code string
		var date time.Time
		var json string
		var 값 *S내재가치_정보

		lb.F확인1(rows.Scan(&code, &date, &json))

		if 에러 = lb.F디코딩(lb.JSON, []byte(json), &값); 에러 != nil {
			lb.New에러with출력("%v %v : 디코딩 에러\n%v", code, date.Format(lb.P일자_형식), 에러)
			continue
		} else if 값 == nil {
			lb.New에러with출력("%v %v : nil 값", code, date.Format(lb.P일자_형식))
			continue
		} else if 값.S내재가치_식별정보 == nil {
			lb.New에러with출력("%v %v : nil 식별정보", code, date.Format(lb.P일자_형식))
			continue
		}

		s.M저장소[값.G키()] = 값
	}

	return
}

func (s *S내재가치_정보_모음) DB저장(db *sql.DB) error {
	return F내재가치_정보_모음_DB저장(db, s.M저장소)
}

type S내재가치_정보 struct {
	*S내재가치_식별정보
	*S재무제표_정보_내용
	*S재무비율_정보_내용
	*S투자지표_정보_내용
}

type S재무제표_정보 struct {
	*S내재가치_식별정보
	*S재무제표_정보_내용
}

type S재무비율_정보 struct {
	*S내재가치_식별정보
	*S재무비율_정보_내용
}

type S투자지표_정보 struct {
	*S내재가치_식별정보
	*S투자지표_정보_내용
}

func New내재가치_식별정보(종목코드 string, 연도 uint16, 월 uint8) *S내재가치_식별정보 {
	s := new(S내재가치_식별정보)
	s.M종목코드 = 종목코드
	s.S연도_월(연도, 월)

	return s
}

type S내재가치_식별정보 struct {
	M종목코드 string
	M연도_월 uint32
}

func (s *S내재가치_식별정보) G키() string {
	return s.M종목코드 + strconv.Itoa(int(s.M연도_월))
}

func (s *S내재가치_식별정보) S연도_월(연도 uint16, 월 uint8) {
	s.M연도_월 = uint32(연도)*100 + uint32(월)
}

func (s *S내재가치_식별정보) G일자() uint32 {
	return s.M연도_월*100 + 1
}

func (s *S내재가치_식별정보) G일자2() time.Time {
	return lb.F확인2(lb.F2포맷된_일자("20060102", strconv.Itoa(int(s.M연도_월*100+1))))
}

type S재무제표_정보_내용 struct {
	M매출액      float64
	M매출총이익    float64
	M연구개발비    float64
	M광고선전비    float64
	M감가상각비    float64
	M영업이익     float64
	M이자비용     float64
	M이자수익     float64
	M세전계속사업이익 float64
	M법인세비용    float64
	M당기순이익    float64
	M자산       float64
	M유동자산     float64 // NCAV 전략에 필요함.
	M비유동자산    float64
	M현금성자산    float64 // EV계산에 필요함.
	M재고자산     float64
	M매출채권     float64
	M부채       float64
	M유동부채     float64
	M비유동부채    float64
	M단기이자발생부채 float64
	M장기이자발생부채 float64
	M자본       float64
	M자본금      float64
	M자본잉여금    float64
	M이익잉여금    float64
	M영업_현금흐름  float64
	M투자_현금흐름  float64
	M재무_현금흐름  float64
	M현금_증가    float64
}

type S재무비율_정보_내용 struct {
	M유동_비율 float64
	//M당좌_비율     float64
	M부채_비율 float64
	//M이자보상배율    float64
	//M매출총이익율    float64
	//M세전계속사업이익율 float64
	M영업이익율 float64
	//EBITDA마진율  float64 // EBITDA가 왜 싫으세요?" (워런버핏, 찰리멍거) https://youtu.be/7Fze3RRyQTw
	ROA       float64
	ROE       float64
	ROIC      float64
	M총자산회전율   float64
	M총부채회전율   float64
	M총자본회전율   float64
	M순운전자본회전율 float64
}

type S투자지표_정보_내용 struct {
	M상장주식수 float64
	EPS    float64
	//EBITDAPS  float64 // EBITDA가 왜 싫으세요?" (워런버핏, 찰리멍거) https://youtu.be/7Fze3RRyQTw
	CFPS      float64
	SPS       float64
	BPS       float64
	DPS       float64
	M배당성향     float64
	M총현금흐름    float64
	M세후영업이익   float64
	M유무형자산상각비 float64
	M총투자      float64
	FCFF      float64
}

func F내재가치_정보_모음_DB읽기(db *sql.DB) (s *S내재가치_정보_모음, 에러 error) {
	s = New내재가치_정보_모음()
	에러 = s.DB읽기(db)

	return s, 에러
}

func F내재가치_정보_모음_DB저장(db *sql.DB, 값_맵 map[string]*S내재가치_정보) (에러 error) {
	var tx *sql.Tx
	defer lb.S예외처리{M에러: &에러, M함수: func() { lb.F조건부_실행(tx != nil, tx.Rollback) }}.S실행()

	lb.F확인1(F내재가치_정보_테이블_생성(db))

	SQL생성 := new(bytes.Buffer)
	SQL생성.WriteString("INSERT IGNORE INTO fundamental_data (")
	SQL생성.WriteString("code,")
	SQL생성.WriteString("date,")
	SQL생성.WriteString("json")
	SQL생성.WriteString(") VALUES (?,?,'')")

	SQL수정 := new(bytes.Buffer)
	SQL수정.WriteString("UPDATE fundamental_data SET json=? ")
	SQL수정.WriteString("WHERE code=? AND date=?")

	txOpts := new(sql.TxOptions)
	txOpts.Isolation = sql.LevelDefault
	txOpts.ReadOnly = false

	tx = lb.F확인2(db.BeginTx(context.TODO(), txOpts))

	stmt생성 := lb.F확인2(tx.Prepare(SQL생성.String()))
	defer stmt생성.Close()

	stmt수정 := lb.F확인2(tx.Prepare(SQL수정.String()))
	defer stmt수정.Close()

	for _, 값 := range 값_맵 {
		json := lb.F확인2(lb.F인코딩(lb.JSON, 값))
		lb.F확인2(stmt생성.Exec(값.M종목코드, 값.G일자()))
		lb.F확인2(stmt수정.Exec(string(json), 값.M종목코드, 값.G일자()))
	}

	tx.Commit()

	return nil
}

func F내재가치_정보_테이블_생성(db *sql.DB) error {
	SQL := new(bytes.Buffer)
	SQL.WriteString("CREATE TABLE IF NOT EXISTS fundamental_data (")
	SQL.WriteString("code CHAR(8) NOT NULL,")
	SQL.WriteString("date DATE NOT NULL,")
	SQL.WriteString("json TEXT NOT NULL,")
	SQL.WriteString("CONSTRAINT PRIMARY KEY (code,date)")
	SQL.WriteString(")")

	_, 에러 := db.Exec(SQL.String())

	return 에러
}
