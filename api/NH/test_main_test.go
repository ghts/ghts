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

package NH

import (
	공용 "github.com/ghts/ghts/common"
	"golang.org/x/sys/windows"
	ini "gopkg.in/ini.v1"

	"os"
	"testing"
)

var ID_테스트용 string
var 암호_테스트용 string
var 공인인증_암호_테스트용 string
var 계좌번호_테스트용 string
var 유효기간_테스트용_계좌 time.Time

func TestMain(m *testing.M) {
	f테스트_설정_읽어오기()

	F테스트_모드_시작()
	defer F테스트_모드_종료()

	os.Exit(m.Run())
}

func f테스트_설정_읽어오기() {
	cfg, 에러 := ini.Load("test_cfg.ini")
	공용.F에러_패닉(에러)

	섹션, 에러 := cfg.GetSection("NH_OpenApi_LogIn_Info")
	공용.F에러_패닉(에러)

	키, 에러 := 섹션.GetKey("ID")
	공용.F에러_패닉(에러)
	ID_테스트용 := 키.String()

	키, 에러 := 섹션.GetKey("PWD")
	공용.F에러_패닉(에러)
	암호_테스트용 := 키.String()

	키, 에러 := 섹션.GetKey("CertPWD")
	공용.F에러_패닉(에러)
	공인인증_암호_테스트용 := 키.String()

	키, 에러 := 섹션.GetKey("TestAccountNo")
	공용.F에러_패닉(에러)
	계좌번호_테스트용 := 키.String()

	키, 에러 := 섹션.GetKey("TestAccountValidUntil")
	공용.F에러_패닉(에러)
	유효기간_문자열 := 키.String()
	유효기간_테스트용_계좌, 에러 := time.Parse(유효기간_문자열, "2006.01.02")
	공용.F에러_패닉(에러)
}
