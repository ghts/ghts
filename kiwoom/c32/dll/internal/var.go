/* Copyright(C) 2015-2020년 김운하 (unha.kim@ghts.org)

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
	"github.com/ghts/ghts/lib"
	"github.com/ghts/ghts/lib/nanomsg"
	"github.com/ghts/ghts/lib/w32"
	"go.nanomsg.org/mangos/v3"
	"sync"
)

var (
	메인_윈도우 w32.HWND

	소켓REP_TR수신   = nanomsg.NewNano소켓XREP_단순형(lib.P주소_키움_C함수_호출)
	소켓PUB_실시간_정보 = nanomsg.NewNano소켓PUB_단순형(lib.P주소_키움_실시간)

	소켓REQ_저장소 = lib.New소켓_저장소(20, func() lib.I소켓_질의 {
		return nanomsg.NewNano소켓REQ_단순형(lib.P주소_키움_C함수_콜백, lib.P30초)
	})

	접속_처리_잠금  sync.Mutex
	api_호출_잠금 sync.Mutex

	Ch작업      = make(chan *lib.S작업, 100)
	Ch로그인     = make(chan bool, 1)
	Ch수신      = make(chan *mangos.Message, 1000)
	Ch질의      = make(chan *lib.S채널_질의_API, 100)
	ch콜백      = make(chan lib.I콜백, 100)
	Ch디버깅_메시지 = make(chan string, 1000)

	전달_도우미_수량 int
	콜백_도우미_수량 int

	Ch모니터링_루틴_종료       = make(chan lib.T신호, 1)
	Ch디버깅_메시지_출력_루틴_종료 = make(chan lib.T신호, 1)
	Ch수신_도우미_종료        = make(chan lib.T신호, 1)
	Ch전달_도우미_종료        = make(chan lib.T신호, 100)
	Ch콜백_도우미_종료        = make(chan lib.T신호, 100)
	Ch함수_호출_도우미_종료     = make(chan lib.T신호, 1)

	메시지_일련번호_생성기 = lib.New안전한_일련번호()
	S메시지_보관소     = New윈도우_메시지_보관소()
)

// 초기화 이후에는 사실상 읽기 전용이어서, 다중 사용에 문제가 없는 값들.
var (
	//설정파일_디렉토리 = filepath.Join(lib.GOPATH(), "src", reflect.TypeOf(S콜백_대기_저장소{}).PkgPath())
	//설정파일_경로   = filepath.Join(설정파일_디렉토리, "config.ini")
	계좌번호_모음 []string
	계좌_비밀번호 string
)
