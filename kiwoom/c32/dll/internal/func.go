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
	"github.com/ghts/ghts/lib/w32"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func F체크(값_모음 ...interface{}) {
	버퍼 := new(bytes.Buffer)

	if len(값_모음) == 0 {
		버퍼.WriteString("Check Point")
	} else {
		버퍼.WriteString(f변수값_문자열(값_모음...))
	}

	버퍼.WriteString(" ")
	버퍼.WriteString(f소스코드_위치(1))

	println(버퍼.String())
}

func f문자열_인코딩_변환(UTF8문자열 string) string {
	EUC_KR문자열, _, _ := transform.String(korean.EUCKR.NewDecoder(), UTF8문자열)

	return EUC_KR문자열
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

func F윈도우_핸들_설정(hWnd w32.HWND) {
	메인_윈도우 = hWnd
}

func F접속() {
	F체크("SendMessage 1")
	w32.SendMessage(메인_윈도우, KM_CONNECT, 0,0)
	F체크("SendMessage 2")

	//질의 := new(lib.S채널_질의_API)
	//질의.M질의값 = lib.New질의값_기본형(lib.TR접속, "")
	//질의.Ch회신값 = make(chan interface{}, 0)
	//질의.Ch에러 = make(chan error, 0)
	//
	//F체크("접속 질의 준비 완료.")
	//
	//Ch질의 <- 질의
	//
	//F체크("접속 채널 질의 전송 완료.")
	//
	//select {
	//case <-질의.Ch회신값:
	//case 에러 := <-질의.Ch에러:
	//	F체크("접속 질의 에러 발생.")
	//	println(에러.Error())
	//case <-lib.F공통_종료_채널():
	//	return
	//}
}

