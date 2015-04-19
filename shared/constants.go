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

const (
	P주소_주소정보 string = "tcp://127.0.0.1:3003"

	P주소_가격정보_입수 = "tcp://127.0.0.1:3004"
	P주소_가격정보_배포 = "tcp://127.0.0.1:3005"
	P주소_가격정보 = "tcp://127.0.0.1:3006"
	
	P주소_테스트_결과_회신 string = "tcp://127.0.0.1:3999"
)

const (
	P메시지_구분_일반 string = "N"
	P메시지_구분_종료 string = "Q"
	P메시지_구분_OK string = "O"
	P메시지_구분_에러 string = "E"
	
	P메시지_구분_GET string = "G"
	P메시지_구분_PUT string = "P"
)

const (
	KRW string = "KRW"
	USD       = "USD"
	EUR       = "EUR"
	CNY       = "CNY"
)

type T비교결과 int

const (
	P같음   T비교결과 = 0
	P더_큼        = 1
	P더_작음       = -1
	P비교불가       = -999
)
