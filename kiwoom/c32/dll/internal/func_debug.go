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
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func F체크(값_모음 ...interface{}) {
	버퍼 := new(bytes.Buffer)

	if len(값_모음) == 0 {
		버퍼.WriteString("체크포인트")
	} else {
		버퍼.WriteString(f변수값_문자열(값_모음...))
	}

	버퍼.WriteString(" ")
	버퍼.WriteString(f소스코드_위치(1))

	//println(버퍼.String())

	Ch디버깅_메시지 <- 버퍼.String()
}

func f소스코드_위치(건너뛰는_단계 int) string {
	건너뛰는_단계++ // 이 메소드를 호출한 함수를 기준으로 0이 되게 하기 위함.

	_, 파일_경로, 행_번호, _ := runtime.Caller(건너뛰는_단계)

	var 파일명 string
	시작점 := strings.Index(파일_경로, "github.com")
	if 시작점 >= 0 && 시작점 < len(파일_경로) {
		파일명 = 파일_경로[시작점:]
	} else {
		파일명 = filepath.Base(파일_경로)
	}

	return 파일명 + ":" + strconv.Itoa(행_번호)
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
	F계좌_수량_테스트()
	F전체_계좌_번호_테스트()
	F사용자_ID_테스트()
	F사용자_이름_테스트()
	F키보드_보안_상태_테스트()
	F방화벽_상태_테스트()
}

func F계좌_수량_테스트() {
	F체크("F계좌_수량_테스트()")

	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P전체_계좌_수량)))
	F체크("Account Qty query ready.")

	Ch질의 <- 질의
	F체크("Account Qerrty query sent.")

	select {
	case 회신값 := <-질의.Ch회신값:
		if 수량, ok := 회신값.(int); !ok {
			F체크(lib.F2문자열("Account Qty Test Failure. Unexpected data type : '%T'", 회신값))
		} else {
			F체크(lib.F2문자열("Account Qty Test Success : '%v'", 수량))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("Account Qty Test Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}

func F전체_계좌_번호_테스트() {
	F체크("Account No Test.")

	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P전체_계좌_번호)))
	F체크("Account No query ready.")

	Ch질의 <- 질의
	F체크("Account No query sent.")

	select {
	case 회신값 := <-질의.Ch회신값:
		if 계좌_번호_모음, ok := 회신값.([]string); !ok {
			F체크(lib.F2문자열("Account No Test Failure. Unexpected data type : '%T'", 회신값))
		} else {
			F체크(lib.F2문자열("Account No Test Success : '%v'", 계좌_번호_모음))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("Account No Test Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}

func F사용자_ID_테스트() {
	F체크("User ID Test.")

	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P사용자_ID)))
	F체크("User ID query ready.")

	Ch질의 <- 질의
	F체크("User ID query sent.")

	select {
	case 회신값 := <-질의.Ch회신값:
		if 사용자_ID, ok := 회신값.(string); !ok {
			F체크(lib.F2문자열("User ID Test Failure. Unexpected data type : '%T'", 회신값))
		} else {
			F체크(lib.F2문자열("User ID Test Success : '%v'", 사용자_ID))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("User ID Test Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}

func F사용자_이름_테스트() {
	F체크("User Name Test.")

	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P사용자_이름)))
	F체크("User Name query ready.")

	Ch질의 <- 질의
	F체크("User Name query sent.")

	select {
	case 회신값 := <-질의.Ch회신값:
		if 사용자_이름, ok := 회신값.(string); !ok {
			F체크(lib.F2문자열("User Name Test Failure. Unexpected data type : '%T'", 회신값))
		} else {
			F체크(lib.F2문자열("User Name Test Success : '%v'", 사용자_이름))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("User Name Test Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}

func F키보드_보안_상태_테스트() {
	F체크("Keyboard Security On/Off Test.")

	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P키보드_보안_상태)))
	F체크("Keyboard Security On/Off query ready.")

	Ch질의 <- 질의
	F체크("Keyboard Security On/Off query sent.")

	select {
	case 회신값 := <-질의.Ch회신값:
		if 키보드_보안_On_Off, ok := 회신값.(bool); !ok {
			F체크(lib.F2문자열("Keyboard Security On/Off Test Failure. Unexpected data type : '%T'", 회신값))
		} else {
			F체크(lib.F2문자열("Keyboard Security On/Off Test Success : '%v'", 키보드_보안_On_Off))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("Keyboard Security On/Off Test Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}

func F방화벽_상태_테스트() {
	F체크("Firewall On/Off Test.")

	질의 := lib.New채널_질의_API(lib.New질의값_정수(kt.TR로그인_정보, "", int(kt.P방화벽_상태)))
	F체크("Firewall On/Off query ready.")

	Ch질의 <- 질의
	F체크("Firewall On/Off query sent.")

	select {
	case 회신값 := <-질의.Ch회신값:
		if 방화벽_상태, ok := 회신값.(kt.T방화벽_상태); !ok {
			F체크(lib.F2문자열("Firewall On/Off Test Failure. Unexpected data type : '%T'", 회신값))
		} else {
			F체크(lib.F2문자열("Firewall On/Off Test Success : '%v'", 방화벽_상태))
		}
	case 에러 := <-질의.Ch에러:
		F체크(lib.F2문자열("Firewall On/Off Test Error : '%v'.", 에러.Error()))
	case <-lib.F공통_종료_채널():
		return
	}
}
