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
	"bytes"
	"errors"
	"math"
	"math/big"
	"math/rand"
	"reflect"
	"time"
)

// 임의값 생성 관련
func F임의값_생성기() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
}

func F임의_참거짓() bool {
	return F임의값_생성기().Intn(2) == 0
}

func F임의_문자열(최소_길이, 최대_길이 int) string {
	if 최소_길이 >= 최대_길이 {
		panic(New에러with출력("최소 길이는 최대 길이보다 작아야 합니다. %v %v", 최소_길이, 최대_길이))
	}

	r := F임의값_생성기()
	길이 := r.Intn(최대_길이-최소_길이) + 최소_길이
	버퍼 := new(bytes.Buffer)
	문자열_모음 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"a", "b", "dll", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "z", "y",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "_", "+"}

	for i := 0; i < 길이; i++ {
		버퍼.WriteString(문자열_모음[r.Intn(len(문자열_모음))])
	}

	return 버퍼.String()
}

func F임의_양의_정수값() int {
	return F임의값_생성기().Int()
}

func F임의_양의_정수8값() int8 {
	return int8(F임의_범위_이내_정수값(0, 127))
}

func F임의_양의_정수64값() int64 {
	return int64(F임의_양의_정수값())
}

func F임의_정수값() int {
	정수값 := F임의_양의_정수값()

	if F임의_참거짓() {
		정수값 *= -1
	}

	return 정수값
}

func F임의_범위_이내_정수값(최소값, 최대값 int) int {
	F조건부_패닉(최소값 > 최대값, "최소값이 최대값보다 큽니다. 최소값 : %v, 최대값 : %v", 최소값, 최대값)

	추가값_비율 := F임의값_생성기().Float64()
	추가값 := int(추가값_비율 * float64(최대값-최소값))

	return 최소값 + 추가값
}

func F임의_정수64값() int64 {
	정수64값 := F임의_양의_정수64값()

	if F임의_참거짓() {
		정수64값 *= -1
	}

	return 정수64값
}

func F임의_범위_이내_정수64값(최소값, 최대값 int64) int64 {
	F조건부_패닉(최소값 > 최대값, "최소값이 최대값보다 큽니다. 최소값 : %v, 최대값 : %v", 최소값, 최대값)

	추가값_비율 := F임의값_생성기().Float64()
	추가값 := int64(추가값_비율 * float64(최대값-최소값))

	return 최소값 + 추가값
}

func F임의_범위_이내_정수64값_모음(길이 int, 최소값, 최대값 int64) []int64 {
	정수64값_모음 := make([]int64, 길이)

	for i := 0; i < 길이; i++ {
		정수64값_모음[i] = F임의_범위_이내_정수64값(최소값, 최대값)
	}

	return 정수64값_모음
}

func F임의_실수64() float64 {
	return F임의값_생성기().Float64() * math.Pow10(F임의_범위_이내_정수값(0, 20))
}

func F임의_범위_이내_실수64값(최소값, 최대값 float64) float64 {
	F조건부_패닉(최소값 > 최대값, "최소값이 최대값보다 큽니다. 최소값 : %v, 최대값 : %v", 최소값, 최대값)

	추가값_비율 := F임의값_생성기().Float64()
	추가값 := 추가값_비율 * (최대값 - 최소값)

	return 최소값 + 추가값
}

func F샘플_통화단위_모음() []T통화 {
	샘플_통화단위_모음 := make([]T통화, 0)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, KRW)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, USD)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, CNY)
	샘플_통화단위_모음 = append(샘플_통화단위_모음, EUR)

	return 샘플_통화단위_모음
}

func F임의_통화단위() T통화 {
	r := F임의값_생성기()
	통화단위_모음 := F샘플_통화단위_모음()

	return 통화단위_모음[r.Intn(len(통화단위_모음))]
}

func F임의_통화값() *S통화 {
	return F임의_통화값_모음(1)[0]
}

func F임의_통화값_모음(수량 int) []*S통화 {
	통화_모음 := make([]*S통화, 수량)
	통화단위_모음 := F샘플_통화단위_모음()
	r := F임의값_생성기()

	for i := 0; i < 수량; i++ {
		통화단위 := 통화단위_모음[r.Intn(len(통화단위_모음))]
		금액 := math.Trunc(r.Float64()*math.Pow10(r.Intn(5))*100) / 100

		통화_모음[i] = New통화(통화단위, 금액)
	}

	return 통화_모음
}

func F임의_시각() time.Time {
	r := F임의값_생성기()

	연도 := r.Intn(200) + 1970
	월 := time.Month(r.Intn(12))
	일 := r.Intn(31)
	시 := r.Intn(24)
	분 := r.Intn(60)
	초 := r.Intn(60)
	나노초 := r.Intn(1000000000)

	return time.Date(연도, 월, 일, 시, 분, 초, 나노초, time.Now().Location())
}

func F테스트용_임의_주소() T주소 {
	소켓_테스트용_주소_중복_방지_잠금.Lock()
	defer 소켓_테스트용_주소_중복_방지_잠금.Unlock()

	for {
		주소 := T주소(F임의값_생성기().Intn(60000))
		주소_문자열 := 주소.TCP주소()

		_, 중복 := 소켓_테스트용_주소_중복_방지_맵[주소_문자열]

		if 중복 {
			continue
		}

		소켓_테스트용_주소_중복_방지_맵[주소_문자열] = S비어있음{}

		return 주소
	}
}

type s샘플_구조체_1 struct {
	M정수   int
	M정수64 int64
	M실수64 float64
	M참거짓  bool
	M문자열  string
	//M시간   time.Time	// msgPack 포맷은 location 데이터를 상실.
}

// 간단한 테스트용 구조체
func F샘플_구조체_1() s샘플_구조체_1 {
	r := F임의값_생성기()

	s := s샘플_구조체_1{}
	s.M정수 = r.Int()
	s.M정수64 = r.Int63()
	s.M실수64 = r.Float64()
	s.M참거짓 = F임의_참거짓()
	s.M문자열 = F임의_문자열(5, 100)
	//s.M시간 = F임의_시각()	// msgPack포맷은 location 데이터를 상실.

	return s
}

type s샘플_구조체_2 struct {
	M슬라이스 []string
	M맵    map[string]int
	M구조체  s샘플_구조체_1
}

// 좀 더 복잡한 테스트용 구조체
func F샘플_구조체_2() s샘플_구조체_2 {
	r := F임의값_생성기()

	s := s샘플_구조체_2{}
	s.M슬라이스 = []string{F임의_문자열(3, 5), F임의_문자열(3, 5)}
	s.M맵 = map[string]int{
		F임의_문자열(5, 10): r.Int(),
		F임의_문자열(5, 10): r.Int(),
		F임의_문자열(5, 10): r.Int()}
	s.M구조체 = F샘플_구조체_1()

	return s
}

func f테스트용_안전한_전달값_모음() []interface{} {
	r := F임의값_생성기()

	바이트_전송값_1 := F확인(New바이트_변환(P변환형식_기본값, F샘플_구조체_1())).(*S바이트_변환)
	F조건부_패닉(바이트_전송값_1 == nil, "바이트_전송값 변환값이 nil임")

	바이트_전송값_2 := F확인(New바이트_변환(P변환형식_기본값, F샘플_구조체_2())).(*S바이트_변환)
	F조건부_패닉(바이트_전송값_2 == nil, "바이트_전송값 변환값이 nil임")

	안전한_전달값_모음 := []interface{}{
		r.Int(), uint(r.Int()), uintptr(r.Int()),
		int8(r.Intn(127)), int16(r.Intn(127)), r.Int31(), r.Int63(),
		uint8(r.Intn(127)), uint16(r.Intn(127)), uint32(r.Int31()), uint64(r.Int63()),
		r.Float32(), r.Float64(), F임의_참거짓(),
		F임의_문자열(5, 10), []byte(F임의_문자열(5, 10)),
		make(chan S비어있음), func() {},
		nil, errors.New(F임의_문자열(5, 10)), F임의_시각(),
		[]string{F임의_문자열(5, 10), F임의_문자열(5, 10)},
		바이트_전송값_1, 바이트_전송값_2}

	인터페이스_모음 := make([]interface{}, 10)
	for i := range 인터페이스_모음 {
		인터페이스_모음[i] = 안전한_전달값_모음[r.Intn(len(안전한_전달값_모음))]
	}

	안전한_전달값_모음 = append(안전한_전달값_모음, 인터페이스_모음)

	return 안전한_전달값_모음
}

func f테스트용_위험한_전달값_모음() []interface{} {
	r := F임의값_생성기()

	구조체_1 := F샘플_구조체_1()
	구조체_2 := F샘플_구조체_2()

	return []interface{}{
		구조체_1, &구조체_1,
		구조체_2, &구조체_2,
		big.NewInt(r.Int63()),
		big.NewRat(r.Int63(), r.Int63()),
		big.NewFloat(r.Float64())}
}

func f테스트용_변환가능한_전달값_모음() []interface{} {
	후보값_모음 := make([]interface{}, 0)
	후보값_모음 = append(후보값_모음, f테스트용_안전한_전달값_모음()...)
	후보값_모음 = append(후보값_모음, f테스트용_위험한_전달값_모음()...)

	변환가능한_값_모음 := make([]interface{}, 0)

	for _, 값 := range 후보값_모음 {
		switch 값.(type) {
		case []interface{}:
			continue
		}

		switch F종류(값) {
		case reflect.Func, reflect.Chan:
			continue
		}

		변환가능한_값_모음 = append(변환가능한_값_모음, 값)
	}

	return 변환가능한_값_모음
}

func f테스트용_변환형식_모음() []T변환 {
	return []T변환{JSON, MsgPack}
}

func F임의_변환_형식() T변환 {
	변환형식_모음 := []T변환{JSON, MsgPack}

	r := F임의값_생성기()

	return 변환형식_모음[r.Intn(len(변환형식_모음)-1)]
}

func F임의_시장_구분() T시장구분 {
	r := F임의값_생성기()
	시장_구분_모음 := []T시장구분{P시장구분_코스피, P시장구분_코스닥, P시장구분_코넥스, P시장구분_ETF}

	return 시장_구분_모음[r.Intn(len(시장_구분_모음))]
}

func F임의_샘플_종목() *S종목 {
	종목_모음 := F샘플_종목_모음_전체()

	return 종목_모음[F임의_범위_이내_정수값(0, len(종목_모음))].G복제본()
}

func F임의_샘플_종목_코스피_주식() *S종목 {
	종목_모음 := F샘플_종목_모음_코스피_주식()

	return 종목_모음[F임의_범위_이내_정수값(0, len(종목_모음))].G복제본()
}

func F임의_샘플_종목_코스닥_주식() *S종목 {
	종목_모음 := F샘플_종목_모음_코스닥_주식()

	return 종목_모음[F임의_범위_이내_정수값(0, len(종목_모음))].G복제본()
}

func F임의_샘플_종목_ETF() *S종목 {
	종목_모음 := F샘플_종목_모음_ETF()

	return 종목_모음[F임의_범위_이내_정수값(0, len(종목_모음))].G복제본()
}

func F중복_문자열_제거(문자열_모음 []string) (중복_제거된_문자열_모음 []string) {
	문자열_맵 := make(map[string]S비어있음)

	for _, 문자열 := range 문자열_모음 {
		문자열_맵[문자열] = S비어있음{}
	}

	중복_제거된_문자열_모음 = make([]string, len(문자열_맵))

	인덱스 := 0
	for 문자열 := range 문자열_맵 {
		중복_제거된_문자열_모음[인덱스] = 문자열
		인덱스++
	}

	return 중복_제거된_문자열_모음
}

func F중복_종목_제거(종목_모음 []*S종목) (중복_제거된_종목_모음 []*S종목) {
	종목_맵 := make(map[string]*S종목)
	for _, 종목 := range 종목_모음 {
		종목_맵[종목.G코드()] = 종목
	}

	중복_제거된_종목_모음 = make([]*S종목, len(종목_맵))
	인덱스 := 0
	for _, 종목 := range 종목_맵 {
		중복_제거된_종목_모음[인덱스] = 종목
		인덱스++
	}

	return 중복_제거된_종목_모음
}

func F종목_추출(종목_모음 []*S종목, 수량 int) (추출_종목_모음 []*S종목) {
	defer S예외처리{M함수: func() { 추출_종목_모음 = nil }}.S실행()

	종목_모음 = F중복_종목_제거(종목_모음)

	F조건부_패닉(len(종목_모음) < 수량, "종목 수량 : %v, 추출 수량 %v", len(종목_모음), 수량)

	r := F임의값_생성기()
	총길이 := len(종목_모음)
	추출_종목_맵 := make(map[string]*S종목)

	for {
		if len(추출_종목_맵) == 수량 {
			break
		}

		추출_종목 := 종목_모음[r.Intn(총길이)]
		추출_종목_맵[추출_종목.G코드()] = 추출_종목
	}

	추출_종목_모음 = make([]*S종목, len(추출_종목_맵))
	인덱스 := 0
	for _, 추출_종목 := range 추출_종목_맵 {
		추출_종목_모음[인덱스] = 추출_종목
		인덱스++
	}

	return 추출_종목_모음
}

func F종목코드_추출(종목_모음 []*S종목, 수량 int) (종목코드_모음 []string) {
	defer S예외처리{M함수: func() { 종목코드_모음 = nil }}.S실행()
	return F2종목코드_모음(F종목_추출(종목_모음, 수량))
}

func F샘플_종목_모음_전체() []*S종목 {
	종목_모음 := make([]*S종목, 0)
	종목_모음 = append(종목_모음, F샘플_종목_모음_코스피_주식()...)
	종목_모음 = append(종목_모음, F샘플_종목_모음_코스닥_주식()...)
	종목_모음 = append(종목_모음, F샘플_종목_모음_ETF()...)

	return 종목_모음
}

func F샘플_종목_모음_코스피_주식() []*S종목 {
	종목_모음 := make([]*S종목, 0)
	종목_모음 = append(종목_모음, New종목("000020", "동화약품", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("000040", "KR모터스", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("000050", "경방", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("000060", "메리츠화재", P시장구분_코스피))

	종목_모음 = append(종목_모음, New종목("005930", "삼성전자", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("005380", "현대차", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("015760", "한국전력", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("028260", "삼성물산", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("005935", "삼성전자우", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("012330", "현대모비스", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("000270", "기아차", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("000660", "SK하이닉스", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("032830", "삼성생명", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("090430", "아모레퍼시픽", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("035420", "NAVER", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("055550", "신한지주", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("051910", "LG화학", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("018260", "삼성에스디에스", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("017670", "SK텔레콤", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("034730", "SK", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("005490", "POSCO", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("000810", "삼성화재", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("051900", "LG생활건강", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("002790", "아모레G", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("003550", "LG", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("047810", "한국항공우주", P시장구분_코스피))
	종목_모음 = append(종목_모음, New종목("066570", "LG전자", P시장구분_코스피))

	return 종목_모음
}

func F샘플_종목_모음_코스닥_주식() []*S종목 {
	종목_모음 := make([]*S종목, 0)
	종목_모음 = append(종목_모음, New종목("068270", "셀트리온", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("035720", "카카오", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("026960", "동서", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("035760", "CJ ENM", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("086900", "메디톡스", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("084990", "헬릭스미스", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("034230", "파라다이스", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("041960", "코미팜", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("078340", "컴투스", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("102940", "코오롱생명과학", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("039030", "이오테크닉스", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("084110", "휴온스글로벌", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("046890", "서울반도체", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("096530", "씨젠", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("036490", "SK머티리얼즈", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("067080", "대화제약", P시장구분_코스닥))
	종목_모음 = append(종목_모음, New종목("034830", "한국토지신탁", P시장구분_코스닥))

	return 종목_모음
}

func F샘플_종목_모음_코스피200_ETF() []*S종목 {
	종목_모음 := []*S종목{
		New종목("069500", "KODEX 200", P시장구분_ETF),
		New종목("114800", "KODEX 인버스", P시장구분_ETF),
		New종목("122630", "KODEX 레버리지", P시장구분_ETF),
		New종목("252670", "KODEX 200 선물인버스2X", P시장구분_ETF),
		New종목("069660", "KOSEF 200", P시장구분_ETF),
		New종목("253250", "KOSEF 200 선물레버리지", P시장구분_ETF),
		New종목("253240", "KOSEF 200 선물인버스", P시장구분_ETF),
		New종목("253230", "KOSEF 200 선물인버스2X", P시장구분_ETF),
		New종목("102110", "TIGER 200", P시장구분_ETF),
		New종목("252710", "TIGER 200 선물인버스2X", P시장구분_ETF),
		New종목("105190", "KINDEX 200", P시장구분_ETF),
		New종목("108590", "TREX 200", P시장구분_ETF),
		New종목("148020", "KBSTAR 200", P시장구분_ETF),
		New종목("252400", "KBSTAR 200 선물레버리지", P시장구분_ETF),
		New종목("252410", "KBSTAR 200 선물인버스", P시장구분_ETF),
		New종목("252420", "KBSTAR 200 선물인버스2X", P시장구분_ETF),
		New종목("152100", "ARIRANG 200", P시장구분_ETF),
		New종목("253150", "ARIRANG 200 선물레버리지", P시장구분_ETF),
		New종목("253160", "ARIRANG 200 선물인버스2X", P시장구분_ETF)}

	return 종목_모음
}

func F샘플_종목_모음_ETF() []*S종목 {
	종목_모음 := make([]*S종목, 0)
	종목_모음 = append(종목_모음, New종목("069500", "KODEX 200", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("069660", "KOSEF 200", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("091160", "KODEX 반도체", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("091170", "KODEX 은행", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("091180", "KODEX 자동차", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("091220", "TIGER 은행", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("091230", "TIGER 반도체", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("098560", "TIGER 방송통신", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("099140", "KODEX China H", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("100910", "KOSEF KRX100", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("102110", "TIGER 200", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("102780", "KODEX 삼성그룹", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("102960", "KODEX 기계장비", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("102970", "KODEX 증권", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("104520", "KOSEF 블루칩", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("104530", "KOSEF 고배당", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("105010", "TIGER 라틴35", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("105190", "KINDEX 200", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("105780", "KBSTAR 5대그룹주", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("108450", "KINDEX 삼성그룹SW", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("108590", "TREX 200", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("114100", "KStar 국고채", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("114260", "KODEX 국고채", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("114460", "KINDEX 국고채", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("114470", "KOSEF 국고채", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("114800", "KODEX 인버스", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("114820", "TIGER 국채3", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("117460", "KODEX 에너지화학", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("117680", "KODEX 철강", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("117690", "TIGER 차이나", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("117700", "KODEX 건설", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("122090", "ARIRANG KOSPI50", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("122260", "KOSEF 통안채", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("122630", "KODEX 레버리지", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("123310", "TIGER 인버스", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("123320", "TIGER 레버리지", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("130680", "TIGER 원유선물(H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("130730", "KOSEF 단기자금", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("131890", "KINDEX 삼성그룹EW", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("132030", "KODEX 골드선물(H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("133690", "TIGER 나스닥100", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("136340", "KStar 우량회사채", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("137610", "TIGER 농산물선물(H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("137930", "마이다스 커버드콜", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("138230", "KOSEF 달러선물", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("138520", "TIGER 삼성그룹", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("138530", "TIGER LG그룹+", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("138540", "TIGER 현대차그룹+", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("138910", "KODEX 구리선물(H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("138920", "KODEX 콩선물(H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("139220", "TIGER 200 건설", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("139230", "TIGER 200 중공업", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("139240", "TIGER 200 철강소재", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("139250", "TIGER 200 에너지화학", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("139260", "TIGER 200 IT", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("139270", "TIGER 200 금융", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("139280", "TIGER 경기방어", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("139290", "TIGER 200 경기소비재", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("139310", "TIGER 금속선물(H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("139320", "TIGER 금은선물(H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("139660", "KOSEF 달러인버스선물", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("140570", "KStar 수출주", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("140580", "KStar 우량업종", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("140700", "KODEX 보험", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("140710", "KODEX 운송", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("140950", "파워 K100", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("143460", "KINDEX 밸류대형", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("143850", "TIGER S&P500선물(H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("143860", "TIGER 헬스케어", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("144600", "KODEX 은선물(H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("145670", "KINDEX 인버스", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("145850", "TREX 펀더멘탈 200", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("147970", "TIGER 모멘텀", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("148020", "KStar 200", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("148070", "KOSEF 10년 국고채", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("150460", "TIGER 중국소비테마", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("152100", "ARIRANG 200", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("152380", "KODEX 10년 국채선물", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("152500", "KINDEX 레버리지", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("152870", "파워 K200", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("153130", "KODEX 단기채권", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("153270", "iKon 100", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("156080", "KODEX MSCI KOREA", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("157450", "TIGER 유동자금", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("157490", "TIGER 소프트웨어", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("157500", "TIGER 증권", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("159800", "마이티 K100", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("160580", "TIGER 구리실물", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("161510", "ARIRANG 고배당주", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("166400", "TIGER 커버드C200", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("167860", "KOSEF 10년 국고채 레버리지", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("168300", "KTOP50", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("168580", "KINDEX 중국본토CSI30", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("169950", "KODEX 중국본토 A50", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("174350", "TIGER 로우볼", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("174360", "KStar 중국본토 대형", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("176710", "파워 국고채", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("176950", "KODEX 인버스국채선물", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("181480", "KINDEX 미국 리츠 부동산", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("182480", "TIGER US리츠(합성 H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("182490", "TIGER 단기선진하이일드", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("183700", "KStar 채권혼합", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("183710", "KStar 주식혼합", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("185680", "KODEX 미국바이오(합성)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("189400", "ARIRANG AC 월드(합성)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("190620", "KINDEX 단기자금", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("192090", "TIGER 차이나A300", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("192720", "파워고배당저변동성", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("195920", "TIGER 일본(합성 H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("195930", "TIGER 유로스탁스50(합성 H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("195970", "ARIRANG 선진국(합성 H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("195980", "ARIRANG 신흥국(합성 H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("196030", "KINDEX 일본레버리지(H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("196230", "KStar 단기통안채", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("200030", "KODEX 미국산업재(합성)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("200250", "KOSEF 인디아(합성)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("203780", "TIGER 나스닥 바이오", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("204450", "KODEX China H 레버리지", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("204480", "TIGER 차이나A레버리지", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("205720", "KINDEX 일본인버스(합성 H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("208470", "SMART MSCI선진국(합성 H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("210780", "TIGER 코스피 고배당", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("211560", "TIGER 배당성장", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("211900", "KODEX 배당성장", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("213610", "KODEX 삼성그룹밸류", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("213630", "ARIRANG 미국고배당주(합성 H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("214980", "KODEX 단기채권 PLUS", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("215620", "흥국 S&P 로우볼", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("217770", "TIGER 원유선물 인버스", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("217780", "TIGER 차이나A 인버스", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("217790", "TIGER 가격조정", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("218420", "KODEX 미국에너지(합성)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("219390", "KStar 미국원유생산기업(합성 H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("219480", "KODEX S&P500선물(H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("219900", "KINDEX 중국본토레버리지", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("220130", "SMART 중국본토 중소형 CSI500(합성 H)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("223190", "KODEX 200 내재가치", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("225030", "TIGER S&P500 인버스 선물", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("225040", "TIGER S&P500 레버리지", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("225050", "TIGER 유로스탁스 레버리지", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("225060", "TIGER 이머징마켓 레버리지", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("225130", "KINDEX 골드선물 레버리지", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("225800", "KOSEF 미국달러선물 레버리지(합성)", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("226380", "KINDEX 한류", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("226490", "KODEX 코스피", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("226980", "KODEX 200 중소형", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("227540", "TIGER 200 헬스케어", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("227550", "TIGER 200 산업재", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("227560", "TIGER 200 생활소비재", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("227570", "TIGER 우량가치", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("227830", "ARIRANG 코스피", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("228790", "TIGER 화장품", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("228800", "TIGER 여행레저", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("228810", "TIGER 미디어컨텐츠", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("228820", "TIGER KTOP30", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("229200", "KODEX 코스닥 150", P시장구분_ETF))
	종목_모음 = append(종목_모음, New종목("229720", "KODEX KTOP30", P시장구분_ETF))
	return 종목_모음
}
