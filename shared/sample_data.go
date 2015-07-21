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

func F샘플_종목_모음() []I종목 {
	샘플_종목_모음 := make([]I종목, 0)
	샘플_종목_모음 = append(샘플_종목_모음, New종목("000020", "동화약품"))
	샘플_종목_모음 = append(샘플_종목_모음, New종목("000030", "우리은행"))
	샘플_종목_모음 = append(샘플_종목_모음, New종목("000040", "KR모터스"))
	샘플_종목_모음 = append(샘플_종목_모음, New종목("000050", "경방"))
	샘플_종목_모음 = append(샘플_종목_모음, New종목("000060", "메리츠화재"))

	return 샘플_종목_모음
}

func F샘플_통화단위_모음() []string {
	샘플_통화단위_모음 := make([]string, 0)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, KRW)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, USD)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, CNY)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, EUR)
	
	return 샘플_통화단위_모음
}
