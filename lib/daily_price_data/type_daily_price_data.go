/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package daily_price_data

import (
	"bytes"
	"context"
	"database/sql"
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/trade"
	"math"
	"sort"
	"strconv"
	"time"
)

func New일일_가격정보(종목코드 string, 일자 time.Time, 시가, 고가, 저가, 종가, 거래량 int64) *S일일_가격정보 {
	if len(종목코드) != 6 {
		panic(lib.New에러("예상과 다른 종목코드 길이 : '%v' '%v'", 종목코드, len(종목코드)))
	} else if 일자 = lib.F2일자(일자); 일자.Before(lib.F지금().AddDate(-40, 0, 0)) {
		panic(lib.New에러("너무 오래된 일자 : '%v'", 일자.Format(lib.P일자_형식)))
	} else if 시가 < 0 {
		panic(lib.New에러("음수 시가 : '%v'", 시가))
	} else if 고가 < 0 {
		panic(lib.New에러("음수 고가 : '%v'", 고가))
	} else if 저가 < 0 {
		panic(lib.New에러("음수 저가 : '%v'", 저가))
	} else if 종가 < 0 {
		panic(lib.New에러("음수 종가 : '%v'", 종가))
	} else if 거래량 < 0 {
		panic(lib.New에러("음수 거래량 : '%v'", 거래량))
	}

	return &S일일_가격정보{
		M종목코드: 종목코드,
		M일자:   uint32(lib.F2정수_단순형(일자.Format("20060102"))),
		M시가:   float64(시가),
		M고가:   float64(고가),
		M저가:   float64(저가),
		M종가:   float64(종가),
		M거래량:  float64(거래량)}
}

type I일일_가격정보 interface {
	G종목코드() string
	G일자() uint32
	G시가() float64
	G고가() float64
	G저가() float64
	G종가() float64
	G거래량() uint64
}

type S일일_가격정보 struct {
	M종목코드 string
	M일자   uint32
	M시가   float64
	M고가   float64
	M저가   float64
	M종가   float64
	M거래량  float64
}

func (s S일일_가격정보) G종목코드() string { return s.M종목코드 }
func (s S일일_가격정보) G일자() uint32   { return s.M일자 }
func (s S일일_가격정보) G시가() float64  { return s.M시가 }
func (s S일일_가격정보) G고가() float64  { return s.M고가 }
func (s S일일_가격정보) G저가() float64  { return s.M저가 }
func (s S일일_가격정보) G종가() float64  { return s.M종가 }
func (s S일일_가격정보) G거래량() float64 { return s.M거래량 }
func (s S일일_가격정보) G키() string {
	return s.M종목코드 + "_" + strconv.Itoa(int(s.M일자))
}
func (s S일일_가격정보) G일자2() time.Time {
	return lib.F2포맷된_일자_단순형("20060102", lib.F2문자열(s.M일자))
}

func New종목별_일일_가격정보_모음_3년치_DB읽기(db *sql.DB, 종목코드 string) (s *S종목별_일일_가격정보_모음, 에러 error) {
	s = new(S종목별_일일_가격정보_모음)

	if 에러 = s.DB읽기with시작일(db, 종목코드, lib.F3년전()); 에러 != nil {
		return nil, 에러
	}

	s.M백테스트 = &(S일일_가격_백테스트_도우미{d: s})

	return s, nil
}

func New종목별_일일_가격정보_모음_DB읽기(db *sql.DB, 종목코드 string) (s *S종목별_일일_가격정보_모음, 에러 error) {
	s = new(S종목별_일일_가격정보_모음)

	if 에러 = s.DB읽기(db, 종목코드); 에러 != nil {
		return nil, 에러
	}

	s.M백테스트 = &(S일일_가격_백테스트_도우미{d: s})

	return s, nil
}

func New종목별_일일_가격정보_모음(값_모음 []*S일일_가격정보) (s *S종목별_일일_가격정보_모음, 에러 error) {
	if len(값_모음) == 0 {
		return nil, lib.New에러("비어 있는 입력값.")
	}

	종목코드 := 값_모음[0].M종목코드

	if 종목코드 == "" || len(종목코드) != 6 {
		return nil, lib.New에러("잘못된 종목코드 : '%v'", 종목코드)
	}

	// 중복 제거를 위한 맵.
	맵 := make(map[uint32]*S일일_가격정보)

	for _, 값 := range 값_모음 {
		if 값.M종목코드 != 종목코드 {
			return nil, lib.New에러("서로 다른 종목코드 : '%v' '%v'", 값.M종목코드, 종목코드)
		}

		맵[값.M일자] = 값
	}

	값_모음 = make([]*S일일_가격정보, len(맵))

	i := 0
	for _, 값 := range 맵 {
		값_모음[i] = 값 // 중복 제거된 값.
		i++
	}

	s = &S종목별_일일_가격정보_모음{M저장소: 값_모음}
	s.S정렬_및_인덱스_설정()

	s.M백테스트 = &(S일일_가격_백테스트_도우미{d: s})

	return s, nil
}

type S종목별_일일_가격정보_모음 struct {
	M저장소  []*S일일_가격정보
	인덱스   map[uint32]int
	M백테스트 *S일일_가격_백테스트_도우미
}

func (s *S종목별_일일_가격정보_모음) DB읽기(db *sql.DB, 종목코드 string) (에러 error) {
	return s.DB읽기with시작일(db, 종목코드, time.Time{})
}

func (s *S종목별_일일_가격정보_모음) DB읽기with시작일(db *sql.DB, 종목코드 string, 시작일 time.Time) (에러 error) {
	종목코드 = trade.F종목코드_보정(종목코드)
	lib.F확인(F일일_가격정보_테이블_생성(db))

	SQL := new(bytes.Buffer)
	SQL.WriteString("SELECT")
	SQL.WriteString(" code,")
	SQL.WriteString(" date,")
	SQL.WriteString(" open,")
	SQL.WriteString(" high,")
	SQL.WriteString(" low,")
	SQL.WriteString(" close,")
	SQL.WriteString(" volume ")
	SQL.WriteString("FROM daily_price ")
	SQL.WriteString("WHERE code=?")
	SQL.WriteString(" AND date>=? ")
	SQL.WriteString("ORDER BY date")

	stmt, 에러 := db.Prepare(SQL.String())
	lib.F확인(에러)
	defer stmt.Close()

	rows, 에러 := stmt.Query(종목코드, 시작일)
	lib.F확인(에러)
	defer rows.Close()

	s.M저장소 = make([]*S일일_가격정보, 0)

	var 일자 time.Time

	for rows.Next() {
		일일_가격정보 := new(S일일_가격정보)

		lib.F확인(rows.Scan(
			&일일_가격정보.M종목코드,
			&일자,
			&일일_가격정보.M시가,
			&일일_가격정보.M고가,
			&일일_가격정보.M저가,
			&일일_가격정보.M종가,
			&일일_가격정보.M거래량))

		일일_가격정보.M일자 = lib.F일자2정수(일자)

		s.M저장소 = append(s.M저장소, 일일_가격정보)
	}

	s.S정렬_및_인덱스_설정()

	return nil
}

func (s *S종목별_일일_가격정보_모음) DB저장(db *sql.DB) (에러 error) {
	var tx *sql.Tx
	defer lib.S예외처리{M에러: &에러, M함수: func() {
		lib.F에러_출력(에러)

		if tx != nil {
			tx.Rollback()
		}
	}}.S실행()

	lib.F확인(F일일_가격정보_테이블_생성(db))

	SQL생성 := new(bytes.Buffer)
	SQL생성.WriteString("INSERT IGNORE INTO daily_price (")
	SQL생성.WriteString("  code,")
	SQL생성.WriteString("  date,")
	SQL생성.WriteString("  open,")
	SQL생성.WriteString("  high,")
	SQL생성.WriteString("  low,")
	SQL생성.WriteString("  close,")
	SQL생성.WriteString("  volume")
	SQL생성.WriteString(") VALUES (?,?,0,0,0,0,0)")

	SQL수정 := new(bytes.Buffer)
	SQL수정.WriteString("UPDATE daily_price SET")
	SQL수정.WriteString("  open=?,")
	SQL수정.WriteString("  high=?,")
	SQL수정.WriteString("  low=?,")
	SQL수정.WriteString("  close=?,")
	SQL수정.WriteString("  volume=? ")
	SQL수정.WriteString("WHERE code=? AND date=?")

	txOpts := new(sql.TxOptions)
	txOpts.Isolation = sql.LevelDefault
	txOpts.ReadOnly = false

	tx, 에러 = db.BeginTx(context.TODO(), txOpts)
	lib.F확인(에러)

	stmt생성, 에러 := tx.Prepare(SQL생성.String())
	lib.F확인(에러)
	defer stmt생성.Close()

	stmt수정, 에러 := tx.Prepare(SQL수정.String())
	lib.F확인(에러)
	defer stmt수정.Close()

	for _, 값 := range s.M저장소 {
		_, 에러 = stmt생성.Exec(값.M종목코드, 값.G일자())
		lib.F확인(에러)

		_, 에러 := stmt수정.Exec(값.M시가, 값.M고가, 값.M저가, 값.M종가, 값.M거래량, 값.M종목코드, 값.G일자())
		lib.F확인(에러)
	}

	return tx.Commit()
}

func (s *S종목별_일일_가격정보_모음) S정렬_및_인덱스_설정() {
	sort.Sort(s)                 // 정렬
	s.인덱스 = make(map[uint32]int) // 인덱스 설정

	for i, 값 := range s.M저장소 {
		s.인덱스[값.M일자] = i
	}
}

func (s S종목별_일일_가격정보_모음) Len() int { return len(s.M저장소) }
func (s S종목별_일일_가격정보_모음) Swap(i, j int) {
	s.M저장소[i], s.M저장소[j] = s.M저장소[j], s.M저장소[i]
}
func (s S종목별_일일_가격정보_모음) Less(i, j int) bool {
	return s.M저장소[i].M일자 < s.M저장소[j].M일자
}
func (s S종목별_일일_가격정보_모음) G종목코드() string {
	if len(s.M저장소) == 0 {
		return ""
	}

	return s.M저장소[0].M종목코드
}
func (s S종목별_일일_가격정보_모음) G인덱스(일자 uint32) (int, error) {
	if 인덱스, 존재함 := s.인덱스[일자]; !존재함 {
		return 0, lib.New에러("[%v] 해당되는 인덱스 없음 : '%v'", s.G종목코드(), 일자)
	} else if 인덱스 < 0 {
		return 0, lib.New에러("음수 인덱스 : '%v'", 인덱스)
	} else if 인덱스 >= len(s.M저장소) {
		return 0, lib.New에러("너무 큰 인덱스 : '%v' '%v'", 인덱스, len(s.M저장소))
	} else {
		return 인덱스, nil
	}
}

func (s S종목별_일일_가격정보_모음) G값(일자 uint32) (*S일일_가격정보, error) {
	if 인덱스, 에러 := s.G인덱스(일자); 에러 != nil {
		return nil, 에러
	} else {
		return s.M저장소[인덱스], nil
	}
}

func (s S종목별_일일_가격정보_모음) G값2(일자 time.Time) (*S일일_가격정보, error) {
	return s.G값(lib.F일자2정수(일자))
}

func (s S종목별_일일_가격정보_모음) G값_모음(시작일, 종료일 uint32) ([]*S일일_가격정보, error) {
	if 시작일 > 종료일 {
		return nil, lib.New에러("시작일과 종료일이 뒤바뀜 : '%v' '%v'", 시작일, 종료일)
	} else if 시작_인덱스, 존재함 := s.인덱스[시작일]; !존재함 {
		return nil, lib.New에러("해당되는 인덱스 없음 : '%v'", 시작일)
	} else if 시작_인덱스 < 0 {
		return nil, lib.New에러("음수 인덱스 : '%v'", 시작_인덱스)
	} else if 시작_인덱스 >= len(s.M저장소) {
		return nil, lib.New에러("너무 큰 인덱스 : '%v' '%v'", 시작_인덱스, len(s.M저장소))
	} else if 종료_인덱스, 존재함 := s.인덱스[종료일]; !존재함 {
		return nil, lib.New에러("해당되는 인덱스 없음 : '%v'", 종료일)
	} else if 종료_인덱스 < 0 {
		return nil, lib.New에러("음수 인덱스 : '%v'", 종료_인덱스)
	} else if 종료_인덱스 >= len(s.M저장소) {
		return nil, lib.New에러("너무 큰 인덱스 : '%v' '%v'", 종료_인덱스, len(s.M저장소))
	} else if 시작_인덱스 > 종료_인덱스 {
		return nil, lib.New에러("'시작_인덱스'와 '종료_인덱스'가 뒤바뀜 : '%v' '%v'", 시작_인덱스, 종료_인덱스)
	} else {
		값_모음 := make([]*S일일_가격정보, 종료_인덱스-시작_인덱스+1)

		for i := 0; i <= (종료_인덱스 - 시작_인덱스); i++ {
			값_모음[i] = s.M저장소[i+시작_인덱스]
		}

		return 값_모음, nil
	}
}

func (s S종목별_일일_가격정보_모음) G최근_값_모음(시작일 uint32) ([]*S일일_가격정보, error) {
	if 시작_인덱스, 존재함 := s.인덱스[시작일]; !존재함 {
		return nil, lib.New에러("해당되는 인덱스 없음 : '%v'", 시작일)
	} else if 시작_인덱스 < 0 {
		return nil, lib.New에러("음수 인덱스 : '%v'", 시작_인덱스)
	} else if 시작_인덱스 >= len(s.M저장소) {
		return nil, lib.New에러("너무 큰 인덱스 : '%v' '%v'", 시작_인덱스, len(s.M저장소))
	} else {
		return s.M저장소[시작_인덱스:], nil
	}
}

func (s S종목별_일일_가격정보_모음) G일자_모음() []time.Time {
	일자_모음 := make([]time.Time, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		일자_모음[i] = 일일_가격정보.G일자2()
	}

	return 일자_모음
}

func (s S종목별_일일_가격정보_모음) g시가_모음() []float64 {
	시가_모음 := make([]float64, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		시가_모음[i] = 일일_가격정보.M시가
	}

	return 시가_모음
}

func (s S종목별_일일_가격정보_모음) g고가_모음() []float64 {
	고가_모음 := make([]float64, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		고가_모음[i] = 일일_가격정보.M고가
	}

	return 고가_모음
}

func (s S종목별_일일_가격정보_모음) g저가_모음() []float64 {
	저가_모음 := make([]float64, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		저가_모음[i] = 일일_가격정보.M저가
	}

	return 저가_모음
}

func (s S종목별_일일_가격정보_모음) G종가_모음() []float64 {
	종가_모음 := make([]float64, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		종가_모음[i] = 일일_가격정보.M종가
	}

	return 종가_모음
}

func (s S종목별_일일_가격정보_모음) g거래량_모음() []float64 {
	거래량_모음 := make([]float64, len(s.M저장소))

	for i := 1; i < len(s.M저장소); i++ {
		거래량_모음[i] = s.M저장소[i].M거래량
	}

	return 거래량_모음
}

// 실제 매매 용도

func (s S종목별_일일_가격정보_모음) G시가() float64 {
	return s.M저장소[len(s.M저장소)-1].M시가
}

func (s S종목별_일일_가격정보_모음) G고가() float64 {
	return s.M저장소[len(s.M저장소)-1].M고가
}

func (s S종목별_일일_가격정보_모음) G저가() float64 {
	return s.M저장소[len(s.M저장소)-1].M저가
}

func (s S종목별_일일_가격정보_모음) G종가() float64 {
	return s.M저장소[len(s.M저장소)-1].M종가
}

func (s S종목별_일일_가격정보_모음) G거래량() float64 {
	return s.M저장소[len(s.M저장소)-1].M거래량
}

func (s S종목별_일일_가격정보_모음) G이전_종가(이전_일수 int) float64 {
	종가_모음 := s.G종가_모음()

	return 종가_모음[len(종가_모음)-이전_일수-1]
}

func (s S종목별_일일_가격정보_모음) G기간_고점_종가(윈도우_크기 int) float64 {
	종가_모음 := s.G종가_모음()

	return lib.F최대값_실수(종가_모음[len(종가_모음)-윈도우_크기-1:])
}

func (s S종목별_일일_가격정보_모음) G기간_고가(윈도우_크기 int) float64 {
	역순_고가_모음 := make([]float64, 윈도우_크기)

	for i := 0; i < 윈도우_크기; i++ {
		역순_고가_모음[i] = s.M저장소[len(s.M저장소)-1-i].M고가
	}

	return lib.F최대값_실수(역순_고가_모음)
}

func (s S종목별_일일_가격정보_모음) G기간_저가(윈도우_크기 int) float64 {
	역순_저가_모음 := make([]float64, 윈도우_크기)

	for i := 0; i < 윈도우_크기; i++ {
		역순_저가_모음[i] = s.M저장소[len(s.M저장소)-1-i].M저가
	}

	return lib.F최소값_실수(역순_저가_모음)
}

func (s S종목별_일일_가격정보_모음) G평균_거래량(윈도우_크기 int) float64 {
	거래량_모음_전체 := s.g거래량_모음()
	거래량_모음 := 거래량_모음_전체[len(거래량_모음_전체)-윈도우_크기:]

	return lib.F평균(거래량_모음)
}

func (s S종목별_일일_가격정보_모음) G평균_거래_금액(윈도우_크기 int) float64 {
	거래량_모음_전체 := s.g거래량_모음()
	거래량_모음 := 거래량_모음_전체[len(거래량_모음_전체)-윈도우_크기:]

	종가_모음_전체 := s.G종가_모음()
	종가_모음 := 종가_모음_전체[len(종가_모음_전체)-윈도우_크기:]

	거래_금액_모음 := make([]float64, 윈도우_크기)

	for i := 0; i < 윈도우_크기; i++ {
		거래_금액_모음[i] = 거래량_모음[i] * 종가_모음[i]
	}

	return lib.F평균(거래_금액_모음)
}

func (s S종목별_일일_가격정보_모음) G거래_금액_민감도(윈도우_크기 int) float64 {
	거래량_모음_전체 := s.g거래량_모음()
	거래량_모음 := 거래량_모음_전체[len(거래량_모음_전체)-윈도우_크기:]

	종가_모음_전체 := s.G종가_모음()
	종가_모음 := 종가_모음_전체[len(종가_모음_전체)-윈도우_크기:]
	전일_종가_모음 := 종가_모음_전체[len(종가_모음_전체)-윈도우_크기-1:]

	민감도_모음 := make([]float64, 윈도우_크기)

	for i := 0; i < 윈도우_크기; i++ {
		일_수익율_절대값 := math.Abs(종가_모음[i]-전일_종가_모음[i]) / 전일_종가_모음[i] * 100
		일_거래_금액 := 거래량_모음[i] * 종가_모음[i]

		민감도_모음[i] = 일_수익율_절대값 / 일_거래_금액
	}

	return lib.F평균(민감도_모음)
}

func (s S종목별_일일_가격정보_모음) G거래_금액_변동성(윈도우_크기 int) float64 {
	거래량_모음_전체 := s.g거래량_모음()
	거래량_모음 := 거래량_모음_전체[len(거래량_모음_전체)-(윈도우_크기+1):]

	종가_모음_전체 := s.G종가_모음()
	종가_모음 := 종가_모음_전체[len(종가_모음_전체)-(윈도우_크기+1):]

	거래_금액_모음 := make([]float64, 윈도우_크기+1)

	for i := 0; i < 윈도우_크기+1; i++ {
		거래_금액_모음[i] = 거래량_모음[i] * 종가_모음[i]
	}

	거래_금액_변동율 := make([]float64, 윈도우_크기)

	for i := 0; i < 윈도우_크기; i++ {
		거래_금액_변동율[i] = (거래_금액_모음[i] - 거래_금액_모음[i+1]) / 거래_금액_모음[i] * 100
	}

	return lib.F표준_편차(거래_금액_변동율)
}

func (s S종목별_일일_가격정보_모음) G일_최대_수익율(윈도우_크기, 샘플_수량 int) float64 {
	종가_모음_전체 := s.G종가_모음()
	종가_모음 := 종가_모음_전체[len(종가_모음_전체)-윈도우_크기:]
	전일_종가_모음 := 종가_모음_전체[len(종가_모음_전체)-윈도우_크기-1:]
	일_수익율_모음 := make([]float64, 윈도우_크기)

	for i := 0; i < 윈도우_크기; i++ {
		일_수익율_모음[i] = (종가_모음[i] - 전일_종가_모음[i]) / 전일_종가_모음[i] * 100
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(일_수익율_모음)))

	return lib.F평균(일_수익율_모음[:샘플_수량])
}

func (s S종목별_일일_가격정보_모음) G총_변동성(윈도우_크기 int) float64 {
	종가_모음_전체 := s.G종가_모음()
	종가_모음 := 종가_모음_전체[len(종가_모음_전체)-윈도우_크기:]
	전일_종가_모음 := 종가_모음_전체[len(종가_모음_전체)-윈도우_크기-1:]
	일_수익율_모음 := make([]float64, 윈도우_크기)

	for i := 0; i < 윈도우_크기; i++ {
		일_수익율_모음[i] = (종가_모음[i] - 전일_종가_모음[i]) / 전일_종가_모음[i] * 100
	}

	return lib.F표준_편차(일_수익율_모음)
}

func (s S종목별_일일_가격정보_모음) SMA(윈도우_크기 int) float64 {
	단순_이동_평균 := trade.F단순_이동_평균(s.G종가_모음(), 윈도우_크기)

	return 단순_이동_평균[len(단순_이동_평균)-1]
}

func (s S종목별_일일_가격정보_모음) EMA(윈도우_크기 int) float64 {
	지수_이동_평균 := trade.F지수_이동_평균(s.G종가_모음(), 윈도우_크기)

	return 지수_이동_평균[len(지수_이동_평균)-1]
}

func (s S종목별_일일_가격정보_모음) VWMA(윈도우_크기 int) float64 {
	// 백테스트에 쓰는 게 아니고, 실제 거래에 사용할 것이기 때문에
	// Look Ahead Bias를 염려해서 '전일 종가' & '전일 거래량'을 사용할 필요가 없다.
	가중_이동_평균 := trade.F가중_이동_평균(s.G종가_모음(), s.g거래량_모음(), 윈도우_크기)

	return 가중_이동_평균[len(가중_이동_평균)-1]
}

func (s S종목별_일일_가격정보_모음) G볼린저_밴드(윈도우_크기 int, 배율 float64) float64 {
	종가_모음 := s.G종가_모음()
	이동_평균_모음 := trade.F이동_평균_도우미(종가_모음, 윈도우_크기, false)
	표준_편차_모음 := trade.F이동_표준_편차_도우미(종가_모음, 윈도우_크기, false)

	이동_평균 := 이동_평균_모음[len(이동_평균_모음)-1]
	표준_편차 := 표준_편차_모음[len(표준_편차_모음)-1]

	return 이동_평균 + 배율*표준_편차
}

func (s S종목별_일일_가격정보_모음) G볼린저_밴드_정보(윈도우_크기 int, 상한배율, 하한배율 float64) (중심값, 상한값, 하한값 float64) {
	종가_모음 := s.G종가_모음()
	이동_평균_모음 := trade.F이동_평균_도우미(종가_모음, 윈도우_크기, false)
	표준_편차_모음 := trade.F이동_표준_편차_도우미(종가_모음, 윈도우_크기, false)

	이동_평균 := 이동_평균_모음[len(이동_평균_모음)-1]
	표준_편차 := 표준_편차_모음[len(표준_편차_모음)-1]

	상한값 = 이동_평균 + 상한배율*표준_편차
	하한값 = 이동_평균 + 하한배율*표준_편차

	//lib.F체크포인트(s.G종목코드(), 이동_평균, 표준_편차)

	return 이동_평균, 상한값, 하한값
}

func (s S종목별_일일_가격정보_모음) G볼린저_밴드_폭(윈도우_크기 int) (중심값, 폭 float64) {
	종가_모음 := s.G종가_모음()
	이동_평균_모음 := trade.F이동_평균_도우미(종가_모음, 윈도우_크기, false)
	표준_편차_모음 := trade.F이동_표준_편차_도우미(종가_모음, 윈도우_크기, false)

	중심값 = 이동_평균_모음[len(이동_평균_모음)-1]
	폭 = 표준_편차_모음[len(표준_편차_모음)-1]

	return 중심값, 폭
}

func (s S종목별_일일_가격정보_모음) ATR(윈도우_크기 int) float64 {
	TrueRange모음 := make([]float64, len(s.M저장소))

	for i := 2; i < len(s.M저장소); i++ {
		고가 := lib.F최대값_실수([]float64{s.M저장소[i].M고가, s.M저장소[i-1].M종가})
		저가 := lib.F최소값_실수([]float64{s.M저장소[i].M저가, s.M저장소[i-1].M종가})

		TrueRange모음[i] = 고가 - 저가
	}

	TrueRange모음[0] = (TrueRange모음[3] + TrueRange모음[4] + TrueRange모음[5]) / 3.0 // 임의로 값을 채워 넣음.
	TrueRange모음[1] = (TrueRange모음[4] + TrueRange모음[5] + TrueRange모음[6]) / 3.0 // 임의로 값을 채워 넣음.

	atr모음 := trade.F지수_이동_평균(TrueRange모음, 윈도우_크기)

	return atr모음[len(atr모음)-1]
}

func (s S종목별_일일_가격정보_모음) vpci_모음(단기, 장기 int) []float64 {
	단기_VMWA_모음 := trade.F가중_이동_평균(s.G종가_모음(), s.g거래량_모음(), 단기)
	장기_VWMA_모음 := trade.F가중_이동_평균(s.G종가_모음(), s.g거래량_모음(), 장기)
	단기_SMA_모음 := trade.F단순_이동_평균(s.G종가_모음(), 단기)
	장기_SMA_모음 := trade.F단순_이동_평균(s.G종가_모음(), 장기)
	단기_거래량_SMA_모음 := trade.F단순_이동_평균(s.g거래량_모음(), 단기)
	장기_거래량_SMA_모음 := trade.F단순_이동_평균(s.g거래량_모음(), 장기)

	return s.VPCI_도우미(
		단기, 장기,
		단기_VMWA_모음, 장기_VWMA_모음,
		단기_SMA_모음, 장기_SMA_모음,
		단기_거래량_SMA_모음, 장기_거래량_SMA_모음)
}

func (s S종목별_일일_가격정보_모음) VPCI(단기, 장기 int) float64 {
	vpci_모음 := s.vpci_모음(단기, 장기)

	return vpci_모음[len(vpci_모음)-1]
}

func (s S종목별_일일_가격정보_모음) G전일_VPCI(단기, 장기 int) float64 {
	vpci_모음 := s.vpci_모음(단기, 장기)

	return vpci_모음[len(vpci_모음)-2]
}

func (s S종목별_일일_가격정보_모음) vpcis_모음(단기, 장기 int) []float64 {
	return trade.F가중_이동_평균(s.vpci_모음(단기, 장기), s.g거래량_모음(), 단기)
}

func (s S종목별_일일_가격정보_모음) VPCIs(단기, 장기 int) float64 {
	VPCIs_모음 := s.vpcis_모음(단기, 장기)

	return VPCIs_모음[len(VPCIs_모음)-1]
}

func (s S종목별_일일_가격정보_모음) G전일_VPCIs(단기, 장기 int) float64 {
	VPCIs_모음 := s.vpcis_모음(단기, 장기)

	return VPCIs_모음[len(VPCIs_모음)-2]
}

func (s S종목별_일일_가격정보_모음) mfi_모음(윈도우_크기 int) []float64 {
	return s.MFI_도우미(
		윈도우_크기,
		s.g고가_모음(),
		s.g저가_모음(),
		s.G종가_모음(),
		s.g거래량_모음())
}

func (s S종목별_일일_가격정보_모음) MFI(윈도우_크기 int) float64 {
	MFI_모음 := s.mfi_모음(윈도우_크기)

	return MFI_모음[len(MFI_모음)-1]
}

func (s S종목별_일일_가격정보_모음) G전일_MFI(윈도우_크기 int) float64 {
	MFI_모음 := s.MFI_도우미(
		윈도우_크기,
		s.g고가_모음(),
		s.g저가_모음(),
		s.G종가_모음(),
		s.g거래량_모음())

	return MFI_모음[len(MFI_모음)-2]
}

func (s S종목별_일일_가격정보_모음) mfis_모음(윈도우_크기_MFI, 윈도우_크기_이동평균 int) []float64 {
	MFI_모음 := s.MFI_도우미(
		윈도우_크기_MFI,
		s.g고가_모음(),
		s.g저가_모음(),
		s.G종가_모음(),
		s.g거래량_모음())

	return trade.F가중_이동_평균(MFI_모음, s.g거래량_모음(), 윈도우_크기_이동평균)
}

func (s S종목별_일일_가격정보_모음) MFIs(윈도우_크기_MFI, 윈도우_크기_이동평균 int) float64 {
	MFIs_모음 := s.mfis_모음(윈도우_크기_MFI, 윈도우_크기_이동평균)

	return MFIs_모음[len(MFIs_모음)-1]
}

func (s S종목별_일일_가격정보_모음) G전일_MFIs(윈도우_크기_MFI, 윈도우_크기_이동평균 int) float64 {
	MFIs_모음 := s.mfis_모음(윈도우_크기_MFI, 윈도우_크기_이동평균)

	return MFIs_모음[len(MFIs_모음)-2]
}

func (s S종목별_일일_가격정보_모음) G월별_추세_점수() float64 {
	return s.G월별_추세_점수_도우미(len(s.M저장소)-1, s.G종가_모음())
}

func (s S종목별_일일_가격정보_모음) G월수익율_변동성_모음() []float64 {
	종가_모음 := s.G종가_모음()
	월수익율_변동성 := make([]float64, s.Len())

	for i := 253; i < len(s.M저장소); i++ {
		월수익율_변동성[i] = s.g월수익율_변동성_도우미(i, 종가_모음)
	}

	return 월수익율_변동성
}

func (s S종목별_일일_가격정보_모음) G월수익율_변동성_평균_모음(윈도우_크기 int) []float64 {
	return trade.F지수_이동_평균(s.G월수익율_변동성_모음(), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) G월수익율_변동성() float64 {
	return s.g월수익율_변동성_도우미(s.Len()-1, s.G종가_모음())
}

func (s S종목별_일일_가격정보_모음) G월수익율_변동성_평균(윈도우_크기 int) float64 {
	종가_모음 := s.G종가_모음()
	월수익율_변동성 := make([]float64, len(s.M저장소))

	for i := 253; i < len(s.M저장소); i++ {
		기준_인덱스 := i
		월수익율_변동성[i] = s.g월수익율_변동성_도우미(기준_인덱스, 종가_모음)
	}

	월수익율_변동성_평균 := trade.F지수_이동_평균(월수익율_변동성, 윈도우_크기)

	return 월수익율_변동성_평균[len(월수익율_변동성_평균)-1]
}

func (s S종목별_일일_가격정보_모음) MFI_도우미(윈도우_크기 int,
	고가_모음 []float64, 저가_모음 []float64, 종가_모음 []float64, 거래량_모음 []float64) []float64 {
	// 참고자료 : https://www.investopedia.com/terms/m/mfi.asp

	if len(고가_모음) != len(저가_모음) ||
		len(고가_모음) != len(종가_모음) ||
		len(고가_모음) != len(거래량_모음) {
		panic(lib.New에러("서로 다른 길이 : %v %v %v %v",
			len(고가_모음), len(저가_모음), len(종가_모음), len(거래량_모음)))
	}

	typical_price := make([]float64, len(종가_모음))
	positive_money_flow := make([]float64, len(종가_모음))
	negative_money_flow := make([]float64, len(종가_모음))
	sum_positive_money_flow := 0.0
	sum_negative_money_flow := 0.0
	mfi := make([]float64, len(종가_모음))

	for i := 1; i < len(종가_모음); i++ {
		typical_price[i] = (고가_모음[i] + 저가_모음[i] + 종가_모음[i]) / 3
		raw_money_flow := typical_price[i] * 거래량_모음[i]

		if typical_price[i] > typical_price[i-1] {
			positive_money_flow[i] = raw_money_flow
		} else {
			negative_money_flow[i] = raw_money_flow
		}

		if i < 윈도우_크기 {
			sum_positive_money_flow += positive_money_flow[i]
			sum_negative_money_flow += negative_money_flow[i]
		} else {
			sum_positive_money_flow += positive_money_flow[i] - positive_money_flow[i-윈도우_크기]
			sum_negative_money_flow += negative_money_flow[i] - negative_money_flow[i-윈도우_크기]
		}

		money_flow_ratio := sum_positive_money_flow / sum_negative_money_flow
		mfi[i] = 100 - 100/(1+money_flow_ratio)
	}

	return mfi

}

func (s S종목별_일일_가격정보_모음) VPCI_도우미(
	단기, 장기 int,
	단기_VMWA_모음, 장기_VWMA_모음,
	단기_SMA_모음, 장기_SMA_모음,
	단기_거래량_SMA_모음, 장기_거래량_SMA_모음 []float64) []float64 {
	// '거래량으로 투자하라'(Buff Dormeier 저) 제 17장
	// http://docs.mta.org/docs/2007DowAward.pdf
	// https://www.tradingview.com/script/lmTqKOsa-Indicator-Volume-Price-Confirmation-Indicator-VPCI/

	VPC := make([]float64, s.Len())
	VPR := make([]float64, s.Len())
	VM := make([]float64, s.Len())
	VPCI := make([]float64, s.Len())

	for i := 0; i < s.Len(); i++ {
		VPC[i] = 장기_VWMA_모음[i] - 장기_SMA_모음[i]
		VPR[i] = 단기_VMWA_모음[i] / 단기_SMA_모음[i]
		VM[i] = 단기_거래량_SMA_모음[i] / 장기_거래량_SMA_모음[i]
		VPCI[i] = VPC[i] * VPR[i] * VM[i]
	}

	return VPCI
}

func (s S종목별_일일_가격정보_모음) G월별_추세_점수_도우미(기준_인덱스 int, 종가_모음 []float64) float64 {
	기준가 := 종가_모음[기준_인덱스]
	추세_점수 := 0.0

	for i := 2; i <= 11; i++ {
		if 기준_인덱스-(i*21) < 0 {
			추세_점수++
			continue
		}

		과거_종가 := 종가_모음[기준_인덱스-(i*21)]

		if 기준가 >= 과거_종가 {
			추세_점수++
		}
	}

	추세_점수 = 추세_점수 / 10

	return 추세_점수
}

func (s S종목별_일일_가격정보_모음) g일일_추세_점수_도우미(종가_모음 []float64, 윈도우_크기, 기준_인덱스 int) float64 {
	if 기준_인덱스 <= 0 || 윈도우_크기 <= 0 {
		return 0.0
	}

	기준가 := 종가_모음[기준_인덱스]
	반복_횟수 := lib.F최소값_정수(윈도우_크기, 기준_인덱스)
	카운터 := 0

	for i := 1; i <= 반복_횟수; i++ {
		if 과거_종가 := 종가_모음[기준_인덱스-i]; 기준가 >= 과거_종가 {
			카운터++
		}
	}

	return float64(카운터) / float64(반복_횟수)
}

func (s S종목별_일일_가격정보_모음) g월수익율_변동성_도우미(기준_인덱스 int, 종가_모음 []float64) float64 {
	월수익율 := make([]float64, 12)

	for j := 0; j < 12; j++ {
		과거_종가1 := 종가_모음[기준_인덱스-(j*21)]
		과거_종가2 := 종가_모음[기준_인덱스-((j+1)*21)]
		월수익율[j] = (과거_종가1 - 과거_종가2) / 과거_종가2
	}

	return lib.F표준_편차(월수익율)
}

// 백테스팅 Look Ahead Bias를 방지하기 위해 1일 전 데이터를 이용함.
type S일일_가격_백테스트_도우미 struct {
	d *S종목별_일일_가격정보_모음 // d는 data의 첫 글자.
}

func (s S일일_가격_백테스트_도우미) G시가_모음() []float64 {
	전일_시가_모음 := make([]float64, s.d.Len())
	전일_시가_모음[0] = (s.d.M저장소[0].M시가 + s.d.M저장소[1].M시가 + s.d.M저장소[2].M시가) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < s.d.Len(); i++ {
		전일_시가_모음[i] = s.d.M저장소[i-1].M시가
	}

	return 전일_시가_모음
}

func (s S일일_가격_백테스트_도우미) G고가_모음() []float64 {
	전일_고가_모음 := make([]float64, s.d.Len())
	전일_고가_모음[0] = (s.d.M저장소[0].M고가 + s.d.M저장소[1].M고가 + s.d.M저장소[2].M고가) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < s.d.Len(); i++ {
		전일_고가_모음[i] = s.d.M저장소[i-1].M고가
	}

	return 전일_고가_모음
}

func (s S일일_가격_백테스트_도우미) G저가_모음() []float64 {
	전일_저가_모음 := make([]float64, s.d.Len())
	전일_저가_모음[0] = (s.d.M저장소[0].M저가 + s.d.M저장소[1].M저가 + s.d.M저장소[2].M저가) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < s.d.Len(); i++ {
		전일_저가_모음[i] = s.d.M저장소[i-1].M저가
	}

	return 전일_저가_모음
}

func (s S일일_가격_백테스트_도우미) G종가_모음() []float64 {
	전일_종가_모음 := make([]float64, s.d.Len())
	전일_종가_모음[0] = (s.d.M저장소[0].M종가 + s.d.M저장소[1].M종가 + s.d.M저장소[2].M종가) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < s.d.Len(); i++ {
		전일_종가_모음[i] = s.d.M저장소[i-1].M종가
	}

	return 전일_종가_모음
}

func (s S일일_가격_백테스트_도우미) G거래량_모음() []float64 {
	전일_거래량_모음 := make([]float64, s.d.Len())
	전일_거래량_모음[0] = (s.d.M저장소[0].M거래량 + s.d.M저장소[1].M거래량 + s.d.M저장소[2].M거래량) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < s.d.Len(); i++ {
		전일_거래량_모음[i] = s.d.M저장소[i-1].M거래량
	}

	return 전일_거래량_모음
}

func (s S일일_가격_백테스트_도우미) G기간_고가_모음(윈도우_크기 int) []float64 {
	고가_모음 := s.G고가_모음()

	return trade.F이동_범위_최대값(고가_모음, 윈도우_크기)
}

func (s S일일_가격_백테스트_도우미) G기간_저가_모음(윈도우_크기 int) []float64 {
	저가_모음 := s.G저가_모음()

	return trade.F이동_범위_최소값(저가_모음, 윈도우_크기)
}

func (s *S일일_가격_백테스트_도우미) SMA(일자 time.Time, 윈도우_크기 int) (SMA float64, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { SMA = 0 }}.S실행()

	if 윈도우_크기 <= 0 {
		return 0, lib.New에러("너무 작은 윈도우 크기 %ㅍ.", 윈도우_크기)
	}

	인덱스, 에러 := s.d.G인덱스(lib.F일자2정수(일자))
	lib.F확인(에러)

	길이 := 인덱스 + 1
	if 길이 < 윈도우_크기 {
		return 0, lib.New에러("너무 짧은 인덱스.")
	}

	종가_합계 := 0.0

	for i := 0; i < 윈도우_크기; i++ {
		종가_합계 += s.d.M저장소[길이-윈도우_크기+i].M종가
	}

	return 종가_합계 / float64(윈도우_크기), nil
}

func (s *S일일_가격_백테스트_도우미) SMA_모음(윈도우_크기 int) []float64 {
	if s.d.Len() <= 윈도우_크기 {
		panic(lib.New에러("값_모음 길이가 너무 짦음. %v %v %v", s.d.G종목코드(), s.d.Len(), 윈도우_크기))
	}

	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return trade.F단순_이동_평균(s.G종가_모음(), 윈도우_크기)
}

func (s *S일일_가격_백테스트_도우미) EMA_모음(윈도우_크기 int) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return trade.F지수_이동_평균(s.G종가_모음(), 윈도우_크기)
}

func (s *S일일_가격_백테스트_도우미) VWMA_모음(윈도우_크기 int) []float64 {
	return trade.F가중_이동_평균(s.G종가_모음(), s.G거래량_모음(), 윈도우_크기)
}

func (s S일일_가격_백테스트_도우미) G볼린저_밴드_모음(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return trade.F볼린저_밴드(s.G종가_모음(), 윈도우_크기, 표준편차_배율)
}

func (s S일일_가격_백테스트_도우미) G볼린저_밴드_폭(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return trade.F볼린저_밴드_폭(s.G종가_모음(), 윈도우_크기, 표준편차_배율)
}

// systrader79가 언급한 볼린저 밴드 폭 지수
func (s S일일_가격_백테스트_도우미) G변동성_Z점수_모음(윈도우_크기 int) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return trade.F이동_Z점수(trade.F이동_표준_편차(s.G종가_모음(), 윈도우_크기), 윈도우_크기)
}

func (s S일일_가격_백테스트_도우미) G변동성_Z점수_최저값_모음(윈도우_크기 int) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return trade.F이동_범위_최소값(trade.F이동_Z점수(trade.F이동_표준_편차(s.G종가_모음(), 윈도우_크기), 윈도우_크기), 20)
}

// systrader79가 언급한 볼린저 밴드 폭 Z점수 보완 지수
func (s S일일_가격_백테스트_도우미) G변동성_고점_대비_비율_모음(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	볼린저_밴드_폭 := trade.F볼린저_밴드_폭(s.G종가_모음(), 윈도우_크기, 표준편차_배율)
	고점 := make([]float64, s.d.Len())
	고점_대비_비율 := make([]float64, s.d.Len())

	for i := 0; i < 윈도우_크기; i++ {
		고점[i] = lib.F최대값_실수(볼린저_밴드_폭[:i])
	}

	for i := 윈도우_크기; i < s.d.Len(); i++ {
		고점[i] = lib.F최대값_실수(볼린저_밴드_폭[i-윈도우_크기+1 : i])
	}

	for i := 0; i < s.d.Len(); i++ {
		고점_대비_비율[i] = 볼린저_밴드_폭[i] / 고점[i]
	}

	return 고점_대비_비율
}

func (s S일일_가격_백테스트_도우미) ATR(일자 time.Time, 윈도우_크기 int) (ATR float64, 에러 error) {
	defer lib.S예외처리{M에러: &에러, M함수: func() { ATR = 0 }}.S실행()

	인덱스, 에러 := s.d.G인덱스(lib.F일자2정수(일자))
	lib.F확인(에러)

	길이 := 인덱스 + 1

	if 길이 < 7 {
		return 0, lib.New에러("너무 짧은 인덱스.")
	}

	TrueRange모음 := make([]float64, 길이)

	for i := 2; i < 길이; i++ {
		고가 := lib.F최대값_실수([]float64{s.d.M저장소[i].M고가, s.d.M저장소[i-1].M종가})
		저가 := lib.F최소값_실수([]float64{s.d.M저장소[i].M저가, s.d.M저장소[i-1].M종가})
		TrueRange모음[i] = 고가 - 저가
	}

	TrueRange모음[0] = (TrueRange모음[3] + TrueRange모음[4] + TrueRange모음[5]) / 3.0 // 임의로 값을 채워 넣음.
	TrueRange모음[1] = (TrueRange모음[4] + TrueRange모음[5] + TrueRange모음[6]) / 3.0 // 임의로 값을 채워 넣음.

	atr모음 := trade.F지수_이동_평균(TrueRange모음, 윈도우_크기)

	return atr모음[len(atr모음)-1], nil
}

func (s *S일일_가격_백테스트_도우미) TrueRange_모음() []float64 {
	// Look Ahead Bias를 방지하기 위해서 하루 늦추어서 전일 True Range 값으로 설정함.
	TrueRange모음 := make([]float64, s.d.Len())

	for i := 2; i < s.d.Len(); i++ {
		고가 := lib.F최대값_실수([]float64{s.d.M저장소[i-1].M고가, s.d.M저장소[i-2].M종가})
		저가 := lib.F최소값_실수([]float64{s.d.M저장소[i-1].M저가, s.d.M저장소[i-2].M종가})

		TrueRange모음[i] = 고가 - 저가
	}

	return TrueRange모음
}

func (s *S일일_가격_백테스트_도우미) ATR_모음(윈도우_크기 int) []float64 {
	TrueRange모음 := s.TrueRange_모음()
	TrueRange모음[0] = (TrueRange모음[3] + TrueRange모음[4] + TrueRange모음[5]) / 3.0 // 임의로 값을 채워 넣음.
	TrueRange모음[1] = (TrueRange모음[4] + TrueRange모음[5] + TrueRange모음[6]) / 3.0 // 임의로 값을 채워 넣음.

	return trade.F지수_이동_평균(TrueRange모음, 윈도우_크기)
}

func (s *S일일_가격_백테스트_도우미) ATR채널_SMA_모음(윈도우_크기 int, atr증분 float64) []float64 {
	ATR채널 := make([]float64, s.d.Len())
	SMA := s.SMA_모음(윈도우_크기)
	ATR := s.ATR_모음(윈도우_크기)

	for i := 0; i < s.d.Len(); i++ {
		ATR채널[i] = SMA[i] + atr증분*ATR[i]
	}

	return ATR채널
}

func (s *S일일_가격_백테스트_도우미) ATR채널_EMA_모음(윈도우_크기 int, atr증분 float64) []float64 {
	ATR채널 := make([]float64, s.d.Len())
	EMA := s.EMA_모음(윈도우_크기)
	ATR := s.ATR_모음(윈도우_크기)

	for i := 0; i < s.d.Len(); i++ {
		ATR채널[i] = EMA[i] + atr증분*ATR[i]
	}

	return ATR채널
}

func (s *S일일_가격_백테스트_도우미) Chandelier청산_가격_모음(윈도우_크기 int, atr하락비율 float64) []float64 {
	Chandelier청산선 := make([]float64, s.d.Len())

	고가 := s.G기간_고가_모음(윈도우_크기)
	ATR := s.ATR_모음(윈도우_크기)

	for i := 0; i < s.d.Len(); i++ {
		Chandelier청산선[i] = 고가[i] - math.Abs(atr하락비율)*ATR[i]
	}

	return Chandelier청산선
}

func (s *S일일_가격_백테스트_도우미) MFI_모음(윈도우_크기 int) []float64 {
	return s.d.MFI_도우미(
		윈도우_크기,
		s.G고가_모음(),
		s.G저가_모음(),
		s.G종가_모음(),
		s.G거래량_모음())
}

func (s *S일일_가격_백테스트_도우미) MFIs_모음(윈도우_크기_MFI, 윈도우_크기_이동평균 int) []float64 {
	return trade.F가중_이동_평균(s.MFI_모음(윈도우_크기_MFI), s.G거래량_모음(), 윈도우_크기_이동평균)
}

func (s *S일일_가격_백테스트_도우미) VPCI_모음(단기, 장기 int) []float64 {
	단기_VMWA_모음 := s.VWMA_모음(단기)
	장기_VWMA_모음 := s.VWMA_모음(장기)
	단기_SMA_모음 := s.SMA_모음(단기)
	장기_SMA_모음 := s.SMA_모음(장기)
	단기_거래량_SMA_모음 := trade.F단순_이동_평균(s.G거래량_모음(), 단기)
	장기_거래량_SMA_모음 := trade.F단순_이동_평균(s.G거래량_모음(), 장기)

	return s.d.VPCI_도우미(
		단기, 장기,
		단기_VMWA_모음, 장기_VWMA_모음,
		단기_SMA_모음, 장기_SMA_모음,
		단기_거래량_SMA_모음, 장기_거래량_SMA_모음)
}

func (s *S일일_가격_백테스트_도우미) VPCIs_모음(단기, 장기 int) []float64 {
	// '거래량으로 투자하라'(Buff Dormeier 저) 제 17장
	return trade.F가중_이동_평균(s.VPCI_모음(단기, 장기), s.G거래량_모음(), 단기)
}

func (s S일일_가격_백테스트_도우미) G월별_추세_점수_모음() []float64 {
	추세_점수 := make([]float64, s.d.Len())

	for i := 232; i < s.d.Len(); i++ {
		기준_인덱스 := i - 1
		추세_점수[i] = s.d.G월별_추세_점수_도우미(기준_인덱스, s.G종가_모음())
	}

	return 추세_점수
}

func (s S일일_가격_백테스트_도우미) G월별_추세_점수_평균_모음(윈도우_크기 int) []float64 {
	return trade.F지수_이동_평균(s.G월별_추세_점수_모음(), 윈도우_크기)
}

func F일일_가격정보_테이블_생성(db *sql.DB) error {
	SQL := new(bytes.Buffer)
	SQL.WriteString("CREATE TABLE IF NOT EXISTS daily_price (")
	SQL.WriteString("code CHAR(8) NOT NULL,")
	SQL.WriteString("date DATE NOT NULL,")
	SQL.WriteString("open DECIMAL(20,3) NOT NULL,")
	SQL.WriteString("high DECIMAL(20,3) NOT NULL,")
	SQL.WriteString("low DECIMAL(20,3) NOT NULL,")
	SQL.WriteString("close DECIMAL(20,3) NOT NULL,")
	SQL.WriteString("volume BIGINT NOT NULL,")
	SQL.WriteString("CONSTRAINT PRIMARY KEY (code,date)")
	SQL.WriteString(")")

	_, 에러 := db.Exec(SQL.String())

	return 에러
}
