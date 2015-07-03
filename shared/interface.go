/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package shared

import (
	dec "github.com/landjur/go-decimal"

	"time"
)

// 회신
type I회신 interface {
	G내용() []string
	G에러() error
}

func New회신(내용 []string, 에러 error) I회신 {
	return s회신{내용: 내용, 에러: 에러}
}

// 종목
type I종목 interface {
	G코드() string
	G이름() string
}

func New종목(코드 string, 이름 string) I종목 {
	종목 := new(s종목)
	종목.코드 = 코드
	종목.이름 = 이름

	return 종목
}

// 통화
type I통화 interface {
	G단위() T통화단위
	G실수값() float64
	G정밀값() *dec.Decimal
	G실수_문자열(소숫점_이하_자릿수 int) string
	G비교(다른_통화 I통화) T비교결과
	G부호() T부호
	G복사본() I통화
	G변경불가() bool
	S동결()
	S더하기(다른_통화 I통화) I통화
	S빼기(다른_통화 I통화) I통화
	S곱하기(다른_통화 I통화) I통화
	S나누기(다른_통화 I통화) I통화
	S금액(금액 string) I통화
	String() string
}

func New원화(금액 string) I통화 { return New통화(KRW, 금액) }
func New달러(금액 string) I통화 { return New통화(USD, 금액) }
func New유로(금액 string) I통화 { return New통화(EUR, 금액) }
func New위안(금액 string) I통화 { return New통화(CNY, 금액) }
func New통화(단위 T통화단위, 금액 string) I통화 {
	정밀값, 에러 := dec.Parse(금액)

	if 에러 != nil {
		F에러_출력(에러)
		return nil
	}

	s := new(s통화)
	s.단위 = 단위
	s.금액 = 정밀값
	s.변경불가 = false

	return s
}

// 가격정보
type I가격정보 interface {
	G종목() I종목
	G가격() I통화
	G시점() time.Time
}

func New가격정보(종목 I종목, 가격 I통화) I가격정보 {
	s := s가격정보{종목: 종목, 가격: 가격.G복사본(), 시점: time.Now()}
	return &s
}
