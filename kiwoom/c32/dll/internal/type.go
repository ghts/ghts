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
	"sync"
	"time"
)

func New윈도우_메시지_항목() *S윈도우_메시지_항목 {
	return &S윈도우_메시지_항목{
		M메시지_일련번호: F메시지_일련번호(),
		Ch회신:      make(chan string, 1),
		M보관_시점:    time.Now()}
}

type S윈도우_메시지_항목 struct {
	M메시지_일련번호 uintptr
	Ch회신      chan string
	M보관_시점    time.Time
}

func New윈도우_메시지_보관소() *S윈도우_메시지_보관소 {
	s := new(S윈도우_메시지_보관소)
	s.보관소 = make(map[uintptr]*S윈도우_메시지_항목)

	return s
}

type S윈도우_메시지_보관소 struct {
	sync.Mutex
	보관소 map[uintptr]*S윈도우_메시지_항목
}

func (s *S윈도우_메시지_보관소) S보관(항목 *S윈도우_메시지_항목) {
	s.Lock()
	defer s.Unlock()

	s.보관소[항목.M메시지_일련번호] = 항목
}

func (s *S윈도우_메시지_보관소) S삭제(일련번호 uintptr) {
	s.Lock()
	defer s.Unlock()

	delete(s.보관소, 일련번호)
}

func (s *S윈도우_메시지_보관소) S회신(일련번호 uintptr, 회신값 string) error {
	defer s.S삭제(일련번호)

	if 항목, ok := s.보관소[일련번호]; !ok {
		return lib.New에러("해당 일련번호의 메시지 보관 항목이 존재하지 않음 : '%v'", 일련번호)
	} else {
		항목.Ch회신 <- 회신값
	}

	return nil
}

func (s *S윈도우_메시지_보관소) S정리() {
	s.Lock()
	defer s.Unlock()

	일분전 := time.Now().Add(-1 * lib.P1분)

	for 일련번호, 값 := range s.보관소 {
		if 값.M보관_시점.Before(일분전) {
			delete(s.보관소, 일련번호)
		}
	}
}
