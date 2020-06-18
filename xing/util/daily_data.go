package util

import (
	"github.com/ghts/ghts/lib"
	xt "github.com/ghts/ghts/xing/base"
	xing "github.com/ghts/ghts/xing/go"

	"database/sql"
	"time"
)

func F당일_일일_가격정보_수집(db *sql.DB) (에러 error) {
	if lib.F한국증시_동시호가_시간임() || lib.F한국증시_정규_거래_시간임() {
		lib.F문자열_출력("장 중에는 정확한 데이터를 수집할 수 없습니다.")
		return
	}

	lib.F일일_가격정보_테이블_생성(db)

	종목코드_모음_전체 := xing.F종목코드_모음_전체()
	당일 := lib.F2정수_일자(xing.F당일())

	for i := 0; i <= len(종목코드_모음_전체)/50; i++ {
		var 종목코드_모음 []string

		if len(종목코드_모음_전체) > (i+1)*50 {
			종목코드_모음 = 종목코드_모음_전체[i*50 : (i+1)*50]
		} else {
			종목코드_모음 = 종목코드_모음_전체[i*50:]
		}

		응답값_맵, 에러 := xing.TrT8407_현물_멀티_현재가_조회(종목코드_모음)
		lib.F확인(에러)

		for 종목코드, 값 := range 응답값_맵 {
			s := new(lib.S일일_가격정보)
			s.M종목코드 = 종목코드
			s.M일자 = 당일
			s.M시가 = float64(값.M시가)
			s.M고가 = float64(값.M고가)
			s.M저가 = float64(값.M저가)
			s.M종가 = float64(값.M현재가)
			s.M거래량 = float64(값.M누적_거래량)

			종목별_일일_가격정보_모음, 에러 := lib.New종목별_일일_가격정보_모음([]*lib.S일일_가격정보{s})
			lib.F확인(에러)

			lib.F확인(종목별_일일_가격정보_모음.DB저장(db))
		}
	}

	return nil
}

func F일주일_일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string) (에러 error) {
	var 시작일, 종료일 time.Time
	var 종목별_일일_가격정보_모음 *lib.S종목별_일일_가격정보_모음

	lib.F일일_가격정보_테이블_생성(db)

	for i, 종목코드 := range 종목코드_모음 {
		for {
			if xing.C32_재시작_실행_중.G값() {
				lib.F대기(lib.P10초)
				continue
			}

			break
		}

		종목별_일일_가격정보_모음, 에러 = lib.New종목별_일일_가격정보_모음_DB읽기(db, 종목코드)
		lib.F확인(에러)

		// 시작일 설정. 데이터 수량이 1개이나 100개이나 소요 시간은 비슷함.
		시작일 = lib.F지금().AddDate(0, 0, -14)

		// 종료일 설정
		if lib.F지금().After(xing.F당일().Add(15*lib.P1시간 + lib.P30분)) {
			종료일 = xing.F당일()
		} else {
			종료일 = xing.F전일()
		}

		f일일_가격정보_수집_도우미(db, 종목코드, 시작일, 종료일, 종목별_일일_가격정보_모음, i)
	}

	return nil
}

func F일일_가격정보_수집(db *sql.DB, 종목코드_모음 []string) (에러 error) {
	var 시작일, 종료일, 마지막_저장일 time.Time
	var 종목별_일일_가격정보_모음 *lib.S종목별_일일_가격정보_모음

	lib.F일일_가격정보_테이블_생성(db)

	for i, 종목코드 := range 종목코드_모음 {
		for {
			if xing.C32_재시작_실행_중.G값() {
				lib.F대기(lib.P10초)
				continue
			}

			break
		}

		종목별_일일_가격정보_모음, 에러 = lib.New종목별_일일_가격정보_모음_DB읽기(db, 종목코드)
		lib.F확인(에러)

		// 시작일 설정
		시작일 = lib.F지금().AddDate(-30, 0, 0)
		if 에러 == nil && len(종목별_일일_가격정보_모음.M저장소) > 0 {
			// lib.S종목별_일일_가격정보_모음 는 일자 순서로 정렬되어 있음.
			마지막_저장일 = 종목별_일일_가격정보_모음.M저장소[len(종목별_일일_가격정보_모음.M저장소)-1].G일자2()
			시작일 = 마지막_저장일.AddDate(0, 0, 1)
		}

		if 시작일.After(xing.F당일()) {
			lib.F문자열_출력("%v [%v] : 최신 데이터 업데이트.", i, 종목코드)
			continue
		}

		// 종료일 설정
		if lib.F지금().After(xing.F당일().Add(15*lib.P1시간 + lib.P30분)) {
			종료일 = xing.F당일()
		} else {
			종료일 = xing.F전일()
		}

		if 시작일.After(종료일) {
			continue
		} else if 시작일.Equal(종료일) { // 시작일과 종료일이 같으면 수천 개의 데이터를 불러오는 현상이 있음.
			시작일 = 시작일.AddDate(0, 0, -1)
		}

		// 데이터 수량이 1개이나 100개이나 소요 시간은 비슷함.
		if 시작일.After(lib.F금일().AddDate(0, 0, -14)) {
			시작일 = lib.F금일().AddDate(0, 0, -14)
		}

		f일일_가격정보_수집_도우미(db, 종목코드, 시작일, 종료일, 종목별_일일_가격정보_모음, i)
	}

	return nil
}

func f일일_가격정보_수집_도우미(db *sql.DB,
	종목코드 string, 시작일, 종료일 time.Time,
	종목별_일일_가격정보_모음 *lib.S종목별_일일_가격정보_모음, i int) {
	// 데이터 수집
	값_모음, 에러 := xing.TrT8413_현물_차트_일주월(종목코드, 시작일, 종료일, xt.P일주월_일)
	if 에러 != nil {
		lib.F에러_출력(에러)
		return
	} else if len(값_모음) == 0 {
		lib.F체크포인트(i, 종목코드, "추가 저장할 데이터가 없음.")
		return // 추가 저장할 데이터가 없음.
	}

	일일_가격정보_슬라이스 := make([]*lib.S일일_가격정보, len(값_모음))

	for i, 일일_데이터 := range 값_모음 {
		일일_가격정보_슬라이스[i] = lib.New일일_가격정보(
			일일_데이터.M종목코드,
			일일_데이터.M일자,
			일일_데이터.M시가,
			일일_데이터.M고가,
			일일_데이터.M저가,
			일일_데이터.M종가,
			일일_데이터.M거래량)
	}

	lib.F문자열_출력("%v %v %v~%v %v개", i, 종목코드, 시작일.Format(lib.P일자_형식), 종료일.Format(lib.P일자_형식), len(값_모음))

	종목별_일일_가격정보_모음, 에러 = lib.New종목별_일일_가격정보_모음(일일_가격정보_슬라이스)
	if 에러 != nil {
		lib.F에러_출력(에러)
		return
	}

	lib.F확인(종목별_일일_가격정보_모음.DB저장(db))
}
