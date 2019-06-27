/* Copyright (C) 2015-2019 김운하(UnHa Kim)  < unha.kim.ghts at gmail dot com >

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

Copyright (C) 2015-2019년 UnHa Kim (< unha.kim.ghts at gmail dot com >)

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

package lib

import (
	"github.com/ugorji/go/codec"

	"bytes"
	"errors"
	"math/big"
	"reflect"
)

func init() {
	f메모()
	f인코딩_디코딩_초기화()
}

func f인코딩_디코딩_초기화() (에러 error) {
	defer S예외처리{M에러: &에러}.S실행()

	json처리기 = new(codec.JsonHandle)
	msgPack처리기 = new(codec.MsgpackHandle)

	// 업데이트 된 MessagePack 규격 적용.
	msgPack처리기.WriteExt = true

	자료형_모음 := []reflect.Type{
		reflect.TypeOf(errors.New("")).Elem(),
		reflect.TypeOf(new(big.Int)).Elem(),
		reflect.TypeOf(new(big.Rat)).Elem(),
		reflect.TypeOf(new(big.Float)).Elem()}

	for i, 자료형 := range 자료형_모음 {
		인덱스 := uint64(i + 1)

		F확인(json처리기.SetInterfaceExt(자료형, 인덱스, s변환기{자료형.String()}))
		F확인(msgPack처리기.SetBytesExt(자료형, 인덱스, s변환기{자료형.String()}))
	}

	return nil
}

func f메모() {
	버퍼 := new(bytes.Buffer)
	버퍼.WriteString("\n")
	버퍼.WriteString("최소 호가 단위가 변했는 지 확인해 볼 것.\n")
	버퍼.WriteString("\n")
	버퍼.WriteString("1. 문제를 잘게 쪼갠다.\n")
	버퍼.WriteString("2. 문제를 단순화 시킨다. 예:) 테스트 데이터를 1개로 한정.\n")
	버퍼.WriteString("3. 내부 실행과정을 추적한다.\n")
	버퍼.WriteString("디버깅에 이보다 달리 더 좋은 방법은 없다.\n")
	버퍼.WriteString("\n")
	버퍼.WriteString("- 테스트 커버리지 확인 후 누락된 테스트 작성할 것.\n")
	버퍼.WriteString("- C언어는 컴파일할 때 -Wall옵션을 사용할 것.\n")
	버퍼.WriteString("- REQ소켓 재활용 방안 검토할 것.\n")
	버퍼.WriteString("- C 및 Go변환 구조체 크기 비교 테스트 할 것.\n")
	버퍼.WriteString("- TestMain()에서 환경변수 설정 가능한 지 확인할 것.\n")

	F메모(버퍼.String())
}
