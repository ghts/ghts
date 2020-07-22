/* Copyright (C) 2015-2020 김운하 (unha.kim@ghts.org)

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

Copyright (C) 2015-2020년 UnHa Kim (unha.kim@ghts.org)

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
	"math/big"
	"testing"
	"time"
)

func TestS변환기(t *testing.T) {
	t.Parallel()

	var i변환기 interface{} = new(s변환기)

	switch i변환기.(type) {
	default:
		F문자열_출력("예상하지 못한 경우.")
		t.FailNow()
	case codec.Ext:
		// OK
	}
}

func TestS바이트_변환_매개체(t *testing.T) {
	t.Parallel()

	변환형식_모음 := f테스트용_변환형식_모음()
	값_모음 := f테스트용_변환가능한_전달값_모음()

	바이트_변환_매개체, 에러 := New바이트_변환(P변환형식_기본값, 1000)
	F테스트_에러없음(t, 에러)
	값_모음 = append(값_모음, 바이트_변환_매개체)

	for _, 변환형식 := range 변환형식_모음 {
		for _, 값 := range 값_모음 {
			바이트_변환_매개체1, 에러 := New바이트_변환(변환형식, 값)
			F테스트_에러없음(t, 에러)

			원래값_바이트_변환값_비교(t, 값, 바이트_변환_매개체1)

			바이트_모음, 에러 := 바이트_변환_매개체1.MarshalBinary()
			F테스트_에러없음(t, 에러)

			바이트_변환_매개체2 := new(S바이트_변환)
			에러 = 바이트_변환_매개체2.UnmarshalBinary(바이트_모음)
			F테스트_에러없음(t, 에러)

			원래값_바이트_변환값_비교(t, 값, 바이트_변환_매개체2)
		}
	}
}

func TestS바이트_변환_매개체_모음(t *testing.T) {
	t.Parallel()

	변환형식_모음 := f테스트용_변환형식_모음()
	//값_모음 := f테스트용_변환가능한_전달값_모음()
	값_모음 := []interface{}{100, "문자열"}

	for _, 변환형식 := range 변환형식_모음 {
		바이트_변환_매개체, 에러 := New바이트_변환_모음(변환형식, 값_모음...)
		F테스트_에러없음(t, 에러)

		바이트_모음, 에러 := 바이트_변환_매개체.MarshalBinary()
		F테스트_에러없음(t, 에러)

		바이트_변환_매개체2 := new(S바이트_변환_모음)
		에러 = 바이트_변환_매개체2.UnmarshalBinary(바이트_모음)
		F테스트_에러없음(t, 에러)

		for i := 0; i < 바이트_변환_매개체2.G수량(); i++ {
			원래값_바이트_변환값_비교(t, 값_모음[i], 바이트_변환_매개체2.M바이트_변환_모음[i])
		}
	}
}

func 원래값_바이트_변환값_비교(t *testing.T, 원래값 interface{}, 바이트_전송값 *S바이트_변환) {
	var 에러 error = nil

	switch 원래값.(type) {
	case nil:
		복원값 := interface{}(nil)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		// 값은 nil이지만, 자료형은 nil이 아님.
		// 진짜 nil을 전달받는 방법은 모르겠음.
		F테스트_참임(t, 바이트_전송값.IsNil())
		F테스트_참임(t, 복원값 == nil)
	case int:
		복원값 := 0
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case uint:
		복원값 := uint(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case uintptr:
		복원값 := uintptr(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case int8:
		복원값 := int8(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case int16:
		복원값 := int16(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case int32:
		복원값 := int32(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case int64:
		복원값 := int64(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case uint8:
		복원값 := uint8(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case uint16:
		복원값 := uint16(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case uint32:
		복원값 := uint32(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case uint64:
		복원값 := uint64(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case float32:
		복원값 := float32(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case float64:
		복원값 := float64(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case bool:
		복원값 := true
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case string:
		복원값 := ""
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case []byte:
		var 복원값 []byte
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case error:
		var 복원값 error
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case time.Time:
		복원값 := time.Time{}
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case []string:
		var 복원값 []string
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case s샘플_구조체_1:
		복원값 := s샘플_구조체_1{}
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case *s샘플_구조체_1:
		복원값 := new(s샘플_구조체_1)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case s샘플_구조체_2:
		복원값 := s샘플_구조체_2{}
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case *s샘플_구조체_2:
		복원값 := new(s샘플_구조체_2)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case *big.Int:
		복원값 := big.NewInt(0)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case *big.Rat:
		복원값 := big.NewRat(1, 1)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case *big.Float:
		복원값 := new(big.Float)
		에러 = 바이트_전송값.G값(&복원값)
		F테스트_에러없음(t, 에러)
		F테스트_같음(t, 원래값, 복원값)
	case *S바이트_변환:
		F테스트_같음(t, 원래값, 바이트_전송값)
	default:
		New에러with출력("예상치 못한 자료형. %T", 원래값)
		t.FailNow()
	}
}

func TestS바이트_변환_S질의값_바이트_변환(t *testing.T) {
	지금 := time.Now()
	값1 := New질의값_바이트_변환(TR구분(20), "test", 지금)

	바이트_변환_모음 := New바이트_변환_단순형(P변환형식_기본값, 값1)
	값2 := 바이트_변환_모음.G해석값_단순형().(*S질의값_바이트_변환)

	F테스트_같음(t, 값1.M구분, 값2.M구분)
	F테스트_같음(t, 값1.M코드, 값2.M코드)

	지금_복원값 := 값2.M바이트_변환.G해석값_단순형().(time.Time)
	F테스트_참임(t, 지금_복원값.Equal(지금), 지금_복원값, 지금)
}

func TestS바이트_변환_S질의값_바이트_변환_모음(t *testing.T) {
	지금 := time.Now()
	내일 := 지금.AddDate(0, 0, 1)
	값1 := New질의값_바이트_변환_모음(TR구분(20), "test", 지금, 내일)

	바이트_변환_모음 := New바이트_변환_모음_단순형(P변환형식_기본값, 값1)
	값2 := 바이트_변환_모음.G해석값_단순형(0).(*S질의값_바이트_변환_모음)

	F테스트_같음(t, 값1.M구분, 값2.M구분)
	F테스트_같음(t, 값1.M코드, 값2.M코드)

	지금_복원값 := 값2.M바이트_변환_모음.G해석값_단순형(0).(time.Time)
	F테스트_참임(t, 지금_복원값.Equal(지금), 지금_복원값, 지금)

	내일_복원값 := 값2.M바이트_변환_모음.G해석값_단순형(1).(time.Time)
	F테스트_참임(t, 내일_복원값.Equal(내일), 내일_복원값, 내일)
}
