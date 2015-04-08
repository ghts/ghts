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

Created on 2015. 3. 22.

@author: UnHa Kim <unha.kim@gh-system.com>

*/

package shared

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

type s종목 struct {
	코드 string
	이름 string
}

func (s *s종목) G코드() string {
	return s.코드
}

func (s *s종목) G이름() string {
	return s.이름
}
