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

# include <stdbool.h>
# include <windef.h>

typedef	BOOL(__stdcall *FuncBOOL)();

// 함수 실행과정에서 에러가 발생했는 지 여부와
// 함수 실행결과 얻은 bool형식의 값을
// 한꺼번에 Go언어로 전달하기 위한 구조체.
// C언어는 Go언어와 달리 복수 반환값을 지원하지 않아서 만들게 되었음.
typedef struct {
	bool Value;
	int ErrorCode;
} ErrBool;

