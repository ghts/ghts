package xt

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	lb "github.com/ghts/ghts/lib"
	"gopkg.in/ini.v1"
)

func F로그인_설정_화일_읽기() (로그인_정보 *S로그인_정보, 에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	로그인_정보_화일_경로 := lb.F확인2(F로그인_설정_화일_경로())
	로그인_정보_문자열 := lb.F확인2(F로그인_정보_문자열_읽기(로그인_정보_화일_경로))
	V로그인_정보, 에러 = F로그인_정보_문자열_해석(로그인_정보_문자열)
	F로그인_설정_화일_경로_설정(로그인_정보_화일_경로) // 자식 프로세스 생성한다면 환경 변수를 통해서 정보 전달.

	return V로그인_정보, 에러
}

func F로그인_설정_화일_경로() (경로 string, 에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	경로_후보_모음 := []string{
		filepath.Join(lb.F확인2(os.Getwd()), P로그인_정보_화일명), /// 현재 디렉토리에 존재
		os.Getenv(P환경변수_설정_화일_경로),                       // 환경 변수를 통해서 부모 프로세스로부터 전달 받은 경우.
	}

	for _, 경로_후보 := range 경로_후보_모음 {
		if lb.F파일_존재함(경로_후보) {
			return 경로_후보, nil
		}
	}

	if 경로, 에러 = lb.F파일_검색(filepath.Join(lb.GOPATH(), "src", "github.com", "ghts"), P로그인_정보_화일명); 에러 == nil {
		return 경로, nil // 개발 환경
	}

	// 최후의 방법으로 Go루트 소스 디렉토리 검색
	return lb.F파일_검색(filepath.Join(lb.GOPATH(), "src"), P로그인_정보_화일명)
}

func F로그인_정보_문자열_읽기(로그인_정보_화일_경로 string) (로그인_정보_문자열 string, 에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	if lb.F파일_없음(로그인_정보_화일_경로) {
		버퍼 := new(bytes.Buffer)
		버퍼.WriteString("Xing 설정화일 찾을 수없음\n")
		버퍼.WriteString("'%v'가 존재하지 않습니다.\n")
		버퍼.WriteString("환경변수 '%v'에 설정화일 경로를 설정하십시오.\n")
		버퍼.WriteString("xing_config.ini.sample을 참조하여 새로 생성하십시오.")

		return "", lb.New에러(버퍼.String(), 로그인_정보_화일_경로, P환경변수_설정_화일_경로)
	}

	로그인_정보_화일 := lb.F확인2(os.Open(로그인_정보_화일_경로))
	defer 로그인_정보_화일.Close()

	바이트_모음 := lb.F확인2(io.ReadAll(로그인_정보_화일))

	return string(바이트_모음), nil
}

func F로그인_정보_문자열_해석(로그인_정보_문자열 string) (로그인_정보 *S로그인_정보, 에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	설정_화일 := lb.F확인2(ini.Load([]byte(로그인_정보_문자열)))
	섹션 := lb.F확인2(설정_화일.GetSection("XingAPI_LogIn_Info"))

	로그인_정보 = new(S로그인_정보)
	로그인_정보.M로그인_ID = lb.F확인2(섹션.GetKey("ID")).String()
	로그인_정보.M로그인_암호 = lb.F확인2(섹션.GetKey("PWD")).String()
	로그인_정보.M인증서_암호 = lb.F확인2(섹션.GetKey("CertPWD")).String()
	로그인_정보.M계좌_비밀번호 = lb.F확인2(섹션.GetKey("AcctPWD")).String()

	// 모의투자 암호는 선택사항.
	if 키, 에러 := 섹션.GetKey("TestPWD"); 에러 != nil {
		lb.F문자열_출력("TestPWD 값 조회 에러 발생.")
	} else if 값 := lb.F2문자열_공백_제거(키.String()); 값 == "" {
		lb.F문자열_출력("비어 있는 TestPWD 값.")
	} else {
		로그인_정보.M모의투자_암호 = 값
	}

	return 로그인_정보, nil
}

// F로그인_설정_화일_경로_설정 : 64비트 프로세스에서 DLL호출 전용 32비트 자식 프로세스를 생성할 때 환경 변수를 통해서 필요한 정보를 전달.
func F로그인_설정_화일_경로_설정(경로 string) {
	if 에러 := os.Setenv(P환경변수_설정_화일_경로, 경로); 에러 != nil {
		panic(에러)
	}
}

// F로그인_정보_환경_변수_설정 : 64비트 프로세스에서 DLL호출 전용 32비트 자식 프로세스를 생성할 때 환경 변수를 통해서 로그인 정보를 전달.
func F로그인_정보_환경_변수_설정(로그인_정보 *S로그인_정보) (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	lb.F확인1(os.Setenv(P환경변수_로그인_ID, 로그인_정보.M로그인_ID))
	lb.F확인1(os.Setenv(P환경변수_로그인_암호, 로그인_정보.M로그인_암호))
	lb.F확인1(os.Setenv(P환경변수_인증서_암호, 로그인_정보.M인증서_암호))
	lb.F확인1(os.Setenv(P환경변수_계좌_비밀번호, 로그인_정보.M계좌_비밀번호))

	로그인_정보.M모의투자_암호 = lb.F2문자열_공백_제거(로그인_정보.M모의투자_암호)
	if 로그인_정보.M모의투자_암호 != "" {
		lb.F확인1(os.Setenv(P환경변수_모의투자_암호, 로그인_정보.M모의투자_암호))
	}

	return nil
}

// F로그인_정보_설정 : 32비트 프로세스에서 환경 변수를 통해서 부모 프로세스의 정보가 전달해 준 정보를 받음.
func F로그인_정보_설정() (에러 error) {
	defer lb.S예외처리{M에러: &에러}.S실행()

	if V로그인_정보 != nil {
		return nil
	} else if 로그인_ID := os.Getenv(P환경변수_로그인_ID); 로그인_ID == "" {
		return lb.New에러("로그인 ID값이 설정되어 있지 않습니다. 환경변수 '%v'에 설정하십시오.", P환경변수_로그인_ID)
	} else if 로그인_암호 := os.Getenv(P환경변수_로그인_암호); 로그인_암호 == "" {
		return lb.New에러("로그인 암호가 설정되어 있지 않습니다. 환경변수 '%v'에 설정하십시오.", P환경변수_로그인_암호)
	} else if 인증서_암호 := os.Getenv(P환경변수_인증서_암호); 인증서_암호 == "" {
		return lb.New에러("인증서 암호가 설정되어 있지 않습니다. 환경변수 '%v'에 설정하십시오.", P환경변수_인증서_암호)
	} else if 계좌_비밀번호 := os.Getenv(P환경변수_계좌_비밀번호); 계좌_비밀번호 == "" {
		return lb.New에러("계좌 암호가 설정되어 있지 않습니다. 환경변수 '%v'에 설정하십시오.", P환경변수_계좌_비밀번호)
	} else {
		모의투자_암호 := os.Getenv(P환경변수_모의투자_암호)
		if 모의투자_암호 == "" {
			lb.F문자열_출력("모의투자 암호가 설정되어 있지 않습니다. 환경변수 '%v'에 설정하십시오.", P환경변수_모의투자_암호)
		}

		V로그인_정보 = new(S로그인_정보)
		V로그인_정보.M로그인_ID = 로그인_ID
		V로그인_정보.M로그인_암호 = 로그인_암호
		V로그인_정보.M인증서_암호 = 인증서_암호
		V로그인_정보.M계좌_비밀번호 = 계좌_비밀번호
		V로그인_정보.M모의투자_암호 = 모의투자_암호

		F로그인_정보_환경_변수_삭제()

		return nil
	}
}

// F로그인_정보_설정 : 64비트 부모 프로세스와 32비트 자식 프로세스 간 정보 전달이 끝나면 메모리에서 흔적 삭제.
func F로그인_정보_환경_변수_삭제() {
	lb.F확인1(os.Setenv(P환경변수_로그인_ID, ""))
	lb.F확인1(os.Setenv(P환경변수_로그인_암호, ""))
	lb.F확인1(os.Setenv(P환경변수_인증서_암호, ""))
	lb.F확인1(os.Setenv(P환경변수_계좌_비밀번호, ""))
	lb.F확인1(os.Setenv(P환경변수_모의투자_암호, ""))
}
