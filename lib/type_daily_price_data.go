/* Copyright (C) 2015-2020 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2020년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

package lib

import (
	"bytes"
	"context"
	"database/sql"
	"math"
	"sort"
	"strconv"
	"time"
)

func New일일_가격정보(종목코드 string, 일자 time.Time, 시가, 고가, 저가, 종가, 거래량 int64) *S일일_가격정보 {
	if len(종목코드) != 6 {
		panic(New에러("예상과 다른 종목코드 길이 : '%v' '%v'", 종목코드, len(종목코드)))
	} else if 일자 = F2일자(일자); 일자.Before(F지금().AddDate(-40, 0, 0)) {
		panic(New에러("너무 오래된 일자 : '%v'", 일자.Format(P일자_형식)))
	} else if 시가 < 0 {
		panic(New에러("음수 시가 : '%v'", 시가))
	} else if 고가 < 0 {
		panic(New에러("음수 고가 : '%v'", 고가))
	} else if 저가 < 0 {
		panic(New에러("음수 저가 : '%v'", 저가))
	} else if 종가 < 0 {
		panic(New에러("음수 종가 : '%v'", 종가))
	} else if 거래량 < 0 {
		panic(New에러("음수 거래량 : '%v'", 거래량))
	}

	return &S일일_가격정보{
		M종목코드: 종목코드,
		M일자:   uint32(F2정수_단순형(일자.Format("20060102"))),
		M시가:   float64(시가),
		M고가:   float64(고가),
		M저가:   float64(저가),
		M종가:   float64(종가),
		M거래량:  uint64(거래량)}
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
	M거래량  uint64
}

func (s S일일_가격정보) G종목코드() string { return s.M종목코드 }
func (s S일일_가격정보) G일자() uint32   { return s.M일자 }
func (s S일일_가격정보) G시가() float64  { return s.M시가 }
func (s S일일_가격정보) G고가() float64  { return s.M고가 }
func (s S일일_가격정보) G저가() float64  { return s.M저가 }
func (s S일일_가격정보) G종가() float64  { return s.M종가 }
func (s S일일_가격정보) G거래량() uint64  { return s.M거래량 }
func (s S일일_가격정보) G키() string {
	return s.M종목코드 + "_" + strconv.Itoa(int(s.M일자))
}
func (s S일일_가격정보) G일자2() time.Time {
	return F2포맷된_일자_단순형("20060102", F2문자열(s.M일자))
}

func New종목별_일일_가격정보_모음_DB읽기(db *sql.DB, 종목코드 string) (s *S종목별_일일_가격정보_모음, 에러 error) {
	s = new(S종목별_일일_가격정보_모음)

	if 에러 = s.DB읽기(db, 종목코드); 에러 != nil {
		return nil, 에러
	}

	return s, nil
}

func New종목별_일일_가격정보_모음(값_모음 []*S일일_가격정보) (s *S종목별_일일_가격정보_모음, 에러 error) {
	if len(값_모음) == 0 {
		return nil, New에러("비어 있는 입력값.")
	}

	종목코드 := 값_모음[0].M종목코드

	if 종목코드 == "" || len(종목코드) != 6 {
		return nil, New에러("잘못된 종목코드 : '%v'", 종목코드)
	}

	// 중복 제거를 위한 맵.
	맵 := make(map[uint32]*S일일_가격정보)

	for _, 값 := range 값_모음 {
		if 값.M종목코드 != 종목코드 {
			return nil, New에러("서로 다른 종목코드 : '%v' '%v'", 값.M종목코드, 종목코드)
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

	return s, nil
}

type S종목별_일일_가격정보_모음 struct {
	M저장소 []*S일일_가격정보
	인덱스  map[uint32]int
}

func (s *S종목별_일일_가격정보_모음) DB읽기(db *sql.DB, 종목코드 string) (에러 error) {
	종목코드 = F종목코드_보정(종목코드)
	F확인(F일일_가격정보_테이블_생성(db))

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
	SQL.WriteString("WHERE code=? ")
	SQL.WriteString("ORDER BY date")

	stmt, 에러 := db.Prepare(SQL.String())
	F확인(에러)
	defer stmt.Close()

	rows, 에러 := stmt.Query(종목코드)
	F확인(에러)
	defer rows.Close()

	s.M저장소 = make([]*S일일_가격정보, 0)

	var 일자 time.Time

	for rows.Next() {
		일일_가격정보 := new(S일일_가격정보)

		F확인(rows.Scan(
			&일일_가격정보.M종목코드,
			&일자,
			&일일_가격정보.M시가,
			&일일_가격정보.M고가,
			&일일_가격정보.M저가,
			&일일_가격정보.M종가,
			&일일_가격정보.M거래량))

		일일_가격정보.M일자 = F2정수_일자(일자)

		s.M저장소 = append(s.M저장소, 일일_가격정보)
	}

	s.S정렬_및_인덱스_설정()

	return nil
}

func (s *S종목별_일일_가격정보_모음) DB저장(db *sql.DB) (에러 error) {
	var tx *sql.Tx
	defer S예외처리{M에러: &에러, M함수: func() {
		if tx != nil {
			tx.Rollback()
		}
	}}.S실행()

	F확인(F일일_가격정보_테이블_생성(db))

	SQL := new(bytes.Buffer)
	SQL.WriteString("INSERT IGNORE INTO daily_price (")
	SQL.WriteString("code,")
	SQL.WriteString("date,")
	SQL.WriteString("open,")
	SQL.WriteString("high,")
	SQL.WriteString("low,")
	SQL.WriteString("close,")
	SQL.WriteString("volume")
	SQL.WriteString(") VALUES (?,?,?,?,?,?,?)")

	txOpts := new(sql.TxOptions)
	txOpts.Isolation = sql.LevelDefault
	txOpts.ReadOnly = false

	tx, 에러 = db.BeginTx(context.TODO(), txOpts)
	F확인(에러)

	stmt, 에러 := tx.Prepare(SQL.String())
	F확인(에러)
	defer stmt.Close()

	for _, 값 := range s.M저장소 {
		_, 에러 = stmt.Exec(
			값.M종목코드,
			값.M일자,
			값.M시가,
			값.M고가,
			값.M저가,
			값.M종가,
			값.M거래량)
		F확인(에러)
	}

	tx.Commit()

	return nil
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
	return s.M저장소[0].M종목코드
}
func (s S종목별_일일_가격정보_모음) G인덱스(일자 uint32) (int, error) {
	if 인덱스, 존재함 := s.인덱스[일자]; !존재함 {
		return 0, New에러("해당되는 인덱스 없음 : '%v'", 일자)
	} else if 인덱스 < 0 {
		return 0, New에러("음수 인덱스 : '%v'", 인덱스)
	} else if 인덱스 >= len(s.M저장소) {
		return 0, New에러("너무 큰 인덱스 : '%v' '%v'", 인덱스, len(s.M저장소))
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
	return s.G값(F2정수_일자(일자))
}

func (s S종목별_일일_가격정보_모음) G값_모음(시작일, 종료일 uint32) ([]*S일일_가격정보, error) {
	if 시작일 > 종료일 {
		return nil, New에러("시작일과 종료일이 뒤바뀜 : '%v' '%v'", 시작일, 종료일)
	} else if 시작_인덱스, 존재함 := s.인덱스[시작일]; !존재함 {
		return nil, New에러("해당되는 인덱스 없음 : '%v'", 시작일)
	} else if 시작_인덱스 < 0 {
		return nil, New에러("음수 인덱스 : '%v'", 시작_인덱스)
	} else if 시작_인덱스 >= len(s.M저장소) {
		return nil, New에러("너무 큰 인덱스 : '%v' '%v'", 시작_인덱스, len(s.M저장소))
	} else if 종료_인덱스, 존재함 := s.인덱스[종료일]; !존재함 {
		return nil, New에러("해당되는 인덱스 없음 : '%v'", 종료일)
	} else if 종료_인덱스 < 0 {
		return nil, New에러("음수 인덱스 : '%v'", 종료_인덱스)
	} else if 종료_인덱스 >= len(s.M저장소) {
		return nil, New에러("너무 큰 인덱스 : '%v' '%v'", 종료_인덱스, len(s.M저장소))
	} else if 시작_인덱스 > 종료_인덱스 {
		return nil, New에러("'시작_인덱스'와 '종료_인덱스'가 뒤바뀜 : '%v' '%v'", 시작_인덱스, 종료_인덱스)
	} else {
		값_모음 := make([]*S일일_가격정보, 종료_인덱스-시작_인덱스+1)

		for i := 0; i <= (종료_인덱스 - 시작_인덱스); i++ {
			값_모음[i] = s.M저장소[i+시작_인덱스]
		}

		return 값_모음, nil
	}
}

func (s S종목별_일일_가격정보_모음) G일자_모음() []time.Time {
	일자_모음 := make([]time.Time, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		일자_모음[i] = 일일_가격정보.G일자2()
	}

	return 일자_모음
}

func (s S종목별_일일_가격정보_모음) G시가_모음() []float64 {
	시가_모음 := make([]float64, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		시가_모음[i] = 일일_가격정보.M시가
	}

	return 시가_모음
}

func (s S종목별_일일_가격정보_모음) G고가_모음() []float64 {
	고가_모음 := make([]float64, len(s.M저장소))

	for i, 일일_가격정보 := range s.M저장소 {
		고가_모음[i] = 일일_가격정보.M고가
	}

	return 고가_모음
}

func (s S종목별_일일_가격정보_모음) G저가_모음() []float64 {
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

func (s S종목별_일일_가격정보_모음) G최근_종가() float64 {
	종가_모음 := s.G종가_모음()

	for i := 0; i < len(종가_모음); i++ {
		if 종가 := 종가_모음[len(종가_모음)-1-i]; 종가 != 0 {
			return 종가
		}
	}

	return 0
}

func (s S종목별_일일_가격정보_모음) G전일_고가_모음() []float64 {
	전일_고가_모음 := make([]float64, len(s.M저장소))
	전일_고가_모음[0] = (s.M저장소[0].M고가 + s.M저장소[1].M고가 + s.M저장소[2].M고가) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < len(s.M저장소); i++ {
		전일_고가_모음[i] = s.M저장소[i-1].M고가
	}

	return 전일_고가_모음
}

func (s S종목별_일일_가격정보_모음) G전일_저가_모음() []float64 {
	전일_저가_모음 := make([]float64, len(s.M저장소))
	전일_저가_모음[0] = (s.M저장소[0].M저가 + s.M저장소[1].M저가 + s.M저장소[2].M저가) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < len(s.M저장소); i++ {
		전일_저가_모음[i] = s.M저장소[i-1].M저가
	}

	return 전일_저가_모음
}

func (s S종목별_일일_가격정보_모음) G전일_종가_모음() []float64 {
	전일_종가_모음 := make([]float64, len(s.M저장소))
	전일_종가_모음[0] = (s.M저장소[0].M종가 + s.M저장소[1].M종가 + s.M저장소[2].M종가) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < len(s.M저장소); i++ {
		전일_종가_모음[i] = s.M저장소[i-1].M종가
	}

	return 전일_종가_모음
}

func (s S종목별_일일_가격정보_모음) G전일_거래량_모음() []float64 {
	전일_거래량 := make([]float64, len(s.M저장소))
	전일_거래량[0] = float64(s.M저장소[0].M거래량+s.M저장소[1].M거래량+s.M저장소[2].M거래량) / 3.0 // 임의로 값을 채워넣음.

	for i := 1; i < len(s.M저장소); i++ {
		전일_거래량[i] = float64(s.M저장소[i-1].M거래량)
	}

	return 전일_거래량
}

func (s S종목별_일일_가격정보_모음) G기간_고가_모음(윈도우_크기 int) []float64 {
	전일_고가_모음 := s.G전일_고가_모음()

	return F이동_범위_최대값(전일_고가_모음, 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) G기간_저가_모음(윈도우_크기 int) []float64 {
	전일_저가_모음 := s.G전일_저가_모음()

	return F이동_범위_최소값(전일_저가_모음, 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) G단순_이동_평균_모음(윈도우_크기 int) []float64 {
	if s.Len() <= 윈도우_크기 {
		panic(New에러("값_모음 길이가 너무 짦음. %v %v %v", s.G종목코드(), s.Len(), 윈도우_크기))
	}

	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F단순_이동_평균(s.G전일_종가_모음(), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) G지수_이동_평균_모음(윈도우_크기 int) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F지수_이동_평균(s.G전일_종가_모음(), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) G거래량_가중_이동_평균_모음(윈도우_크기 int) []float64 {
	return F가중_이동_평균(s.G전일_종가_모음(), s.G전일_거래량_모음(), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) G볼린저_밴드_모음(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F볼린저_밴드(s.G전일_종가_모음(), 윈도우_크기, 표준편차_배율)
}

func (s S종목별_일일_가격정보_모음) G볼린저_밴드_폭(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F볼린저_밴드_폭(s.G전일_종가_모음(), 윈도우_크기, 표준편차_배율)
}

// systrader79가 언급한 볼린저 밴드 폭 지수
func (s S종목별_일일_가격정보_모음) G변동성_Z점수_모음(윈도우_크기 int) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F이동_Z점수(F이동_표준_편차(s.G전일_종가_모음(), 윈도우_크기), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) G변동성_Z점수_최저값_모음(윈도우_크기 int) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F이동_범위_최소값(F이동_Z점수(F이동_표준_편차(s.G전일_종가_모음(), 윈도우_크기), 윈도우_크기), 20)
}

// systrader79가 언급한 볼린저 밴드 폭 Z점수 보완 지수
func (s S종목별_일일_가격정보_모음) G변동성_고점_대비_비율_모음(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	볼린저_밴드_폭 := F볼린저_밴드_폭(s.G전일_종가_모음(), 윈도우_크기, 표준편차_배율)
	고점 := make([]float64, len(s.M저장소))
	고점_대비_비율 := make([]float64, len(s.M저장소))

	for i := 0; i < 윈도우_크기; i++ {
		고점[i] = F최대값_실수(볼린저_밴드_폭[:i])
	}

	for i := 윈도우_크기; i < len(s.M저장소); i++ {
		고점[i] = F최대값_실수(볼린저_밴드_폭[i-윈도우_크기+1 : i])
	}

	for i := 0; i < len(s.M저장소); i++ {
		고점_대비_비율[i] = 볼린저_밴드_폭[i] / 고점[i]
	}

	return 고점_대비_비율
}

func (s S종목별_일일_가격정보_모음) G지수_볼린저_밴드_모음(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F지수_볼린저_밴드(s.G전일_종가_모음(), 윈도우_크기, 표준편차_배율)
}

func (s S종목별_일일_가격정보_모음) G지수_볼린저_밴드_폭_모음(윈도우_크기 int, 표준편차_배율 float64) []float64 {
	// Look Ahead Bias를 방지하기 위해서 전일 종가 기준으로 함.
	return F지수_볼린저_밴드_폭(s.G전일_종가_모음(), 윈도우_크기, 표준편차_배율)
}

func (s S종목별_일일_가격정보_모음) TrueRange_모음() []float64 {
	// Look Ahead Bias를 방지하기 위해서 하루 늦추어서 전일 True Range 값으로 설정함.
	TrueRange모음 := make([]float64, len(s.M저장소))

	TrueRange모음[0] = 0.0
	TrueRange모음[1] = 0.0

	for i := 2; i < len(s.M저장소); i++ {
		값1 := s.M저장소[i-2].M고가 - s.M저장소[i-1].M저가
		값2 := math.Abs(s.M저장소[i-1].M고가 - s.M저장소[i-2].M종가)
		값3 := math.Abs(s.M저장소[i-2].M종가 - s.M저장소[i-1].M저가)

		TrueRange모음[i] = math.Max(값1, math.Max(값2, 값3))
	}

	return TrueRange모음
}

func (s S종목별_일일_가격정보_모음) ATR_모음(윈도우_크기 int) []float64 {
	TrueRange모음 := s.TrueRange_모음()
	TrueRange모음[0] = (TrueRange모음[3] + TrueRange모음[4] + TrueRange모음[5]) / 3.0 // 임의로 값을 채워 넣음.
	TrueRange모음[1] = (TrueRange모음[4] + TrueRange모음[5] + TrueRange모음[6]) / 3.0 // 임의로 값을 채워 넣음.

	return F지수_이동_평균(TrueRange모음, 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) ATR채널_SMA_모음(윈도우_크기 int, atr증분 float64) []float64 {
	ATR채널 := make([]float64, len(s.M저장소))
	SMA := s.G단순_이동_평균_모음(윈도우_크기)
	ATR := s.ATR_모음(윈도우_크기)

	for i := 0; i < len(s.M저장소); i++ {
		ATR채널[i] = SMA[i] + atr증분*ATR[i]
	}

	return ATR채널
}

func (s S종목별_일일_가격정보_모음) ATR채널_EMA_모음(윈도우_크기 int, atr증분 float64) []float64 {
	ATR채널 := make([]float64, len(s.M저장소))
	EMA := s.G지수_이동_평균_모음(윈도우_크기)
	ATR := s.ATR_모음(윈도우_크기)

	for i := 0; i < len(s.M저장소); i++ {
		ATR채널[i] = EMA[i] + atr증분*ATR[i]
	}

	return ATR채널
}

func (s S종목별_일일_가격정보_모음) Chandelier청산_가격_모음(윈도우_크기 int, atr하락비율 float64) []float64 {
	Chandelier청산선 := make([]float64, len(s.M저장소))

	고가 := s.G기간_고가_모음(윈도우_크기)
	ATR := s.ATR_모음(윈도우_크기)

	for i := 0; i < len(s.M저장소); i++ {
		Chandelier청산선[i] = 고가[i] - math.Abs(atr하락비율)*ATR[i]
	}

	return Chandelier청산선
}

func (s S종목별_일일_가격정보_모음) MFI_모음(윈도우_크기 int) []float64 {
	// 참고자료 : https://www.investopedia.com/terms/m/mfi.asp
	// 해당 웹페이지에 나온 공식을 그대로 적용하다보니 모든 단어들이 영어임.

	high := s.G전일_고가_모음()
	low := s.G전일_저가_모음()
	close := s.G전일_종가_모음()
	volume := s.G전일_거래량_모음()
	typical_price := make([]float64, len(s.M저장소))
	positive_money_flow := make([]float64, len(s.M저장소))
	negative_money_flow := make([]float64, len(s.M저장소))
	sum_positive_money_flow := 0.0
	sum_negative_money_flow := 0.0
	mfi := make([]float64, len(s.M저장소))

	for i := 1; i < len(s.M저장소); i++ {
		typical_price[i] = (high[i] + low[i] + close[i]) / 3
		raw_money_flow := typical_price[i] * volume[i]

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

func (s S종목별_일일_가격정보_모음) VPCI_모음(단기, 장기 int) []float64 {
	// '거래량으로 투자하라'(Buff Dormeier 저) 제 17장
	// http://docs.mta.org/docs/2007DowAward.pdf
	// https://www.tradingview.com/script/lmTqKOsa-Indicator-Volume-Price-Confirmation-Indicator-VPCI/
	단기_VMWA := s.G거래량_가중_이동_평균_모음(단기)
	장기_VWMA := s.G거래량_가중_이동_평균_모음(장기)
	단기_SMA := s.G단순_이동_평균_모음(단기)
	장기_SMA := s.G단순_이동_평균_모음(장기)
	단기_거래량_SMA := F단순_이동_평균(s.G전일_거래량_모음(), 단기)
	장기_거래량_SMA := F단순_이동_평균(s.G전일_거래량_모음(), 장기)

	VPC := make([]float64, len(s.M저장소))
	VPR := make([]float64, len(s.M저장소))
	VM := make([]float64, len(s.M저장소))
	VPCI := make([]float64, len(s.M저장소))

	for i := 0; i < len(s.M저장소); i++ {
		VPC[i] = 장기_VWMA[i] - 장기_SMA[i]
		VPR[i] = 단기_VMWA[i] / 단기_SMA[i]
		VM[i] = 단기_거래량_SMA[i] / 장기_거래량_SMA[i]
		VPCI[i] = VPC[i] * VPR[i] * VM[i]
	}

	return VPCI
}

func (s S종목별_일일_가격정보_모음) VPCIs_모음(단기, 장기 int) []float64 {
	// '거래량으로 투자하라'(Buff Dormeier 저) 제 17장
	return F가중_이동_평균(s.VPCI_모음(단기, 장기), s.G전일_거래량_모음(), 단기)
}

func (s S종목별_일일_가격정보_모음) G추세_점수_모음() []float64 {
	추세_점수 := make([]float64, len(s.M저장소))

	for i := 0; i < 232; i++ {
		추세_점수[i] = 0.0
	}

	for i := 232; i < len(s.M저장소); i++ {
		기준_인덱스 := i - 1
		합계 := 0
		전일_종가 := s.M저장소[기준_인덱스].M종가

		for i := 2; i <= 11; i++ {
			과거_종가 := s.M저장소[기준_인덱스-(i*21)].M종가

			if 전일_종가 >= 과거_종가 {
				합계++
			}
		}

		추세_점수[i] = float64(합계) / 10.0
	}

	return 추세_점수
}

func (s S종목별_일일_가격정보_모음) G추세_점수_평균_모음(윈도우_크기 int) []float64 {
	return F지수_이동_평균(s.G추세_점수_모음(), 윈도우_크기)
}

func (s S종목별_일일_가격정보_모음) G최근_추세_점수() float64 {
	종가_모음 := s.G종가_모음()
	추세_점수_모음 := s.G추세_점수_모음()

	for i := 0; i < len(종가_모음); i++ {
		if 종가_모음[len(종가_모음)-1-i] != 0 {
			return 추세_점수_모음[len(추세_점수_모음)-1-i]
		}
	}

	return 0
}

//func (s S종목별_일일_가격정보_모음) G추세_점수값(전일 uint32) (float64, error) {
//	기준_인덱스, 에러 := s.G인덱스(전일)
//	if 에러 != nil {
//		return 0, 에러
//	}
//
//	합계 := 0
//
//	전일_종가 := s.M저장소[기준_인덱스].M종가
//
//	for i := 2; i <= 11; i++ {
//		과거_종가 := s.M저장소[기준_인덱스-(i*21)].M종가
//
//		if 전일_종가 >= 과거_종가 {
//			합계++
//		}
//	}
//
//	return float64(합계) / 10.0, nil
//}
//
//func (s S종목별_일일_가격정보_모음) G추세_점수값2(전일 time.Time) (float64, error) {
//	return s.G추세_점수값(F2정수_일자(전일))
//}

func (s S종목별_일일_가격정보_모음) G월수익율_변동성() float64 {
	월변동성_모음 := s.G월수익율_변동성_모음()

	return 월변동성_모음[len(월변동성_모음)-1]
}

func (s S종목별_일일_가격정보_모음) G월수익율_변동성_평균(윈도우_크기 int) float64 {
	월변동성_평균_모음 := s.G월수익율_변동성_평균_모음(윈도우_크기)

	return 월변동성_평균_모음[len(월변동성_평균_모음)-1]
}

func (s S종목별_일일_가격정보_모음) G월수익율_변동성_모음() []float64 {
	월수익율_변동성 := make([]float64, len(s.M저장소))

	for i := 0; i < 253; i++ {
		월수익율_변동성[i] = 0.0
	}

	for i := 253; i < len(s.M저장소); i++ {
		기준_인덱스 := i - 1 // Look Ahead Bias 방지를 위해서 전일 기준으로 시작.
		월수익율 := make([]float64, 12)

		for j := 0; j < 12; j++ {
			과거_종가1 := s.M저장소[기준_인덱스-(j*21)].M종가
			과거_종가2 := s.M저장소[기준_인덱스-((j+1)*21)].M종가
			월수익율[j] = (과거_종가1 - 과거_종가2) / 과거_종가2
		}

		월수익율_변동성[i] = F표준_편차(월수익율)
	}

	return 월수익율_변동성
}

func (s S종목별_일일_가격정보_모음) G월수익율_변동성_평균_모음(윈도우_크기 int) []float64 {
	return F지수_이동_평균(s.G월수익율_변동성_모음(), 윈도우_크기)
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
