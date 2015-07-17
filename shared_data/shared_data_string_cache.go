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

package shared_data

import (
	공용 "github.com/ghts/ghts/shared"
)

func f문자열_캐시_질의_처리(문자열_캐시_맵 map[string][]string, 질의 공용.I질의) {
	switch 질의.G구분() {
	case 공용.P메시지_GET:
		에러 := 질의.G검사(공용.P메시지_GET, 1)
		if 에러 != nil {
			질의.S회신(에러)
			break
		}

		문자열_모음, 존재함 := 문자열_캐시_맵[질의.G내용(0)]

		if !존재함 {
			에러 = 공용.F에러_생성("존재하지 않는 값. %s", 질의.G내용(0))
			질의.S회신(에러)
			break
		}

		회신_데이터 := 공용.F문자열_모음2인터페이스_모음(문자열_모음)

		질의.S회신(nil, 회신_데이터...)
	case 공용.P메시지_SET:
		if 질의.G길이() < 2 {
			질의.S회신(공용.F에러_생성("설정할 값이 없음.\n%v", 질의))
			break
		}

		문자열_캐시_맵[질의.G내용(0)] = 질의.G내용_전체()[1:]

		질의.S회신(nil)
	case 공용.P메시지_DEL:
		에러 := 질의.G검사(공용.P메시지_DEL, 1)
		if 에러 != nil {
			질의.S회신(에러)
			break
		}

		delete(문자열_캐시_맵, 질의.G내용(0))

		질의.S회신(nil)
	default:
		에러 := 공용.F에러_생성("예상치 못한 메시지 구분 %s.\n%v", 질의.G구분(), 질의)
		공용.F에러_출력(에러)
		질의.S회신(에러)
	}
}
