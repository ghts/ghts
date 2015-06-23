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
	"strconv"
)

type T주소 int

func (this T주소) String() string { return "tcp://127.0.0.1:" + strconv.Itoa(3001+int(this)) }

const (
	P주소_주소정보 T주소 = iota
	P주소_종목정보
	P주소_가격정보_입수
	P주소_가격정보_배포
	P주소_가격정보
	P주소_테스트_결과_회신 T주소 = 998 // 테스트 결과 회신 주소는 3999번 포트로 고정
)

const (
	P메시지_구분_일반  = "N"
	P메시지_구분_종료  = "Q"
	P메시지_구분_OK  = "O"
	P메시지_구분_에러  = "E"
	P메시지_구분_GET = "G"
	P메시지_구분_PUT = "P"
)

type T통화단위 string

func (this T통화단위) String() string { return string(this) }

const (
	KRW T통화단위 = "KRW"
	USD T통화단위 = "USD"
	EUR T통화단위 = "EUR"
	CNY T통화단위 = "CNY"
)

type T비교결과 int

const (
	P같음   T비교결과 = 0
	P큼    T비교결과 = -1
	P작음   T비교결과 = 1
	P비교불가 T비교결과 = -999
)

type T부호 int

const (
	P양수 T부호 = 1
	P영  T부호 = 0
	P음수 T부호 = -1
)
