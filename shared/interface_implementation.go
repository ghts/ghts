/* This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>.

@author: UnHa Kim <unha.kim.ghts@gmail.com> */

package shared

import (
	dec "github.com/wayn3h0/go-decimal"

	"bytes"
	"strconv"
	"sync"
	"time"
)

type S비어있는_구조체 struct{}

// 안전한 bool
type s안전한_bool struct {
	sync.RWMutex
	값 bool
}

func (this *s안전한_bool) G값() bool {
	this.RLock() // Go언어의 Embedded Lock
	defer this.RUnlock()

	return this.값
}

func (this *s안전한_bool) S값(값 bool) error {
	this.Lock()
	defer this.Unlock()

	if this.값 == 값 {
		return F에러_생성("이미 %v임.", 값)
	} else {
		this.값 = 값
		return nil
	}
}

// 안전한 string
type s안전한_string struct {
	sync.RWMutex
	값 string
}

func (this *s안전한_string) G값() string {
	this.RLock() // Go언어의 Embedded Lock
	defer this.RUnlock()

	return this.값
}

func (this *s안전한_string) S값(값 string) {
	this.Lock()
	defer this.Unlock()

	this.값 = 값
}

// 기본 메시지
type s기본_메시지 struct {
	구분 string
	내용 []string
}

func (this s기본_메시지) G구분() string {
	return this.구분
}

func (this s기본_메시지) G내용(인덱스 int) string {
	if 인덱스 >= len(this.내용) {
		F문자열_및_호출경로_출력("인덱스 입력값은 '길이'보다 작아야 함 : 길이 %v, 입력값 %v", len(this.내용), 인덱스)
		panic("무효한 인덱스")
	}

	return this.내용[인덱스]
}

func (this s기본_메시지) G내용_전체() []string {
	return this.내용
}

func (this s기본_메시지) G길이() int {
	return len(this.내용)
}

func (this s기본_메시지) String() string {
	버퍼 := new(bytes.Buffer)

	버퍼.WriteString("구분 : " + this.구분 + "\n")
	버퍼.WriteString("길이 : " + strconv.Itoa(this.G길이()) + "\n")

	if len(this.내용) == 0 {
		버퍼.WriteString("내용 없음. len(내용) == 0. \n")
	} else {
		버퍼.WriteString("내용\n")

		for i := 0; i < len(this.내용); i++ {
			버퍼.WriteString(strconv.Itoa(i) + " : " + this.내용[i] + "\n")
		}
	}

	return 버퍼.String()
}

// 질의 메시지
type s질의_메시지 struct {
	s기본_메시지 // Go언어 구조체 embedding(임베딩) 기능. 상속 비스무리함.
	회신_채널   chan I회신
}

func (this s질의_메시지) G검사(메시지_구분 string, 질의_길이 int) error {
	if this.G구분() == 메시지_구분 &&
		this.G길이() == 질의_길이 {
		return nil
	}

	에러 := F에러_생성("잘못된 질의 메시지.\n%s", this.String())
	F에러_출력(에러)

	this.회신_채널 <- New회신(에러, P메시지_에러)

	return 에러
}

func (this s질의_메시지) G회신(질의_채널 chan I질의, 타임아웃 time.Duration) I회신 {
	질의_채널 <- this

	select {
	case 회신 := <-this.회신_채널:
		return 회신
	case <-time.After(타임아웃):
		return New회신(F에러_생성("I질의.G회신() 타임아웃.\n%v", this))
	}
}

func (this s질의_메시지) S회신(에러 error, 내용 ...interface{}) {
	this.회신_채널 <- New회신(에러, 내용...)
}

// 회신 메시지
type s회신_메시지 struct {
	s기본_메시지 // Go언어 구조체 embedding(임베딩)
	에러      error
}

func (this s회신_메시지) G에러() error {
	return this.에러
}

// 종목
type s종목 struct {
	코드 string
	이름 string
}

func (this *s종목) G코드() string {
	return this.코드
}

func (this *s종목) G이름() string {
	return this.이름
}

func (this *s종목) String() string {
	return this.코드 + " " + this.이름
}

// 통화
type s통화 struct {
	단위   string
	금액   *dec.Decimal
	변경불가 bool
}

func (this *s통화) G단위() string   { return this.단위 }
func (this *s통화) G실수값() float64 { return this.금액.Float() }
func (this *s통화) G정밀값() *dec.Decimal {
	정밀값, _ := dec.Parse(this.금액.String())

	return 정밀값
}
func (this *s통화) G문자열값() string { return this.금액.String() }
func (this *s통화) G문자열값_고정소숫점(소숫점_이하_자릿수 int) string {
	return this.금액.FloatString(소숫점_이하_자릿수)
}

func (this *s통화) G비교(다른_통화 I통화) int {
	switch {
	case this.단위 != 다른_통화.G단위():
		return P비교불가
	default:
		return this.금액.Cmp(다른_통화.G정밀값())
	}
}

func (this *s통화) G부호() int {
	return this.금액.Sign()
}

func (this *s통화) G복사본() I통화 {
	s := new(s통화)
	s.단위 = this.G단위()
	s.금액 = this.G정밀값()
	s.변경불가 = false

	return s
}

func (this *s통화) G변경불가() bool {
	return this.변경불가
}

func (this *s통화) S동결() {
	this.변경불가 = true
}

func (this *s통화) S더하기(값 float64) I통화 {
	if this.변경불가 {
		panic(F에러_생성("변경불가능한 값입니다."))
		return this
	}

	this.금액 = this.금액.Add(dec.New(값))

	return this
}

func (this *s통화) S빼기(값 float64) I통화 {
	if this.변경불가 {
		panic(F에러_생성("변경불가능한 값입니다."))
		return this
	}

	this.금액 = this.금액.Sub(dec.New(값))

	return this
}

func (this *s통화) S곱하기(값 float64) I통화 {
	if this.변경불가 {
		panic(F에러_생성("변경불가능한 값입니다."))
		return this
	}

	this.금액 = this.금액.Mul(dec.New(값))

	return this
}

func (this *s통화) S나누기(값 float64) (I통화, error) {
	switch {
	case this.변경불가:
		panic(F에러_생성("변경불가능한 값입니다."))
		return this, F에러_생성("변경불가능한 값입니다.")
	case 값 == 0.0:
		//panic(F에러_생성("분모가 0인 나눗셈 불가."))
		return nil, F에러_생성("분모가 0인 나눗셈 불가.")
	default:
		this.금액 = this.금액.Div(dec.New(값))
		return this, nil
	}
}

func (this *s통화) S금액(금액 float64) I통화 {
	if this.변경불가 {
		panic("변경불가능한 값입니다.")
		return this
	}

	this.금액 = dec.New(금액)

	return this
}

func (this *s통화) String() string {
	return this.단위 + " " + this.금액.String()
}

// 가격정보
type s가격정보 struct {
	종목코드 string
	가격   I통화
	시점   time.Time
}

func (this *s가격정보) G종목코드() string  { return this.종목코드 }
func (this *s가격정보) G가격() I통화       { return this.가격.G복사본() }
func (this *s가격정보) G시점() time.Time { return this.시점 }
