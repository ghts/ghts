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
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestF문자열_검색_복수_정규식(t *testing.T) {
	t.Parallel()

	검색_대상 := "aabbcc <span>xxx2006.01.02xxx</span> ddeeff"
	정규식_문자열_모음 := []string{
		`<span>.*[0-9]{4}.[0-9]{1,2}.[0-9]{1,2}.*</span>`,
		`[0-9]{4}.[0-9]{1,2}.[0-9]{1,2}`}

	검색_결과 := F정규식_검색(검색_대상, 정규식_문자열_모음)

	F테스트_같음(t, 검색_결과, "2006.01.02")
}

func TestF절대값(t *testing.T) {
	t.Parallel()

	F테스트_같음(t, F절대값_실수(-1), 1.0)
	F테스트_같음(t, F절대값_실수(1), 1.0)
	F테스트_같음(t, F절대값_실수(int64(-1)), 1.0)
	F테스트_같음(t, F절대값_실수(int64(1)), 1.0)
	F테스트_같음(t, F절대값_실수(float32(-1.0)), 1.0)
	F테스트_같음(t, F절대값_실수(float32(1.0)), 1.0)
	F테스트_같음(t, F절대값_실수(float64(-1.0)), 1.0)
	F테스트_같음(t, F절대값_실수(float64(1.0)), 1.0)
}

func TestF문자열_복사(t *testing.T) {
	t.Parallel()

	F테스트_같음(t, F문자열_복사("12 34 "), "12 34 ")
}

// 이하 최대 스레드 수량 관련 함수

func TestF단일_스레드_모드(t *testing.T) {
	최대_스레드_수량_원본 := runtime.GOMAXPROCS(-1)
	defer func() {
		runtime.GOMAXPROCS(최대_스레드_수량_원본)
	}()

	runtime.GOMAXPROCS(2)
	F단일_스레드_모드()

	F테스트_같음(t, runtime.GOMAXPROCS(-1), 1)
}

func TestF멀티_스레드_모드(t *testing.T) {
	최대_스레드_수량_원본 := runtime.GOMAXPROCS(-1)
	defer func() {
		runtime.GOMAXPROCS(최대_스레드_수량_원본)
	}()

	runtime.GOMAXPROCS(1)
	F멀티_스레드_모드()

	F테스트_같음(t, runtime.GOMAXPROCS(-1), runtime.NumCPU())
}

func TestF단일_스레드_모드임(t *testing.T) {
	최대_스레드_수량_원본 := runtime.GOMAXPROCS(-1)
	defer func() {
		runtime.GOMAXPROCS(최대_스레드_수량_원본)
	}()

	F단일_스레드_모드()
	F테스트_참임(t, F단일_스레드_모드임())

	F멀티_스레드_모드()
	F테스트_거짓임(t, F단일_스레드_모드임())
}

func TestF멀티_스레드_모드임(t *testing.T) {
	최대_스레드_수량_원본 := runtime.GOMAXPROCS(-1)
	defer func() {
		runtime.GOMAXPROCS(최대_스레드_수량_원본)
	}()

	F단일_스레드_모드()
	F테스트_거짓임(t, F멀티_스레드_모드임())

	F멀티_스레드_모드()
	F테스트_참임(t, F멀티_스레드_모드임())
}

func TestF실행파일_검색(t *testing.T) {
	var 파일명 string

	switch runtime.GOOS {
	case "windows":
		파일명 = "go.exe"
	default:
		파일명 = "go"
	}

	파일경로, 에러 := F실행파일_검색(파일명)
	F테스트_에러없음(t, 에러)
	F테스트_다름(t, strings.TrimSpace(파일경로), "")
	F테스트_참임(t, strings.HasSuffix(파일경로, 파일명))
}

func TestF파일경로_검색(t *testing.T) {
	var 파일명 string

	switch runtime.GOOS {
	case "windows":
		파일명 = "go.exe"
	default:
		파일명 = "go"
	}

	파일경로, 에러 := F파일_검색(filepath.Join(GOROOT(), "bin"), 파일명)
	F테스트_에러없음(t, 에러)
	F테스트_다름(t, strings.TrimSpace(파일경로), "")
	F테스트_참임(t, strings.HasSuffix(파일경로, 파일명))
}

func TestCSV파일에_값_저장_및_읽기(t *testing.T) {
	원본 := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"}}

	const 파일명 = "csv_test.csv"
	CSV쓰기(원본, 파일명, nil)
	복제본, 에러 := CSV읽기(파일명, ",", nil)
	F테스트_에러없음(t, 에러)
	F테스트_같음(t, len(복제본), len(원본))
	F테스트_같음(t, len(복제본[0]), len(원본[0]))

	for i := 0; i < len(원본); i++ {
		for j := 0; j < len(원본[i]); j++ {
			F테스트_같음(t, 원본[i][j], 복제본[i][j])
		}
	}
}

func TestF평균_표준편차(t *testing.T) {
	값_모음 := make([]float64, 100)

	for i := 0; i < 100; i++ {
		값_모음[i] = float64(i + 1)
	}

	F테스트_참임(t, F평균(값_모음)-50.5 < 0.001)
	F테스트_참임(t, F표준_편차(값_모음)-29.01149 < 0.001)
}
