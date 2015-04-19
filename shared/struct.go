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

@author: UnHa Kim <unha.kim@gh-system.com> */

package shared

import (
	dec "github.com/landjur/go-decimal"
	"math/big"
	"strings"
	"time"
)

// 종목
type I종목 interface {
	G코드() string
	G이름() string
}

func New종목(코드 string, 이름 string) I종목 {
	종목 := new(종목)
	종목.코드 = 코드
	종목.이름 = 이름

	return 종목
}

type 종목 struct {
	코드 string
	이름 string
}

func (this *종목) G코드() string {
	return this.코드
}

func (this *종목) G이름() string {
	return this.이름
}

// 가격정보
type I가격정보 interface {
	G종목코드() string
	G가격() I통화
	G시점() time.Time
}

func New가격정보(종목코드 string, 가격 I통화) I가격정보 {
	s := 가격정보{종목코드: 종목코드, 가격: 가격.G복사본(), 시점: time.Now()}
	return &s
}

type 가격정보 struct {
	종목코드 string
	가격 I통화
	시점 time.Time
}

func (this *가격정보) G종목코드() string       { return this.종목코드 }
func (this *가격정보) G가격() I통화       { return this.가격.G복사본() }
func (this *가격정보) G시점() time.Time { return this.시점 }

// 통화
type I통화 interface {
	G단위() string
	G실수값() float64
	G정밀값() *dec.Decimal
	G실수_문자열(소숫점_이하_자릿수 int) string
	G비교(다른_통화 I통화) T비교결과
	G부호() int
	G복사본() I통화
	G변경불가() bool
	S동결()
	S더하기(다른_통화 I통화) I통화
	S빼기(다른_통화 I통화) I통화
	S곱하기(다른_통화 I통화) I통화
	S나누기(다른_통화 I통화) I통화
	S금액(금액 string)
	String() string
}

func New원화(금액 string) I통화 { return New통화(KRW, 금액) }
func New달러(금액 string) I통화 { return New통화(USD, 금액) }
func New유로(금액 string) I통화 { return New통화(EUR, 금액) }
func New위안(금액 string) I통화 { return New통화(CNY, 금액) }
func New통화(단위 string, 금액 string) I통화 {
	정밀값, 에러 := dec.Parse(금액)
	
	if 에러 != nil {
	    F문자열_출력(에러.Error())
	    return nil
	}
	
	s := new(통화)
	s.단위 = 단위
	s.금액 = 정밀값
	s.변경불가 = false

	return s
}

type 통화 struct {
	단위   string
	금액   *dec.Decimal
	변경불가 bool
}

func (this *통화) G단위() string    { return this.단위 }
func (this *통화) G실수값() float64 { return this.금액.Float64() }
func (this *통화) G정밀값() *dec.Decimal {
	// 참조형이므로 그대로 주지 않고, 복사본을 준다.
	정밀값, 에러 := dec.Parse(this.금액.String())

	if 에러 != nil {
		panic(에러.Error())
	}

	return 정밀값
}
func (this *통화) G실수_문자열(소숫점_이하_자릿수 int) string {
	return this.금액.FloatString(소숫점_이하_자릿수)
}

func (this *통화) G비교(다른_통화 I통화) T비교결과 {
	switch {
	case this.단위 != 다른_통화.G단위():
		return P비교불가
	default:
		return T비교결과(this.금액.Cmp(다른_통화.G정밀값()))
	}
}

func (this *통화) G부호() int {
	return this.금액.Sign()
}

func (this *통화) G복사본() I통화 {
    s := new(통화)
    s.단위 = this.G단위()
    s.금액 = this.G정밀값()
    s.변경불가 = false
    
	return s
}

func (this *통화) G변경불가() bool {
	return this.변경불가
}

func (this *통화) S동결() {
	this.변경불가 = true
}

func (this *통화) S더하기(다른_통화 I통화) I통화 {
	if this.변경불가 {
		panic("변경불가능한 값입니다.")
	}

	다른_통화_금액 := 다른_통화.G정밀값()

	if this.단위 != 다른_통화.G단위() ||
		this.금액 == nil ||
		다른_통화_금액 == nil {
		this.금액 = nil
	} else {
		this.금액 = this.금액.Add(다른_통화_금액)
	}

	return this
}

func (this *통화) S빼기(다른_통화 I통화) I통화 {
	if this.변경불가 {
		panic("변경불가능한 값입니다.")
	}

	다른_통화_금액 := 다른_통화.G정밀값()

	if this.단위 != 다른_통화.G단위() ||
		this.금액 == nil ||
		다른_통화_금액 == nil {
		this.금액 = nil
	} else {
		this.금액 = this.금액.Sub(다른_통화_금액)
	}

	return this
}

func (this *통화) S곱하기(다른_통화 I통화) I통화 {
	if this.변경불가 {
		panic("변경불가능한 값입니다.")
	}

	다른_통화_금액 := 다른_통화.G정밀값()

	if this.단위 != 다른_통화.G단위() ||
		this.금액 == nil ||
		다른_통화_금액 == nil {
		this.금액 = nil
	} else {
		this.금액 = this.금액.Mul(다른_통화_금액)
	}

	return this
}

func (this *통화) S나누기(다른_통화 I통화) I통화 {
	if this.변경불가 {
		panic("변경불가능한 값입니다.")
	}

	다른_통화_금액 := 다른_통화.G정밀값()

	if this.단위 != 다른_통화.G단위() ||
		this.금액 == nil ||
		다른_통화_금액 == nil {
		this.금액 = nil

		return this
	}

	분자, 변환성공1 := new(big.Rat).SetString(this.String())
	분모, 변환성공2 := new(big.Rat).SetString(다른_통화_금액.String())

	// 변환에 실패하거나, 분모가 0이 되면 안 됨.
	if !변환성공1 || !변환성공2 || 분모.Cmp(big.NewRat(0, 1)) == 0 {
		this.금액 = nil

		return nil
	}

	결과값 := new(big.Rat).Quo(분자, 분모)

	// 소숫점 이하 1000자리 정도면 충분히 정밀하지 않을까?
	문자열 := 결과값.FloatString(1000)

	for strings.HasSuffix(문자열, "0") {
		문자열 = strings.TrimSuffix(문자열, "0")
	}

	if strings.HasSuffix(문자열, ".") {
		문자열 = strings.TrimSuffix(문자열, ".")
	}

	정밀값, 에러 := dec.Parse(문자열)

	if 에러 != nil {
		this.금액 = nil

		return nil
	}

	this.금액 = 정밀값

	return this
}

func (this *통화) S금액(금액 string) {
    if this.변경불가 {
		panic("변경불가능한 값입니다.")
	}
    
    정밀값, 에러 := dec.Parse(금액)
    
    if 에러 != nil {
        panic(에러.Error())
    }
    
    this.금액 = 정밀값
}

func (this *통화) String() string {
	return string(this.단위) + " " + this.금액.String()
}