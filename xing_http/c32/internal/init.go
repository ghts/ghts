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

package x32_http

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	"time"
)

func init() {
	lib.TR구분_String = xt.TR구분_String
}

func F초기화() {
	ch초기화 := make(chan lib.T신호)
	go Go루틴_관리(ch초기화)
	<-ch초기화

	F서버_접속_및_로그인()
	f계좌_리스트_설정()
	f초기화_TR전송_제한()
}

func f종료_질의_송신() {
	defer lib.S예외처리{}.S실행()

	select {
	case <-xt.New질의(lib.New질의값_기본형(xt.TR종료, ""), Ch질의).Ch응답:
	case <-time.After(lib.P10초):
	}

	F종료_대기()
}

func F종료_대기() {
	<-Ch_HTTP_모듈_종료
	<-Ch함수_호출_모듈_종료
	<-Ch콜백_처리_모듈_종료
	<-Ch관리_모듈_종료
}

func F소켓_정리() error {
	lib.F패닉억제_호출(소켓PUB_실시간_정보.Close)

	for {
		if lib.F포트_닫힘_확인(xt.F주소_실시간()) {
			break
		}
	}

	lib.F대기(lib.P3초) // 소켓이 정리될 시간적 여유를 둠.

	return nil
}

// 접속 채널 질의 송신. 테스트용 접속이나 재접속 용도.
func F서버_접속_및_로그인() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	lib.F조건부_패닉(!lib.F인터넷에_접속됨(), "서버 접속이 불가 : 인터넷 접속을 확인하십시오.")

	select {
	case 응답 := <-xt.New질의(lib.New질의값_기본형(xt.TR접속_및_로그인, ""), Ch질의).Ch응답:
		if 응답.Error() != nil {
			lib.F문자열_출력("서버 접속 실패.")
			return 에러
		}
	case <-time.After(lib.P30초):
		return lib.New에러("접속 타임아웃")
	}

	select {
	case 로그인_여부 := <-Ch로그인:
		if !로그인_여부 {
			return lib.New에러with출력("로그인 실패.")
		}
	case <-time.After(lib.P30초):
		return lib.New에러with출력("로그인 타임아웃")
	}

	return nil
}

func f초기화_TR전송_제한() (에러 error) {
	defer lib.S예외처리{M에러: &에러}.S실행()

	TR코드_모음 := []string{
		xt.TR현물계좌_총평가_CSPAQ12200,
		xt.TR현물계좌_잔고내역_조회_CSPAQ12300,
		xt.TR현물계좌_주문체결내역_조회_CSPAQ13700,
		xt.TR현물계좌_예수금_주문가능금액_CSPAQ22200,
		xt.TR현물_정상_주문_CSPAT00600,
		xt.TR현물_정정_주문_CSPAT00700,
		xt.TR현물_취소_주문_CSPAT00800,
		xt.TR현물_당일_매매일지_t0150,
		xt.TR현물_일자별_매매일지_t0151,
		xt.TR시간_조회_t0167,
		xt.TR현물_체결_미체결_조회_t0425,
		xt.TR현물_호가_조회_t1101,
		xt.TR현물_시세_조회_t1102,
		xt.TR현물_기간별_조회_t1305,
		xt.TR현물_당일_전일_분틱_조회_t1310,
		xt.TR관리_불성실_투자유의_조회_t1404,
		xt.TR투자경고_매매정지_정리매매_조회_t1405,
		xt.TR_ETF_시간별_추이_t1902,
		xt.TR_ETF_LP호가_조회_t1906,
		xt.TR재무순위_종합_t3341,
		xt.TR현물_멀티_현재가_조회_t8407,
		xt.TR현물_차트_틱_t8411,
		xt.TR현물_차트_분_t8412,
		xt.TR현물_차트_일주월_t8413,
		xt.TR증시_주변_자금_추이_t8428,
		xt.TR현물_종목_조회_t8436,
		//xt.TR선물옵션_주문체결내역조회_CFOAQ00600,
		//xt.TR선물옵션_정상주문_CFOAT00100,
		//xt.TR선물옵션_정정주문_CFOAT00200,
		//xt.TR선물옵션_취소주문_CFOAT00300,
		//xt.TR선물옵션_예탁금_증거금_조회_CFOBQ10500,
		//xt.TR선물옵션_미결제약정_현황_CFOFQ02400,
		//xt.TR선물옵션_주문체결내역조회_CFOAQ00600,
		//xt.TR선물옵션_정상주문_CFOAT00100,
		//xt.TR선물옵션_정정주문_CFOAT00200,
		//xt.TR선물옵션_취소주문_CFOAT00300,
		//xt.TR선물옵션_예탁금_증거금_조회_CFOBQ10500,
		//xt.TR선물옵션_미결제약정_현황_CFOFQ02400,
		//xt.TR선물옵션_체결_미체결_조회_t0434,
		//xt.TR기업정보_요약_t3320,
		//xt.TR지수선물_마스터_조회_t8432,
	}

	TR코드_모음 = lib.F중복_문자열_제거(TR코드_모음)

	for {
		var 응답 *xt.S응답

		select {
		case 응답 = <-xt.New질의(lib.New질의값_문자열_모음(xt.TR코드별_전송_제한, "", TR코드_모음), Ch질의).Ch응답:
		case <-time.After(lib.P10초):
			return lib.New에러("f초기화_TR전송_제한() 타임아웃.")
		}

		if 응답.Error() != nil {
			lib.F에러_출력(응답.Error())
		} else if _, ok := 응답.V.(*xt.TR코드별_전송_제한_정보_모음); !ok {
			lib.New에러with출력("예상하지 못한 자료형. '%T'", 응답.V)
		} else if 전송_제한_정보_모음 = 응답.V.(*xt.TR코드별_전송_제한_정보_모음); len(TR코드_모음) != len(전송_제한_정보_모음.M맵) {
			lib.New에러with출력("서로 다른 길이 : '%v' '%v'", len(TR코드_모음), len(전송_제한_정보_모음.M맵))
		} else {
			정상 := false
			for _, 전송_제한_정보 := range 전송_제한_정보_모음.M맵 {
				if 전송_제한_정보.M초당_전송_제한 > 0 {
					정상 = true
					break
				}
			}

			if 정상 {
				break
			} else {
				lib.F대기(lib.P1초)
			}
		}
	}

	for TR코드, 전송_제한_정보 := range 전송_제한_정보_모음.M맵 {
		if 전송_제한_정보.M초_베이스 > 0 {
			if 전송_권한, 존재함 := tr코드별_전송_제한_1초[TR코드]; 존재함 {
				전송_권한.S수량_간격_변경(1, lib.P1초*time.Duration(전송_제한_정보.M초_베이스))
				tr코드별_전송_제한_1초[TR코드] = 전송_권한
			} else {
				tr코드별_전송_제한_1초[TR코드] = lib.New전송_권한(TR코드, 1, lib.P1초*time.Duration(전송_제한_정보.M초_베이스))
			}
		} else if 전송_제한_정보.M초당_전송_제한 > 0 {
			if 전송_권한, 존재함 := tr코드별_전송_제한_1초[TR코드]; 존재함 {
				전송_권한.S수량_간격_변경(전송_제한_정보.M초당_전송_제한, lib.P1초)
				tr코드별_전송_제한_1초[TR코드] = 전송_권한
			} else {
				tr코드별_전송_제한_1초[TR코드] = lib.New전송_권한(TR코드, 전송_제한_정보.M초당_전송_제한, lib.P1초)
			}
		}

		if 전송_제한_정보.M10분당_전송_제한 > 0 {
			if 전송_권한, 존재함 := tr코드별_전송_제한_10분[TR코드]; 존재함 {
				전송_권한.S수량_간격_변경(전송_제한_정보.M10분당_전송_제한, lib.P10분)
				tr코드별_전송_제한_10분[TR코드] = 전송_권한
			} else {
				전송_권한 = lib.New전송_권한(TR코드, 전송_제한_정보.M10분당_전송_제한, lib.P10분)

				// 지난 10분간 이미 전송한 수량을 반영.
				for i := 0; i < 전송_제한_정보.M10분간_전송한_수량; i++ {
					전송_권한.G획득()
				}

				tr코드별_전송_제한_10분[TR코드] = 전송_권한
			}
		}
	}

	return nil
}
