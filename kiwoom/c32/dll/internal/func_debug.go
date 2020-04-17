/* Copyright(C) 2015-2020년 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

키움증권 API는 OCX규격으로 작성되어 있는 데,
OCX규격은 Go언어로 직접 사용하기에 기술적 난이도가 높아서,
손쉽게 다룰 수 있게 도와주는 Qt 라이브러리의 오픈소스 버전을 사용하였습니다.

Qt 라이브러리의 오픈소스 버전의 경우
GHTS의 대부분에서 사용하고 있는 GNU LGPL v 2.1보다
좀 더 강력하고 엄격한 소스코드 공개 의무가 있는 GNU GPL v2를 사용해야 합니다.

이는 개발 난이도 경감을 위한 개발자의 필요에 의한 것이며,
사용자에게 GPL v2의 소스코드 공개 의무를 강제하려는 의도는 아닙니다.

키움증권 API 호출 모듈에 적용된 GPL v2이 LGPL v2보다 더 엄격하긴 합니다만,
키움증권 API 호출 모듈을 애초 의도된 사용법대로 '소켓을 통해서 호출'하여 사용하는 경우에는
GPL v2에서 규정하는 '하나의 단일 소프트웨어' 규정에 포함되지 않기에
사용자가 작성한 소스코드는 GPL v2의 소스코드 공개 의무가 적용되지 않습니다.

다만, 키움증권 API 호출 모듈 그 자체를 수정하거나 타인에게 배포할 경우,
GPL v2 규정에 따른 소스코드 공개 의무가 발생할 수 있습니다.

---------------------------------------------------------

이 프로그램은 자유 소프트웨어입니다.
소프트웨어의 피양도자는 자유 소프트웨어 재단이 공표한 GNU GPL v2
규정에 따라 프로그램을 개작하거나 재배포할 수 있습니다.

이 프로그램은 유용하게 사용될 수 있으리라는 희망에서 배포되고 있지만,
특정한 목적에 적합하다거나, 이익을 안겨줄 수 있다는 묵시적인 보증을 포함한
어떠한 형태의 보증도 제공하지 않습니다.
보다 자세한 사항에 대해서는 GNU GPL v2를 참고하시기 바랍니다.
GNU GPL v2는 이 프로그램과 함께 제공됩니다.

만약, 이 문서가 누락되어 있다면 자유 소프트웨어 재단으로 문의하시기 바랍니다.
(자유 소프트웨어 재단 : Free Software Foundation, Inc.,
59 Temple Place - Suite 330, Boston, MA 02111-1307, USA) */

package k32

import (
	"bytes"
	"fmt"
	kt "github.com/ghts/ghts/kiwoom/base"
	"github.com/ghts/ghts/lib"
)

func F체크(값_모음 ...interface{}) {
	버퍼 := new(bytes.Buffer)

	if len(값_모음) == 0 {
		버퍼.WriteString("체크포인트")
	} else {
		버퍼.WriteString(f변수값_문자열(값_모음...))
	}

	버퍼.WriteString(" ")
	버퍼.WriteString(lib.F소스코드_위치(1))

	Ch디버깅_메시지 <- 버퍼.String()
}

func f변수값_문자열(값_모음 ...interface{}) string {
	버퍼 := new(bytes.Buffer)

	for i, _ := range 값_모음 {
		if i == 0 {
			버퍼.WriteString("'%v'")
		} else {
			버퍼.WriteString(", '%v'")
		}
	}

	return fmt.Sprintf(버퍼.String(), 값_모음...)
}

func F로그인_정보_테스트() {
	//F계좌_수량_테스트()
	//F전체_계좌_번호_테스트()
	//F사용자_ID_테스트()
	//F사용자_이름_테스트()
	//F키보드_보안_상태_테스트()
	//F방화벽_상태_테스트()
	F접속_상태_테스트()
}

func F계좌_수량_테스트() {
	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P전체_계좌_수량)))

	Ch질의 <- 질의

	select {
	case 회신값 := <-질의.Ch회신값:
		if 수량, ok := 회신값.(int); !ok {
			F체크(lib.F2문자열("F계좌_수량_테스트() 예상하지 못한 자료형 : '%T'", 회신값))
		} else {
			F체크(lib.F2문자열("F계좌_수량_테스트() OK : '%v'", 수량))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("F계좌_수량_테스트() Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}

func F전체_계좌_번호_테스트() {
	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P전체_계좌_번호)))

	Ch질의 <- 질의

	select {
	case 회신값 := <-질의.Ch회신값:
		if 계좌_번호_모음, ok := 회신값.([]string); !ok {
			F체크(lib.F2문자열("F전체_계좌_번호_테스트() 예상하지 못한 자료형 : '%T'", 회신값))
		} else {
			F체크(lib.F2문자열("F전체_계좌_번호_테스트() OK. '%v'", 계좌_번호_모음))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("F전체_계좌_번호_테스트() Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}

func F사용자_ID_테스트() {
	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P사용자_ID)))

	Ch질의 <- 질의

	select {
	case 회신값 := <-질의.Ch회신값:
		if 사용자_ID, ok := 회신값.(string); !ok {
			F체크(lib.F2문자열("F사용자_ID_테스트() 예상하지 못한 자료형 : '%T'", 회신값))
		} else if 사용자_ID == "" {
			F체크(lib.F2문자열("F사용자_ID_테스트() Error. 비어있는 회신값."))
		} else {
			F체크(lib.F2문자열("F사용자_ID_테스트() OK : '%v'", 회신값))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("F사용자_ID_테스트() Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}

func F사용자_이름_테스트() {
	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P사용자_이름)))

	Ch질의 <- 질의

	select {
	case 회신값 := <-질의.Ch회신값:
		if 사용자_이름, ok := 회신값.(string); !ok {
			F체크(lib.F2문자열("F사용자_이름_테스트() 예상하지 못한 자료형 : '%T'", 회신값))
		} else if 사용자_이름 == "" {
			F체크(lib.F2문자열("F사용자_이름_테스트() Error. 비어있는 회신값."))
		} else {
			F체크(lib.F2문자열("F사용자_이름_테스트() OK : '%v'", 사용자_이름))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("F사용자_이름_테스트() Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}

func F키보드_보안_상태_테스트() {
	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P키보드_보안_상태)))

	Ch질의 <- 질의

	select {
	case 회신값 := <-질의.Ch회신값:
		if 키보드_보안_On_Off, ok := 회신값.(bool); !ok {
			F체크(lib.F2문자열("F키보드_보안_상태_테스트() 예상하지 못한 자료형 : '%T'", 회신값))
		} else {
			F체크(lib.F2문자열("F키보드_보안_상태_테스트() OK : '%v'", 키보드_보안_On_Off))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("F키보드_보안_상태_테스트() Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}

func F방화벽_상태_테스트() {
	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P방화벽_상태)))

	Ch질의 <- 질의

	select {
	case 회신값 := <-질의.Ch회신값:
		if 방화벽_상태, ok := 회신값.(kt.T방화벽_상태); !ok {
			F체크(lib.F2문자열("F방화벽_상태_테스트() 예상하지 못한 자료형 : '%T'", 회신값))
		} else {
			F체크(lib.F2문자열("F방화벽_상태_테스트() OK : '%v'", 방화벽_상태))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("F방화벽_상태_테스트() Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}

func F접속_상태_테스트() {
	질의 := lib.New채널_질의_API(lib.New질의값_기본형(kt.TR접속_상태, ""))

	Ch질의 <- 질의

	select {
	case 회신값 := <-질의.Ch회신값:
		if 접속_상태, ok := 회신값.(bool); !ok {
			F체크(lib.F2문자열("F접속_상태_테스트() 예상하지 못한 자료형 : '%T'", 회신값))
		} else {
			F체크(lib.F2문자열("F접속_상태_테스트() OK : '%v'", 접속_상태))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("F접속_상태_테스트() Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}