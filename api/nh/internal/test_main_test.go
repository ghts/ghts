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

package internal

import (
	공용 "github.com/ghts/ghts/common"
	ini "gopkg.in/ini.v1"

	"os"
	"strings"
	"testing"
	"time"
)

var 테스트용_ID string
var 테스트용_암호 string
var 테스트용_공인인증_암호 string
var 테스트용_계좌_번호 string
var 테스트용_계좌_유효기간 time.Time

func TestMain(m *testing.M) {
	공용.F테스트_모드_시작()
	defer 공용.F테스트_모드_종료()

	f테스트_설정_읽어오기()

	if NH_OpenAPI_Go루틴_실행_중.G값() {
		공용.New질의(공용.P메시지_종료).S질의(Ch종료).G회신()
	}

	초기화_대기 := make(chan bool)
	go F_NH_OpenAPI_Go루틴(초기화_대기)
	<-초기화_대기

	os.Exit(m.Run())
}

func f테스트_설정_읽어오기() {
	cfg, 에러 := ini.Load("config.ini")
	공용.F에러_패닉(에러)

	섹션, 에러 := cfg.GetSection("NH_OpenApi_LogIn_Info")
	공용.F에러_패닉(에러)

	키_ID, 에러 := 섹션.GetKey("ID")
	공용.F에러_패닉(에러)
	테스트용_ID = 키_ID.String()

	키_PWD, 에러 := 섹션.GetKey("PWD")
	공용.F에러_패닉(에러)
	테스트용_암호 = 키_PWD.String()

	//키_CertPWD, 에러 := 섹션.GetKey("CertPWD")
	//공용.F에러_패닉(에러)
	//테스트용_공인인증_암호 = 키_CertPWD.String()
	테스트용_공인인증_암호 = ""

	키_TestAccountNo, 에러 := 섹션.GetKey("TestAccountNo")
	공용.F에러_패닉(에러)
	테스트용_계좌_번호 = strings.Replace(키_TestAccountNo.String(), "-", "", -1)

	키_TestAccountValidUntil, 에러 := 섹션.GetKey("TestAccountValidUntil")
	공용.F에러_패닉(에러)
	유효기간_문자열 := 키_TestAccountValidUntil.String()
	테스트용_계좌_유효기간, 에러 = time.Parse("2006.01.02", 유효기간_문자열)
	공용.F에러_패닉(에러)
}
