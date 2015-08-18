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
	"time"
)

const (
	P시간_형식 string = time.RFC3339Nano

	P타임아웃_Go  time.Duration = 10 * time.Second
	P타임아웃_zmq time.Duration = 10 * time.Second
)

const (
	P주소_주소정보   string = "tcp://127.0.0.1:3001"
	P주소_테스트_결과        = "tcp://127.0.0.1:3002"
)

const (
	P주소명_주소정보   string = "주소정보"
	P주소명_테스트_결과        = "테스트_결과"

	P주소명_종목정보    = "종목정보"
	P주소명_가격정보    = "가격정보"
	P주소명_가격정보_입수 = "가격정보_입수"
	P주소명_가격정보_배포 = "가격정보_배포"
)

const (
	// 질의 메시지 구분
	P메시지_GET = "G"
	P메시지_SET = "S"
	P메시지_DEL = "D"
	P메시지_초기화 = "I" // 주로 테스트에서만 사용함.
	P메시지_종료  = "Q" // 주로 zmq와 테스트에서만 사용함.

	// 회신 메시지 구분
	P메시지_OK = "O"
	P메시지_에러 = "E"
)

const (
	KRW string = "KRW"
	USD        = "USD"
	EUR        = "EUR"
	CNY        = "CNY"
)

const (
	P같음   int = 0
	P큼        = -1
	P작음       = 1
	P비교불가     = -999
)

const (
	P양수 int = 1
	P영  int = 0
	P음수 int = -1
)

const (
	P포지션_롱 string = "L"
	P포지션_숏 string = "S"
)
