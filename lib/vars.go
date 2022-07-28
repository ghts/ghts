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

package lib

import (
	"github.com/ugorji/go/codec"
	"sync"
	"time"
)

var (
	ch공통_종료_채널 = make(chan T신호)

	json처리기    *codec.JsonHandle    = nil
	msgPack처리기 *codec.MsgpackHandle = nil

	한국증시_최근_개장일       = New안전한_시각(time.Time{})
	한국증시_최근_개장일_갱신_시점 = New안전한_시각(time.Time{})
	한국증기_최근_개장일_질의_잠금 sync.Mutex

	파일경로_맵_잠금 = sync.RWMutex{}
	파일경로_맵    = make(map[string]string)

	실행경로_수정_잠금 = new(sync.Mutex)

	// 이하 테스트 관련 함수 모음
	인터넷_접속_확인_잠금 sync.Mutex
	인터넷_접속됨      = true
	인터넷_접속_확인_완료 = false

	테스트_모드         = New안전한_bool(false)
	문자열_출력_일시정지_모드 = New안전한_bool(false)

	화면_출력_잠금 sync.Mutex

	문자열_출력_중복_방지_잠금 = new(sync.Mutex)
	문자열_출력_중복_방지_맵  = make(map[string]S비어있음)

	소켓_테스트용_주소_중복_방지_잠금 = new(sync.Mutex)
	소켓_테스트용_주소_중복_방지_맵  = make(map[string]S비어있음)

	체크포인트_잠금 = new(sync.Mutex)

	한국 = time.FixedZone("UTC+9", 9*60*60)
)
