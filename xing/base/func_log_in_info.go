/* Copyright (C) 2015-2022 김운하 (unha.kim@ghts.org)

이 파일은 GHTS의 일부입니다.

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU LGPL 2.1판
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU LGPL 2.1판을 참고하시기 바랍니다.
GNU LGPL 2.1판은 이 프로그램과 함께 제공됩니다.
만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA)

Copyright (C) 2015-2022년 UnHa Kim (unha.kim@ghts.org)

This file is part of GHTS.

GHTS is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 2.1 of the License.

GHTS is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with GHTS.  If not, see <http://www.gnu.org/licenses/>. */

package xt

import (
	"bytes"
	"github.com/ghts/ghts/lib"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

func F로그인_설정_화일_경로_설정(경로 string) {
	os.Setenv(P환경변수_설정_화일_경로, 경로)
}

func F로그인_설정_화일_경로() string {
	if 현재_디렉토리, 에러 := os.Getwd(); 에러 == nil {
		if 경로 := filepath.Join(현재_디렉토리, "xing_config.ini"); lib.F파일_존재함(경로) {
			return 경로
		}
	}

	return os.Getenv(P환경변수_설정_화일_경로)
}

func F로그인_설정_화일_읽기() (로그인_정보 *S로그인_정보, 에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	설정_화일_경로 := F로그인_설정_화일_경로()

	if lib.F파일_없음(설정_화일_경로) {
		버퍼 := new(bytes.Buffer)
		버퍼.WriteString("Xing 설정화일 찾을 수없음\n")
		버퍼.WriteString("'%v'가 존재하지 않습니다.\n")
		버퍼.WriteString("환경변수 '%v'에 설정화일 경로를 설정하십시오.\n")
		버퍼.WriteString("xing_config.ini.sample을 참조하여 새로 생성하십시오.")

		return nil, lib.New에러(버퍼.String(), 설정_화일_경로, P환경변수_설정_화일_경로)
	}

	cfg파일 := lib.F확인(ini.Load(설정_화일_경로)).(*ini.File)
	섹션 := lib.F확인(cfg파일.GetSection("XingAPI_LogIn_Info")).(*ini.Section)

	로그인_정보 = new(S로그인_정보)
	로그인_정보.M로그인_ID = lib.F확인(섹션.GetKey("ID")).(*ini.Key).String()
	로그인_정보.M로그인_암호 = lib.F확인(섹션.GetKey("PWD")).(*ini.Key).String()
	로그인_정보.M인증서_암호 = lib.F확인(섹션.GetKey("CertPWD")).(*ini.Key).String()
	로그인_정보.M계좌_비밀번호 = lib.F확인(섹션.GetKey("AcctPWD")).(*ini.Key).String()

	// 모의투자 암호는 선택사항.
	if 키, 에러 := 섹션.GetKey("TestPWD"); 에러 != nil {
		lib.F문자열_출력("TestPWD 값 조회 에러 발생.")
	} else if 값 := 키.String(); 값 == "" {
		lib.F문자열_출력("비어 있는 TestPWD 값.")
	} else {
		로그인_정보.M모의투자_암호 = 값
	}

	return 로그인_정보, nil
}

func F로그인_정보_환경_변수_설정(로그인_정보 *S로그인_정보) (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F확인(os.Setenv(P환경변수_로그인_ID, 로그인_정보.M로그인_ID))
	lib.F확인(os.Setenv(P환경변수_로그인_암호, 로그인_정보.M로그인_암호))
	lib.F확인(os.Setenv(P환경변수_인증서_암호, 로그인_정보.M인증서_암호))
	lib.F확인(os.Setenv(P환경변수_계좌_비밀번호, 로그인_정보.M계좌_비밀번호))
	lib.F확인(os.Setenv(P환경변수_모의투자_암호, 로그인_정보.M모의투자_암호))

	return nil
}

func F로그인_정보_환경_변수_읽기() (로그인_정보 *S로그인_정보, 에러 error) {
	if 로그인_ID := os.Getenv(P환경변수_로그인_ID); 로그인_ID == "" {
		return nil,
			lib.New에러("로그인 ID값이 설정되어 있지 않습니다. 환경변수 '%v'에 설정하십시오.", P환경변수_로그인_ID)
	} else if 로그인_암호 := os.Getenv(P환경변수_로그인_암호); 로그인_암호 == "" {
		return nil, lib.New에러("로그인 암호가 설정되어 있지 않습니다. 환경변수 '%v'에 설정하십시오.", P환경변수_로그인_암호)
	} else if 인증서_암호 := os.Getenv(P환경변수_인증서_암호); 인증서_암호 == "" {
		return nil, lib.New에러("인증서 암호가 설정되어 있지 않습니다. 환경변수 '%v'에 설정하십시오.", P환경변수_인증서_암호)
	} else if 계좌_비밀번호 := os.Getenv(P환경변수_계좌_비밀번호); 계좌_비밀번호 == "" {
		return nil, lib.New에러("계좌 암호가 설정되어 있지 않습니다. 환경변수 '%v'에 설정하십시오.", P환경변수_계좌_비밀번호)
	} else {
		모의투자_암호 := os.Getenv(P환경변수_모의투자_암호)
		if 모의투자_암호 == "" {
			lib.F문자열_출력("모의투자 암호가 설정되어 있지 않습니다. 환경변수 '%v'에 설정하십시오.", P환경변수_모의투자_암호)
		}

		로그인_정보 = new(S로그인_정보)
		로그인_정보.M로그인_ID = 로그인_ID
		로그인_정보.M로그인_암호 = 로그인_암호
		로그인_정보.M인증서_암호 = 인증서_암호
		로그인_정보.M계좌_비밀번호 = 계좌_비밀번호
		로그인_정보.M모의투자_암호 = 모의투자_암호

		return 로그인_정보, nil
	}
}

func F로그인_정보_환경_변수_삭제() {
	os.Setenv(P환경변수_로그인_ID, "")
	os.Setenv(P환경변수_로그인_암호, "")
	os.Setenv(P환경변수_인증서_암호, "")
	os.Setenv(P환경변수_계좌_비밀번호, "")
	os.Setenv(P환경변수_모의투자_암호, "")
}
