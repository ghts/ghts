package xt

import (
	"github.com/ghts/ghts/lib"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func F전일_당일_설정(전일값, 당일값 time.Time) {
	전일 = lib.New안전한_시각(전일값)
	당일 = lib.New안전한_시각(당일값)
}

func F당일() time.Time {
	if 당일 == nil || 당일.G값().Equal(time.Time{}) {
		panic("Xing API가 초기화 되지 않았습니다.")
	}

	return 당일.G값()
}

func F전일() time.Time {
	if 전일 == nil || 전일.G값().Equal(time.Time{}) {
		panic("Xing API가 초기화 되지 않았습니다.")
	}

	return 전일.G값()
}

func f속성값_초기화(질의값 interface{}) interface{} {
	값 := reflect.ValueOf(질의값).Elem()

	for i := 0; i < 값.NumField(); i++ {
		switch {
		case !strings.HasPrefix(값.Type().Field(i).Name, "X_"),
			값.Field(i).Kind() != reflect.Uint8,
			!값.Field(i).CanSet():
			continue
		}

		값.Field(i).SetUint(0x20)
	}

	return 값.Interface()
}

func F서버_구분() T서버_구분 {
	if 구분값, 에러 := lib.F2정수(os.Getenv(P환경변수_서버_구분)); 에러 == nil && T서버_구분(구분값) == P서버_실거래 {
		return P서버_실거래
	} else {
		return P서버_모의투자
	}
}

func F서버_구분_설정(서버_구분 T서버_구분) {
	os.Setenv(P환경변수_서버_구분, strconv.Itoa(int(서버_구분)))
}

func F주소_DLL32() lib.T주소 {
	for {
		// 환경변수를 통하면 자동으로 자식 프로세스에 같은 값이 전달된다.
		if 주소, 에러 := lib.F2정수(os.Getenv(P환경변수_주소_TR)); 에러 != nil {
			F주소_설정()
		} else {
			return lib.T주소(주소)
		}
	}
}

func F주소_콜백() lib.T주소 {
	for {
		// 환경변수를 통하면 자동으로 자식 프로세스에 같은 값이 전달된다.
		if 주소, 에러 := lib.F2정수(os.Getenv(P환경변수_주소_콜백)); 에러 != nil {
			F주소_설정()
		} else {
			return lib.T주소(주소)
		}
	}
}

func F주소_실시간() lib.T주소 {
	for {
		// 환경변수를 통하면 자동으로 자식 프로세스에 같은 값이 전달된다.
		if 주소, 에러 := lib.F2정수(os.Getenv(P환경변수_주소_실시간)); 에러 != nil {
			F주소_설정()
		} else {
			return lib.T주소(주소)
		}
	}
}

func F주소_설정() {
	주소_설정_잠금.Lock()
	defer 주소_설정_잠금.Unlock()

	if 주소_설정_완료.G값() {
		return
	} else {
		defer 주소_설정_완료.S값(true)
	}

	for {
		임의_포트_번호 := lib.F임의_범위_이내_정수값(0, 30000)
		주소_TR := lib.T주소(임의_포트_번호)
		주소_콜백 := lib.T주소(임의_포트_번호 + 1)
		주소_실시간 := lib.T주소(임의_포트_번호 + 2)

		if lib.F포트_닫힘_확인(주소_TR) &&
			lib.F포트_닫힘_확인(주소_콜백) &&
			lib.F포트_닫힘_확인(주소_실시간) {
			// 환경변수를 통하면 자동으로 자식 프로세스에 같은 값이 전달된다.
			os.Setenv(P환경변수_주소_TR, strconv.Itoa(int(주소_TR)))
			os.Setenv(P환경변수_주소_콜백, strconv.Itoa(int(주소_콜백)))
			os.Setenv(P환경변수_주소_실시간, strconv.Itoa(int(주소_실시간)))

			return
		}
	}
}

func F주소_재설정() {
	주소_설정_완료.S값(false)
	F주소_설정()
}
